// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeAccessTokenFactory struct {
	CreateAccessTokenStub        func(string, db.Claims) error
	createAccessTokenMutex       sync.RWMutex
	createAccessTokenArgsForCall []struct {
		arg1 string
		arg2 db.Claims
	}
	createAccessTokenReturns struct {
		result1 error
	}
	createAccessTokenReturnsOnCall map[int]struct {
		result1 error
	}
	GetAccessTokenStub        func(string) (db.AccessToken, bool, error)
	getAccessTokenMutex       sync.RWMutex
	getAccessTokenArgsForCall []struct {
		arg1 string
	}
	getAccessTokenReturns struct {
		result1 db.AccessToken
		result2 bool
		result3 error
	}
	getAccessTokenReturnsOnCall map[int]struct {
		result1 db.AccessToken
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccessTokenFactory) CreateAccessToken(arg1 string, arg2 db.Claims) error {
	fake.createAccessTokenMutex.Lock()
	ret, specificReturn := fake.createAccessTokenReturnsOnCall[len(fake.createAccessTokenArgsForCall)]
	fake.createAccessTokenArgsForCall = append(fake.createAccessTokenArgsForCall, struct {
		arg1 string
		arg2 db.Claims
	}{arg1, arg2})
	stub := fake.CreateAccessTokenStub
	fakeReturns := fake.createAccessTokenReturns
	fake.recordInvocation("CreateAccessToken", []interface{}{arg1, arg2})
	fake.createAccessTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeAccessTokenFactory) CreateAccessTokenCallCount() int {
	fake.createAccessTokenMutex.RLock()
	defer fake.createAccessTokenMutex.RUnlock()
	return len(fake.createAccessTokenArgsForCall)
}

func (fake *FakeAccessTokenFactory) CreateAccessTokenCalls(stub func(string, db.Claims) error) {
	fake.createAccessTokenMutex.Lock()
	defer fake.createAccessTokenMutex.Unlock()
	fake.CreateAccessTokenStub = stub
}

func (fake *FakeAccessTokenFactory) CreateAccessTokenArgsForCall(i int) (string, db.Claims) {
	fake.createAccessTokenMutex.RLock()
	defer fake.createAccessTokenMutex.RUnlock()
	argsForCall := fake.createAccessTokenArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAccessTokenFactory) CreateAccessTokenReturns(result1 error) {
	fake.createAccessTokenMutex.Lock()
	defer fake.createAccessTokenMutex.Unlock()
	fake.CreateAccessTokenStub = nil
	fake.createAccessTokenReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAccessTokenFactory) CreateAccessTokenReturnsOnCall(i int, result1 error) {
	fake.createAccessTokenMutex.Lock()
	defer fake.createAccessTokenMutex.Unlock()
	fake.CreateAccessTokenStub = nil
	if fake.createAccessTokenReturnsOnCall == nil {
		fake.createAccessTokenReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createAccessTokenReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAccessTokenFactory) GetAccessToken(arg1 string) (db.AccessToken, bool, error) {
	fake.getAccessTokenMutex.Lock()
	ret, specificReturn := fake.getAccessTokenReturnsOnCall[len(fake.getAccessTokenArgsForCall)]
	fake.getAccessTokenArgsForCall = append(fake.getAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetAccessTokenStub
	fakeReturns := fake.getAccessTokenReturns
	fake.recordInvocation("GetAccessToken", []interface{}{arg1})
	fake.getAccessTokenMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAccessTokenFactory) GetAccessTokenCallCount() int {
	fake.getAccessTokenMutex.RLock()
	defer fake.getAccessTokenMutex.RUnlock()
	return len(fake.getAccessTokenArgsForCall)
}

func (fake *FakeAccessTokenFactory) GetAccessTokenCalls(stub func(string) (db.AccessToken, bool, error)) {
	fake.getAccessTokenMutex.Lock()
	defer fake.getAccessTokenMutex.Unlock()
	fake.GetAccessTokenStub = stub
}

func (fake *FakeAccessTokenFactory) GetAccessTokenArgsForCall(i int) string {
	fake.getAccessTokenMutex.RLock()
	defer fake.getAccessTokenMutex.RUnlock()
	argsForCall := fake.getAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAccessTokenFactory) GetAccessTokenReturns(result1 db.AccessToken, result2 bool, result3 error) {
	fake.getAccessTokenMutex.Lock()
	defer fake.getAccessTokenMutex.Unlock()
	fake.GetAccessTokenStub = nil
	fake.getAccessTokenReturns = struct {
		result1 db.AccessToken
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAccessTokenFactory) GetAccessTokenReturnsOnCall(i int, result1 db.AccessToken, result2 bool, result3 error) {
	fake.getAccessTokenMutex.Lock()
	defer fake.getAccessTokenMutex.Unlock()
	fake.GetAccessTokenStub = nil
	if fake.getAccessTokenReturnsOnCall == nil {
		fake.getAccessTokenReturnsOnCall = make(map[int]struct {
			result1 db.AccessToken
			result2 bool
			result3 error
		})
	}
	fake.getAccessTokenReturnsOnCall[i] = struct {
		result1 db.AccessToken
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAccessTokenFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createAccessTokenMutex.RLock()
	defer fake.createAccessTokenMutex.RUnlock()
	fake.getAccessTokenMutex.RLock()
	defer fake.getAccessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAccessTokenFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.AccessTokenFactory = new(FakeAccessTokenFactory)
