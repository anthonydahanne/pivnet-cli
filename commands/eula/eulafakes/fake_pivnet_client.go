// Code generated by counterfeiter. DO NOT EDIT.
package eulafakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v7"
	"github.com/pivotal-cf/pivnet-cli/v2/commands/eula"
)

type FakePivnetClient struct {
	AcceptEULAStub        func(string, int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		arg1 string
		arg2 int
	}
	acceptEULAReturns struct {
		result1 error
	}
	acceptEULAReturnsOnCall map[int]struct {
		result1 error
	}
	EULAStub        func(string) (pivnet.EULA, error)
	eULAMutex       sync.RWMutex
	eULAArgsForCall []struct {
		arg1 string
	}
	eULAReturns struct {
		result1 pivnet.EULA
		result2 error
	}
	eULAReturnsOnCall map[int]struct {
		result1 pivnet.EULA
		result2 error
	}
	EULAsStub        func() ([]pivnet.EULA, error)
	eULAsMutex       sync.RWMutex
	eULAsArgsForCall []struct {
	}
	eULAsReturns struct {
		result1 []pivnet.EULA
		result2 error
	}
	eULAsReturnsOnCall map[int]struct {
		result1 []pivnet.EULA
		result2 error
	}
	ReleaseForVersionStub        func(string, string) (pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		arg1 string
		arg2 string
	}
	releaseForVersionReturns struct {
		result1 pivnet.Release
		result2 error
	}
	releaseForVersionReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) AcceptEULA(arg1 string, arg2 int) error {
	fake.acceptEULAMutex.Lock()
	ret, specificReturn := fake.acceptEULAReturnsOnCall[len(fake.acceptEULAArgsForCall)]
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("AcceptEULA", []interface{}{arg1, arg2})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.acceptEULAReturns
	return fakeReturns.result1
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULACalls(stub func(string, int) error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = stub
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	argsForCall := fake.acceptEULAArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AcceptEULAReturnsOnCall(i int, result1 error) {
	fake.acceptEULAMutex.Lock()
	defer fake.acceptEULAMutex.Unlock()
	fake.AcceptEULAStub = nil
	if fake.acceptEULAReturnsOnCall == nil {
		fake.acceptEULAReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acceptEULAReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) EULA(arg1 string) (pivnet.EULA, error) {
	fake.eULAMutex.Lock()
	ret, specificReturn := fake.eULAReturnsOnCall[len(fake.eULAArgsForCall)]
	fake.eULAArgsForCall = append(fake.eULAArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("EULA", []interface{}{arg1})
	fake.eULAMutex.Unlock()
	if fake.EULAStub != nil {
		return fake.EULAStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.eULAReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) EULACallCount() int {
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	return len(fake.eULAArgsForCall)
}

func (fake *FakePivnetClient) EULACalls(stub func(string) (pivnet.EULA, error)) {
	fake.eULAMutex.Lock()
	defer fake.eULAMutex.Unlock()
	fake.EULAStub = stub
}

func (fake *FakePivnetClient) EULAArgsForCall(i int) string {
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	argsForCall := fake.eULAArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePivnetClient) EULAReturns(result1 pivnet.EULA, result2 error) {
	fake.eULAMutex.Lock()
	defer fake.eULAMutex.Unlock()
	fake.EULAStub = nil
	fake.eULAReturns = struct {
		result1 pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAReturnsOnCall(i int, result1 pivnet.EULA, result2 error) {
	fake.eULAMutex.Lock()
	defer fake.eULAMutex.Unlock()
	fake.EULAStub = nil
	if fake.eULAReturnsOnCall == nil {
		fake.eULAReturnsOnCall = make(map[int]struct {
			result1 pivnet.EULA
			result2 error
		})
	}
	fake.eULAReturnsOnCall[i] = struct {
		result1 pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAs() ([]pivnet.EULA, error) {
	fake.eULAsMutex.Lock()
	ret, specificReturn := fake.eULAsReturnsOnCall[len(fake.eULAsArgsForCall)]
	fake.eULAsArgsForCall = append(fake.eULAsArgsForCall, struct {
	}{})
	fake.recordInvocation("EULAs", []interface{}{})
	fake.eULAsMutex.Unlock()
	if fake.EULAsStub != nil {
		return fake.EULAsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.eULAsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) EULAsCallCount() int {
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	return len(fake.eULAsArgsForCall)
}

func (fake *FakePivnetClient) EULAsCalls(stub func() ([]pivnet.EULA, error)) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = stub
}

func (fake *FakePivnetClient) EULAsReturns(result1 []pivnet.EULA, result2 error) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = nil
	fake.eULAsReturns = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) EULAsReturnsOnCall(i int, result1 []pivnet.EULA, result2 error) {
	fake.eULAsMutex.Lock()
	defer fake.eULAsMutex.Unlock()
	fake.EULAsStub = nil
	if fake.eULAsReturnsOnCall == nil {
		fake.eULAsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.EULA
			result2 error
		})
	}
	fake.eULAsReturnsOnCall[i] = struct {
		result1 []pivnet.EULA
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersion(arg1 string, arg2 string) (pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	ret, specificReturn := fake.releaseForVersionReturnsOnCall[len(fake.releaseForVersionArgsForCall)]
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ReleaseForVersion", []interface{}{arg1, arg2})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.releaseForVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionCalls(stub func(string, string) (pivnet.Release, error)) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = stub
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	argsForCall := fake.releaseForVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseForVersionReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.releaseForVersionMutex.Lock()
	defer fake.releaseForVersionMutex.Unlock()
	fake.ReleaseForVersionStub = nil
	if fake.releaseForVersionReturnsOnCall == nil {
		fake.releaseForVersionReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.releaseForVersionReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.eULAMutex.RLock()
	defer fake.eULAMutex.RUnlock()
	fake.eULAsMutex.RLock()
	defer fake.eULAsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ eula.PivnetClient = new(FakePivnetClient)
