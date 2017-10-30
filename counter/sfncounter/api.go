// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sfncounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
)

// CreateActivityRequest is a passthrough to the underlying CreateActivityRequest.
// It will increment the count of requests made to CreateActivity.
func (c *SFN) CreateActivityRequest(input *sfn.CreateActivityInput) (req *request.Request, output *sfn.CreateActivityOutput) {
	c.inc("CreateActivity")
	return c.svc.CreateActivityRequest(input)
}

// CreateActivity is a passthrough to the underlying CreateActivity method.
// It will increment the count of requests made to CreateActivity.
func (c *SFN) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
	c.inc("CreateActivity")
	return c.svc.CreateActivity(input)
}

// CreateActivityWithContext is a passthrough to the underlying CreateActivityWithContext method.
// It will increment the count of requests made to CreateActivity.
func (c *SFN) CreateActivityWithContext(ctx aws.Context, input *sfn.CreateActivityInput, opts ...request.Option) (*sfn.CreateActivityOutput, error) {
	c.inc("CreateActivity")
	return c.svc.CreateActivityWithContext(ctx, input, opts...)
}

// CreateStateMachineRequest is a passthrough to the underlying CreateStateMachineRequest.
// It will increment the count of requests made to CreateStateMachine.
func (c *SFN) CreateStateMachineRequest(input *sfn.CreateStateMachineInput) (req *request.Request, output *sfn.CreateStateMachineOutput) {
	c.inc("CreateStateMachine")
	return c.svc.CreateStateMachineRequest(input)
}

// CreateStateMachine is a passthrough to the underlying CreateStateMachine method.
// It will increment the count of requests made to CreateStateMachine.
func (c *SFN) CreateStateMachine(input *sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error) {
	c.inc("CreateStateMachine")
	return c.svc.CreateStateMachine(input)
}

// CreateStateMachineWithContext is a passthrough to the underlying CreateStateMachineWithContext method.
// It will increment the count of requests made to CreateStateMachine.
func (c *SFN) CreateStateMachineWithContext(ctx aws.Context, input *sfn.CreateStateMachineInput, opts ...request.Option) (*sfn.CreateStateMachineOutput, error) {
	c.inc("CreateStateMachine")
	return c.svc.CreateStateMachineWithContext(ctx, input, opts...)
}

// DeleteActivityRequest is a passthrough to the underlying DeleteActivityRequest.
// It will increment the count of requests made to DeleteActivity.
func (c *SFN) DeleteActivityRequest(input *sfn.DeleteActivityInput) (req *request.Request, output *sfn.DeleteActivityOutput) {
	c.inc("DeleteActivity")
	return c.svc.DeleteActivityRequest(input)
}

// DeleteActivity is a passthrough to the underlying DeleteActivity method.
// It will increment the count of requests made to DeleteActivity.
func (c *SFN) DeleteActivity(input *sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error) {
	c.inc("DeleteActivity")
	return c.svc.DeleteActivity(input)
}

// DeleteActivityWithContext is a passthrough to the underlying DeleteActivityWithContext method.
// It will increment the count of requests made to DeleteActivity.
func (c *SFN) DeleteActivityWithContext(ctx aws.Context, input *sfn.DeleteActivityInput, opts ...request.Option) (*sfn.DeleteActivityOutput, error) {
	c.inc("DeleteActivity")
	return c.svc.DeleteActivityWithContext(ctx, input, opts...)
}

// DeleteStateMachineRequest is a passthrough to the underlying DeleteStateMachineRequest.
// It will increment the count of requests made to DeleteStateMachine.
func (c *SFN) DeleteStateMachineRequest(input *sfn.DeleteStateMachineInput) (req *request.Request, output *sfn.DeleteStateMachineOutput) {
	c.inc("DeleteStateMachine")
	return c.svc.DeleteStateMachineRequest(input)
}

