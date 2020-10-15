package test

import (
	"time"

	costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"
)

// RunBuilder handles construction of runs
type RunBuilder interface {
	WithID(id int64) RunBuilder
	WithMeasuredTimeUTC(time time.Time) RunBuilder
	WithClusterCPUMillicore(cpu int) RunBuilder
	WithClusterMemoryMegaByte(mem int) RunBuilder
	WithResources(rr []costModels.RequiredResources) RunBuilder
	BuildRun() *costModels.Run
}

// RunBuilderStruct instance variables
type RunBuilderStruct struct {
	id   int64
	time time.Time
	cpu  int
	mem  int
	rr   []costModels.RequiredResources
}

// WithID sets id
func (rb *RunBuilderStruct) WithID(id int64) RunBuilder {
	rb.id = id
	return rb
}

// WithMeasuredTimeUTC sets measured time
func (rb *RunBuilderStruct) WithMeasuredTimeUTC(time time.Time) RunBuilder {
	rb.time = time
	return rb
}

// WithClusterCPUMillicore sets cpu millicore
func (rb *RunBuilderStruct) WithClusterCPUMillicore(cpu int) RunBuilder {
	rb.cpu = cpu
	return rb
}

// WithClusterMemoryMegaByte sets memory
func (rb *RunBuilderStruct) WithClusterMemoryMegaByte(mem int) RunBuilder {
	rb.mem = mem
	return rb
}

// WithResources sets resources
func (rb *RunBuilderStruct) WithResources(resources []costModels.RequiredResources) RunBuilder {
	rb.rr = resources
	return rb
}

// BuildRun builds the run
func (rb *RunBuilderStruct) BuildRun() *costModels.Run {
	return &costModels.Run{
		ClusterCPUMillicore:   rb.cpu,
		ClusterMemoryMegaByte: rb.mem,
		ID:                    rb.id,
		MeasuredTimeUTC:       rb.time,
		Resources:             rb.rr,
	}
}

// NewRunBuilder Constructor for run builder
func NewRunBuilder() RunBuilder {
	return &RunBuilderStruct{}
}

// ARun Constructor for RunBuilder with test data
func ARun() RunBuilder {
	resources := AListOfRequiredResources()

	builder := NewRunBuilder().
		WithID(1).
		WithMeasuredTimeUTC(time.Now()).
		WithResources(resources.BuildResources()).
		WithClusterCPUMillicore(1000).
		WithClusterMemoryMegaByte(750)

	return builder
}
