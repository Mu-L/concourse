// Code generated by counterfeiter. DO NOT EDIT.
package baggageclaimfakes

import (
	"context"
	"io"
	"sync"

	"github.com/concourse/concourse/worker/baggageclaim"
)

type FakeVolume struct {
	DestroyStub        func(context.Context) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 context.Context
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	GetPrivilegedStub        func(context.Context) (bool, error)
	getPrivilegedMutex       sync.RWMutex
	getPrivilegedArgsForCall []struct {
		arg1 context.Context
	}
	getPrivilegedReturns struct {
		result1 bool
		result2 error
	}
	getPrivilegedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	GetStreamInP2pUrlStub        func(context.Context, string) (string, error)
	getStreamInP2pUrlMutex       sync.RWMutex
	getStreamInP2pUrlArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getStreamInP2pUrlReturns struct {
		result1 string
		result2 error
	}
	getStreamInP2pUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
	}
	pathReturns struct {
		result1 string
	}
	pathReturnsOnCall map[int]struct {
		result1 string
	}
	PropertiesStub        func(context.Context) (baggageclaim.VolumeProperties, error)
	propertiesMutex       sync.RWMutex
	propertiesArgsForCall []struct {
		arg1 context.Context
	}
	propertiesReturns struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	propertiesReturnsOnCall map[int]struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}
	SetPrivilegedStub        func(context.Context, bool) error
	setPrivilegedMutex       sync.RWMutex
	setPrivilegedArgsForCall []struct {
		arg1 context.Context
		arg2 bool
	}
	setPrivilegedReturns struct {
		result1 error
	}
	setPrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	SetPropertyStub        func(context.Context, string, string) error
	setPropertyMutex       sync.RWMutex
	setPropertyArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	setPropertyReturns struct {
		result1 error
	}
	setPropertyReturnsOnCall map[int]struct {
		result1 error
	}
	StreamInStub        func(context.Context, string, baggageclaim.Encoding, float64, io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 float64
		arg5 io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	streamInReturnsOnCall map[int]struct {
		result1 error
	}
	StreamOutStub        func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	streamOutReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StreamP2pOutStub        func(context.Context, string, string, baggageclaim.Encoding) error
	streamP2pOutMutex       sync.RWMutex
	streamP2pOutArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}
	streamP2pOutReturns struct {
		result1 error
	}
	streamP2pOutReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolume) Destroy(arg1 context.Context) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.DestroyStub
	fakeReturns := fake.destroyReturns
	fake.recordInvocation("Destroy", []interface{}{arg1})
	fake.destroyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeVolume) DestroyCalls(stub func(context.Context) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeVolume) DestroyArgsForCall(i int) context.Context {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) DestroyReturnsOnCall(i int, result1 error) {
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

func (fake *FakeVolume) GetPrivileged(arg1 context.Context) (bool, error) {
	fake.getPrivilegedMutex.Lock()
	ret, specificReturn := fake.getPrivilegedReturnsOnCall[len(fake.getPrivilegedArgsForCall)]
	fake.getPrivilegedArgsForCall = append(fake.getPrivilegedArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetPrivilegedStub
	fakeReturns := fake.getPrivilegedReturns
	fake.recordInvocation("GetPrivileged", []interface{}{arg1})
	fake.getPrivilegedMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) GetPrivilegedCallCount() int {
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	return len(fake.getPrivilegedArgsForCall)
}

func (fake *FakeVolume) GetPrivilegedCalls(stub func(context.Context) (bool, error)) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = stub
}

func (fake *FakeVolume) GetPrivilegedArgsForCall(i int) context.Context {
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	argsForCall := fake.getPrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) GetPrivilegedReturns(result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	fake.getPrivilegedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetPrivilegedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.getPrivilegedMutex.Lock()
	defer fake.getPrivilegedMutex.Unlock()
	fake.GetPrivilegedStub = nil
	if fake.getPrivilegedReturnsOnCall == nil {
		fake.getPrivilegedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.getPrivilegedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetStreamInP2pUrl(arg1 context.Context, arg2 string) (string, error) {
	fake.getStreamInP2pUrlMutex.Lock()
	ret, specificReturn := fake.getStreamInP2pUrlReturnsOnCall[len(fake.getStreamInP2pUrlArgsForCall)]
	fake.getStreamInP2pUrlArgsForCall = append(fake.getStreamInP2pUrlArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStreamInP2pUrlStub
	fakeReturns := fake.getStreamInP2pUrlReturns
	fake.recordInvocation("GetStreamInP2pUrl", []interface{}{arg1, arg2})
	fake.getStreamInP2pUrlMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) GetStreamInP2pUrlCallCount() int {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	return len(fake.getStreamInP2pUrlArgsForCall)
}

func (fake *FakeVolume) GetStreamInP2pUrlCalls(stub func(context.Context, string) (string, error)) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = stub
}

func (fake *FakeVolume) GetStreamInP2pUrlArgsForCall(i int) (context.Context, string) {
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	argsForCall := fake.getStreamInP2pUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) GetStreamInP2pUrlReturns(result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	fake.getStreamInP2pUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) GetStreamInP2pUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getStreamInP2pUrlMutex.Lock()
	defer fake.getStreamInP2pUrlMutex.Unlock()
	fake.GetStreamInP2pUrlStub = nil
	if fake.getStreamInP2pUrlReturnsOnCall == nil {
		fake.getStreamInP2pUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getStreamInP2pUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	stub := fake.HandleStub
	fakeReturns := fake.handleReturns
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Path() string {
	fake.pathMutex.Lock()
	ret, specificReturn := fake.pathReturnsOnCall[len(fake.pathArgsForCall)]
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
	}{})
	stub := fake.PathStub
	fakeReturns := fake.pathReturns
	fake.recordInvocation("Path", []interface{}{})
	fake.pathMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolume) PathCalls(stub func() string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = stub
}

func (fake *FakeVolume) PathReturns(result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) PathReturnsOnCall(i int, result1 string) {
	fake.pathMutex.Lock()
	defer fake.pathMutex.Unlock()
	fake.PathStub = nil
	if fake.pathReturnsOnCall == nil {
		fake.pathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeVolume) Properties(arg1 context.Context) (baggageclaim.VolumeProperties, error) {
	fake.propertiesMutex.Lock()
	ret, specificReturn := fake.propertiesReturnsOnCall[len(fake.propertiesArgsForCall)]
	fake.propertiesArgsForCall = append(fake.propertiesArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.PropertiesStub
	fakeReturns := fake.propertiesReturns
	fake.recordInvocation("Properties", []interface{}{arg1})
	fake.propertiesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) PropertiesCallCount() int {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	return len(fake.propertiesArgsForCall)
}

func (fake *FakeVolume) PropertiesCalls(stub func(context.Context) (baggageclaim.VolumeProperties, error)) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = stub
}

func (fake *FakeVolume) PropertiesArgsForCall(i int) context.Context {
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	argsForCall := fake.propertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVolume) PropertiesReturns(result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	fake.propertiesReturns = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) PropertiesReturnsOnCall(i int, result1 baggageclaim.VolumeProperties, result2 error) {
	fake.propertiesMutex.Lock()
	defer fake.propertiesMutex.Unlock()
	fake.PropertiesStub = nil
	if fake.propertiesReturnsOnCall == nil {
		fake.propertiesReturnsOnCall = make(map[int]struct {
			result1 baggageclaim.VolumeProperties
			result2 error
		})
	}
	fake.propertiesReturnsOnCall[i] = struct {
		result1 baggageclaim.VolumeProperties
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) SetPrivileged(arg1 context.Context, arg2 bool) error {
	fake.setPrivilegedMutex.Lock()
	ret, specificReturn := fake.setPrivilegedReturnsOnCall[len(fake.setPrivilegedArgsForCall)]
	fake.setPrivilegedArgsForCall = append(fake.setPrivilegedArgsForCall, struct {
		arg1 context.Context
		arg2 bool
	}{arg1, arg2})
	stub := fake.SetPrivilegedStub
	fakeReturns := fake.setPrivilegedReturns
	fake.recordInvocation("SetPrivileged", []interface{}{arg1, arg2})
	fake.setPrivilegedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPrivilegedCallCount() int {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	return len(fake.setPrivilegedArgsForCall)
}

func (fake *FakeVolume) SetPrivilegedCalls(stub func(context.Context, bool) error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = stub
}

func (fake *FakeVolume) SetPrivilegedArgsForCall(i int) (context.Context, bool) {
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	argsForCall := fake.setPrivilegedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeVolume) SetPrivilegedReturns(result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	fake.setPrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPrivilegedReturnsOnCall(i int, result1 error) {
	fake.setPrivilegedMutex.Lock()
	defer fake.setPrivilegedMutex.Unlock()
	fake.SetPrivilegedStub = nil
	if fake.setPrivilegedReturnsOnCall == nil {
		fake.setPrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetProperty(arg1 context.Context, arg2 string, arg3 string) error {
	fake.setPropertyMutex.Lock()
	ret, specificReturn := fake.setPropertyReturnsOnCall[len(fake.setPropertyArgsForCall)]
	fake.setPropertyArgsForCall = append(fake.setPropertyArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.SetPropertyStub
	fakeReturns := fake.setPropertyReturns
	fake.recordInvocation("SetProperty", []interface{}{arg1, arg2, arg3})
	fake.setPropertyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) SetPropertyCallCount() int {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	return len(fake.setPropertyArgsForCall)
}

func (fake *FakeVolume) SetPropertyCalls(stub func(context.Context, string, string) error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = stub
}

func (fake *FakeVolume) SetPropertyArgsForCall(i int) (context.Context, string, string) {
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	argsForCall := fake.setPropertyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolume) SetPropertyReturns(result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	fake.setPropertyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) SetPropertyReturnsOnCall(i int, result1 error) {
	fake.setPropertyMutex.Lock()
	defer fake.setPropertyMutex.Unlock()
	fake.SetPropertyStub = nil
	if fake.setPropertyReturnsOnCall == nil {
		fake.setPropertyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPropertyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamIn(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding, arg4 float64, arg5 io.Reader) error {
	fake.streamInMutex.Lock()
	ret, specificReturn := fake.streamInReturnsOnCall[len(fake.streamInArgsForCall)]
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
		arg4 float64
		arg5 io.Reader
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.StreamInStub
	fakeReturns := fake.streamInReturns
	fake.recordInvocation("StreamIn", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.streamInMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeVolume) StreamInCalls(stub func(context.Context, string, baggageclaim.Encoding, float64, io.Reader) error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = stub
}

func (fake *FakeVolume) StreamInArgsForCall(i int) (context.Context, string, baggageclaim.Encoding, float64, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	argsForCall := fake.streamInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeVolume) StreamInReturns(result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamInReturnsOnCall(i int, result1 error) {
	fake.streamInMutex.Lock()
	defer fake.streamInMutex.Unlock()
	fake.StreamInStub = nil
	if fake.streamInReturnsOnCall == nil {
		fake.streamInReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamInReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamOut(arg1 context.Context, arg2 string, arg3 baggageclaim.Encoding) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	ret, specificReturn := fake.streamOutReturnsOnCall[len(fake.streamOutArgsForCall)]
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 baggageclaim.Encoding
	}{arg1, arg2, arg3})
	stub := fake.StreamOutStub
	fakeReturns := fake.streamOutReturns
	fake.recordInvocation("StreamOut", []interface{}{arg1, arg2, arg3})
	fake.streamOutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolume) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeVolume) StreamOutCalls(stub func(context.Context, string, baggageclaim.Encoding) (io.ReadCloser, error)) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = stub
}

func (fake *FakeVolume) StreamOutArgsForCall(i int) (context.Context, string, baggageclaim.Encoding) {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	argsForCall := fake.streamOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolume) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamOutReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.streamOutMutex.Lock()
	defer fake.streamOutMutex.Unlock()
	fake.StreamOutStub = nil
	if fake.streamOutReturnsOnCall == nil {
		fake.streamOutReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.streamOutReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeVolume) StreamP2pOut(arg1 context.Context, arg2 string, arg3 string, arg4 baggageclaim.Encoding) error {
	fake.streamP2pOutMutex.Lock()
	ret, specificReturn := fake.streamP2pOutReturnsOnCall[len(fake.streamP2pOutArgsForCall)]
	fake.streamP2pOutArgsForCall = append(fake.streamP2pOutArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 baggageclaim.Encoding
	}{arg1, arg2, arg3, arg4})
	stub := fake.StreamP2pOutStub
	fakeReturns := fake.streamP2pOutReturns
	fake.recordInvocation("StreamP2pOut", []interface{}{arg1, arg2, arg3, arg4})
	fake.streamP2pOutMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVolume) StreamP2pOutCallCount() int {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	return len(fake.streamP2pOutArgsForCall)
}

func (fake *FakeVolume) StreamP2pOutCalls(stub func(context.Context, string, string, baggageclaim.Encoding) error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = stub
}

func (fake *FakeVolume) StreamP2pOutArgsForCall(i int) (context.Context, string, string, baggageclaim.Encoding) {
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	argsForCall := fake.streamP2pOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeVolume) StreamP2pOutReturns(result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	fake.streamP2pOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) StreamP2pOutReturnsOnCall(i int, result1 error) {
	fake.streamP2pOutMutex.Lock()
	defer fake.streamP2pOutMutex.Unlock()
	fake.StreamP2pOutStub = nil
	if fake.streamP2pOutReturnsOnCall == nil {
		fake.streamP2pOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.streamP2pOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.getPrivilegedMutex.RLock()
	defer fake.getPrivilegedMutex.RUnlock()
	fake.getStreamInP2pUrlMutex.RLock()
	defer fake.getStreamInP2pUrlMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.propertiesMutex.RLock()
	defer fake.propertiesMutex.RUnlock()
	fake.setPrivilegedMutex.RLock()
	defer fake.setPrivilegedMutex.RUnlock()
	fake.setPropertyMutex.RLock()
	defer fake.setPropertyMutex.RUnlock()
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	fake.streamP2pOutMutex.RLock()
	defer fake.streamP2pOutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolume) recordInvocation(key string, args []interface{}) {
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

var _ baggageclaim.Volume = new(FakeVolume)
