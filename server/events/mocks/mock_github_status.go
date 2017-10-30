// Automatically generated by pegomock. DO NOT EDIT!
// Source: github.com/hootsuite/atlantis/server/events (interfaces: GHStatusUpdater)

package mocks

import (
	events "github.com/hootsuite/atlantis/server/events"
	models "github.com/hootsuite/atlantis/server/events/models"
	pegomock "github.com/petergtz/pegomock"
	"reflect"
)

type MockGHStatusUpdater struct {
	fail func(message string, callerSkip ...int)
}

func NewMockGHStatusUpdater() *MockGHStatusUpdater {
	return &MockGHStatusUpdater{fail: pegomock.GlobalFailHandler}
}

func (mock *MockGHStatusUpdater) Update(repo models.Repo, pull models.PullRequest, status events.Status, cmd *events.Command) error {
	params := []pegomock.Param{repo, pull, status, cmd}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Update", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockGHStatusUpdater) UpdateProjectResult(ctx *events.CommandContext, res events.CommandResponse) error {
	params := []pegomock.Param{ctx, res}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProjectResult", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockGHStatusUpdater) VerifyWasCalledOnce() *VerifierGHStatusUpdater {
	return &VerifierGHStatusUpdater{mock, pegomock.Times(1), nil}
}

func (mock *MockGHStatusUpdater) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierGHStatusUpdater {
	return &VerifierGHStatusUpdater{mock, invocationCountMatcher, nil}
}

func (mock *MockGHStatusUpdater) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierGHStatusUpdater {
	return &VerifierGHStatusUpdater{mock, invocationCountMatcher, inOrderContext}
}

type VerifierGHStatusUpdater struct {
	mock                   *MockGHStatusUpdater
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierGHStatusUpdater) Update(repo models.Repo, pull models.PullRequest, status events.Status, cmd *events.Command) *GHStatusUpdater_Update_OngoingVerification {
	params := []pegomock.Param{repo, pull, status, cmd}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Update", params)
	return &GHStatusUpdater_Update_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GHStatusUpdater_Update_OngoingVerification struct {
	mock              *MockGHStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *GHStatusUpdater_Update_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, events.Status, *events.Command) {
	repo, pull, status, cmd := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], status[len(status)-1], cmd[len(cmd)-1]
}

func (c *GHStatusUpdater_Update_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []events.Status, _param3 []*events.Command) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]events.Status, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(events.Status)
		}
		_param3 = make([]*events.Command, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(*events.Command)
		}
	}
	return
}

func (verifier *VerifierGHStatusUpdater) UpdateProjectResult(ctx *events.CommandContext, res events.CommandResponse) *GHStatusUpdater_UpdateProjectResult_OngoingVerification {
	params := []pegomock.Param{ctx, res}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProjectResult", params)
	return &GHStatusUpdater_UpdateProjectResult_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type GHStatusUpdater_UpdateProjectResult_OngoingVerification struct {
	mock              *MockGHStatusUpdater
	methodInvocations []pegomock.MethodInvocation
}

func (c *GHStatusUpdater_UpdateProjectResult_OngoingVerification) GetCapturedArguments() (*events.CommandContext, events.CommandResponse) {
	ctx, res := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], res[len(res)-1]
}

func (c *GHStatusUpdater_UpdateProjectResult_OngoingVerification) GetAllCapturedArguments() (_param0 []*events.CommandContext, _param1 []events.CommandResponse) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*events.CommandContext, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(*events.CommandContext)
		}
		_param1 = make([]events.CommandResponse, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(events.CommandResponse)
		}
	}
	return
}