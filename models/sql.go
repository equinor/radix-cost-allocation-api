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

// SQLClient used to perform sql queries
type SQLClient struct {
	Server   string
	Port     int
	Database string
	UserID   string
	Password string
	db       *sql.DB
}

// NewSQLClient create sqlclient and setup the db connection
func NewSQLClient(server, database string, port int, userID, password string) SQLClient {
	sqlClient := SQLClient{
		Server:   server,
		Database: database,
		Port:     port,
		UserID:   userID,
		Password: password,
	}
	sqlClient.db = sqlClient.setupDBConnection()
	return sqlClient
}

// GetLatestRun fetches the last run and joins them with resources logged for that run
func (sqlClient SQLClient) GetLatestRun() (costModels.Run, error) {
	var requiredResources []costModels.RequiredResources
	var lastRun costModels.Run

	ctx, err := sqlClient.verifyConnection()
	if err != nil {
		return costModels.Run{}, err
	}

	query :=
		" WITH temptable AS " +
			" (SELECT [run_id],[wbs],[application],[cpu_millicores],[memory_mega_bytes],[replicas] " +
			" FROM [cost].[required_resources] " +
			" WHERE [run_id] IN ( SELECT MAX(run_id) FROM [cost].[required_resources] )) " +
			" SELECT rr.application, rr.cpu_millicores, rr.memory_mega_bytes, rr.wbs, rr.replicas, r.cluster_cpu_millicores, r.cluster_memory_mega_bytes, r.measured_time_utc FROM [cost].[runs] r " +
			" INNER JOIN temptable rr ON r.[id] = rr.[run_id] "

	rows, err := sqlClient.db.QueryContext(ctx, query)

	if err != nil {
		return costModels.Run{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var application, environment, component, wbs string
		var cpuMillicores, memoryMegaBytes, replicas, clusterCPUMillicore, clusterMemoryMegabyte int

		err := rows.Scan(
			&application,
			&environment,
			&component,
			&cpuMillicores,
			&memoryMegaBytes,
			&replicas,
			&clusterCPUMillicore,
			&clusterMemoryMegabyte,
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

// GetRunsBetweenTimes get all runs with its resources between from and to time
func (sqlClient SQLClient) GetRunsBetweenTimes(from, to *time.Time) ([]costModels.Run, error) {
	runsResources := map[int64]*[]costModels.RequiredResources{}
	runs := map[int64]costModels.Run{}
	ctx, err := sqlClient.verifyConnection()
	if err != nil {
		return nil, err
	}

	tsql := "SELECT r.id run_id, r.measured_time_utc," +
		" COALESCE(r.cluster_cpu_millicores, 0) AS cluster_cpu_millicores," +
		" COALESCE(r.cluster_memory_mega_bytes, 0) AS cluster_memory_mega_bytes," +
		" rr.id, rr.wbs, rr.application, rr.environment, rr.component, rr.cpu_millicores, rr.memory_mega_bytes, rr.replicas" +
		" FROM [cost].[runs] r" +
		" JOIN [cost].[required_resources] rr ON r.id = rr.run_id" +
		" WHERE measured_time_utc BETWEEN @from AND @to;"

	// Execute query
	rows, err := sqlClient.db.QueryContext(ctx, tsql, sql.Named("from", from), sql.Named("to", to))
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

// SaveRequiredResources inserts all required resources under run.Resources
func (sqlClient SQLClient) SaveRequiredResources(run costModels.Run) error {
	tsql := `INSERT INTO cost.required_resources (run_id, wbs, application, environment, component, cpu_millicores, memory_mega_bytes, replicas) 
	VALUES (@runId, @wbs, @application, @environment, @component, @cpuMillicores, @memoryMegaBytes, @replicas); select convert(bigint, SCOPE_IDENTITY());`
	for _, req := range run.Resources {
		_, err := sqlClient.execSQL(tsql,
			sql.Named("runId", run.ID),
			sql.Named("wbs", req.WBS),
			sql.Named("application", req.Application),
			sql.Named("environment", req.Environment),
			sql.Named("component", req.Component),
			sql.Named("cpuMillicores", req.CPUMillicore),
			sql.Named("memoryMegaBytes", req.MemoryMegaBytes),
			sql.Named("replicas", req.Replicas))
		if err != nil {
			return fmt.Errorf("Failed to insert req resources %v", err)
		}
	}
	return nil
}

// SaveRun inserts a new run, returns id
func (sqlClient SQLClient) SaveRun(measuredTime time.Time, clusterCPUMillicores, clusterMemoryMegaBytes int) (int64, error) {
	tsql := `INSERT INTO cost.runs (measured_time_utc, cluster_cpu_millicores, cluster_memory_mega_bytes) 
	VALUES (@measuredTimeUTC, @clusterCPUMillicores, @clusterMemoryMegaBytes); select convert(bigint, SCOPE_IDENTITY());`
	return sqlClient.execSQL(tsql,
		sql.Named("measuredTimeUTC", measuredTime),
		sql.Named("clusterCPUMillicores", clusterCPUMillicores),
		sql.Named("clusterMemoryMegaBytes", clusterMemoryMegaBytes))
}

// Close the underlying db connection
func (sqlClient SQLClient) Close() {
	sqlClient.db.Close()
}

// SetupDBConnection sets up db connection
func (sqlClient SQLClient) setupDBConnection() *sql.DB {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		sqlClient.Server, sqlClient.UserID, sqlClient.Password, sqlClient.Port, sqlClient.Database)

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

func (sqlClient SQLClient) execSQL(tsql string, args ...interface{}) (int64, error) {
	ctx, err := sqlClient.verifyConnection()
	if err != nil {
		return -1, err
	}

	stmt, err := sqlClient.db.Prepare(tsql)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, args...)
	var newID int64
	err = row.Scan(&newID)
	if err != nil {
		return -1, err
	}

	return newID, nil
}

func (sqlClient SQLClient) verifyConnection() (context.Context, error) {
	ctx := context.Background()
	var err error

	if sqlClient.db == nil {
		err = errors.New("CreateRun: db is null")
		return ctx, err
	}

	// Check if database is alive.
	err = sqlClient.db.PingContext(ctx)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}
