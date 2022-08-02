// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/envelopeprocessor"
	"code.cloudfoundry.org/app-autoscaler/src/autoscaler/models"
	"code.cloudfoundry.org/go-loggregator/v8/rpc/loggregator_v2"
)

type FakeEnvelopeProcessor struct {
	GetGaugeMetricsStub        func([]*loggregator_v2.Envelope, int64) ([]models.AppInstanceMetric, error)
	getGaugeMetricsMutex       sync.RWMutex
	getGaugeMetricsArgsForCall []struct {
		arg1 []*loggregator_v2.Envelope
		arg2 int64
	}
	getGaugeMetricsReturns struct {
		result1 []models.AppInstanceMetric
		result2 error
	}
	getGaugeMetricsReturnsOnCall map[int]struct {
		result1 []models.AppInstanceMetric
		result2 error
	}
	GetTimerMetricsStub        func([]*loggregator_v2.Envelope, string, int64) []models.AppInstanceMetric
	getTimerMetricsMutex       sync.RWMutex
	getTimerMetricsArgsForCall []struct {
		arg1 []*loggregator_v2.Envelope
		arg2 string
		arg3 int64
	}
	getTimerMetricsReturns struct {
		result1 []models.AppInstanceMetric
	}
	getTimerMetricsReturnsOnCall map[int]struct {
		result1 []models.AppInstanceMetric
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetrics(arg1 []*loggregator_v2.Envelope, arg2 int64) ([]models.AppInstanceMetric, error) {
	var arg1Copy []*loggregator_v2.Envelope
	if arg1 != nil {
		arg1Copy = make([]*loggregator_v2.Envelope, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getGaugeMetricsMutex.Lock()
	ret, specificReturn := fake.getGaugeMetricsReturnsOnCall[len(fake.getGaugeMetricsArgsForCall)]
	fake.getGaugeMetricsArgsForCall = append(fake.getGaugeMetricsArgsForCall, struct {
		arg1 []*loggregator_v2.Envelope
		arg2 int64
	}{arg1Copy, arg2})
	stub := fake.GetGaugeMetricsStub
	fakeReturns := fake.getGaugeMetricsReturns
	fake.recordInvocation("GetGaugeMetrics", []interface{}{arg1Copy, arg2})
	fake.getGaugeMetricsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetricsCallCount() int {
	fake.getGaugeMetricsMutex.RLock()
	defer fake.getGaugeMetricsMutex.RUnlock()
	return len(fake.getGaugeMetricsArgsForCall)
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetricsCalls(stub func([]*loggregator_v2.Envelope, int64) ([]models.AppInstanceMetric, error)) {
	fake.getGaugeMetricsMutex.Lock()
	defer fake.getGaugeMetricsMutex.Unlock()
	fake.GetGaugeMetricsStub = stub
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetricsArgsForCall(i int) ([]*loggregator_v2.Envelope, int64) {
	fake.getGaugeMetricsMutex.RLock()
	defer fake.getGaugeMetricsMutex.RUnlock()
	argsForCall := fake.getGaugeMetricsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetricsReturns(result1 []models.AppInstanceMetric, result2 error) {
	fake.getGaugeMetricsMutex.Lock()
	defer fake.getGaugeMetricsMutex.Unlock()
	fake.GetGaugeMetricsStub = nil
	fake.getGaugeMetricsReturns = struct {
		result1 []models.AppInstanceMetric
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvelopeProcessor) GetGaugeMetricsReturnsOnCall(i int, result1 []models.AppInstanceMetric, result2 error) {
	fake.getGaugeMetricsMutex.Lock()
	defer fake.getGaugeMetricsMutex.Unlock()
	fake.GetGaugeMetricsStub = nil
	if fake.getGaugeMetricsReturnsOnCall == nil {
		fake.getGaugeMetricsReturnsOnCall = make(map[int]struct {
			result1 []models.AppInstanceMetric
			result2 error
		})
	}
	fake.getGaugeMetricsReturnsOnCall[i] = struct {
		result1 []models.AppInstanceMetric
		result2 error
	}{result1, result2}
}

func (fake *FakeEnvelopeProcessor) GetTimerMetrics(arg1 []*loggregator_v2.Envelope, arg2 string, arg3 int64) []models.AppInstanceMetric {
	var arg1Copy []*loggregator_v2.Envelope
	if arg1 != nil {
		arg1Copy = make([]*loggregator_v2.Envelope, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getTimerMetricsMutex.Lock()
	ret, specificReturn := fake.getTimerMetricsReturnsOnCall[len(fake.getTimerMetricsArgsForCall)]
	fake.getTimerMetricsArgsForCall = append(fake.getTimerMetricsArgsForCall, struct {
		arg1 []*loggregator_v2.Envelope
		arg2 string
		arg3 int64
	}{arg1Copy, arg2, arg3})
	stub := fake.GetTimerMetricsStub
	fakeReturns := fake.getTimerMetricsReturns
	fake.recordInvocation("GetTimerMetrics", []interface{}{arg1Copy, arg2, arg3})
	fake.getTimerMetricsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEnvelopeProcessor) GetTimerMetricsCallCount() int {
	fake.getTimerMetricsMutex.RLock()
	defer fake.getTimerMetricsMutex.RUnlock()
	return len(fake.getTimerMetricsArgsForCall)
}

func (fake *FakeEnvelopeProcessor) GetTimerMetricsCalls(stub func([]*loggregator_v2.Envelope, string, int64) []models.AppInstanceMetric) {
	fake.getTimerMetricsMutex.Lock()
	defer fake.getTimerMetricsMutex.Unlock()
	fake.GetTimerMetricsStub = stub
}

func (fake *FakeEnvelopeProcessor) GetTimerMetricsArgsForCall(i int) ([]*loggregator_v2.Envelope, string, int64) {
	fake.getTimerMetricsMutex.RLock()
	defer fake.getTimerMetricsMutex.RUnlock()
	argsForCall := fake.getTimerMetricsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeEnvelopeProcessor) GetTimerMetricsReturns(result1 []models.AppInstanceMetric) {
	fake.getTimerMetricsMutex.Lock()
	defer fake.getTimerMetricsMutex.Unlock()
	fake.GetTimerMetricsStub = nil
	fake.getTimerMetricsReturns = struct {
		result1 []models.AppInstanceMetric
	}{result1}
}

func (fake *FakeEnvelopeProcessor) GetTimerMetricsReturnsOnCall(i int, result1 []models.AppInstanceMetric) {
	fake.getTimerMetricsMutex.Lock()
	defer fake.getTimerMetricsMutex.Unlock()
	fake.GetTimerMetricsStub = nil
	if fake.getTimerMetricsReturnsOnCall == nil {
		fake.getTimerMetricsReturnsOnCall = make(map[int]struct {
			result1 []models.AppInstanceMetric
		})
	}
	fake.getTimerMetricsReturnsOnCall[i] = struct {
		result1 []models.AppInstanceMetric
	}{result1}
}

func (fake *FakeEnvelopeProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getGaugeMetricsMutex.RLock()
	defer fake.getGaugeMetricsMutex.RUnlock()
	fake.getTimerMetricsMutex.RLock()
	defer fake.getTimerMetricsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnvelopeProcessor) recordInvocation(key string, args []interface{}) {
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

var _ envelopeprocessor.EnvelopeProcessor = new(FakeEnvelopeProcessor)