// DeleteStateMachine is a passthrough to the underlying DeleteStateMachine method.
// It will increment the count of requests made to DeleteStateMachine.
func (c *SFN) DeleteStateMachine(input *sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error) {
	c.inc("DeleteStateMachine")
	return c.svc.DeleteStateMachine(input)
}

// DeleteStateMachineWithContext is a passthrough to the underlying DeleteStateMachineWithContext method.
// It will increment the count of requests made to DeleteStateMachine.
func (c *SFN) DeleteStateMachineWithContext(ctx aws.Context, input *sfn.DeleteStateMachineInput, opts ...request.Option) (*sfn.DeleteStateMachineOutput, error) {
	c.inc("DeleteStateMachine")
	return c.svc.DeleteStateMachineWithContext(ctx, input, opts...)
}

// DescribeActivityRequest is a passthrough to the underlying DescribeActivityRequest.
// It will increment the count of requests made to DescribeActivity.
func (c *SFN) DescribeActivityRequest(input *sfn.DescribeActivityInput) (req *request.Request, output *sfn.DescribeActivityOutput) {
	c.inc("DescribeActivity")
	return c.svc.DescribeActivityRequest(input)
}

// DescribeActivity is a passthrough to the underlying DescribeActivity method.
// It will increment the count of requests made to DescribeActivity.
func (c *SFN) DescribeActivity(input *sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error) {
	c.inc("DescribeActivity")
	return c.svc.DescribeActivity(input)
}

// DescribeActivityWithContext is a passthrough to the underlying DescribeActivityWithContext method.
// It will increment the count of requests made to DescribeActivity.
func (c *SFN) DescribeActivityWithContext(ctx aws.Context, input *sfn.DescribeActivityInput, opts ...request.Option) (*sfn.DescribeActivityOutput, error) {
	c.inc("DescribeActivity")
	return c.svc.DescribeActivityWithContext(ctx, input, opts...)
}

// DescribeExecutionRequest is a passthrough to the underlying DescribeExecutionRequest.
// It will increment the count of requests made to DescribeExecution.
func (c *SFN) DescribeExecutionRequest(input *sfn.DescribeExecutionInput) (req *request.Request, output *sfn.DescribeExecutionOutput) {
	c.inc("DescribeExecution")
	return c.svc.DescribeExecutionRequest(input)
}

// DescribeExecution is a passthrough to the underlying DescribeExecution method.
// It will increment the count of requests made to DescribeExecution.
func (c *SFN) DescribeExecution(input *sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error) {
	c.inc("DescribeExecution")
	return c.svc.DescribeExecution(input)
}

// DescribeExecutionWithContext is a passthrough to the underlying DescribeExecutionWithContext method.
// It will increment the count of requests made to DescribeExecution.
func (c *SFN) DescribeExecutionWithContext(ctx aws.Context, input *sfn.DescribeExecutionInput, opts ...request.Option) (*sfn.DescribeExecutionOutput, error) {
	c.inc("DescribeExecution")
	return c.svc.DescribeExecutionWithContext(ctx, input, opts...)
}

// DescribeStateMachineRequest is a passthrough to the underlying DescribeStateMachineRequest.
// It will increment the count of requests made to DescribeStateMachine.
func (c *SFN) DescribeStateMachineRequest(input *sfn.DescribeStateMachineInput) (req *request.Request, output *sfn.DescribeStateMachineOutput) {
	c.inc("DescribeStateMachine")
	return c.svc.DescribeStateMachineRequest(input)
}

// DescribeStateMachine is a passthrough to the underlying DescribeStateMachine method.
// It will increment the count of requests made to DescribeStateMachine.
func (c *SFN) DescribeStateMachine(input *sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error) {
	c.inc("DescribeStateMachine")
	return c.svc.DescribeStateMachine(input)
}

// DescribeStateMachineWithContext is a passthrough to the underlying DescribeStateMachineWithContext method.
// It will increment the count of requests made to DescribeStateMachine.
func (c *SFN) DescribeStateMachineWithContext(ctx aws.Context, input *sfn.DescribeStateMachineInput, opts ...request.Option) (*sfn.DescribeStateMachineOutput, error) {
	c.inc("DescribeStateMachine")
	return c.svc.DescribeStateMachineWithContext(ctx, input, opts...)
}

