// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/metricsgateway/helpers"

	"code.cloudfoundry.org/go-loggregator/v9/rpc/loggregator_v2"
)

type FakeWSHelper struct {
	SetupConnStub        func() error
	setupConnMutex       sync.RWMutex
	setupConnArgsForCall []struct{}
	setupConnReturns     struct {
		result1 error
	}
	CloseConnStub        func() error
	closeConnMutex       sync.RWMutex
	closeConnArgsForCall []struct{}
	closeConnReturns     struct {
		result1 error
	}
	WriteStub        func(envelope *loggregator_v2.Envelope) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		envelope *loggregator_v2.Envelope
	}
	writeReturns struct {
		result1 error
	}
	ReadStub        func() error
	readMutex       sync.RWMutex
	readArgsForCall []struct{}
	readReturns     struct {
		result1 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWSHelper) SetupConn() error {
	fake.setupConnMutex.Lock()
	fake.setupConnArgsForCall = append(fake.setupConnArgsForCall, struct{}{})
	fake.recordInvocation("SetupConn", []interface{}{})
	fake.setupConnMutex.Unlock()
	if fake.SetupConnStub != nil {
		return fake.SetupConnStub()
	}
	return fake.setupConnReturns.result1
}

func (fake *FakeWSHelper) SetupConnCallCount() int {
	fake.setupConnMutex.RLock()
	defer fake.setupConnMutex.RUnlock()
	return len(fake.setupConnArgsForCall)
}

func (fake *FakeWSHelper) SetupConnReturns(result1 error) {
	fake.SetupConnStub = nil
	fake.setupConnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWSHelper) CloseConn() error {
	fake.closeConnMutex.Lock()
	fake.closeConnArgsForCall = append(fake.closeConnArgsForCall, struct{}{})
	fake.recordInvocation("CloseConn", []interface{}{})
	fake.closeConnMutex.Unlock()
	if fake.CloseConnStub != nil {
		return fake.CloseConnStub()
	}
	return fake.closeConnReturns.result1
}

func (fake *FakeWSHelper) CloseConnCallCount() int {
	fake.closeConnMutex.RLock()
	defer fake.closeConnMutex.RUnlock()
	return len(fake.closeConnArgsForCall)
}

func (fake *FakeWSHelper) CloseConnReturns(result1 error) {
	fake.CloseConnStub = nil
	fake.closeConnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWSHelper) Write(envelope *loggregator_v2.Envelope) error {
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		envelope *loggregator_v2.Envelope
	}{envelope})
	fake.recordInvocation("Write", []interface{}{envelope})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(envelope)
	}
	return fake.writeReturns.result1
}

func (fake *FakeWSHelper) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeWSHelper) WriteArgsForCall(i int) *loggregator_v2.Envelope {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].envelope
}

func (fake *FakeWSHelper) WriteReturns(result1 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWSHelper) Read() error {
	fake.readMutex.Lock()
	fake.readArgsForCall = append(fake.readArgsForCall, struct{}{})
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub()
	}
	return fake.readReturns.result1
}

func (fake *FakeWSHelper) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeWSHelper) ReadReturns(result1 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWSHelper) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	return fake.pingReturns.result1
}

func (fake *FakeWSHelper) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeWSHelper) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWSHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setupConnMutex.RLock()
	defer fake.setupConnMutex.RUnlock()
	fake.closeConnMutex.RLock()
	defer fake.closeConnMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWSHelper) recordInvocation(key string, args []interface{}) {
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

var _ helpers.WSHelper = new(FakeWSHelper)
