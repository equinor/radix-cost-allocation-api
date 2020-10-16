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

// RunBuilderList handles construction of run lists
type RunBuilderList interface {
	WithBuilders(builders []RunBuilder) RunBuilderList
	BuildRuns() []costModels.Run
}

// RunBuilderStruct instance variables
type RunBuilderStruct struct {
	id   int64
	time time.Time
	cpu  int
	mem  int
	rr   []costModels.RequiredResources
}

// RunBuilderListStruct instance variables
type RunBuilderListStruct struct {
	builders []RunBuilder
}

// WithBuilders sets the runs
func (rbl *RunBuilderListStruct) WithBuilders(builders []RunBuilder) RunBuilderList {
	rbl.builders = builders
	return rbl
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

// BuildRuns build the run list
func (rbl *RunBuilderListStruct) BuildRuns() []costModels.Run {
	var runs []costModels.Run

	for _, r := range rbl.builders {
		runs = append(runs, *r.BuildRun())
	}

	return runs
}

// NewRunBuilder Constructor for run builder
func NewRunBuilder() RunBuilder {
	return &RunBuilderStruct{}
}

// NewRunBuilderList constructor for run builder list
func NewRunBuilderList() RunBuilderList {
	return &RunBuilderListStruct{}
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

// AListOfRuns constructor for RunBuilderList with test data
func AListOfRuns() RunBuilderList {
	builders := []RunBuilder{
		NewRunBuilder().
			WithID(1).
			WithMeasuredTimeUTC(time.Now()).
			WithClusterCPUMillicore(1000).
			WithClusterMemoryMegaByte(750),
		NewRunBuilder().
			WithID(1).
			WithMeasuredTimeUTC(time.Now()).
			WithClusterCPUMillicore(1000).
			WithClusterMemoryMegaByte(750),
		NewRunBuilder().
			WithID(1).
			WithMeasuredTimeUTC(time.Now()).
			WithClusterCPUMillicore(1000).
			WithClusterMemoryMegaByte(750),
	}

	for index := range builders {
		resources := AListOfRequiredResources().BuildResources()
		builders[index].WithResources(resources)
	}

	builderList := NewRunBuilderList().WithBuilders(builders)
	return builderList
}
