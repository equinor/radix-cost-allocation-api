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

	assert.Equal(t, float64(3000), adjustCostPeriod(c, time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC)).Cost)
	assert.Equal(t, float64(500), adjustCostPeriod(c, time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC), time.Date(2020, 1, 2, 12, 0, 0, 0, time.UTC)).Cost)
}

func Test_findNodePoolCostByPoolId(t *testing.T) {
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

	t.Run("multiple cost", func(t *testing.T) {
		t.Parallel()
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
	})

	t.Run("join past outside cost with inside", func(t *testing.T) {
		t.Parallel()
		/*
			Range (25)         |--------------------
			c1 (1)     |-
			c2 (3)                  |---
		*/

		day := time.Duration(24 * time.Hour)
		periodDuration := time.Duration(20 * day)
		periodFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		periodTo := periodFrom.Add(periodDuration)
		c1From := periodFrom.Add(-8 * day)
		c1To := c1From.Add(1 * day)
		c2From := periodFrom.Add(5 * day)
		c2To := c2From.Add(3 * day)

		cost := []models.NodePoolCostDto{
			{Id: 1, Cost: 100, FromDate: c1From, ToDate: c1To},
			{Id: 2, Cost: 300, FromDate: c2From, ToDate: c2To},
		}

		expect := []models.NodePoolCostDto{
			{Id: 1, FromDate: periodFrom, ToDate: c2From, Cost: 500},
			{Id: 2, FromDate: c2From, ToDate: periodTo, Cost: 1500},
		}

		actual := adjustNodePoolCostTimeRange(periodFrom, periodTo, cost)
		assert.ElementsMatch(t, expect, actual)
	})

	t.Run("join future outside cost with inside - only inside is used", func(t *testing.T) {
		t.Parallel()
		/*
			Range (25)         |--------------------
			c1 (1)                                   |-
			c2 (3)                  |---
		*/

		day := time.Duration(24 * time.Hour)
		periodDuration := time.Duration(20 * day)
		periodFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		periodTo := periodFrom.Add(periodDuration)
		c1From := periodFrom.Add(22 * day)
		c1To := c1From.Add(1 * day)
		c2From := periodFrom.Add(5 * day)
		c2To := c2From.Add(3 * day)

		cost := []models.NodePoolCostDto{
			{Id: 1, Cost: 100, FromDate: c1From, ToDate: c1To},
			{Id: 2, Cost: 300, FromDate: c2From, ToDate: c2To},
		}

		expect := []models.NodePoolCostDto{
			{Id: 2, FromDate: periodFrom, ToDate: periodTo, Cost: 2000},
		}

		actual := adjustNodePoolCostTimeRange(periodFrom, periodTo, cost)
		assert.ElementsMatch(t, expect, actual)
	})

	t.Run("use past outside cost", func(t *testing.T) {
		t.Parallel()
		/*
			Range (25)         |--------------------
			c1 (1)          |-
		*/

		day := time.Duration(24 * time.Hour)
		periodDuration := time.Duration(20 * day)
		periodFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		periodTo := periodFrom.Add(periodDuration)
		c1From := periodFrom.Add(-3 * day)
		c1To := c1From.Add(1 * day)

		cost := []models.NodePoolCostDto{
			{Id: 1, Cost: 100, FromDate: c1From, ToDate: c1To},
		}

		expect := []models.NodePoolCostDto{
			{Id: 1, FromDate: periodFrom, ToDate: periodTo, Cost: 2000},
		}

		actual := adjustNodePoolCostTimeRange(periodFrom, periodTo, cost)
		assert.ElementsMatch(t, expect, actual)
	})

	t.Run("use future outside cost", func(t *testing.T) {
		t.Parallel()
		/*
			Range (25)         |--------------------
			c1 (1)                                   |-
		*/

		day := time.Duration(24 * time.Hour)
		periodDuration := time.Duration(20 * day)
		periodFrom := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		periodTo := periodFrom.Add(periodDuration)
		c1From := periodFrom.Add(22 * day)
		c1To := c1From.Add(1 * day)

		cost := []models.NodePoolCostDto{
			{Id: 1, Cost: 100, FromDate: c1From, ToDate: c1To},
		}

		expect := []models.NodePoolCostDto{
			{Id: 1, FromDate: periodFrom, ToDate: periodTo, Cost: 2000},
		}

		actual := adjustNodePoolCostTimeRange(periodFrom, periodTo, cost)
		assert.ElementsMatch(t, expect, actual)
	})
}

