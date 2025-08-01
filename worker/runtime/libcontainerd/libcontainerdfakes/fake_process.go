// Code generated by counterfeiter. DO NOT EDIT.
package libcontainerdfakes

import (
	"context"
	"sync"
	"syscall"

	"github.com/containerd/containerd/v2/client"
	"github.com/containerd/containerd/v2/pkg/cio"
)

type FakeProcess struct {
	CloseIOStub        func(context.Context, ...client.IOCloserOpts) error
	closeIOMutex       sync.RWMutex
	closeIOArgsForCall []struct {
		arg1 context.Context
		arg2 []client.IOCloserOpts
	}
	closeIOReturns struct {
		result1 error
	}
	closeIOReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, ...client.ProcessDeleteOpts) (*client.ExitStatus, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 []client.ProcessDeleteOpts
	}
	deleteReturns struct {
		result1 *client.ExitStatus
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *client.ExitStatus
		result2 error
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	IOStub        func() cio.IO
	iOMutex       sync.RWMutex
	iOArgsForCall []struct {
	}
	iOReturns struct {
		result1 cio.IO
	}
	iOReturnsOnCall map[int]struct {
		result1 cio.IO
	}
	KillStub        func(context.Context, syscall.Signal, ...client.KillOpts) error
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		arg1 context.Context
		arg2 syscall.Signal
		arg3 []client.KillOpts
	}
	killReturns struct {
		result1 error
	}
	killReturnsOnCall map[int]struct {
		result1 error
	}
	PidStub        func() uint32
	pidMutex       sync.RWMutex
	pidArgsForCall []struct {
	}
	pidReturns struct {
		result1 uint32
	}
	pidReturnsOnCall map[int]struct {
		result1 uint32
	}
	ResizeStub        func(context.Context, uint32, uint32) error
	resizeMutex       sync.RWMutex
	resizeArgsForCall []struct {
		arg1 context.Context
		arg2 uint32
		arg3 uint32
	}
	resizeReturns struct {
		result1 error
	}
	resizeReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(context.Context) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StatusStub        func(context.Context) (client.Status, error)
	statusMutex       sync.RWMutex
	statusArgsForCall []struct {
		arg1 context.Context
	}
	statusReturns struct {
		result1 client.Status
		result2 error
	}
	statusReturnsOnCall map[int]struct {
		result1 client.Status
		result2 error
	}
	WaitStub        func(context.Context) (<-chan client.ExitStatus, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
		arg1 context.Context
	}
	waitReturns struct {
		result1 <-chan client.ExitStatus
		result2 error
	}
	waitReturnsOnCall map[int]struct {
		result1 <-chan client.ExitStatus
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProcess) CloseIO(arg1 context.Context, arg2 ...client.IOCloserOpts) error {
	fake.closeIOMutex.Lock()
	ret, specificReturn := fake.closeIOReturnsOnCall[len(fake.closeIOArgsForCall)]
	fake.closeIOArgsForCall = append(fake.closeIOArgsForCall, struct {
		arg1 context.Context
		arg2 []client.IOCloserOpts
	}{arg1, arg2})
	stub := fake.CloseIOStub
	fakeReturns := fake.closeIOReturns
	fake.recordInvocation("CloseIO", []interface{}{arg1, arg2})
	fake.closeIOMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) CloseIOCallCount() int {
	fake.closeIOMutex.RLock()
	defer fake.closeIOMutex.RUnlock()
	return len(fake.closeIOArgsForCall)
}

func (fake *FakeProcess) CloseIOCalls(stub func(context.Context, ...client.IOCloserOpts) error) {
	fake.closeIOMutex.Lock()
	defer fake.closeIOMutex.Unlock()
	fake.CloseIOStub = stub
}

