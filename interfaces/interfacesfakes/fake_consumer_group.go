// Code generated by counterfeiter. DO NOT EDIT.
package interfacesfakes

import (
	"context"
	"sync"

	"github.com/IBM/sarama"
	"github.com/ehsaniara/delay-box/interfaces"
)

type FakeConsumerGroup struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ConsumeStub        func(context.Context, []string, sarama.ConsumerGroupHandler) error
	consumeMutex       sync.RWMutex
	consumeArgsForCall []struct {
		arg1 context.Context
		arg2 []string
		arg3 sarama.ConsumerGroupHandler
	}
	consumeReturns struct {
		result1 error
	}
	consumeReturnsOnCall map[int]struct {
		result1 error
	}
	ErrorsStub        func() <-chan error
	errorsMutex       sync.RWMutex
	errorsArgsForCall []struct {
	}
	errorsReturns struct {
		result1 <-chan error
	}
	errorsReturnsOnCall map[int]struct {
		result1 <-chan error
	}
	PauseStub        func(map[string][]int32)
	pauseMutex       sync.RWMutex
	pauseArgsForCall []struct {
		arg1 map[string][]int32
	}
	PauseAllStub        func()
	pauseAllMutex       sync.RWMutex
	pauseAllArgsForCall []struct {
	}
	ResumeStub        func(map[string][]int32)
	resumeMutex       sync.RWMutex
	resumeArgsForCall []struct {
		arg1 map[string][]int32
	}
	ResumeAllStub        func()
	resumeAllMutex       sync.RWMutex
	resumeAllArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConsumerGroup) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumerGroup) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConsumerGroup) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeConsumerGroup) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumerGroup) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumerGroup) Consume(arg1 context.Context, arg2 []string, arg3 sarama.ConsumerGroupHandler) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.consumeMutex.Lock()
	ret, specificReturn := fake.consumeReturnsOnCall[len(fake.consumeArgsForCall)]
	fake.consumeArgsForCall = append(fake.consumeArgsForCall, struct {
		arg1 context.Context
		arg2 []string
		arg3 sarama.ConsumerGroupHandler
	}{arg1, arg2Copy, arg3})
	stub := fake.ConsumeStub
	fakeReturns := fake.consumeReturns
	fake.recordInvocation("Consume", []interface{}{arg1, arg2Copy, arg3})
	fake.consumeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumerGroup) ConsumeCallCount() int {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	return len(fake.consumeArgsForCall)
}

func (fake *FakeConsumerGroup) ConsumeCalls(stub func(context.Context, []string, sarama.ConsumerGroupHandler) error) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = stub
}

