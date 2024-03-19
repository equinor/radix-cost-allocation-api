package repository

import (
	"fmt"
	"time"

	commongorm "github.com/equinor/radix-common/pkg/gorm"
	"github.com/equinor/radix-cost-allocation-api/models"
	"github.com/microsoft/go-mssqldb/azuread"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRepository struct {
	db *gorm.DB
}

func OpenGormSqlServerDB(server, database string, port int) (*gorm.DB, error) {
	dsn := fmt.Sprintf("server=%s;database=%s;port=%d;fedauth=ActiveDirectoryDefault", server, database, port)

	dialector := sqlserver.New(sqlserver.Config{
		DriverName: azuread.DriverName,
		DSN:        dsn,
	})

	return gorm.Open(dialector, &gorm.Config{
		DisableAutomaticPing: false,
		Logger:               commongorm.NewLogger(),
	})
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) GetContainers(from, to time.Time) ([]models.ContainerDto, error) {
	var containers []models.ContainerDto

	tx := r.db.
		Preload("Node.NodePool").
		Where(clause.Lt{Column: "started_at", Value: to}).
		Where(clause.Gt{Column: "last_known_running_at", Value: from}).
		Find(&containers)
	return containers, tx.Error
}

func (r *gormRepository) GetNodePoolCost() ([]models.NodePoolCostDto, error) {
	var cost []models.NodePoolCostDto

	tx := r.db.
		Preload(clause.Associations).
		Find(&cost)
	return cost, tx.Error
}

func (r *gormRepository) GetNodePools() ([]models.NodePoolDto, error) {
	var nodes []models.NodePoolDto

	tx := r.db.
		Find(&nodes)
	return nodes, tx.Error
}
