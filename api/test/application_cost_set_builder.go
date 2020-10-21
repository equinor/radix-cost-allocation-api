package test

import (
	"fmt"
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
)

// ApplicationCostSetBuilder handles construction of applicationcostset
type ApplicationCostSetBuilder interface {
	WithFromTime(from time.Time) ApplicationCostSetBuilder
	WithToTime(to time.Time) ApplicationCostSetBuilder
	WithApplicationCosts(applicationCosts []ApplicationCostBuilder) ApplicationCostSetBuilder
	WithTotalRequestedCPU(totalCPU int) ApplicationCostSetBuilder
	WithTotalRequestedMemory(totalMemory int) ApplicationCostSetBuilder
	BuildApplicationCostSet() *costModels.ApplicationCostSet
}

// ApplicationCostBuilder handles construction of applicationcost
type ApplicationCostBuilder interface {
	WithName(name string) ApplicationCostBuilder
	WithOwner(owner string) ApplicationCostBuilder
	WithCreator(creator string) ApplicationCostBuilder
	WithWBS(wbs string) ApplicationCostBuilder
	WithCostPercentageByCPU(percentage float64) ApplicationCostBuilder
	WithCostPercentageByMemory(percentage float64) ApplicationCostBuilder
	WithComment(comment string) ApplicationCostBuilder
	WithCost(cost float64) ApplicationCostBuilder
	WithCurrency(currency string) ApplicationCostBuilder
	BuildApplicationCost() *costModels.ApplicationCost
}

// ApplicationCostSetBuilderStruct instance variables
type ApplicationCostSetBuilderStruct struct {
	From                 time.Time
	To                   time.Time
	ApplicationCosts     []ApplicationCostBuilder
	TotalRequestedCPU    int
	TotalRequestedMemory int
}

// ApplicationCostBuilderStruct instance variables
type ApplicationCostBuilderStruct struct {
	Name                   string
	Owner                  string
	Creator                string
	WBS                    string
	CostPercentageByCPU    float64
	CostPercentageByMemory float64
	Comment                string
	Cost                   float64
	Currency               string
}

// WithName sets name
func (acb *ApplicationCostBuilderStruct) WithName(name string) ApplicationCostBuilder {
	acb.Name = name
	return acb
}

// WithCreator sets creator
func (acb *ApplicationCostBuilderStruct) WithCreator(creator string) ApplicationCostBuilder {
	acb.Creator = creator
	return acb
}

// WithOwner sets owner
func (acb *ApplicationCostBuilderStruct) WithOwner(owner string) ApplicationCostBuilder {
	acb.Owner = owner
	return acb
}

// WithWBS sets wbs
func (acb *ApplicationCostBuilderStruct) WithWBS(wbs string) ApplicationCostBuilder {
	acb.WBS = wbs
	return acb
}

// WithCostPercentageByCPU sets cost percentage by CPU
func (acb *ApplicationCostBuilderStruct) WithCostPercentageByCPU(percentage float64) ApplicationCostBuilder {
	acb.CostPercentageByCPU = percentage
	return acb
}

// WithCostPercentageByMemory sets cost percentage by memory
func (acb *ApplicationCostBuilderStruct) WithCostPercentageByMemory(percentage float64) ApplicationCostBuilder {
	acb.CostPercentageByMemory = percentage
	return acb
}

// WithComment sets comment
func (acb *ApplicationCostBuilderStruct) WithComment(comment string) ApplicationCostBuilder {
	acb.Comment = comment
	return acb
}

// WithCost sets cost
func (acb *ApplicationCostBuilderStruct) WithCost(cost float64) ApplicationCostBuilder {
	acb.Cost = cost
	return acb
}

// WithCurrency sets currency
func (acb *ApplicationCostBuilderStruct) WithCurrency(currency string) ApplicationCostBuilder {
	acb.Currency = currency
	return acb
}

