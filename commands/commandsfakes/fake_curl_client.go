// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/v2/commands"
)

type FakeCurlClient struct {
	MakeRequestStub        func(string, string, string) error
	makeRequestMutex       sync.RWMutex
	makeRequestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	makeRequestReturns struct {
		result1 error
	}
	makeRequestReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCurlClient) MakeRequest(arg1 string, arg2 string, arg3 string) error {
	fake.makeRequestMutex.Lock()
	ret, specificReturn := fake.makeRequestReturnsOnCall[len(fake.makeRequestArgsForCall)]
	fake.makeRequestArgsForCall = append(fake.makeRequestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("MakeRequest", []interface{}{arg1, arg2, arg3})
	fake.makeRequestMutex.Unlock()
	if fake.MakeRequestStub != nil {
		return fake.MakeRequestStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.makeRequestReturns
	return fakeReturns.result1
}

func (fake *FakeCurlClient) MakeRequestCallCount() int {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	return len(fake.makeRequestArgsForCall)
}

func (fake *FakeCurlClient) MakeRequestCalls(stub func(string, string, string) error) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = stub
}

func (fake *FakeCurlClient) MakeRequestArgsForCall(i int) (string, string, string) {
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	argsForCall := fake.makeRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCurlClient) MakeRequestReturns(result1 error) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = nil
	fake.makeRequestReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCurlClient) MakeRequestReturnsOnCall(i int, result1 error) {
	fake.makeRequestMutex.Lock()
	defer fake.makeRequestMutex.Unlock()
	fake.MakeRequestStub = nil
	if fake.makeRequestReturnsOnCall == nil {
		fake.makeRequestReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.makeRequestReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCurlClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.makeRequestMutex.RLock()
	defer fake.makeRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCurlClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.CurlClient = new(FakeCurlClient)
