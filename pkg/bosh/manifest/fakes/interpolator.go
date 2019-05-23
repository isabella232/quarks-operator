// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	manifest "code.cloudfoundry.org/cf-operator/pkg/bosh/manifest"
)

type FakeInterpolator struct {
	BuildOpsStub        func([]byte) error
	buildOpsMutex       sync.RWMutex
	buildOpsArgsForCall []struct {
		arg1 []byte
	}
	buildOpsReturns struct {
		result1 error
	}
	buildOpsReturnsOnCall map[int]struct {
		result1 error
	}
	InterpolateStub        func([]byte) ([]byte, error)
	interpolateMutex       sync.RWMutex
	interpolateArgsForCall []struct {
		arg1 []byte
	}
	interpolateReturns struct {
		result1 []byte
		result2 error
	}
	interpolateReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterpolator) BuildOps(arg1 []byte) error {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.buildOpsMutex.Lock()
	ret, specificReturn := fake.buildOpsReturnsOnCall[len(fake.buildOpsArgsForCall)]
	fake.buildOpsArgsForCall = append(fake.buildOpsArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("BuildOps", []interface{}{arg1Copy})
	fake.buildOpsMutex.Unlock()
	if fake.BuildOpsStub != nil {
		return fake.BuildOpsStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.buildOpsReturns
	return fakeReturns.result1
}

func (fake *FakeInterpolator) BuildOpsCallCount() int {
	fake.buildOpsMutex.RLock()
	defer fake.buildOpsMutex.RUnlock()
	return len(fake.buildOpsArgsForCall)
}

func (fake *FakeInterpolator) BuildOpsCalls(stub func([]byte) error) {
	fake.buildOpsMutex.Lock()
	defer fake.buildOpsMutex.Unlock()
	fake.BuildOpsStub = stub
}

func (fake *FakeInterpolator) BuildOpsArgsForCall(i int) []byte {
	fake.buildOpsMutex.RLock()
	defer fake.buildOpsMutex.RUnlock()
	argsForCall := fake.buildOpsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInterpolator) BuildOpsReturns(result1 error) {
	fake.buildOpsMutex.Lock()
	defer fake.buildOpsMutex.Unlock()
	fake.BuildOpsStub = nil
	fake.buildOpsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterpolator) BuildOpsReturnsOnCall(i int, result1 error) {
	fake.buildOpsMutex.Lock()
	defer fake.buildOpsMutex.Unlock()
	fake.BuildOpsStub = nil
	if fake.buildOpsReturnsOnCall == nil {
		fake.buildOpsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.buildOpsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInterpolator) Interpolate(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.interpolateMutex.Lock()
	ret, specificReturn := fake.interpolateReturnsOnCall[len(fake.interpolateArgsForCall)]
	fake.interpolateArgsForCall = append(fake.interpolateArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Interpolate", []interface{}{arg1Copy})
	fake.interpolateMutex.Unlock()
	if fake.InterpolateStub != nil {
		return fake.InterpolateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.interpolateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterpolator) InterpolateCallCount() int {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	return len(fake.interpolateArgsForCall)
}

func (fake *FakeInterpolator) InterpolateCalls(stub func([]byte) ([]byte, error)) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = stub
}

func (fake *FakeInterpolator) InterpolateArgsForCall(i int) []byte {
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	argsForCall := fake.interpolateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInterpolator) InterpolateReturns(result1 []byte, result2 error) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = nil
	fake.interpolateReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpolator) InterpolateReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.interpolateMutex.Lock()
	defer fake.interpolateMutex.Unlock()
	fake.InterpolateStub = nil
	if fake.interpolateReturnsOnCall == nil {
		fake.interpolateReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.interpolateReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeInterpolator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildOpsMutex.RLock()
	defer fake.buildOpsMutex.RUnlock()
	fake.interpolateMutex.RLock()
	defer fake.interpolateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterpolator) recordInvocation(key string, args []interface{}) {
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

var _ manifest.Interpolator = new(FakeInterpolator)