func Test_isContainerRunningInNodePoolCost(t *testing.T) {
	/*
		cost       |--------------------|
		c1   |-----|
		c2   |----------|
		c3   |-------------------------------|
		c4              |---------|
		c5                           |-------|
		c6                              |----|
	*/

	cost := models.NodePoolCostDto{FromDate: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), ToDate: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}
	c1 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}
	c2 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC)}
	c3 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC)}
	c4 := models.ContainerDto{StartedAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 14, 0, 0, 0, 0, time.UTC)}
	c5 := models.ContainerDto{StartedAt: time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)}
	c6 := models.ContainerDto{StartedAt: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)}

	assert.False(t, isContainerRunningInNodePoolCost(cost, c1))
	assert.True(t, isContainerRunningInNodePoolCost(cost, c2))
	assert.True(t, isContainerRunningInNodePoolCost(cost, c3))
	assert.True(t, isContainerRunningInNodePoolCost(cost, c4))
	assert.True(t, isContainerRunningInNodePoolCost(cost, c5))
	assert.False(t, isContainerRunningInNodePoolCost(cost, c6))
}

func Test_getContainerDurationInNodePoolCost(t *testing.T) {
	/*
		cost       |--------------------|
		c1   |-----|
		c2   |----------|
		c3   |-------------------------------|
		c4              |---------|
		c5                           |-------|
		c6                              |----|
	*/
	day := 24 * time.Hour
	cost := models.NodePoolCostDto{FromDate: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), ToDate: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC)}
	c1 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC)}
	c2 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC)}
	c3 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC)}
	c4 := models.ContainerDto{StartedAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)}
	c5 := models.ContainerDto{StartedAt: time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)}
	c6 := models.ContainerDto{StartedAt: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC)}

	assert.Equal(t, day*0, getContainerDurationInNodePoolCost(cost, c1))
	assert.Equal(t, day*2, getContainerDurationInNodePoolCost(cost, c2))
	assert.Equal(t, day*10, getContainerDurationInNodePoolCost(cost, c3))
	assert.Equal(t, day*3, getContainerDurationInNodePoolCost(cost, c4))
	assert.Equal(t, day*2, getContainerDurationInNodePoolCost(cost, c5))
	assert.Equal(t, day*0, getContainerDurationInNodePoolCost(cost, c6))
}

