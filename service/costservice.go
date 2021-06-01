package service

import (
	"time"

	"github.com/equinor/radix-cost-allocation-api/models"
)

type CostService interface {
	GetCostForPeriod(from, to time.Time) (*models.ApplicationCostSet, error)
	GetFutureCost(appName string) (*models.ApplicationCost, error)
}