// BuildApplicationCost builds the ApplicationCost model
func (acb *ApplicationCostBuilderStruct) BuildApplicationCost() *costModels.ApplicationCost {
	return &costModels.ApplicationCost{
		Comment:                acb.Comment,
		Cost:                   acb.Cost,
		CostPercentageByCPU:    acb.CostPercentageByCPU,
		CostPercentageByMemory: acb.CostPercentageByMemory,
		Creator:                acb.Creator,
		WBS:                    acb.WBS,
		Owner:                  acb.Owner,
		Currency:               acb.Currency,
		Name:                   acb.Name,
	}
}

// WithFromTime sets from time
func (acsb *ApplicationCostSetBuilderStruct) WithFromTime(from time.Time) ApplicationCostSetBuilder {
	acsb.From = from
	return acsb
}

// WithToTime sets to time
func (acsb *ApplicationCostSetBuilderStruct) WithToTime(to time.Time) ApplicationCostSetBuilder {
	acsb.To = to
	return acsb
}

// WithApplicationCosts sets application costs
func (acsb *ApplicationCostSetBuilderStruct) WithApplicationCosts(applicationCosts []ApplicationCostBuilder) ApplicationCostSetBuilder {
	acsb.ApplicationCosts = applicationCosts
	return acsb
}

// WithTotalRequestedCPU sets total requested CPU
func (acsb *ApplicationCostSetBuilderStruct) WithTotalRequestedCPU(totalCPU int) ApplicationCostSetBuilder {
	acsb.TotalRequestedCPU = totalCPU
	return acsb
}

// WithTotalRequestedMemory sets total requested memory
func (acsb *ApplicationCostSetBuilderStruct) WithTotalRequestedMemory(totalMemory int) ApplicationCostSetBuilder {
	acsb.TotalRequestedMemory = totalMemory
	return acsb
}

// BuildApplicationCostSet builds the ApplicationCostSet model
func (acsb *ApplicationCostSetBuilderStruct) BuildApplicationCostSet() *costModels.ApplicationCostSet {
	appCosts := make([]costModels.ApplicationCost, 0)
	for _, appCost := range acsb.ApplicationCosts {
		builtAppCost := appCost.BuildApplicationCost()
		appCosts = append(appCosts, *builtAppCost)
	}

	return &costModels.ApplicationCostSet{
		From:                 acsb.From,
		To:                   acsb.To,
		ApplicationCosts:     appCosts,
		TotalRequestedCPU:    acsb.TotalRequestedCPU,
		TotalRequestedMemory: acsb.TotalRequestedMemory,
	}
}

// NewApplicationCostBuilder constructor for application cost builder
func NewApplicationCostBuilder() ApplicationCostBuilder {
	return &ApplicationCostBuilderStruct{}
}

// NewApplicationCostSetBuilder constructor for application cost set builder
func NewApplicationCostSetBuilder() ApplicationCostSetBuilder {
	return &ApplicationCostSetBuilderStruct{}
}

// AnApplicationCost constructor for ApplicationCostBuilder with test data
func AnApplicationCost() ApplicationCostBuilder {
	builder := NewApplicationCostBuilder().
		WithName("any-app").
		WithWBS("A.BCD.00.999").
		WithComment("Test comment").
		WithCost(2459).
		WithCostPercentageByCPU(1.2).
		WithCostPercentageByMemory(2.2).
		WithCreator("Radix").
		WithCurrency("NOK").
		WithOwner("Radix")

	return builder
}

// AnApplicationCostSet constructor for ApplicationCostSetBuilder with test data
func AnApplicationCostSet() ApplicationCostSetBuilder {
	numOfAppCosts := 4
	appCostBuilders := make([]ApplicationCostBuilder, 0)
	for i := 0; i < numOfAppCosts; i++ {
		appCostBuilders = append(appCostBuilders, AnApplicationCost().WithName(fmt.Sprintf("any-app-%d", i)))
	}

	builder := NewApplicationCostSetBuilder().
		WithFromTime(time.Now()).
		WithToTime(time.Now()).
		WithTotalRequestedCPU(100).
		WithTotalRequestedMemory(200).
		WithApplicationCosts(appCostBuilders)

	return builder
}
