// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/worker"
)

type FakeImageFactory struct {
	GetImageStub        func(lager.Logger, worker.Worker, worker.VolumeClient, worker.ImageSpec, int, worker.ImageFetchingDelegate, creds.VersionedResourceTypes, worker.Client) (worker.Image, error)
	getImageMutex       sync.RWMutex
	getImageArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Worker
		arg3 worker.VolumeClient
		arg4 worker.ImageSpec
		arg5 int
		arg6 worker.ImageFetchingDelegate
		arg7 creds.VersionedResourceTypes
		arg8 worker.Client
	}
	getImageReturns struct {
		result1 worker.Image
		result2 error
	}
	getImageReturnsOnCall map[int]struct {
		result1 worker.Image
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFactory) GetImage(arg1 lager.Logger, arg2 worker.Worker, arg3 worker.VolumeClient, arg4 worker.ImageSpec, arg5 int, arg6 worker.ImageFetchingDelegate, arg7 creds.VersionedResourceTypes, arg8 worker.Client) (worker.Image, error) {
	fake.getImageMutex.Lock()
	ret, specificReturn := fake.getImageReturnsOnCall[len(fake.getImageArgsForCall)]
	fake.getImageArgsForCall = append(fake.getImageArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Worker
		arg3 worker.VolumeClient
		arg4 worker.ImageSpec
		arg5 int
		arg6 worker.ImageFetchingDelegate
		arg7 creds.VersionedResourceTypes
		arg8 worker.Client
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.recordInvocation("GetImage", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8})
	fake.getImageMutex.Unlock()
	if fake.GetImageStub != nil {
		return fake.GetImageStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getImageReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageFactory) GetImageCallCount() int {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	return len(fake.getImageArgsForCall)
}

func (fake *FakeImageFactory) GetImageCalls(stub func(lager.Logger, worker.Worker, worker.VolumeClient, worker.ImageSpec, int, worker.ImageFetchingDelegate, creds.VersionedResourceTypes, worker.Client) (worker.Image, error)) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = stub
}

func (fake *FakeImageFactory) GetImageArgsForCall(i int) (lager.Logger, worker.Worker, worker.VolumeClient, worker.ImageSpec, int, worker.ImageFetchingDelegate, creds.VersionedResourceTypes, worker.Client) {
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	argsForCall := fake.getImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7, argsForCall.arg8
}

func (fake *FakeImageFactory) GetImageReturns(result1 worker.Image, result2 error) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = nil
	fake.getImageReturns = struct {
		result1 worker.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFactory) GetImageReturnsOnCall(i int, result1 worker.Image, result2 error) {
	fake.getImageMutex.Lock()
	defer fake.getImageMutex.Unlock()
	fake.GetImageStub = nil
	if fake.getImageReturnsOnCall == nil {
		fake.getImageReturnsOnCall = make(map[int]struct {
			result1 worker.Image
			result2 error
		})
	}
	fake.getImageReturnsOnCall[i] = struct {
		result1 worker.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeImageFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getImageMutex.RLock()
	defer fake.getImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.ImageFactory = new(FakeImageFactory)