func (fake *FakeConsumerGroup) ConsumeArgsForCall(i int) (context.Context, []string, sarama.ConsumerGroupHandler) {
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	argsForCall := fake.consumeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeConsumerGroup) ConsumeReturns(result1 error) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = nil
	fake.consumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumerGroup) ConsumeReturnsOnCall(i int, result1 error) {
	fake.consumeMutex.Lock()
	defer fake.consumeMutex.Unlock()
	fake.ConsumeStub = nil
	if fake.consumeReturnsOnCall == nil {
		fake.consumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.consumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumerGroup) Errors() <-chan error {
	fake.errorsMutex.Lock()
	ret, specificReturn := fake.errorsReturnsOnCall[len(fake.errorsArgsForCall)]
	fake.errorsArgsForCall = append(fake.errorsArgsForCall, struct {
	}{})
	stub := fake.ErrorsStub
	fakeReturns := fake.errorsReturns
	fake.recordInvocation("Errors", []interface{}{})
	fake.errorsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumerGroup) ErrorsCallCount() int {
	fake.errorsMutex.RLock()
	defer fake.errorsMutex.RUnlock()
	return len(fake.errorsArgsForCall)
}

func (fake *FakeConsumerGroup) ErrorsCalls(stub func() <-chan error) {
	fake.errorsMutex.Lock()
	defer fake.errorsMutex.Unlock()
	fake.ErrorsStub = stub
}

func (fake *FakeConsumerGroup) ErrorsReturns(result1 <-chan error) {
	fake.errorsMutex.Lock()
	defer fake.errorsMutex.Unlock()
	fake.ErrorsStub = nil
	fake.errorsReturns = struct {
		result1 <-chan error
	}{result1}
}

func (fake *FakeConsumerGroup) ErrorsReturnsOnCall(i int, result1 <-chan error) {
	fake.errorsMutex.Lock()
	defer fake.errorsMutex.Unlock()
	fake.ErrorsStub = nil
	if fake.errorsReturnsOnCall == nil {
		fake.errorsReturnsOnCall = make(map[int]struct {
			result1 <-chan error
		})
	}
	fake.errorsReturnsOnCall[i] = struct {
		result1 <-chan error
	}{result1}
}

func (fake *FakeConsumerGroup) Pause(arg1 map[string][]int32) {
	fake.pauseMutex.Lock()
	fake.pauseArgsForCall = append(fake.pauseArgsForCall, struct {
		arg1 map[string][]int32
	}{arg1})
	stub := fake.PauseStub
	fake.recordInvocation("Pause", []interface{}{arg1})
	fake.pauseMutex.Unlock()
	if stub != nil {
		fake.PauseStub(arg1)
	}
}

func (fake *FakeConsumerGroup) PauseCallCount() int {
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	return len(fake.pauseArgsForCall)
}

func (fake *FakeConsumerGroup) PauseCalls(stub func(map[string][]int32)) {
	fake.pauseMutex.Lock()
	defer fake.pauseMutex.Unlock()
	fake.PauseStub = stub
}

func (fake *FakeConsumerGroup) PauseArgsForCall(i int) map[string][]int32 {
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	argsForCall := fake.pauseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumerGroup) PauseAll() {
	fake.pauseAllMutex.Lock()
	fake.pauseAllArgsForCall = append(fake.pauseAllArgsForCall, struct {
	}{})
	stub := fake.PauseAllStub
	fake.recordInvocation("PauseAll", []interface{}{})
	fake.pauseAllMutex.Unlock()
	if stub != nil {
		fake.PauseAllStub()
	}
}

func (fake *FakeConsumerGroup) PauseAllCallCount() int {
	fake.pauseAllMutex.RLock()
	defer fake.pauseAllMutex.RUnlock()
	return len(fake.pauseAllArgsForCall)
}

func (fake *FakeConsumerGroup) PauseAllCalls(stub func()) {
	fake.pauseAllMutex.Lock()
	defer fake.pauseAllMutex.Unlock()
	fake.PauseAllStub = stub
}

func (fake *FakeConsumerGroup) Resume(arg1 map[string][]int32) {
	fake.resumeMutex.Lock()
	fake.resumeArgsForCall = append(fake.resumeArgsForCall, struct {
		arg1 map[string][]int32
	}{arg1})
	stub := fake.ResumeStub
	fake.recordInvocation("Resume", []interface{}{arg1})
	fake.resumeMutex.Unlock()
	if stub != nil {
		fake.ResumeStub(arg1)
	}
}

func (fake *FakeConsumerGroup) ResumeCallCount() int {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return len(fake.resumeArgsForCall)
}

func (fake *FakeConsumerGroup) ResumeCalls(stub func(map[string][]int32)) {
	fake.resumeMutex.Lock()
	defer fake.resumeMutex.Unlock()
	fake.ResumeStub = stub
}

func (fake *FakeConsumerGroup) ResumeArgsForCall(i int) map[string][]int32 {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	argsForCall := fake.resumeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumerGroup) ResumeAll() {
	fake.resumeAllMutex.Lock()
	fake.resumeAllArgsForCall = append(fake.resumeAllArgsForCall, struct {
	}{})
	stub := fake.ResumeAllStub
	fake.recordInvocation("ResumeAll", []interface{}{})
	fake.resumeAllMutex.Unlock()
	if stub != nil {
		fake.ResumeAllStub()
	}
}

func (fake *FakeConsumerGroup) ResumeAllCallCount() int {
	fake.resumeAllMutex.RLock()
	defer fake.resumeAllMutex.RUnlock()
	return len(fake.resumeAllArgsForCall)
}

func (fake *FakeConsumerGroup) ResumeAllCalls(stub func()) {
	fake.resumeAllMutex.Lock()
	defer fake.resumeAllMutex.Unlock()
	fake.ResumeAllStub = stub
}

func (fake *FakeConsumerGroup) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.consumeMutex.RLock()
	defer fake.consumeMutex.RUnlock()
	fake.errorsMutex.RLock()
	defer fake.errorsMutex.RUnlock()
	fake.pauseMutex.RLock()
	defer fake.pauseMutex.RUnlock()
	fake.pauseAllMutex.RLock()
	defer fake.pauseAllMutex.RUnlock()
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	fake.resumeAllMutex.RLock()
	defer fake.resumeAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConsumerGroup) recordInvocation(key string, args []interface{}) {
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

var _ interfaces.ConsumerGroup = new(FakeConsumerGroup)
