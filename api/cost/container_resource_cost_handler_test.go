package cost

import (
	"sort"
	"testing"
	"time"

	models "github.com/equinor/radix-cost-allocation-api/models"
	mockrepository "github.com/equinor/radix-cost-allocation-api/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetTotalCost(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mockrepository.NewMockRepository(ctrl)

	from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC)
	h := NewContainerResourceCostHandler(repo)
	cost, err := h.GetTotalCost(from, to)
	assert.Nil(t, err)
	assert.NotNil(t, cost)
}

func Test_NodePoolCostByFromAndTo(t *testing.T) {
	from1 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to1 := time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)

	from2 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to2 := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

	from3 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	to3 := time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)

	from4 := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	to4 := time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC)

	from5 := time.Date(2019, 1, 2, 0, 0, 0, 0, time.UTC)
	to5 := time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)

	values := []models.NodePoolCostDto{
		{FromDate: from1, ToDate: to1},
		{FromDate: from2, ToDate: to2},
		{FromDate: from3, ToDate: to3},
		{FromDate: from4, ToDate: to4},
		{FromDate: from5, ToDate: to5},
	}

	expected := []models.NodePoolCostDto{
		{FromDate: from5, ToDate: to5},
		{FromDate: from2, ToDate: to2},
		{FromDate: from3, ToDate: to3},
		{FromDate: from1, ToDate: to1},
		{FromDate: from4, ToDate: to4},
	}

	sort.Sort(SortByFromAndTo(values))
	assert.Equal(t, expected, values)
}

func Test_isCostConnected(t *testing.T) {
	/*
	 c1 |----------|
	 c2       |----------|
	 c3             |----------|
	 c4                  |----------|
	 c1 overlaps c2
	 c1 does not overlap c3
	 c2 overlaps c3
	*/

	duration := time.Duration(10 * time.Hour)
	c1From := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	c1To := c1From.Add(duration)
	c2From := c1From.Add(duration / 2)
	c2To := c2From.Add(duration)
	c3From := c1To
	c3To := c3From.Add(duration)
	c4From := c2To
	c4To := c4From.Add(duration)

	c1 := models.NodePoolCostDto{FromDate: c1From, ToDate: c1To}
	c2 := models.NodePoolCostDto{FromDate: c2From, ToDate: c2To}
	c3 := models.NodePoolCostDto{FromDate: c3From, ToDate: c3To}
	c4 := models.NodePoolCostDto{FromDate: c4From, ToDate: c4To}

	assert.False(t, isCostConnected(c1, c2))
	assert.True(t, isCostConnected(c1, c3))
	assert.False(t, isCostConnected(c2, c3))
	assert.False(t, isCostConnected(c1, c4))
}

func Test_isCostEncapsulated(t *testing.T) {
	/*
		c1  |------|
		c2  |----------|
		c3     |-----|
		c4     |-----|
	*/

	c1From := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	c1To := c1From.Add(6 * 24 * time.Hour)
	c2From := c1From
	c2To := c2From.Add(10 * 24 * time.Hour)
	c3From := c1From.Add(3 * 24 * time.Hour)
	c3To := c3From.Add(5 * 24 * time.Hour)
	c4From := c3From
	c4To := c3To
	c1 := models.NodePoolCostDto{FromDate: c1From, ToDate: c1To}
	c2 := models.NodePoolCostDto{FromDate: c2From, ToDate: c2To}
	c3 := models.NodePoolCostDto{FromDate: c3From, ToDate: c3To}
	c4 := models.NodePoolCostDto{FromDate: c4From, ToDate: c4To}

	assert.True(t, isCostEncapsulated(c1, c2))
	assert.False(t, isCostEncapsulated(c1, c3))
	assert.False(t, isCostEncapsulated(c2, c1))
	assert.True(t, isCostEncapsulated(c3, c2))
	assert.True(t, isCostEncapsulated(c3, c4))
}

func Test_adjustCostPeriod(t *testing.T) {

	c := models.NodePoolCostDto{
		FromDate: time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
		ToDate:   time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
		Cost:     1000,
	}

	assert.Equal(t, int32(3000), adjustCostPeriod(c, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)).Cost)
	assert.Equal(t, int32(500), adjustCostPeriod(c, time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)).Cost)

}

func Test_filterNodePoolCostByPoolId(t *testing.T) {
	cost := []models.NodePoolCostDto{
		{NodePoolId: 1},
		{NodePoolId: 2},
		{NodePoolId: 2},
	}

	actual := filterNodePoolCostByPoolId(cost, 2)
	expected := []models.NodePoolCostDto{
		{NodePoolId: 2},
		{NodePoolId: 2},
	}
	assert.ElementsMatch(t, expected, actual)
}

func Test_adjustNodePoolCostTimeRange(t *testing.T) {
	/*
		Range (25)         |--------------------
		c1 (1)     |-
		c2 (3)            |---
		c3 (5)              |-----
		c4 (3)                  |---
		c5 (8)                  |--------
		c6 (2)                             |--
		c7 (10)                               |----------
		c8 (1)                                         |-
	*/
	day := time.Duration(24 * time.Hour)
	periodDuration := time.Duration(20 * day)
	periodFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	periodTo := periodFrom.Add(periodDuration)
	c1From := periodFrom.Add(-8 * day)
	c1To := c1From.Add(1 * day)
	c2From := periodFrom.Add(-1 * day)
	c2To := c2From.Add(3 * day)
	c3From := periodFrom.Add(1 * day)
	c3To := c3From.Add(5 * day)
	c4From := periodFrom.Add(5 * day)
	c4To := c4From.Add(3 * day)
	c5From := periodFrom.Add(5 * day)
	c5To := c5From.Add(8 * day)
	c6From := periodFrom.Add(16 * day)
	c6To := c6From.Add(2 * day)
	c7From := periodFrom.Add(19 * day)
	c7To := c7From.Add(10 * day)
	c8From := periodFrom.Add(28 * day)
	c8To := c8From.Add(1 * day)

	cost := []models.NodePoolCostDto{
		{Id: 1, Cost: 100, FromDate: c1From, ToDate: c1To},
		{Id: 2, Cost: 300, FromDate: c2From, ToDate: c2To},
		{Id: 3, Cost: 500, FromDate: c3From, ToDate: c3To},
		{Id: 4, Cost: 300, FromDate: c4From, ToDate: c4To},
		{Id: 5, Cost: 800, FromDate: c5From, ToDate: c5To},
		{Id: 6, Cost: 200, FromDate: c6From, ToDate: c6To},
		{Id: 7, Cost: 1000, FromDate: c7From, ToDate: c7To},
		{Id: 8, Cost: 100, FromDate: c8From, ToDate: c8To},
	}

	expect := []models.NodePoolCostDto{
		{Id: 2, FromDate: periodFrom, ToDate: c3From, Cost: 100},
		{Id: 3, FromDate: c3From, ToDate: c4From, Cost: 400},
		{Id: 5, FromDate: c5From, ToDate: c6From, Cost: 1100},
		{Id: 6, FromDate: c6From, ToDate: c7From, Cost: 300},
		{Id: 7, FromDate: c7From, ToDate: periodTo, Cost: 100},
	}

	actual := adjustNodePoolCostTimeRange(periodFrom, periodTo, cost)
	assert.ElementsMatch(t, expect, actual)
}