// GetActivityTaskRequest is a passthrough to the underlying GetActivityTaskRequest.
// It will increment the count of requests made to GetActivityTask.
func (c *SFN) GetActivityTaskRequest(input *sfn.GetActivityTaskInput) (req *request.Request, output *sfn.GetActivityTaskOutput) {
	c.inc("GetActivityTask")
	return c.svc.GetActivityTaskRequest(input)
}

// GetActivityTask is a passthrough to the underlying GetActivityTask method.
// It will increment the count of requests made to GetActivityTask.
func (c *SFN) GetActivityTask(input *sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error) {
	c.inc("GetActivityTask")
	return c.svc.GetActivityTask(input)
}

// GetActivityTaskWithContext is a passthrough to the underlying GetActivityTaskWithContext method.
// It will increment the count of requests made to GetActivityTask.
func (c *SFN) GetActivityTaskWithContext(ctx aws.Context, input *sfn.GetActivityTaskInput, opts ...request.Option) (*sfn.GetActivityTaskOutput, error) {
	c.inc("GetActivityTask")
	return c.svc.GetActivityTaskWithContext(ctx, input, opts...)
}

// GetExecutionHistoryRequest is a passthrough to the underlying GetExecutionHistoryRequest.
// It will increment the count of requests made to GetExecutionHistory.
func (c *SFN) GetExecutionHistoryRequest(input *sfn.GetExecutionHistoryInput) (req *request.Request, output *sfn.GetExecutionHistoryOutput) {
	c.inc("GetExecutionHistory")
	return c.svc.GetExecutionHistoryRequest(input)
}

// GetExecutionHistory is a passthrough to the underlying GetExecutionHistory method.
// It will increment the count of requests made to GetExecutionHistory.
func (c *SFN) GetExecutionHistory(input *sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error) {
	c.inc("GetExecutionHistory")
	return c.svc.GetExecutionHistory(input)
}

// GetExecutionHistoryWithContext is a passthrough to the underlying GetExecutionHistoryWithContext method.
// It will increment the count of requests made to GetExecutionHistory.
func (c *SFN) GetExecutionHistoryWithContext(ctx aws.Context, input *sfn.GetExecutionHistoryInput, opts ...request.Option) (*sfn.GetExecutionHistoryOutput, error) {
	c.inc("GetExecutionHistory")
	return c.svc.GetExecutionHistoryWithContext(ctx, input, opts...)
}

// GetExecutionHistoryPages is a passthrough to the underlying GetExecutionHistoryPages method.
// It will increment the count of requests made to GetExecutionHistory on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use GetExecutionHistoryPagesWithContext to avoid this.
func (c *SFN) GetExecutionHistoryPages(input *sfn.GetExecutionHistoryInput, fn func(*sfn.GetExecutionHistoryOutput, bool) bool) error {
	wrappedFn := func(page *sfn.GetExecutionHistoryOutput, lastPage bool) bool {
		c.inc("GetExecutionHistory")
		return fn(page, lastPage)
	}
	return c.svc.GetExecutionHistoryPages(input, wrappedFn)
}

// GetExecutionHistoryPagesWithContext is a passthrough to the underlying GetExecutionHistoryPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to GetExecutionHistory when applied to the request.
func (c *SFN) GetExecutionHistoryPagesWithContext(ctx aws.Context, input *sfn.GetExecutionHistoryInput, fn func(*sfn.GetExecutionHistoryOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("GetExecutionHistory"))
	return c.svc.GetExecutionHistoryPagesWithContext(ctx, input, fn, opts...)
}

// ListActivitiesRequest is a passthrough to the underlying ListActivitiesRequest.
// It will increment the count of requests made to ListActivities.
func (c *SFN) ListActivitiesRequest(input *sfn.ListActivitiesInput) (req *request.Request, output *sfn.ListActivitiesOutput) {
	c.inc("ListActivities")
	return c.svc.ListActivitiesRequest(input)
}