func (fake *FakeProcess) CloseIOArgsForCall(i int) (context.Context, []client.IOCloserOpts) {
	fake.closeIOMutex.RLock()
	defer fake.closeIOMutex.RUnlock()
	argsForCall := fake.closeIOArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProcess) CloseIOReturns(result1 error) {
	fake.closeIOMutex.Lock()
	defer fake.closeIOMutex.Unlock()
	fake.CloseIOStub = nil
	fake.closeIOReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) CloseIOReturnsOnCall(i int, result1 error) {
	fake.closeIOMutex.Lock()
	defer fake.closeIOMutex.Unlock()
	fake.CloseIOStub = nil
	if fake.closeIOReturnsOnCall == nil {
		fake.closeIOReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeIOReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) Delete(arg1 context.Context, arg2 ...client.ProcessDeleteOpts) (*client.ExitStatus, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 []client.ProcessDeleteOpts
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcess) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeProcess) DeleteCalls(stub func(context.Context, ...client.ProcessDeleteOpts) (*client.ExitStatus, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeProcess) DeleteArgsForCall(i int) (context.Context, []client.ProcessDeleteOpts) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProcess) DeleteReturns(result1 *client.ExitStatus, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *client.ExitStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) DeleteReturnsOnCall(i int, result1 *client.ExitStatus, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *client.ExitStatus
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *client.ExitStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) ID() string {
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

func (fake *FakeProcess) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeProcess) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeProcess) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeProcess) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeProcess) IO() cio.IO {
	fake.iOMutex.Lock()
	ret, specificReturn := fake.iOReturnsOnCall[len(fake.iOArgsForCall)]
	fake.iOArgsForCall = append(fake.iOArgsForCall, struct {
	}{})
	stub := fake.IOStub
	fakeReturns := fake.iOReturns
	fake.recordInvocation("IO", []interface{}{})
	fake.iOMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) IOCallCount() int {
	fake.iOMutex.RLock()
	defer fake.iOMutex.RUnlock()
	return len(fake.iOArgsForCall)
}

func (fake *FakeProcess) IOCalls(stub func() cio.IO) {
	fake.iOMutex.Lock()
	defer fake.iOMutex.Unlock()
	fake.IOStub = stub
}

func (fake *FakeProcess) IOReturns(result1 cio.IO) {
	fake.iOMutex.Lock()
	defer fake.iOMutex.Unlock()
	fake.IOStub = nil
	fake.iOReturns = struct {
		result1 cio.IO
	}{result1}
}

func (fake *FakeProcess) IOReturnsOnCall(i int, result1 cio.IO) {
	fake.iOMutex.Lock()
	defer fake.iOMutex.Unlock()
	fake.IOStub = nil
	if fake.iOReturnsOnCall == nil {
		fake.iOReturnsOnCall = make(map[int]struct {
			result1 cio.IO
		})
	}
	fake.iOReturnsOnCall[i] = struct {
		result1 cio.IO
	}{result1}
}

func (fake *FakeProcess) Kill(arg1 context.Context, arg2 syscall.Signal, arg3 ...client.KillOpts) error {
	fake.killMutex.Lock()
	ret, specificReturn := fake.killReturnsOnCall[len(fake.killArgsForCall)]
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		arg1 context.Context
		arg2 syscall.Signal
		arg3 []client.KillOpts
	}{arg1, arg2, arg3})
	stub := fake.KillStub
	fakeReturns := fake.killReturns
	fake.recordInvocation("Kill", []interface{}{arg1, arg2, arg3})
	fake.killMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeProcess) KillCalls(stub func(context.Context, syscall.Signal, ...client.KillOpts) error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = stub
}

