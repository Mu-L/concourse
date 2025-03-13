// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeResourceCache struct {
	BaseResourceTypeStub        func() *db.UsedBaseResourceType
	baseResourceTypeMutex       sync.RWMutex
	baseResourceTypeArgsForCall []struct {
	}
	baseResourceTypeReturns struct {
		result1 *db.UsedBaseResourceType
	}
	baseResourceTypeReturnsOnCall map[int]struct {
		result1 *db.UsedBaseResourceType
	}
	DestroyStub        func(db.Tx) (bool, error)
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 db.Tx
	}
	destroyReturns struct {
		result1 bool
		result2 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	ResourceConfigStub        func() db.ResourceConfig
	resourceConfigMutex       sync.RWMutex
	resourceConfigArgsForCall []struct {
	}
	resourceConfigReturns struct {
		result1 db.ResourceConfig
	}
	resourceConfigReturnsOnCall map[int]struct {
		result1 db.ResourceConfig
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceCache) BaseResourceType() *db.UsedBaseResourceType {
	fake.baseResourceTypeMutex.Lock()
	ret, specificReturn := fake.baseResourceTypeReturnsOnCall[len(fake.baseResourceTypeArgsForCall)]
	fake.baseResourceTypeArgsForCall = append(fake.baseResourceTypeArgsForCall, struct {
	}{})
	stub := fake.BaseResourceTypeStub
	fakeReturns := fake.baseResourceTypeReturns
	fake.recordInvocation("BaseResourceType", []interface{}{})
	fake.baseResourceTypeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResourceCache) BaseResourceTypeCallCount() int {
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	return len(fake.baseResourceTypeArgsForCall)
}

func (fake *FakeResourceCache) BaseResourceTypeCalls(stub func() *db.UsedBaseResourceType) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = stub
}

func (fake *FakeResourceCache) BaseResourceTypeReturns(result1 *db.UsedBaseResourceType) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = nil
	fake.baseResourceTypeReturns = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceCache) BaseResourceTypeReturnsOnCall(i int, result1 *db.UsedBaseResourceType) {
	fake.baseResourceTypeMutex.Lock()
	defer fake.baseResourceTypeMutex.Unlock()
	fake.BaseResourceTypeStub = nil
	if fake.baseResourceTypeReturnsOnCall == nil {
		fake.baseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.UsedBaseResourceType
		})
	}
	fake.baseResourceTypeReturnsOnCall[i] = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceCache) Destroy(arg1 db.Tx) (bool, error) {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 db.Tx
	}{arg1})
	stub := fake.DestroyStub
	fakeReturns := fake.destroyReturns
	fake.recordInvocation("Destroy", []interface{}{arg1})
	fake.destroyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceCache) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeResourceCache) DestroyCalls(stub func(db.Tx) (bool, error)) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeResourceCache) DestroyArgsForCall(i int) db.Tx {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceCache) DestroyReturns(result1 bool, result2 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCache) DestroyReturnsOnCall(i int, result1 bool, result2 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceCache) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResourceCache) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceCache) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeResourceCache) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceCache) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceCache) ResourceConfig() db.ResourceConfig {
	fake.resourceConfigMutex.Lock()
	ret, specificReturn := fake.resourceConfigReturnsOnCall[len(fake.resourceConfigArgsForCall)]
	fake.resourceConfigArgsForCall = append(fake.resourceConfigArgsForCall, struct {
	}{})
	stub := fake.ResourceConfigStub
	fakeReturns := fake.resourceConfigReturns
	fake.recordInvocation("ResourceConfig", []interface{}{})
	fake.resourceConfigMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResourceCache) ResourceConfigCallCount() int {
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	return len(fake.resourceConfigArgsForCall)
}

func (fake *FakeResourceCache) ResourceConfigCalls(stub func() db.ResourceConfig) {
	fake.resourceConfigMutex.Lock()
	defer fake.resourceConfigMutex.Unlock()
	fake.ResourceConfigStub = stub
}

func (fake *FakeResourceCache) ResourceConfigReturns(result1 db.ResourceConfig) {
	fake.resourceConfigMutex.Lock()
	defer fake.resourceConfigMutex.Unlock()
	fake.ResourceConfigStub = nil
	fake.resourceConfigReturns = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeResourceCache) ResourceConfigReturnsOnCall(i int, result1 db.ResourceConfig) {
	fake.resourceConfigMutex.Lock()
	defer fake.resourceConfigMutex.Unlock()
	fake.ResourceConfigStub = nil
	if fake.resourceConfigReturnsOnCall == nil {
		fake.resourceConfigReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfig
		})
	}
	fake.resourceConfigReturnsOnCall[i] = struct {
		result1 db.ResourceConfig
	}{result1}
}

func (fake *FakeResourceCache) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeResourceCache) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceCache) VersionCalls(stub func() atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeResourceCache) VersionReturns(result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceCache) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceCache) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.baseResourceTypeMutex.RLock()
	defer fake.baseResourceTypeMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.resourceConfigMutex.RLock()
	defer fake.resourceConfigMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceCache) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceCache = new(FakeResourceCache)