func Test_getContainerResourcesUsageInNodePoolCost(t *testing.T) {
	/*
		cost       |--------------------|
		c1   |-----|
		c2   |----------|
		c3   |-------------------------------|
		c4              |---------|
		c5                           |-------|
		c6                              |----|
	*/
	cpuReq, memReq, nodePoolId := int64(15), int64(5), int32(1)
	day := 24 * time.Hour
	cost := models.NodePoolCostDto{FromDate: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), ToDate: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), NodePoolId: nodePoolId}
	c1 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	c2 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	c3 := models.ContainerDto{StartedAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	c4 := models.ContainerDto{StartedAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	c5 := models.ContainerDto{StartedAt: time.Date(2020, 1, 18, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	c6 := models.ContainerDto{StartedAt: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC), MemoryRequestedBytes: memReq, CpuRequestedMillicores: cpuReq, Node: &models.NodeDto{NodePoolId: &nodePoolId}}
	cMissingNode := models.ContainerDto{ContainerId: "c_id"}
	cMissingNodePoolId := models.ContainerDto{ContainerId: "c_id", Node: &models.NodeDto{}}

	cpuSec, memSec, _ := getContainerResourcesUsageInNodePoolCost(cost, c1)
	assert.Equal(t, float64(0), cpuSec)
	assert.Equal(t, float64(0), memSec)
	cpuSec, memSec, _ = getContainerResourcesUsageInNodePoolCost(cost, c2)
	assert.Equal(t, (2*day).Seconds()*float64(cpuReq), cpuSec)
	assert.Equal(t, (2*day).Seconds()*float64(memReq), memSec)
	cpuSec, memSec, _ = getContainerResourcesUsageInNodePoolCost(cost, c3)
	assert.Equal(t, (10*day).Seconds()*float64(cpuReq), cpuSec)
	assert.Equal(t, (10*day).Seconds()*float64(memReq), memSec)
	cpuSec, memSec, _ = getContainerResourcesUsageInNodePoolCost(cost, c4)
	assert.Equal(t, (3*day).Seconds()*float64(cpuReq), cpuSec)
	assert.Equal(t, (3*day).Seconds()*float64(memReq), memSec)
	cpuSec, memSec, _ = getContainerResourcesUsageInNodePoolCost(cost, c5)
	assert.Equal(t, (2*day).Seconds()*float64(cpuReq), cpuSec)
	assert.Equal(t, (2*day).Seconds()*float64(memReq), memSec)
	cpuSec, memSec, _ = getContainerResourcesUsageInNodePoolCost(cost, c6)
	assert.Equal(t, float64(0), cpuSec)
	assert.Equal(t, float64(0), memSec)
	_, _, err := getContainerResourcesUsageInNodePoolCost(cost, cMissingNode)
	assert.Equal(t, ContainerMissingNodeError("c_id"), err)
	_, _, err = getContainerResourcesUsageInNodePoolCost(cost, cMissingNodePoolId)
	assert.Equal(t, ContainerMissingNodePoolIdError("c_id"), err)
}

func Test_getAllocatedResourcesForNodePoolCost(t *testing.T) {
	/*
		cost         |----------------------|
		c1    |--|
		c2         |----|
		c3         |----------------------------|
		c4                 |----|
		c5                                |-----|
		c6                                        |----|
		other node pool:
		c7             |-----|

	*/
	day := time.Hour * 24
	nodePoolId1 := int32(1)
	nodePoolId2 := int32(2)
	cost := models.NodePoolCostDto{FromDate: time.Date(2020, 1, 10, 0, 0, 0, 0, time.UTC), ToDate: time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC),
		NodePoolId: 1, Cost: 1000, Currency: "NOK"}
	c1 := models.ContainerDto{ContainerId: "c1", StartedAt: time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 7, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 11, MemoryRequestedBytes: 21, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c2 := models.ContainerDto{ContainerId: "c2", StartedAt: time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 12, MemoryRequestedBytes: 22, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c3 := models.ContainerDto{ContainerId: "c3", StartedAt: time.Date(2020, 1, 8, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 23, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 13, MemoryRequestedBytes: 23, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c4 := models.ContainerDto{ContainerId: "c4", StartedAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 14, MemoryRequestedBytes: 24, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c5 := models.ContainerDto{ContainerId: "c5", StartedAt: time.Date(2020, 1, 16, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 15, MemoryRequestedBytes: 25, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c6 := models.ContainerDto{ContainerId: "c6", StartedAt: time.Date(2020, 1, 22, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 24, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 16, MemoryRequestedBytes: 26, Node: &models.NodeDto{NodePoolId: &nodePoolId1}}
	c7 := models.ContainerDto{ContainerId: "c7", StartedAt: time.Date(2020, 1, 12, 0, 0, 0, 0, time.UTC), LastKnownRunningAt: time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC),
		CpuRequestedMillicores: 17, MemoryRequestedBytes: 27, Node: &models.NodeDto{NodePoolId: &nodePoolId2}}
	actual, err := getAllocatedResourcesForNodePoolCost(cost, []models.ContainerDto{c1, c2, c3, c4, c5, c6, c7})
	expectedTotalCpu := 2*day*12 + 10*day*13 + 3*day*14 + 4*day*15 // c2+c3+c4+c5
	expectedTotalMem := 2*day*22 + 10*day*23 + 3*day*24 + 4*day*25 // c2+c3+c4+c5
	expectedContainerResources := []ContainerResourceUsage{
		{ContainerId: "c2", CPUMillicoreSeconds: (2 * day * 12).Seconds(), MemoryBytesSeconds: (2 * day * 22).Seconds()},
		{ContainerId: "c3", CPUMillicoreSeconds: (10 * day * 13).Seconds(), MemoryBytesSeconds: (10 * day * 23).Seconds()},
		{ContainerId: "c4", CPUMillicoreSeconds: (3 * day * 14).Seconds(), MemoryBytesSeconds: (3 * day * 24).Seconds()},
		{ContainerId: "c5", CPUMillicoreSeconds: (4 * day * 15).Seconds(), MemoryBytesSeconds: (4 * day * 25).Seconds()},
	}
	assert.Nil(t, err)
	assert.Equal(t, expectedTotalCpu.Seconds(), actual.CPUMillicoreSeconds)
	assert.Equal(t, expectedTotalMem.Seconds(), actual.MemoryBytesSeconds)
	assert.Equal(t, float64(1000), actual.Cost)
	assert.Equal(t, "NOK", actual.Currency)
	assert.Len(t, actual.ContainerResources, 4)
	assert.ElementsMatch(t, expectedContainerResources, actual.ContainerResources)
}

func Test_calculateContainerResourceCost(t *testing.T) {
	actual := calculateContainerResourceCost(1000, 2000, 100, 50, 5000)
	assert.Equal(t, float64(312.5), actual)
}

func Test_calculateNodePoolContainerResourceCost(t *testing.T) {
	cost := NodePoolCostAllocatedResources{
		Cost:                10000,
		Currency:            "NOK",
		CPUMillicoreSeconds: 200,
		MemoryBytesSeconds:  1000,
		ContainerResources: []ContainerResourceUsage{
			{ContainerId: "c1", CPUMillicoreSeconds: 50, MemoryBytesSeconds: 900},
			{ContainerId: "c2", CPUMillicoreSeconds: 150, MemoryBytesSeconds: 100},
		},
	}
	expected := []ContainerCost{
		{ContainerId: "c1", Cost: Cost{Value: 5750, Currency: "NOK"}},
		{ContainerId: "c2", Cost: Cost{Value: 4250, Currency: "NOK"}},
	}
	actual := calculateNodePoolContainerResourceCost(cost)
	assert.ElementsMatch(t, expected, actual)
}

func Test_aggregateContainerCost(t *testing.T) {
	cost1 := ContainerCost{ContainerId: "c1", Cost: Cost{Value: 100, Currency: "NOK"}}
	cost2 := ContainerCost{ContainerId: "c1", Cost: Cost{Value: 100, Currency: "NOK"}}
	cost3 := ContainerCost{ContainerId: "c1", Cost: Cost{Value: 100, Currency: "NOK"}}
	cost4 := ContainerCost{ContainerId: "c1", Cost: Cost{Value: 100, Currency: "NOK"}}
	cost5 := ContainerCost{ContainerId: "c2", Cost: Cost{Value: 100, Currency: "NOK"}}
	cost6 := ContainerCost{ContainerId: "c2", Cost: Cost{Value: 100, Currency: "NOK"}}
	c1 := models.ContainerDto{ContainerId: "c1"}
	c2 := models.ContainerDto{ContainerId: "c2"}

	expected := []ContainerTotalCost{
		{Container: &c1, Cost: Cost{Currency: "NOK", Value: 400}},
		{Container: &c2, Cost: Cost{Currency: "NOK", Value: 200}},
	}

	actual, err := aggregateContainerCost([]ContainerCost{cost1, cost2, cost3, cost4, cost5, cost6}, []models.ContainerDto{c1, c2})
	assert.Nil(t, err)
	assert.ElementsMatch(t, expected, actual)

}
