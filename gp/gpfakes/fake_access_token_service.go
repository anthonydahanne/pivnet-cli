// Code generated by counterfeiter. DO NOT EDIT.
package gpfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/gp"
)

type FakeAccessTokenService struct {
	AccessTokenStub        func() (string, error)
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct {
	}
	accessTokenReturns struct {
		result1 string
		result2 error
	}
	accessTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccessTokenService) AccessToken() (string, error) {
	fake.accessTokenMutex.Lock()
	ret, specificReturn := fake.accessTokenReturnsOnCall[len(fake.accessTokenArgsForCall)]
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("AccessToken", []interface{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.accessTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAccessTokenService) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeAccessTokenService) AccessTokenCalls(stub func() (string, error)) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = stub
}

func (fake *FakeAccessTokenService) AccessTokenReturns(result1 string, result2 error) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAccessTokenService) AccessTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.accessTokenMutex.Lock()
	defer fake.accessTokenMutex.Unlock()
	fake.AccessTokenStub = nil
	if fake.accessTokenReturnsOnCall == nil {
		fake.accessTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.accessTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeAccessTokenService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccessTokenService) recordInvocation(key string, args []interface{}) {
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

var _ gp.AccessTokenService = new(FakeAccessTokenService)
