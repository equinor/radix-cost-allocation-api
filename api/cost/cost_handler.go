package cost

import (
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
)

type CostHandler interface {
	GetFutureCost(appName string) (*costModels.ApplicationCost, error)
	GetTotalCost(fromTime, toTime *time.Time, appName *string) (*costModels.ApplicationCostSet, error)
}
