// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	context "context"
	sync "sync"

	client "github.com/operator-framework/operator-registry/pkg/client"
	registry "github.com/operator-framework/operator-registry/pkg/registry"
)

type FakeInterface struct {
	GetBundleStub        func(context.Context, string, string, string) (*registry.Bundle, error)
	getBundleMutex       sync.RWMutex
	getBundleArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getBundleReturns struct {
		result1 *registry.Bundle
		result2 error
	}
	getBundleReturnsOnCall map[int]struct {
		result1 *registry.Bundle
		result2 error
	}
	GetBundleInPackageChannelStub        func(context.Context, string, string) (*registry.Bundle, error)
	getBundleInPackageChannelMutex       sync.RWMutex
	getBundleInPackageChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	getBundleInPackageChannelReturns struct {
		result1 *registry.Bundle
		result2 error
	}
	getBundleInPackageChannelReturnsOnCall map[int]struct {
		result1 *registry.Bundle
		result2 error
	}
	GetBundleThatProvidesStub        func(context.Context, string, string, string) (*registry.Bundle, error)
	getBundleThatProvidesMutex       sync.RWMutex
	getBundleThatProvidesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getBundleThatProvidesReturns struct {
		result1 *registry.Bundle
		result2 error
	}
	getBundleThatProvidesReturnsOnCall map[int]struct {
		result1 *registry.Bundle
		result2 error
	}
	GetReplacementBundleInPackageChannelStub        func(context.Context, string, string, string) (*registry.Bundle, error)
	getReplacementBundleInPackageChannelMutex       sync.RWMutex
	getReplacementBundleInPackageChannelArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}
	getReplacementBundleInPackageChannelReturns struct {
		result1 *registry.Bundle
		result2 error
	}
	getReplacementBundleInPackageChannelReturnsOnCall map[int]struct {
		result1 *registry.Bundle
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterface) GetBundle(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*registry.Bundle, error) {
	fake.getBundleMutex.Lock()
	ret, specificReturn := fake.getBundleReturnsOnCall[len(fake.getBundleArgsForCall)]
	fake.getBundleArgsForCall = append(fake.getBundleArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetBundle", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBundleMutex.Unlock()
	if fake.GetBundleStub != nil {
		return fake.GetBundleStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) GetBundleCallCount() int {
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	return len(fake.getBundleArgsForCall)
}

func (fake *FakeInterface) GetBundleCalls(stub func(context.Context, string, string, string) (*registry.Bundle, error)) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = stub
}

func (fake *FakeInterface) GetBundleArgsForCall(i int) (context.Context, string, string, string) {
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	argsForCall := fake.getBundleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterface) GetBundleReturns(result1 *registry.Bundle, result2 error) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = nil
	fake.getBundleReturns = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetBundleReturnsOnCall(i int, result1 *registry.Bundle, result2 error) {
	fake.getBundleMutex.Lock()
	defer fake.getBundleMutex.Unlock()
	fake.GetBundleStub = nil
	if fake.getBundleReturnsOnCall == nil {
		fake.getBundleReturnsOnCall = make(map[int]struct {
			result1 *registry.Bundle
			result2 error
		})
	}
	fake.getBundleReturnsOnCall[i] = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetBundleInPackageChannel(arg1 context.Context, arg2 string, arg3 string) (*registry.Bundle, error) {
	fake.getBundleInPackageChannelMutex.Lock()
	ret, specificReturn := fake.getBundleInPackageChannelReturnsOnCall[len(fake.getBundleInPackageChannelArgsForCall)]
	fake.getBundleInPackageChannelArgsForCall = append(fake.getBundleInPackageChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetBundleInPackageChannel", []interface{}{arg1, arg2, arg3})
	fake.getBundleInPackageChannelMutex.Unlock()
	if fake.GetBundleInPackageChannelStub != nil {
		return fake.GetBundleInPackageChannelStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleInPackageChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) GetBundleInPackageChannelCallCount() int {
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	return len(fake.getBundleInPackageChannelArgsForCall)
}

func (fake *FakeInterface) GetBundleInPackageChannelCalls(stub func(context.Context, string, string) (*registry.Bundle, error)) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = stub
}

func (fake *FakeInterface) GetBundleInPackageChannelArgsForCall(i int) (context.Context, string, string) {
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	argsForCall := fake.getBundleInPackageChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeInterface) GetBundleInPackageChannelReturns(result1 *registry.Bundle, result2 error) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = nil
	fake.getBundleInPackageChannelReturns = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetBundleInPackageChannelReturnsOnCall(i int, result1 *registry.Bundle, result2 error) {
	fake.getBundleInPackageChannelMutex.Lock()
	defer fake.getBundleInPackageChannelMutex.Unlock()
	fake.GetBundleInPackageChannelStub = nil
	if fake.getBundleInPackageChannelReturnsOnCall == nil {
		fake.getBundleInPackageChannelReturnsOnCall = make(map[int]struct {
			result1 *registry.Bundle
			result2 error
		})
	}
	fake.getBundleInPackageChannelReturnsOnCall[i] = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetBundleThatProvides(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*registry.Bundle, error) {
	fake.getBundleThatProvidesMutex.Lock()
	ret, specificReturn := fake.getBundleThatProvidesReturnsOnCall[len(fake.getBundleThatProvidesArgsForCall)]
	fake.getBundleThatProvidesArgsForCall = append(fake.getBundleThatProvidesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetBundleThatProvides", []interface{}{arg1, arg2, arg3, arg4})
	fake.getBundleThatProvidesMutex.Unlock()
	if fake.GetBundleThatProvidesStub != nil {
		return fake.GetBundleThatProvidesStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getBundleThatProvidesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) GetBundleThatProvidesCallCount() int {
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	return len(fake.getBundleThatProvidesArgsForCall)
}

func (fake *FakeInterface) GetBundleThatProvidesCalls(stub func(context.Context, string, string, string) (*registry.Bundle, error)) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = stub
}

func (fake *FakeInterface) GetBundleThatProvidesArgsForCall(i int) (context.Context, string, string, string) {
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	argsForCall := fake.getBundleThatProvidesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterface) GetBundleThatProvidesReturns(result1 *registry.Bundle, result2 error) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = nil
	fake.getBundleThatProvidesReturns = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetBundleThatProvidesReturnsOnCall(i int, result1 *registry.Bundle, result2 error) {
	fake.getBundleThatProvidesMutex.Lock()
	defer fake.getBundleThatProvidesMutex.Unlock()
	fake.GetBundleThatProvidesStub = nil
	if fake.getBundleThatProvidesReturnsOnCall == nil {
		fake.getBundleThatProvidesReturnsOnCall = make(map[int]struct {
			result1 *registry.Bundle
			result2 error
		})
	}
	fake.getBundleThatProvidesReturnsOnCall[i] = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannel(arg1 context.Context, arg2 string, arg3 string, arg4 string) (*registry.Bundle, error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	ret, specificReturn := fake.getReplacementBundleInPackageChannelReturnsOnCall[len(fake.getReplacementBundleInPackageChannelArgsForCall)]
	fake.getReplacementBundleInPackageChannelArgsForCall = append(fake.getReplacementBundleInPackageChannelArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetReplacementBundleInPackageChannel", []interface{}{arg1, arg2, arg3, arg4})
	fake.getReplacementBundleInPackageChannelMutex.Unlock()
	if fake.GetReplacementBundleInPackageChannelStub != nil {
		return fake.GetReplacementBundleInPackageChannelStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReplacementBundleInPackageChannelReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannelCallCount() int {
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	return len(fake.getReplacementBundleInPackageChannelArgsForCall)
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannelCalls(stub func(context.Context, string, string, string) (*registry.Bundle, error)) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = stub
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannelArgsForCall(i int) (context.Context, string, string, string) {
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	argsForCall := fake.getReplacementBundleInPackageChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannelReturns(result1 *registry.Bundle, result2 error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = nil
	fake.getReplacementBundleInPackageChannelReturns = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetReplacementBundleInPackageChannelReturnsOnCall(i int, result1 *registry.Bundle, result2 error) {
	fake.getReplacementBundleInPackageChannelMutex.Lock()
	defer fake.getReplacementBundleInPackageChannelMutex.Unlock()
	fake.GetReplacementBundleInPackageChannelStub = nil
	if fake.getReplacementBundleInPackageChannelReturnsOnCall == nil {
		fake.getReplacementBundleInPackageChannelReturnsOnCall = make(map[int]struct {
			result1 *registry.Bundle
			result2 error
		})
	}
	fake.getReplacementBundleInPackageChannelReturnsOnCall[i] = struct {
		result1 *registry.Bundle
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBundleMutex.RLock()
	defer fake.getBundleMutex.RUnlock()
	fake.getBundleInPackageChannelMutex.RLock()
	defer fake.getBundleInPackageChannelMutex.RUnlock()
	fake.getBundleThatProvidesMutex.RLock()
	defer fake.getBundleThatProvidesMutex.RUnlock()
	fake.getReplacementBundleInPackageChannelMutex.RLock()
	defer fake.getReplacementBundleInPackageChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterface) recordInvocation(key string, args []interface{}) {
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

var _ client.Interface = new(FakeInterface)