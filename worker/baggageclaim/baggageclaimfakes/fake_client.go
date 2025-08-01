// Code generated by counterfeiter. DO NOT EDIT.
package baggageclaimfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim"
)

type FakeClient struct {
	CreateVolumeStub        func(context.Context, string, baggageclaim.VolumeSpec) (baggageclaim.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.VolumeSpec
	}
	createVolumeReturns struct {
		result1 baggageclaim.Volume
		result2 error
	}
	createVolumeReturnsOnCall map[int]struct {
		result1 baggageclaim.Volume
		result2 error
	}
	DestroyVolumeStub        func(context.Context, string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	destroyVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	DestroyVolumesStub        func(context.Context, []string) error
	destroyVolumesMutex       sync.RWMutex
	destroyVolumesArgsForCall []struct {
		arg1 context.Context
		arg2 []string
	}
	destroyVolumesReturns struct {
		result1 error
	}
	destroyVolumesReturnsOnCall map[int]struct {
		result1 error
	}
	ListVolumesStub        func(context.Context, baggageclaim.VolumeProperties) (baggageclaim.Volumes, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 context.Context
		arg2 baggageclaim.VolumeProperties
	}
	listVolumesReturns struct {
		result1 baggageclaim.Volumes
		result2 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 baggageclaim.Volumes
		result2 error
	}
	LookupVolumeStub        func(context.Context, string) (baggageclaim.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 baggageclaim.Volume
		result2 bool
		result3 error
	}
	lookupVolumeReturnsOnCall map[int]struct {
		result1 baggageclaim.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) CreateVolume(arg1 context.Context, arg2 string, arg3 baggageclaim.VolumeSpec) (baggageclaim.Volume, error) {
	fake.createVolumeMutex.Lock()
	ret, specificReturn := fake.createVolumeReturnsOnCall[len(fake.createVolumeArgsForCall)]
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.VolumeSpec
	}{arg1, arg2, arg3})
	stub := fake.CreateVolumeStub
	fakeReturns := fake.createVolumeReturns
	fake.recordInvocation("CreateVolume", []interface{}{arg1, arg2, arg3})
	fake.createVolumeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeClient) CreateVolumeCalls(stub func(context.Context, string, baggageclaim.VolumeSpec) (baggageclaim.Volume, error)) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = stub
}

func (fake *FakeClient) CreateVolumeArgsForCall(i int) (context.Context, string, baggageclaim.VolumeSpec) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	argsForCall := fake.createVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeClient) CreateVolumeReturns(result1 baggageclaim.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) CreateVolumeReturnsOnCall(i int, result1 baggageclaim.Volume, result2 error) {
	fake.createVolumeMutex.Lock()
	defer fake.createVolumeMutex.Unlock()
	fake.CreateVolumeStub = nil
	if fake.createVolumeReturnsOnCall == nil {
		fake.createVolumeReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Volume
			result2 error
		})
	}
	fake.createVolumeReturnsOnCall[i] = struct {
		result1 baggageclaim.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) DestroyVolume(arg1 context.Context, arg2 string) error {
	fake.destroyVolumeMutex.Lock()
	ret, specificReturn := fake.destroyVolumeReturnsOnCall[len(fake.destroyVolumeArgsForCall)]
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DestroyVolumeStub
	fakeReturns := fake.destroyVolumeReturns
	fake.recordInvocation("DestroyVolume", []interface{}{arg1, arg2})
	fake.destroyVolumeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeClient) DestroyVolumeCalls(stub func(context.Context, string) error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = stub
}

