// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeImageReferenceClient struct {
	AddToReleaseStub        func(string, int, string) error
	addToReleaseMutex       sync.RWMutex
	addToReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 string
	}
	addToReleaseReturns struct {
		result1 error
	}
	addToReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	CreateStub        func(pivnet.CreateImageReferenceConfig) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 pivnet.CreateImageReferenceConfig
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string, int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
		arg2 int
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string, string, int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 int
	}
	getReturns struct {
		result1 error
	}
	getReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(string, string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listReturns struct {
		result1 error
	}
	listReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFromReleaseStub        func(string, int, string) error
	removeFromReleaseMutex       sync.RWMutex
	removeFromReleaseArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 string
	}
	removeFromReleaseReturns struct {
		result1 error
	}
	removeFromReleaseReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(string, int, *string, *string, *string, *[]string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 *string
		arg4 *string
		arg5 *string
		arg6 *[]string
	}
	updateReturns struct {
		result1 error
	}
	updateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageReferenceClient) AddToRelease(arg1 string, arg2 int, arg3 string) error {
	fake.addToReleaseMutex.Lock()
	ret, specificReturn := fake.addToReleaseReturnsOnCall[len(fake.addToReleaseArgsForCall)]
	fake.addToReleaseArgsForCall = append(fake.addToReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddToRelease", []interface{}{arg1, arg2, arg3})
	fake.addToReleaseMutex.Unlock()
	if fake.AddToReleaseStub != nil {
		return fake.AddToReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addToReleaseReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) AddToReleaseCallCount() int {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	return len(fake.addToReleaseArgsForCall)
}

func (fake *FakeImageReferenceClient) AddToReleaseCalls(stub func(string, int, string) error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = stub
}

func (fake *FakeImageReferenceClient) AddToReleaseArgsForCall(i int) (string, int, string) {
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	argsForCall := fake.addToReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageReferenceClient) AddToReleaseReturns(result1 error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = nil
	fake.addToReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) AddToReleaseReturnsOnCall(i int, result1 error) {
	fake.addToReleaseMutex.Lock()
	defer fake.addToReleaseMutex.Unlock()
	fake.AddToReleaseStub = nil
	if fake.addToReleaseReturnsOnCall == nil {
		fake.addToReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addToReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Create(arg1 pivnet.CreateImageReferenceConfig) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 pivnet.CreateImageReferenceConfig
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeImageReferenceClient) CreateCalls(stub func(pivnet.CreateImageReferenceConfig) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeImageReferenceClient) CreateArgsForCall(i int) pivnet.CreateImageReferenceConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImageReferenceClient) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Delete(arg1 string, arg2 int) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeImageReferenceClient) DeleteCalls(stub func(string, int) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeImageReferenceClient) DeleteArgsForCall(i int) (string, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageReferenceClient) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Get(arg1 string, arg2 string, arg3 int) error {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeImageReferenceClient) GetCalls(stub func(string, string, int) error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeImageReferenceClient) GetArgsForCall(i int) (string, string, int) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageReferenceClient) GetReturns(result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) GetReturnsOnCall(i int, result1 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) List(arg1 string, arg2 string) error {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("List", []interface{}{arg1, arg2})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeImageReferenceClient) ListCalls(stub func(string, string) error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeImageReferenceClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImageReferenceClient) ListReturns(result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) ListReturnsOnCall(i int, result1 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) RemoveFromRelease(arg1 string, arg2 int, arg3 string) error {
	fake.removeFromReleaseMutex.Lock()
	ret, specificReturn := fake.removeFromReleaseReturnsOnCall[len(fake.removeFromReleaseArgsForCall)]
	fake.removeFromReleaseArgsForCall = append(fake.removeFromReleaseArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("RemoveFromRelease", []interface{}{arg1, arg2, arg3})
	fake.removeFromReleaseMutex.Unlock()
	if fake.RemoveFromReleaseStub != nil {
		return fake.RemoveFromReleaseStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeFromReleaseReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) RemoveFromReleaseCallCount() int {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	return len(fake.removeFromReleaseArgsForCall)
}

func (fake *FakeImageReferenceClient) RemoveFromReleaseCalls(stub func(string, int, string) error) {
	fake.removeFromReleaseMutex.Lock()
	defer fake.removeFromReleaseMutex.Unlock()
	fake.RemoveFromReleaseStub = stub
}

func (fake *FakeImageReferenceClient) RemoveFromReleaseArgsForCall(i int) (string, int, string) {
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	argsForCall := fake.removeFromReleaseArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageReferenceClient) RemoveFromReleaseReturns(result1 error) {
	fake.removeFromReleaseMutex.Lock()
	defer fake.removeFromReleaseMutex.Unlock()
	fake.RemoveFromReleaseStub = nil
	fake.removeFromReleaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) RemoveFromReleaseReturnsOnCall(i int, result1 error) {
	fake.removeFromReleaseMutex.Lock()
	defer fake.removeFromReleaseMutex.Unlock()
	fake.RemoveFromReleaseStub = nil
	if fake.removeFromReleaseReturnsOnCall == nil {
		fake.removeFromReleaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFromReleaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Update(arg1 string, arg2 int, arg3 *string, arg4 *string, arg5 *string, arg6 *[]string) error {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 *string
		arg4 *string
		arg5 *string
		arg6 *[]string
	}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3, arg4, arg5, arg6)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1
}

func (fake *FakeImageReferenceClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeImageReferenceClient) UpdateCalls(stub func(string, int, *string, *string, *string, *[]string) error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeImageReferenceClient) UpdateArgsForCall(i int) (string, int, *string, *string, *string, *[]string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6
}

func (fake *FakeImageReferenceClient) UpdateReturns(result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) UpdateReturnsOnCall(i int, result1 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addToReleaseMutex.RLock()
	defer fake.addToReleaseMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeFromReleaseMutex.RLock()
	defer fake.removeFromReleaseMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageReferenceClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ImageReferenceClient = new(FakeImageReferenceClient)
