// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3SetDropletActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	SetApplicationDropletByApplicationNameAndSpaceStub        func(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error)
	setApplicationDropletByApplicationNameAndSpaceMutex       sync.RWMutex
	setApplicationDropletByApplicationNameAndSpaceArgsForCall []struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}
	setApplicationDropletByApplicationNameAndSpaceReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	setApplicationDropletByApplicationNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3SetDropletActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3SetDropletActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3SetDropletActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3SetDropletActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletByApplicationNameAndSpace(appName string, spaceGUID string, dropletGUID string) (v3action.Warnings, error) {
	fake.setApplicationDropletByApplicationNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletByApplicationNameAndSpaceReturnsOnCall[len(fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall)]
	fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall = append(fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall, struct {
		appName     string
		spaceGUID   string
		dropletGUID string
	}{appName, spaceGUID, dropletGUID})
	fake.recordInvocation("SetApplicationDropletByApplicationNameAndSpace", []interface{}{appName, spaceGUID, dropletGUID})
	fake.setApplicationDropletByApplicationNameAndSpaceMutex.Unlock()
	if fake.SetApplicationDropletByApplicationNameAndSpaceStub != nil {
		return fake.SetApplicationDropletByApplicationNameAndSpaceStub(appName, spaceGUID, dropletGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setApplicationDropletByApplicationNameAndSpaceReturns.result1, fake.setApplicationDropletByApplicationNameAndSpaceReturns.result2
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletByApplicationNameAndSpaceCallCount() int {
	fake.setApplicationDropletByApplicationNameAndSpaceMutex.RLock()
	defer fake.setApplicationDropletByApplicationNameAndSpaceMutex.RUnlock()
	return len(fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall)
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletByApplicationNameAndSpaceArgsForCall(i int) (string, string, string) {
	fake.setApplicationDropletByApplicationNameAndSpaceMutex.RLock()
	defer fake.setApplicationDropletByApplicationNameAndSpaceMutex.RUnlock()
	return fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall[i].appName, fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall[i].spaceGUID, fake.setApplicationDropletByApplicationNameAndSpaceArgsForCall[i].dropletGUID
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletByApplicationNameAndSpaceReturns(result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletByApplicationNameAndSpaceStub = nil
	fake.setApplicationDropletByApplicationNameAndSpaceReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3SetDropletActor) SetApplicationDropletByApplicationNameAndSpaceReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.SetApplicationDropletByApplicationNameAndSpaceStub = nil
	if fake.setApplicationDropletByApplicationNameAndSpaceReturnsOnCall == nil {
		fake.setApplicationDropletByApplicationNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletByApplicationNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV3SetDropletActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.setApplicationDropletByApplicationNameAndSpaceMutex.RLock()
	defer fake.setApplicationDropletByApplicationNameAndSpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3SetDropletActor) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ v3.V3SetDropletActor = new(FakeV3SetDropletActor)