func (fake *FakeClient) DestroyVolumeArgsForCall(i int) (context.Context, string) {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	argsForCall := fake.destroyVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DestroyVolumeReturns(result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyVolumeReturnsOnCall(i int, result1 error) {
	fake.destroyVolumeMutex.Lock()
	defer fake.destroyVolumeMutex.Unlock()
	fake.DestroyVolumeStub = nil
	if fake.destroyVolumeReturnsOnCall == nil {
		fake.destroyVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyVolumes(arg1 context.Context, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.destroyVolumesMutex.Lock()
	ret, specificReturn := fake.destroyVolumesReturnsOnCall[len(fake.destroyVolumesArgsForCall)]
	fake.destroyVolumesArgsForCall = append(fake.destroyVolumesArgsForCall, struct {
		arg1 context.Context
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.DestroyVolumesStub
	fakeReturns := fake.destroyVolumesReturns
	fake.recordInvocation("DestroyVolumes", []interface{}{arg1, arg2Copy})
	fake.destroyVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) DestroyVolumesCallCount() int {
	fake.destroyVolumesMutex.RLock()
	defer fake.destroyVolumesMutex.RUnlock()
	return len(fake.destroyVolumesArgsForCall)
}

func (fake *FakeClient) DestroyVolumesCalls(stub func(context.Context, []string) error) {
	fake.destroyVolumesMutex.Lock()
	defer fake.destroyVolumesMutex.Unlock()
	fake.DestroyVolumesStub = stub
}

func (fake *FakeClient) DestroyVolumesArgsForCall(i int) (context.Context, []string) {
	fake.destroyVolumesMutex.RLock()
	defer fake.destroyVolumesMutex.RUnlock()
	argsForCall := fake.destroyVolumesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DestroyVolumesReturns(result1 error) {
	fake.destroyVolumesMutex.Lock()
	defer fake.destroyVolumesMutex.Unlock()
	fake.DestroyVolumesStub = nil
	fake.destroyVolumesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyVolumesReturnsOnCall(i int, result1 error) {
	fake.destroyVolumesMutex.Lock()
	defer fake.destroyVolumesMutex.Unlock()
	fake.DestroyVolumesStub = nil
	if fake.destroyVolumesReturnsOnCall == nil {
		fake.destroyVolumesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyVolumesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) ListVolumes(arg1 context.Context, arg2 baggageclaim.VolumeProperties) (baggageclaim.Volumes, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 context.Context
		arg2 baggageclaim.VolumeProperties
	}{arg1, arg2})
	stub := fake.ListVolumesStub
	fakeReturns := fake.listVolumesReturns
	fake.recordInvocation("ListVolumes", []interface{}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeClient) ListVolumesCalls(stub func(context.Context, baggageclaim.VolumeProperties) (baggageclaim.Volumes, error)) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = stub
}

func (fake *FakeClient) ListVolumesArgsForCall(i int) (context.Context, baggageclaim.VolumeProperties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	argsForCall := fake.listVolumesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ListVolumesReturns(result1 baggageclaim.Volumes, result2 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 baggageclaim.Volumes
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListVolumesReturnsOnCall(i int, result1 baggageclaim.Volumes, result2 error) {
	fake.listVolumesMutex.Lock()
	defer fake.listVolumesMutex.Unlock()
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Volumes
			result2 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 baggageclaim.Volumes
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) LookupVolume(arg1 context.Context, arg2 string) (baggageclaim.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	ret, specificReturn := fake.lookupVolumeReturnsOnCall[len(fake.lookupVolumeArgsForCall)]
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.LookupVolumeStub
	fakeReturns := fake.lookupVolumeReturns
	fake.recordInvocation("LookupVolume", []interface{}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeClient) LookupVolumeCalls(stub func(context.Context, string) (baggageclaim.Volume, bool, error)) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = stub
}

func (fake *FakeClient) LookupVolumeArgsForCall(i int) (context.Context, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	argsForCall := fake.lookupVolumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) LookupVolumeReturns(result1 baggageclaim.Volume, result2 bool, result3 error) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 baggageclaim.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) LookupVolumeReturnsOnCall(i int, result1 baggageclaim.Volume, result2 bool, result3 error) {
	fake.lookupVolumeMutex.Lock()
	defer fake.lookupVolumeMutex.Unlock()
	fake.LookupVolumeStub = nil
	if fake.lookupVolumeReturnsOnCall == nil {
		fake.lookupVolumeReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.Volume
			result2 bool
			result3 error
		})
	}
	fake.lookupVolumeReturnsOnCall[i] = struct {
		result1 baggageclaim.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
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

var _ baggageclaim.Client = new(FakeClient)