// ListActivities is a passthrough to the underlying ListActivities method.
// It will increment the count of requests made to ListActivities.
func (c *SFN) ListActivities(input *sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error) {
	c.inc("ListActivities")
	return c.svc.ListActivities(input)
}

// ListActivitiesWithContext is a passthrough to the underlying ListActivitiesWithContext method.
// It will increment the count of requests made to ListActivities.
func (c *SFN) ListActivitiesWithContext(ctx aws.Context, input *sfn.ListActivitiesInput, opts ...request.Option) (*sfn.ListActivitiesOutput, error) {
	c.inc("ListActivities")
	return c.svc.ListActivitiesWithContext(ctx, input, opts...)
}

// ListActivitiesPages is a passthrough to the underlying ListActivitiesPages method.
// It will increment the count of requests made to ListActivities on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListActivitiesPagesWithContext to avoid this.
func (c *SFN) ListActivitiesPages(input *sfn.ListActivitiesInput, fn func(*sfn.ListActivitiesOutput, bool) bool) error {
	wrappedFn := func(page *sfn.ListActivitiesOutput, lastPage bool) bool {
		c.inc("ListActivities")
		return fn(page, lastPage)
	}
	return c.svc.ListActivitiesPages(input, wrappedFn)
}

// ListActivitiesPagesWithContext is a passthrough to the underlying ListActivitiesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListActivities when applied to the request.
func (c *SFN) ListActivitiesPagesWithContext(ctx aws.Context, input *sfn.ListActivitiesInput, fn func(*sfn.ListActivitiesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListActivities"))
	return c.svc.ListActivitiesPagesWithContext(ctx, input, fn, opts...)
}

// ListExecutionsRequest is a passthrough to the underlying ListExecutionsRequest.
// It will increment the count of requests made to ListExecutions.
func (c *SFN) ListExecutionsRequest(input *sfn.ListExecutionsInput) (req *request.Request, output *sfn.ListExecutionsOutput) {
	c.inc("ListExecutions")
	return c.svc.ListExecutionsRequest(input)
}

// ListExecutions is a passthrough to the underlying ListExecutions method.
// It will increment the count of requests made to ListExecutions.
func (c *SFN) ListExecutions(input *sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error) {
	c.inc("ListExecutions")
	return c.svc.ListExecutions(input)
}

// ListExecutionsWithContext is a passthrough to the underlying ListExecutionsWithContext method.
// It will increment the count of requests made to ListExecutions.
func (c *SFN) ListExecutionsWithContext(ctx aws.Context, input *sfn.ListExecutionsInput, opts ...request.Option) (*sfn.ListExecutionsOutput, error) {
	c.inc("ListExecutions")
	return c.svc.ListExecutionsWithContext(ctx, input, opts...)
}

// ListExecutionsPages is a passthrough to the underlying ListExecutionsPages method.
// It will increment the count of requests made to ListExecutions on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListExecutionsPagesWithContext to avoid this.
func (c *SFN) ListExecutionsPages(input *sfn.ListExecutionsInput, fn func(*sfn.ListExecutionsOutput, bool) bool) error {
	wrappedFn := func(page *sfn.ListExecutionsOutput, lastPage bool) bool {
		c.inc("ListExecutions")
		return fn(page, lastPage)
	}
	return c.svc.ListExecutionsPages(input, wrappedFn)
}

// ListExecutionsPagesWithContext is a passthrough to the underlying ListExecutionsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListExecutions when applied to the request.
func (c *SFN) ListExecutionsPagesWithContext(ctx aws.Context, input *sfn.ListExecutionsInput, fn func(*sfn.ListExecutionsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListExecutions"))
	return c.svc.ListExecutionsPagesWithContext(ctx, input, fn, opts...)
}

// ListStateMachinesRequest is a passthrough to the underlying ListStateMachinesRequest.
// It will increment the count of requests made to ListStateMachines.
func (c *SFN) ListStateMachinesRequest(input *sfn.ListStateMachinesInput) (req *request.Request, output *sfn.ListStateMachinesOutput) {
	c.inc("ListStateMachines")
	return c.svc.ListStateMachinesRequest(input)
}

// ListStateMachines is a passthrough to the underlying ListStateMachines method.
// It will increment the count of requests made to ListStateMachines.
func (c *SFN) ListStateMachines(input *sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error) {
	c.inc("ListStateMachines")
	return c.svc.ListStateMachines(input)
}

// ListStateMachinesWithContext is a passthrough to the underlying ListStateMachinesWithContext method.
// It will increment the count of requests made to ListStateMachines.
func (c *SFN) ListStateMachinesWithContext(ctx aws.Context, input *sfn.ListStateMachinesInput, opts ...request.Option) (*sfn.ListStateMachinesOutput, error) {
	c.inc("ListStateMachines")
	return c.svc.ListStateMachinesWithContext(ctx, input, opts...)
}

// ListStateMachinesPages is a passthrough to the underlying ListStateMachinesPages method.
// It will increment the count of requests made to ListStateMachines on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListStateMachinesPagesWithContext to avoid this.
func (c *SFN) ListStateMachinesPages(input *sfn.ListStateMachinesInput, fn func(*sfn.ListStateMachinesOutput, bool) bool) error {
	wrappedFn := func(page *sfn.ListStateMachinesOutput, lastPage bool) bool {
		c.inc("ListStateMachines")
		return fn(page, lastPage)
	}
	return c.svc.ListStateMachinesPages(input, wrappedFn)
}

// ListStateMachinesPagesWithContext is a passthrough to the underlying ListStateMachinesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListStateMachines when applied to the request.
func (c *SFN) ListStateMachinesPagesWithContext(ctx aws.Context, input *sfn.ListStateMachinesInput, fn func(*sfn.ListStateMachinesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListStateMachines"))
	return c.svc.ListStateMachinesPagesWithContext(ctx, input, fn, opts...)
}

// SendTaskFailureRequest is a passthrough to the underlying SendTaskFailureRequest.
// It will increment the count of requests made to SendTaskFailure.
func (c *SFN) SendTaskFailureRequest(input *sfn.SendTaskFailureInput) (req *request.Request, output *sfn.SendTaskFailureOutput) {
	c.inc("SendTaskFailure")
	return c.svc.SendTaskFailureRequest(input)
}

// SendTaskFailure is a passthrough to the underlying SendTaskFailure method.
// It will increment the count of requests made to SendTaskFailure.
func (c *SFN) SendTaskFailure(input *sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error) {
	c.inc("SendTaskFailure")
	return c.svc.SendTaskFailure(input)
}

// SendTaskFailureWithContext is a passthrough to the underlying SendTaskFailureWithContext method.
// It will increment the count of requests made to SendTaskFailure.
func (c *SFN) SendTaskFailureWithContext(ctx aws.Context, input *sfn.SendTaskFailureInput, opts ...request.Option) (*sfn.SendTaskFailureOutput, error) {
	c.inc("SendTaskFailure")
	return c.svc.SendTaskFailureWithContext(ctx, input, opts...)
}

// SendTaskHeartbeatRequest is a passthrough to the underlying SendTaskHeartbeatRequest.
// It will increment the count of requests made to SendTaskHeartbeat.
func (c *SFN) SendTaskHeartbeatRequest(input *sfn.SendTaskHeartbeatInput) (req *request.Request, output *sfn.SendTaskHeartbeatOutput) {
	c.inc("SendTaskHeartbeat")
	return c.svc.SendTaskHeartbeatRequest(input)
}

// SendTaskHeartbeat is a passthrough to the underlying SendTaskHeartbeat method.
// It will increment the count of requests made to SendTaskHeartbeat.
func (c *SFN) SendTaskHeartbeat(input *sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error) {
	c.inc("SendTaskHeartbeat")
	return c.svc.SendTaskHeartbeat(input)
}

// SendTaskHeartbeatWithContext is a passthrough to the underlying SendTaskHeartbeatWithContext method.
// It will increment the count of requests made to SendTaskHeartbeat.
func (c *SFN) SendTaskHeartbeatWithContext(ctx aws.Context, input *sfn.SendTaskHeartbeatInput, opts ...request.Option) (*sfn.SendTaskHeartbeatOutput, error) {
	c.inc("SendTaskHeartbeat")
	return c.svc.SendTaskHeartbeatWithContext(ctx, input, opts...)
}

// SendTaskSuccessRequest is a passthrough to the underlying SendTaskSuccessRequest.
// It will increment the count of requests made to SendTaskSuccess.
func (c *SFN) SendTaskSuccessRequest(input *sfn.SendTaskSuccessInput) (req *request.Request, output *sfn.SendTaskSuccessOutput) {
	c.inc("SendTaskSuccess")
	return c.svc.SendTaskSuccessRequest(input)
}

// SendTaskSuccess is a passthrough to the underlying SendTaskSuccess method.
// It will increment the count of requests made to SendTaskSuccess.
func (c *SFN) SendTaskSuccess(input *sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error) {
	c.inc("SendTaskSuccess")
	return c.svc.SendTaskSuccess(input)
}

// SendTaskSuccessWithContext is a passthrough to the underlying SendTaskSuccessWithContext method.
// It will increment the count of requests made to SendTaskSuccess.
func (c *SFN) SendTaskSuccessWithContext(ctx aws.Context, input *sfn.SendTaskSuccessInput, opts ...request.Option) (*sfn.SendTaskSuccessOutput, error) {
	c.inc("SendTaskSuccess")
	return c.svc.SendTaskSuccessWithContext(ctx, input, opts...)
}

// StartExecutionRequest is a passthrough to the underlying StartExecutionRequest.
// It will increment the count of requests made to StartExecution.
func (c *SFN) StartExecutionRequest(input *sfn.StartExecutionInput) (req *request.Request, output *sfn.StartExecutionOutput) {
	c.inc("StartExecution")
	return c.svc.StartExecutionRequest(input)
}

// StartExecution is a passthrough to the underlying StartExecution method.
// It will increment the count of requests made to StartExecution.
func (c *SFN) StartExecution(input *sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error) {
	c.inc("StartExecution")
	return c.svc.StartExecution(input)
}

// StartExecutionWithContext is a passthrough to the underlying StartExecutionWithContext method.
// It will increment the count of requests made to StartExecution.
func (c *SFN) StartExecutionWithContext(ctx aws.Context, input *sfn.StartExecutionInput, opts ...request.Option) (*sfn.StartExecutionOutput, error) {
	c.inc("StartExecution")
	return c.svc.StartExecutionWithContext(ctx, input, opts...)
}

// StopExecutionRequest is a passthrough to the underlying StopExecutionRequest.
// It will increment the count of requests made to StopExecution.
func (c *SFN) StopExecutionRequest(input *sfn.StopExecutionInput) (req *request.Request, output *sfn.StopExecutionOutput) {
	c.inc("StopExecution")
	return c.svc.StopExecutionRequest(input)
}

// StopExecution is a passthrough to the underlying StopExecution method.
// It will increment the count of requests made to StopExecution.
func (c *SFN) StopExecution(input *sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error) {
	c.inc("StopExecution")
	return c.svc.StopExecution(input)
}

// StopExecutionWithContext is a passthrough to the underlying StopExecutionWithContext method.
// It will increment the count of requests made to StopExecution.
func (c *SFN) StopExecutionWithContext(ctx aws.Context, input *sfn.StopExecutionInput, opts ...request.Option) (*sfn.StopExecutionOutput, error) {
	c.inc("StopExecution")
	return c.svc.StopExecutionWithContext(ctx, input, opts...)
}
