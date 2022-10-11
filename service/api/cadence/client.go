package cadence

import (
	"context"
	"fmt"
	"github.com/cadence-oss/iwf-server/gen/iwfidl"
	"github.com/cadence-oss/iwf-server/service"
	"github.com/cadence-oss/iwf-server/service/api"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/client"
)

type cadenceClient struct {
	cClient   client.Client
	closeFunc func()
}

func NewCadenceClient(cClient client.Client, closeFunc func()) api.UnifiedClient {
	return &cadenceClient{
		cClient:   cClient,
		closeFunc: closeFunc,
	}
}

func (t *cadenceClient) Close() {
	t.closeFunc()
}

func (t *cadenceClient) ExecuteWorkflow(ctx context.Context, options api.StartWorkflowOptions, workflow interface{}, args ...interface{}) (runId string, err error) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                           options.ID,
		TaskList:                     options.TaskQueue,
		ExecutionStartToCloseTimeout: options.WorkflowRunTimeout,
	}

	run, err := t.cClient.ExecuteWorkflow(ctx, workflowOptions, workflow, args...)
	if err != nil {
		return "", err
	}
	return run.GetRunID(), nil
}

func (t *cadenceClient) SignalWorkflow(ctx context.Context, workflowID string, runID string, signalName string, arg interface{}) error {
	return t.cClient.SignalWorkflow(ctx, workflowID, runID, signalName, arg)
}

func (t *cadenceClient) ListWorkflow(ctx context.Context, request *api.ListWorkflowExecutionsRequest) (*api.ListWorkflowExecutionsResponse, error) {
	listReq := &shared.ListWorkflowExecutionsRequest{
		PageSize: &request.PageSize,
		Query:    &request.Query,
	}
	resp, err := t.cClient.ListWorkflow(ctx, listReq)
	if err != nil {
		return nil, err
	}
	var executions []iwfidl.WorkflowSearchResponseEntry
	for _, exe := range resp.GetExecutions() {
		executions = append(executions, iwfidl.WorkflowSearchResponseEntry{
			WorkflowId:    *exe.Execution.WorkflowId,
			WorkflowRunId: *exe.Execution.RunId,
		})
	}
	return &api.ListWorkflowExecutionsResponse{
		Executions: executions,
	}, nil
}

func (t *cadenceClient) QueryWorkflow(ctx context.Context, valuePtr interface{}, workflowID string, runID string, queryType string, args ...interface{}) error {
	qres, err := t.cClient.QueryWorkflow(ctx, workflowID, runID, queryType, args...)
	if err != nil {
		return err
	}
	return qres.Get(valuePtr)
}

func (t *cadenceClient) DescribeWorkflowExecution(ctx context.Context, workflowID, runID string) (*api.DescribeWorkflowExecutionResponse, error) {
	resp, err := t.cClient.DescribeWorkflowExecution(ctx, workflowID, runID)
	if err != nil {
		return nil, err
	}
	status, err := mapToIwfWorkflowStatus(resp.GetWorkflowExecutionInfo().CloseStatus)
	if err != nil {
		return nil, err
	}
	return &api.DescribeWorkflowExecutionResponse{
		RunId:  resp.GetWorkflowExecutionInfo().GetExecution().GetRunId(),
		Status: status,
	}, nil
}

func mapToIwfWorkflowStatus(status *shared.WorkflowExecutionCloseStatus) (string, error) {
	if status == nil {
		return service.WorkflowStatusRunning, nil
	}

	switch *status {
	case shared.WorkflowExecutionCloseStatusCanceled:
		return service.WorkflowStatusCanceled, nil
	case shared.WorkflowExecutionCloseStatusContinuedAsNew:
		return service.WorkflowStatusContinueAsNew, nil
	case shared.WorkflowExecutionCloseStatusFailed:
		return service.WorkflowStatusFailed, nil
	case shared.WorkflowExecutionCloseStatusTimedOut:
		return service.WorkflowStatusTimeout, nil
	case shared.WorkflowExecutionCloseStatusTerminated:
		return service.WorkflowStatusTerminated, nil
	default:
		return "", fmt.Errorf("not supported status %s", status)
	}
}

func (t *cadenceClient) GetWorkflowResult(ctx context.Context, valuePtr interface{}, workflowID string, runID string) error {
	run := t.cClient.GetWorkflow(ctx, workflowID, runID)
	return run.Get(ctx, valuePtr)
}
