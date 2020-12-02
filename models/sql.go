package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"

	"log"
	"time"
)

// CostRepository interface
type CostRepository interface {
	GetLatestRun() (costModels.Run, error)
	GetRunsBetweenTimes(from, to *time.Time) ([]costModels.Run, error)
}

// DBCredentials hold credentials for database
type DBCredentials struct {
	Server   string
	Port     int
	Database string
	UserID   string
	Password string
}

// SQLCostRepository struct defines a connection to DB
type SQLCostRepository struct {
	db *sql.DB
}

// NewSQLCostRepository initializes new connection to database
func NewSQLCostRepository(creds *DBCredentials) *SQLCostRepository {
	dbConnection := creds.setupDBConnection()
	return &SQLCostRepository{
		dbConnection,
	}
}

// GetLatestRun fetches the last run and joins them with resources logged for that run
func (dbCon *SQLCostRepository) GetLatestRun() (costModels.Run, error) {
	var requiredResources []costModels.RequiredResources
	var lastRun costModels.Run

	ctx, err := dbCon.verifyConnection()
	if err != nil {
		return costModels.Run{}, err
	}

	query :=
		"set nocount on; " +
			" WITH temptable AS " +
			" (SELECT [run_id],[wbs],[application],[cpu_millicores],[memory_mega_bytes],[replicas] " +
			" FROM [cost].[required_resources] " +
			" WHERE [run_id] IN ( SELECT MAX(run_id) FROM [cost].[required_resources] ))" +
			" SELECT rr.application, rr.cpu_millicores, rr.memory_mega_bytes, rr.wbs, rr.replicas, r.cluster_cpu_millicores, r.cluster_memory_mega_bytes, r.measured_time_utc FROM [cost].[runs] r " +
			" INNER JOIN temptable rr ON r.[id] = rr.[run_id] "

	// Create new connection to database
	connection, err := dbCon.db.Conn(ctx)
	if err != nil {
		log.Fatal("Error creating connection to DB", err.Error())
	}
	defer connection.Close()

	rows, err := connection.QueryContext(ctx, query)

	if err != nil {
		return costModels.Run{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var application, environment, component, wbs string
		var cpuMillicores, memoryMegaBytes, replicas, clusterCPUMillicore, clusterMemoryMegabyte int
		var measuredTime time.Time

		err := rows.Scan(
			&application,
			&cpuMillicores,
			&memoryMegaBytes,
			&wbs,
			&replicas,
			&clusterCPUMillicore,
			&clusterMemoryMegabyte,
			&measuredTime,
		)

		if err != nil {
			return costModels.Run{}, err
		}

		resource := costModels.RequiredResources{
			Application:     application,
			Environment:     environment,
			Component:       component,
			CPUMillicore:    cpuMillicores,
			MemoryMegaBytes: memoryMegaBytes,
			Replicas:        replicas,
			WBS:             wbs,
		}

		lastRun.ClusterCPUMillicore = clusterCPUMillicore
		lastRun.ClusterMemoryMegaByte = clusterMemoryMegabyte

		requiredResources = append(requiredResources, resource)
	}

	lastRun.Resources = requiredResources

	return lastRun, nil

}

// GetRunsBetweenTimes get all runs with its resources between from and to time, and optionally a specific application
//
// If appName is nil then runs for all applications are returned
func (dbCon *SQLCostRepository) GetRunsBetweenTimes(from, to *time.Time) ([]costModels.Run, error) {
	runsResources := map[int64]*[]costModels.RequiredResources{}
	runs := map[int64]costModels.Run{}
	ctx, err := dbCon.verifyConnection()
	if err != nil {
		return nil, err
	}

	tsql :=
		"set nocount on; " +
			" SELECT r.id run_id, r.measured_time_utc," +
			" COALESCE(r.cluster_cpu_millicores, 0) AS cluster_cpu_millicores," +
			" COALESCE(r.cluster_memory_mega_bytes, 0) AS cluster_memory_mega_bytes," +
			" rr.id, rr.wbs, rr.application, rr.environment, rr.component, rr.cpu_millicores, rr.memory_mega_bytes, rr.replicas" +
			" FROM [cost].[runs] r" +
			" JOIN [cost].[required_resources] rr ON r.id = rr.run_id" +
			" WHERE r.measured_time_utc BETWEEN @from AND @to"

	// Create new connection to database
	connection, err := dbCon.db.Conn(ctx)
	if err != nil {
		log.Fatal("Error creating connection to DB", err.Error())
	}
	defer connection.Close()

	// Execute query
	args := []interface{}{
		sql.Named("from", from),
		sql.Named("to", to),
	}
	rows, err := connection.QueryContext(ctx, tsql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set.
	for rows.Next() {
		var measuredTimeUTC time.Time
		var runID, id int64
		var clusterCPUMillicores, clusterMemoryMegaBytes, cpuMillicores, memoryMegaBytes, replicas int
		var wbs, application, environment, component string
		var run costModels.Run

		// Get values from row.
		err := rows.Scan(&runID,
			&measuredTimeUTC,
			&clusterCPUMillicores,
			&clusterMemoryMegaBytes,
			&id,
			&wbs,
			&application,
			&environment,
			&component,
			&cpuMillicores,
			&memoryMegaBytes,
			&replicas,
		)
		if err != nil {
			return nil, err
		}

		resource := costModels.RequiredResources{
			ID:              id,
			WBS:             wbs,
			Application:     application,
			Environment:     environment,
			Component:       component,
			CPUMillicore:    cpuMillicores,
			MemoryMegaBytes: memoryMegaBytes,
			Replicas:        replicas,
		}

		if run = runs[runID]; run.ID == 0 {
			resources := []costModels.RequiredResources{resource}
			runsResources[runID] = &resources
			run = costModels.Run{
				ID:                    runID,
				MeasuredTimeUTC:       measuredTimeUTC,
				ClusterCPUMillicore:   clusterCPUMillicores,
				ClusterMemoryMegaByte: clusterMemoryMegaBytes,
			}
			runs[runID] = run
		} else {
			resources := *runsResources[runID]
			resources = append(resources, resource)
			runsResources[runID] = &resources
		}
	}

	runsAsArray := make([]costModels.Run, len(runs))
	runEntryIndex := 0
	for key, val := range runs {
		val.Resources = *runsResources[key]
		runsAsArray[runEntryIndex] = val
		runEntryIndex++
	}

	return runsAsArray, nil
}

// CloseDB closes the underlying db connection - Only to be called when API exits
func (dbCon *SQLCostRepository) CloseDB() {
	dbCon.db.Close()
}

// SetupDBConnection sets up db connection
func (creds *DBCredentials) setupDBConnection() *sql.DB {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		creds.Server, creds.UserID, creds.Password, creds.Port, creds.Database)

	var err error

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!\n")
	return db
}

func (dbCon *SQLCostRepository) verifyConnection() (context.Context, error) {
	ctx := context.Background()
	var err error

	if dbCon.db == nil {
		err = errors.New("CreateRun: db is null")
		return ctx, err
	}

	// Check if database is alive.
	err = dbCon.db.PingContext(ctx)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
