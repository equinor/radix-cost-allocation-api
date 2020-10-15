package test

import costModels "github.com/equinor/radix-cost-allocation-api/api/cost/models"

// ResourceBuilder construct RequiredResources
type ResourceBuilder interface {
	WithID(id int64) ResourceBuilder
	WithWBS(string string) ResourceBuilder
	WithApplication(application string) ResourceBuilder
	WithEnvironment(environment string) ResourceBuilder
	WithComponent(component string) ResourceBuilder
	WithCPUMillicore(cpu int) ResourceBuilder
	WithMemoryMegaBytes(mem int) ResourceBuilder
	WithReplicas(replicas int) ResourceBuilder
	BuildResource() *costModels.RequiredResources
	GetCPUMillicore() int
	GetMemoryMegaByte() int
}

// ResourceBuilderList constructs lists of RequiredResources
type ResourceBuilderList interface {
	WithBuilders(builders []ResourceBuilder) ResourceBuilderList
	BuildResources() []costModels.RequiredResources
	GetResourcesCPUMillicore() int
	GetResourcesMemoryMegabyte() int
}

// ResourceBuilderStruct Instance variables
type ResourceBuilderStruct struct {
	id              int64
	wbs             string
	application     string
	environment     string
	component       string
	cpuMillicore    int
	memoryMegaBytes int
	replicas        int
}

// ResourceBuilderListStruct Instance variables
type ResourceBuilderListStruct struct {
	builders []ResourceBuilder
}

// WithBuilders set builders
func (rbl *ResourceBuilderListStruct) WithBuilders(builders []ResourceBuilder) ResourceBuilderList {
	rbl.builders = builders
	return rbl
}

// BuildResources builds all resources in builders list
func (rbl *ResourceBuilderListStruct) BuildResources() []costModels.RequiredResources {
	var resources []costModels.RequiredResources

	for _, rr := range rbl.builders {
		resources = append(resources, *rr.BuildResource())
	}

	return resources
}

// GetResourcesCPUMillicore finds the total requested cpu for all required resources in list
func (rbl *ResourceBuilderListStruct) GetResourcesCPUMillicore() int {
	totalCPU := 0

	for _, rr := range rbl.builders {
		totalCPU += rr.GetCPUMillicore()
	}

	return totalCPU
}

// GetResourcesMemoryMegabyte finds the total requested memory for all required resources in list
func (rbl *ResourceBuilderListStruct) GetResourcesMemoryMegabyte() int {
	totalMem := 0

	for _, rr := range rbl.builders {
		totalMem += rr.GetMemoryMegaByte()
	}

	return totalMem
}

// GetCPUMillicore returns cpu millicore
func (rb *ResourceBuilderStruct) GetCPUMillicore() int {
	return rb.cpuMillicore
}

// GetMemoryMegaByte returns memory
func (rb *ResourceBuilderStruct) GetMemoryMegaByte() int {
	return rb.memoryMegaBytes
}

// WithID sets ID
func (rb *ResourceBuilderStruct) WithID(id int64) ResourceBuilder {
	rb.id = id
	return rb
}

// WithWBS sets wbs
func (rb *ResourceBuilderStruct) WithWBS(wbs string) ResourceBuilder {
	rb.wbs = wbs
	return rb
}

// WithApplication sets application
func (rb *ResourceBuilderStruct) WithApplication(app string) ResourceBuilder {
	rb.application = app
	return rb
}

// WithEnvironment sets environment
func (rb *ResourceBuilderStruct) WithEnvironment(env string) ResourceBuilder {
	rb.environment = env
	return rb
}

// WithComponent sets component
func (rb *ResourceBuilderStruct) WithComponent(component string) ResourceBuilder {
	rb.component = component
	return rb
}

// WithCPUMillicore sets cpumillicore
func (rb *ResourceBuilderStruct) WithCPUMillicore(cpu int) ResourceBuilder {
	rb.cpuMillicore = cpu
	return rb
}

// WithMemoryMegaBytes sets memory
func (rb *ResourceBuilderStruct) WithMemoryMegaBytes(mem int) ResourceBuilder {
	rb.memoryMegaBytes = mem
	return rb
}

// WithReplicas sets replicas
func (rb *ResourceBuilderStruct) WithReplicas(replicas int) ResourceBuilder {
	rb.replicas = replicas
	return rb
}

// BuildResource builds the resource
func (rb *ResourceBuilderStruct) BuildResource() *costModels.RequiredResources {
	resource := &costModels.RequiredResources{
		Application:     rb.application,
		CPUMillicore:    rb.cpuMillicore,
		Component:       rb.component,
		Environment:     rb.environment,
		ID:              rb.id,
		MemoryMegaBytes: rb.memoryMegaBytes,
		Replicas:        rb.replicas,
		WBS:             rb.wbs,
	}

	return resource
}

func NewResourceBuilderList() ResourceBuilderList {
	return &ResourceBuilderListStruct{}
}

// NewResourceBuilder constructor for resource builder
func NewResourceBuilder() ResourceBuilder {
	return &ResourceBuilderStruct{}
}

// ARequiredResource constructor for resource with test data
func ARequiredResource() ResourceBuilder {
	builder := NewResourceBuilder().
		WithApplication("any-app").
		WithComponent("server").
		WithWBS("A.BCD.00.999").
		WithID(1).
		WithReplicas(1).
		WithCPUMillicore(20).
		WithMemoryMegaBytes(50).
		WithEnvironment("dev")

	return builder
}

// AListOfRequiredResources constructor for lost of resourcebuilders containing test data
func AListOfRequiredResources() ResourceBuilderList {
	builders := []ResourceBuilder{
		NewResourceBuilder().
			WithApplication("any-app").
			WithComponent("frontend").
			WithWBS("A.BCD.00.999").
			WithID(1).
			WithReplicas(1).
			WithCPUMillicore(20).
			WithMemoryMegaBytes(50).
			WithEnvironment("dev"),
		NewResourceBuilder().
			WithApplication("any-app").
			WithComponent("server").
			WithWBS("A.BCD.00.999").
			WithID(2).
			WithReplicas(1).
			WithCPUMillicore(100).
			WithMemoryMegaBytes(150).
			WithEnvironment("dev"),
		NewResourceBuilder().
			WithApplication("last-app").
			WithComponent("api").
			WithWBS("A.BCD.00.888").
			WithID(3).
			WithReplicas(1).
			WithCPUMillicore(10).
			WithMemoryMegaBytes(25).
			WithEnvironment("dev"),
	}

	builderList := NewResourceBuilderList().
		WithBuilders(builders)

	return builderList
}
