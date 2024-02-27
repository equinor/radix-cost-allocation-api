package repository

import (
	"fmt"
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type gormRepository struct {
	db *gorm.DB
}

func GetSqlServerDsn(server, database, userID, password string, port int) string {
	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s",
		server, userID, password, database)

	if port > 0 {
		dsn = fmt.Sprintf("%s;port=%d", dsn, port)
	}

	return dsn
}

func OpenGormSqlServerDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(sqlserver.Open(dsn), &gorm.Config{DisableAutomaticPing: true, Logger: NewLogger()})
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
