// Code generated by counterfeiter. DO NOT EDIT.
package libcontainerdfakes

import (
	"context"
	"sync"

	"github.com/concourse/concourse/worker/runtime/libcontainerd"
	"github.com/containerd/containerd/v2/client"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeClient struct {
	ContainersStub        func(context.Context, ...string) ([]client.Container, error)
	containersMutex       sync.RWMutex
	containersArgsForCall []struct {
		arg1 context.Context
		arg2 []string
	}
	containersReturns struct {
		result1 []client.Container
		result2 error
	}
	containersReturnsOnCall map[int]struct {
		result1 []client.Container
		result2 error
	}
	DestroyStub        func(context.Context, string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	GetContainerStub        func(context.Context, string) (client.Container, error)
	getContainerMutex       sync.RWMutex
	getContainerArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getContainerReturns struct {
		result1 client.Container
		result2 error
	}
	getContainerReturnsOnCall map[int]struct {
		result1 client.Container
		result2 error
	}
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct {
	}
	initReturns struct {
		result1 error
	}
	initReturnsOnCall map[int]struct {
		result1 error
	}
	NewContainerStub        func(context.Context, string, map[string]string, *specs.Spec) (client.Container, error)
	newContainerMutex       sync.RWMutex
	newContainerArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
		arg4 *specs.Spec
	}
	newContainerReturns struct {
		result1 client.Container
		result2 error
	}
	newContainerReturnsOnCall map[int]struct {
		result1 client.Container
		result2 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	VersionStub        func(context.Context) error
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
		arg1 context.Context
	}
	versionReturns struct {
		result1 error
	}
	versionReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Containers(arg1 context.Context, arg2 ...string) ([]client.Container, error) {
	fake.containersMutex.Lock()
	ret, specificReturn := fake.containersReturnsOnCall[len(fake.containersArgsForCall)]
	fake.containersArgsForCall = append(fake.containersArgsForCall, struct {
		arg1 context.Context
		arg2 []string
	}{arg1, arg2})
	stub := fake.ContainersStub
	fakeReturns := fake.containersReturns
	fake.recordInvocation("Containers", []interface{}{arg1, arg2})
	fake.containersMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) ContainersCallCount() int {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return len(fake.containersArgsForCall)
}

func (fake *FakeClient) ContainersCalls(stub func(context.Context, ...string) ([]client.Container, error)) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = stub
}

func (fake *FakeClient) ContainersArgsForCall(i int) (context.Context, []string) {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	argsForCall := fake.containersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) ContainersReturns(result1 []client.Container, result2 error) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = nil
	fake.containersReturns = struct {
		result1 []client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ContainersReturnsOnCall(i int, result1 []client.Container, result2 error) {
	fake.containersMutex.Lock()
	defer fake.containersMutex.Unlock()
	fake.ContainersStub = nil
	if fake.containersReturnsOnCall == nil {
		fake.containersReturnsOnCall = make(map[int]struct {
			result1 []client.Container
			result2 error
		})
	}
	fake.containersReturnsOnCall[i] = struct {
		result1 []client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Destroy(arg1 context.Context, arg2 string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DestroyStub
	fakeReturns := fake.destroyReturns
	fake.recordInvocation("Destroy", []interface{}{arg1, arg2})
	fake.destroyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeClient) DestroyCalls(stub func(context.Context, string) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeClient) DestroyArgsForCall(i int) (context.Context, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) GetContainer(arg1 context.Context, arg2 string) (client.Container, error) {
	fake.getContainerMutex.Lock()
	ret, specificReturn := fake.getContainerReturnsOnCall[len(fake.getContainerArgsForCall)]
	fake.getContainerArgsForCall = append(fake.getContainerArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetContainerStub
	fakeReturns := fake.getContainerReturns
	fake.recordInvocation("GetContainer", []interface{}{arg1, arg2})
	fake.getContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) GetContainerCallCount() int {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	return len(fake.getContainerArgsForCall)
}

func (fake *FakeClient) GetContainerCalls(stub func(context.Context, string) (client.Container, error)) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = stub
}

func (fake *FakeClient) GetContainerArgsForCall(i int) (context.Context, string) {
	fake.getContainerMutex.RLock()
	defer fake.getContainerMutex.RUnlock()
	argsForCall := fake.getContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeClient) GetContainerReturns(result1 client.Container, result2 error) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = nil
	fake.getContainerReturns = struct {
		result1 client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetContainerReturnsOnCall(i int, result1 client.Container, result2 error) {
	fake.getContainerMutex.Lock()
	defer fake.getContainerMutex.Unlock()
	fake.GetContainerStub = nil
	if fake.getContainerReturnsOnCall == nil {
		fake.getContainerReturnsOnCall = make(map[int]struct {
			result1 client.Container
			result2 error
		})
	}
	fake.getContainerReturnsOnCall[i] = struct {
		result1 client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Init() error {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
	}{})
	stub := fake.InitStub
	fakeReturns := fake.initReturns
	fake.recordInvocation("Init", []interface{}{})
	fake.initMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeClient) InitCalls(stub func() error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeClient) InitReturns(result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) InitReturnsOnCall(i int, result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) NewContainer(arg1 context.Context, arg2 string, arg3 map[string]string, arg4 *specs.Spec) (client.Container, error) {
	fake.newContainerMutex.Lock()
	ret, specificReturn := fake.newContainerReturnsOnCall[len(fake.newContainerArgsForCall)]
	fake.newContainerArgsForCall = append(fake.newContainerArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]string
		arg4 *specs.Spec
	}{arg1, arg2, arg3, arg4})
	stub := fake.NewContainerStub
	fakeReturns := fake.newContainerReturns
	fake.recordInvocation("NewContainer", []interface{}{arg1, arg2, arg3, arg4})
	fake.newContainerMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeClient) NewContainerCallCount() int {
	fake.newContainerMutex.RLock()
	defer fake.newContainerMutex.RUnlock()
	return len(fake.newContainerArgsForCall)
}

func (fake *FakeClient) NewContainerCalls(stub func(context.Context, string, map[string]string, *specs.Spec) (client.Container, error)) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = stub
}

func (fake *FakeClient) NewContainerArgsForCall(i int) (context.Context, string, map[string]string, *specs.Spec) {
	fake.newContainerMutex.RLock()
	defer fake.newContainerMutex.RUnlock()
	argsForCall := fake.newContainerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeClient) NewContainerReturns(result1 client.Container, result2 error) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = nil
	fake.newContainerReturns = struct {
		result1 client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) NewContainerReturnsOnCall(i int, result1 client.Container, result2 error) {
	fake.newContainerMutex.Lock()
	defer fake.newContainerMutex.Unlock()
	fake.NewContainerStub = nil
	if fake.newContainerReturnsOnCall == nil {
		fake.newContainerReturnsOnCall = make(map[int]struct {
			result1 client.Container
			result2 error
		})
	}
	fake.newContainerReturnsOnCall[i] = struct {
		result1 client.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeClient) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeClient) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) Version(arg1 context.Context) error {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{arg1})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeClient) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeClient) VersionCalls(stub func(context.Context) error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeClient) VersionArgsForCall(i int) context.Context {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	argsForCall := fake.versionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeClient) VersionReturns(result1 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeClient) VersionReturnsOnCall(i int, result1 error) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
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

var _ libcontainerd.Client = new(FakeClient)
