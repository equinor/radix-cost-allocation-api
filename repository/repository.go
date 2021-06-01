package repository

import (
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
)

type Repository interface {
	GetNodePools() ([]models.NodePoolDto, error)
	GetContainers(from, to time.Time) ([]models.ContainerDto, error)
	GetNodePoolCost() ([]models.NodePoolCostDto, error)
}
