// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"github.com/concourse/concourse/atc/db"
)

type FakeUserFactory struct {
	CreateOrUpdateUserStub        func(string, string, string) error
	createOrUpdateUserMutex       sync.RWMutex
	createOrUpdateUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createOrUpdateUserReturns struct {
		result1 error
	}
	createOrUpdateUserReturnsOnCall map[int]struct {
		result1 error
	}
	GetAllUsersStub        func() ([]db.User, error)
	getAllUsersMutex       sync.RWMutex
	getAllUsersArgsForCall []struct {
	}
	getAllUsersReturns struct {
		result1 []db.User
		result2 error
	}
	getAllUsersReturnsOnCall map[int]struct {
		result1 []db.User
		result2 error
	}
	GetAllUsersByLoginDateStub        func(time.Time) ([]db.User, error)
	getAllUsersByLoginDateMutex       sync.RWMutex
	getAllUsersByLoginDateArgsForCall []struct {
		arg1 time.Time
	}
	getAllUsersByLoginDateReturns struct {
		result1 []db.User
		result2 error
	}
	getAllUsersByLoginDateReturnsOnCall map[int]struct {
		result1 []db.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserFactory) CreateOrUpdateUser(arg1 string, arg2 string, arg3 string) error {
	fake.createOrUpdateUserMutex.Lock()
	ret, specificReturn := fake.createOrUpdateUserReturnsOnCall[len(fake.createOrUpdateUserArgsForCall)]
	fake.createOrUpdateUserArgsForCall = append(fake.createOrUpdateUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.CreateOrUpdateUserStub
	fakeReturns := fake.createOrUpdateUserReturns
	fake.recordInvocation("CreateOrUpdateUser", []interface{}{arg1, arg2, arg3})
	fake.createOrUpdateUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserFactory) CreateOrUpdateUserCallCount() int {
	fake.createOrUpdateUserMutex.RLock()
	defer fake.createOrUpdateUserMutex.RUnlock()
	return len(fake.createOrUpdateUserArgsForCall)
}

func (fake *FakeUserFactory) CreateOrUpdateUserCalls(stub func(string, string, string) error) {
	fake.createOrUpdateUserMutex.Lock()
	defer fake.createOrUpdateUserMutex.Unlock()
	fake.CreateOrUpdateUserStub = stub
}

func (fake *FakeUserFactory) CreateOrUpdateUserArgsForCall(i int) (string, string, string) {
	fake.createOrUpdateUserMutex.RLock()
	defer fake.createOrUpdateUserMutex.RUnlock()
	argsForCall := fake.createOrUpdateUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserFactory) CreateOrUpdateUserReturns(result1 error) {
	fake.createOrUpdateUserMutex.Lock()
	defer fake.createOrUpdateUserMutex.Unlock()
	fake.CreateOrUpdateUserStub = nil
	fake.createOrUpdateUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserFactory) CreateOrUpdateUserReturnsOnCall(i int, result1 error) {
	fake.createOrUpdateUserMutex.Lock()
	defer fake.createOrUpdateUserMutex.Unlock()
	fake.CreateOrUpdateUserStub = nil
	if fake.createOrUpdateUserReturnsOnCall == nil {
		fake.createOrUpdateUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createOrUpdateUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserFactory) GetAllUsers() ([]db.User, error) {
	fake.getAllUsersMutex.Lock()
	ret, specificReturn := fake.getAllUsersReturnsOnCall[len(fake.getAllUsersArgsForCall)]
	fake.getAllUsersArgsForCall = append(fake.getAllUsersArgsForCall, struct {
	}{})
	stub := fake.GetAllUsersStub
	fakeReturns := fake.getAllUsersReturns
	fake.recordInvocation("GetAllUsers", []interface{}{})
	fake.getAllUsersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserFactory) GetAllUsersCallCount() int {
	fake.getAllUsersMutex.RLock()
	defer fake.getAllUsersMutex.RUnlock()
	return len(fake.getAllUsersArgsForCall)
}

func (fake *FakeUserFactory) GetAllUsersCalls(stub func() ([]db.User, error)) {
	fake.getAllUsersMutex.Lock()
	defer fake.getAllUsersMutex.Unlock()
	fake.GetAllUsersStub = stub
}

func (fake *FakeUserFactory) GetAllUsersReturns(result1 []db.User, result2 error) {
	fake.getAllUsersMutex.Lock()
	defer fake.getAllUsersMutex.Unlock()
	fake.GetAllUsersStub = nil
	fake.getAllUsersReturns = struct {
		result1 []db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFactory) GetAllUsersReturnsOnCall(i int, result1 []db.User, result2 error) {
	fake.getAllUsersMutex.Lock()
	defer fake.getAllUsersMutex.Unlock()
	fake.GetAllUsersStub = nil
	if fake.getAllUsersReturnsOnCall == nil {
		fake.getAllUsersReturnsOnCall = make(map[int]struct {
			result1 []db.User
			result2 error
		})
	}
	fake.getAllUsersReturnsOnCall[i] = struct {
		result1 []db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFactory) GetAllUsersByLoginDate(arg1 time.Time) ([]db.User, error) {
	fake.getAllUsersByLoginDateMutex.Lock()
	ret, specificReturn := fake.getAllUsersByLoginDateReturnsOnCall[len(fake.getAllUsersByLoginDateArgsForCall)]
	fake.getAllUsersByLoginDateArgsForCall = append(fake.getAllUsersByLoginDateArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	stub := fake.GetAllUsersByLoginDateStub
	fakeReturns := fake.getAllUsersByLoginDateReturns
	fake.recordInvocation("GetAllUsersByLoginDate", []interface{}{arg1})
	fake.getAllUsersByLoginDateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserFactory) GetAllUsersByLoginDateCallCount() int {
	fake.getAllUsersByLoginDateMutex.RLock()
	defer fake.getAllUsersByLoginDateMutex.RUnlock()
	return len(fake.getAllUsersByLoginDateArgsForCall)
}

func (fake *FakeUserFactory) GetAllUsersByLoginDateCalls(stub func(time.Time) ([]db.User, error)) {
	fake.getAllUsersByLoginDateMutex.Lock()
	defer fake.getAllUsersByLoginDateMutex.Unlock()
	fake.GetAllUsersByLoginDateStub = stub
}

func (fake *FakeUserFactory) GetAllUsersByLoginDateArgsForCall(i int) time.Time {
	fake.getAllUsersByLoginDateMutex.RLock()
	defer fake.getAllUsersByLoginDateMutex.RUnlock()
	argsForCall := fake.getAllUsersByLoginDateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserFactory) GetAllUsersByLoginDateReturns(result1 []db.User, result2 error) {
	fake.getAllUsersByLoginDateMutex.Lock()
	defer fake.getAllUsersByLoginDateMutex.Unlock()
	fake.GetAllUsersByLoginDateStub = nil
	fake.getAllUsersByLoginDateReturns = struct {
		result1 []db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFactory) GetAllUsersByLoginDateReturnsOnCall(i int, result1 []db.User, result2 error) {
	fake.getAllUsersByLoginDateMutex.Lock()
	defer fake.getAllUsersByLoginDateMutex.Unlock()
	fake.GetAllUsersByLoginDateStub = nil
	if fake.getAllUsersByLoginDateReturnsOnCall == nil {
		fake.getAllUsersByLoginDateReturnsOnCall = make(map[int]struct {
			result1 []db.User
			result2 error
		})
	}
	fake.getAllUsersByLoginDateReturnsOnCall[i] = struct {
		result1 []db.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createOrUpdateUserMutex.RLock()
	defer fake.createOrUpdateUserMutex.RUnlock()
	fake.getAllUsersMutex.RLock()
	defer fake.getAllUsersMutex.RUnlock()
	fake.getAllUsersByLoginDateMutex.RLock()
	defer fake.getAllUsersByLoginDateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.UserFactory = new(FakeUserFactory)
