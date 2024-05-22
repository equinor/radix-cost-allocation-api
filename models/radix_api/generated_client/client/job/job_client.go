// Code generated by go-swagger; DO NOT EDIT.

package job

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new job API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new job API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new job API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for job API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetJobComponentDeployments(params *GetJobComponentDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobComponentDeploymentsOK, error)

	CopyBatch(params *CopyBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyBatchOK, error)

	CopyJob(params *CopyJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyJobOK, error)

	DeleteBatch(params *DeleteBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBatchNoContent, error)

	DeleteJob(params *DeleteJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteJobNoContent, error)

	GetBatch(params *GetBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBatchOK, error)

	GetBatches(params *GetBatchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBatchesOK, error)

	GetJob(params *GetJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobOK, error)

	GetJobPayload(params *GetJobPayloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobPayloadOK, error)

	GetJobs(params *GetJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobsOK, error)

	JobLog(params *JobLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JobLogOK, error)

	RestartBatch(params *RestartBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartBatchNoContent, error)

	RestartJob(params *RestartJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartJobNoContent, error)

	StopBatch(params *StopBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopBatchNoContent, error)

	StopJob(params *StopJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopJobNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetJobComponentDeployments gets list of deployments for the job component
*/
func (a *Client) GetJobComponentDeployments(params *GetJobComponentDeploymentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobComponentDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobComponentDeploymentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetJobComponentDeployments",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJobComponentDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobComponentDeploymentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetJobComponentDeployments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CopyBatch creates a copy of existing scheduled batch with optional changes
*/
func (a *Client) CopyBatch(params *CopyBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "copyBatch",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CopyBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CopyBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for copyBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CopyJob creates a copy of existing scheduled job with optional changes
*/
func (a *Client) CopyJob(params *CopyJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CopyJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCopyJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "copyJob",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CopyJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CopyJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for copyJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteBatch deletes batch
*/
func (a *Client) DeleteBatch(params *DeleteBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteBatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteBatch",
		Method:             "DELETE",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteJob deletes job
*/
func (a *Client) DeleteJob(params *DeleteJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteJob",
		Method:             "DELETE",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBatch gets list of scheduled batches
*/
func (a *Client) GetBatch(params *GetBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBatch",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetBatches gets list of scheduled batches
*/
func (a *Client) GetBatches(params *GetBatchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetBatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getBatches",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBatchesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetBatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getBatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJob gets list of scheduled jobs
*/
func (a *Client) GetJob(params *GetJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJob",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJobPayload gets payload of a scheduled job
*/
func (a *Client) GetJobPayload(params *GetJobPayloadParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobPayloadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobPayloadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobPayload",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/payload",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJobPayloadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobPayloadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobPayload: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetJobs gets list of scheduled jobs
*/
func (a *Client) GetJobs(params *GetJobsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJobsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getJobs",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetJobsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetJobsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getJobs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
JobLog gets log from a scheduled job
*/
func (a *Client) JobLog(params *JobLogParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*JobLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewJobLogParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "jobLog",
		Method:             "GET",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/scheduledjobs/{scheduledJobName}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &JobLogReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*JobLogOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for jobLog: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestartBatch restarts a scheduled or stopped batch
*/
func (a *Client) RestartBatch(params *RestartBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartBatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartBatch",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartBatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restartBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestartJob restarts a running or stopped scheduled job
*/
func (a *Client) RestartJob(params *RestartJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestartJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "restartJob",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restartJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopBatch stops scheduled batch
*/
func (a *Client) StopBatch(params *StopBatchParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopBatchNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopBatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopBatch",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/batches/{batchName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopBatchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopBatchNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopBatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
StopJob stops scheduled job
*/
func (a *Client) StopJob(params *StopJobParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*StopJobNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopJobParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "stopJob",
		Method:             "POST",
		PathPattern:        "/applications/{appName}/environments/{envName}/jobcomponents/{jobComponentName}/jobs/{jobName}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopJobReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StopJobNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for stopJob: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