func (fake *FakeProcess) KillArgsForCall(i int) (context.Context, syscall.Signal, []client.KillOpts) {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	argsForCall := fake.killArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeProcess) KillReturns(result1 error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) KillReturnsOnCall(i int, result1 error) {
	fake.killMutex.Lock()
	defer fake.killMutex.Unlock()
	fake.KillStub = nil
	if fake.killReturnsOnCall == nil {
		fake.killReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.killReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) Pid() uint32 {
	fake.pidMutex.Lock()
	ret, specificReturn := fake.pidReturnsOnCall[len(fake.pidArgsForCall)]
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct {
	}{})
	stub := fake.PidStub
	fakeReturns := fake.pidReturns
	fake.recordInvocation("Pid", []interface{}{})
	fake.pidMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

func (fake *FakeProcess) PidCalls(stub func() uint32) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = stub
}

func (fake *FakeProcess) PidReturns(result1 uint32) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeProcess) PidReturnsOnCall(i int, result1 uint32) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	if fake.pidReturnsOnCall == nil {
		fake.pidReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.pidReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeProcess) Resize(arg1 context.Context, arg2 uint32, arg3 uint32) error {
	fake.resizeMutex.Lock()
	ret, specificReturn := fake.resizeReturnsOnCall[len(fake.resizeArgsForCall)]
	fake.resizeArgsForCall = append(fake.resizeArgsForCall, struct {
		arg1 context.Context
		arg2 uint32
		arg3 uint32
	}{arg1, arg2, arg3})
	stub := fake.ResizeStub
	fakeReturns := fake.resizeReturns
	fake.recordInvocation("Resize", []interface{}{arg1, arg2, arg3})
	fake.resizeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) ResizeCallCount() int {
	fake.resizeMutex.RLock()
	defer fake.resizeMutex.RUnlock()
	return len(fake.resizeArgsForCall)
}

func (fake *FakeProcess) ResizeCalls(stub func(context.Context, uint32, uint32) error) {
	fake.resizeMutex.Lock()
	defer fake.resizeMutex.Unlock()
	fake.ResizeStub = stub
}

func (fake *FakeProcess) ResizeArgsForCall(i int) (context.Context, uint32, uint32) {
	fake.resizeMutex.RLock()
	defer fake.resizeMutex.RUnlock()
	argsForCall := fake.resizeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeProcess) ResizeReturns(result1 error) {
	fake.resizeMutex.Lock()
	defer fake.resizeMutex.Unlock()
	fake.ResizeStub = nil
	fake.resizeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) ResizeReturnsOnCall(i int, result1 error) {
	fake.resizeMutex.Lock()
	defer fake.resizeMutex.Unlock()
	fake.ResizeStub = nil
	if fake.resizeReturnsOnCall == nil {
		fake.resizeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resizeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) Start(arg1 context.Context) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProcess) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeProcess) StartCalls(stub func(context.Context) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeProcess) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProcess) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProcess) Status(arg1 context.Context) (client.Status, error) {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StatusStub
	fakeReturns := fake.statusReturns
	fake.recordInvocation("Status", []interface{}{arg1})
	fake.statusMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcess) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeProcess) StatusCalls(stub func(context.Context) (client.Status, error)) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = stub
}

func (fake *FakeProcess) StatusArgsForCall(i int) context.Context {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	argsForCall := fake.statusArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProcess) StatusReturns(result1 client.Status, result2 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 client.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) StatusReturnsOnCall(i int, result1 client.Status, result2 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 client.Status
			result2 error
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 client.Status
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) Wait(arg1 context.Context) (<-chan client.ExitStatus, error) {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.WaitStub
	fakeReturns := fake.waitReturns
	fake.recordInvocation("Wait", []interface{}{arg1})
	fake.waitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProcess) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeProcess) WaitCalls(stub func(context.Context) (<-chan client.ExitStatus, error)) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeProcess) WaitArgsForCall(i int) context.Context {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	argsForCall := fake.waitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProcess) WaitReturns(result1 <-chan client.ExitStatus, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 <-chan client.ExitStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) WaitReturnsOnCall(i int, result1 <-chan client.ExitStatus, result2 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 <-chan client.ExitStatus
			result2 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 <-chan client.ExitStatus
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProcess) recordInvocation(key string, args []interface{}) {
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

var _ client.Process = new(FakeProcess)
