// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openservicemesh/osm/pkg/configurator (interfaces: Configurator)

// Package configurator is a generated GoMock package.
package configurator

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockConfigurator is a mock of Configurator interface
type MockConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockConfiguratorMockRecorder
}

// MockConfiguratorMockRecorder is the mock recorder for MockConfigurator
type MockConfiguratorMockRecorder struct {
	mock *MockConfigurator
}

// NewMockConfigurator creates a new mock instance
func NewMockConfigurator(ctrl *gomock.Controller) *MockConfigurator {
	mock := &MockConfigurator{ctrl: ctrl}
	mock.recorder = &MockConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConfigurator) EXPECT() *MockConfiguratorMockRecorder {
	return m.recorder
}

// GetConfigMap mocks base method
func (m *MockConfigurator) GetConfigMap() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap
func (mr *MockConfiguratorMockRecorder) GetConfigMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockConfigurator)(nil).GetConfigMap))
}

// GetEnvoyLogLevel mocks base method
func (m *MockConfigurator) GetEnvoyLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvoyLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEnvoyLogLevel indicates an expected call of GetEnvoyLogLevel
func (mr *MockConfiguratorMockRecorder) GetEnvoyLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvoyLogLevel", reflect.TypeOf((*MockConfigurator)(nil).GetEnvoyLogLevel))
}

// GetOSMNamespace mocks base method
func (m *MockConfigurator) GetOSMNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOSMNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetOSMNamespace indicates an expected call of GetOSMNamespace
func (mr *MockConfiguratorMockRecorder) GetOSMNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOSMNamespace", reflect.TypeOf((*MockConfigurator)(nil).GetOSMNamespace))
}

// GetOutboundIPRangeExclusionList mocks base method
func (m *MockConfigurator) GetOutboundIPRangeExclusionList() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutboundIPRangeExclusionList")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetOutboundIPRangeExclusionList indicates an expected call of GetOutboundIPRangeExclusionList
func (mr *MockConfiguratorMockRecorder) GetOutboundIPRangeExclusionList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutboundIPRangeExclusionList", reflect.TypeOf((*MockConfigurator)(nil).GetOutboundIPRangeExclusionList))
}

// GetServiceCertValidityPeriod mocks base method
func (m *MockConfigurator) GetServiceCertValidityPeriod() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceCertValidityPeriod")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetServiceCertValidityPeriod indicates an expected call of GetServiceCertValidityPeriod
func (mr *MockConfiguratorMockRecorder) GetServiceCertValidityPeriod() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceCertValidityPeriod", reflect.TypeOf((*MockConfigurator)(nil).GetServiceCertValidityPeriod))
}

// GetTracingEndpoint mocks base method
func (m *MockConfigurator) GetTracingEndpoint() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingEndpoint")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingEndpoint indicates an expected call of GetTracingEndpoint
func (mr *MockConfiguratorMockRecorder) GetTracingEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingEndpoint", reflect.TypeOf((*MockConfigurator)(nil).GetTracingEndpoint))
}

// GetTracingHost mocks base method
func (m *MockConfigurator) GetTracingHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTracingHost indicates an expected call of GetTracingHost
func (mr *MockConfiguratorMockRecorder) GetTracingHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingHost", reflect.TypeOf((*MockConfigurator)(nil).GetTracingHost))
}

// GetTracingPort mocks base method
func (m *MockConfigurator) GetTracingPort() uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTracingPort")
	ret0, _ := ret[0].(uint32)
	return ret0
}

// GetTracingPort indicates an expected call of GetTracingPort
func (mr *MockConfiguratorMockRecorder) GetTracingPort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTracingPort", reflect.TypeOf((*MockConfigurator)(nil).GetTracingPort))
}

// IsDebugServerEnabled mocks base method
func (m *MockConfigurator) IsDebugServerEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDebugServerEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDebugServerEnabled indicates an expected call of IsDebugServerEnabled
func (mr *MockConfiguratorMockRecorder) IsDebugServerEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDebugServerEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsDebugServerEnabled))
}

// IsEgressEnabled mocks base method
func (m *MockConfigurator) IsEgressEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEgressEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsEgressEnabled indicates an expected call of IsEgressEnabled
func (mr *MockConfiguratorMockRecorder) IsEgressEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEgressEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsEgressEnabled))
}

// IsPermissiveTrafficPolicyMode mocks base method
func (m *MockConfigurator) IsPermissiveTrafficPolicyMode() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPermissiveTrafficPolicyMode")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPermissiveTrafficPolicyMode indicates an expected call of IsPermissiveTrafficPolicyMode
func (mr *MockConfiguratorMockRecorder) IsPermissiveTrafficPolicyMode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPermissiveTrafficPolicyMode", reflect.TypeOf((*MockConfigurator)(nil).IsPermissiveTrafficPolicyMode))
}

// IsPrivilegedInitContainer mocks base method
func (m *MockConfigurator) IsPrivilegedInitContainer() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrivilegedInitContainer")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrivilegedInitContainer indicates an expected call of IsPrivilegedInitContainer
func (mr *MockConfiguratorMockRecorder) IsPrivilegedInitContainer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrivilegedInitContainer", reflect.TypeOf((*MockConfigurator)(nil).IsPrivilegedInitContainer))
}

// IsPrometheusScrapingEnabled mocks base method
func (m *MockConfigurator) IsPrometheusScrapingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPrometheusScrapingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsPrometheusScrapingEnabled indicates an expected call of IsPrometheusScrapingEnabled
func (mr *MockConfiguratorMockRecorder) IsPrometheusScrapingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPrometheusScrapingEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsPrometheusScrapingEnabled))
}

// IsTracingEnabled mocks base method
func (m *MockConfigurator) IsTracingEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTracingEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTracingEnabled indicates an expected call of IsTracingEnabled
func (mr *MockConfiguratorMockRecorder) IsTracingEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTracingEnabled", reflect.TypeOf((*MockConfigurator)(nil).IsTracingEnabled))
}

// UseHTTPSIngress mocks base method
func (m *MockConfigurator) UseHTTPSIngress() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseHTTPSIngress")
	ret0, _ := ret[0].(bool)
	return ret0
}

// UseHTTPSIngress indicates an expected call of UseHTTPSIngress
func (mr *MockConfiguratorMockRecorder) UseHTTPSIngress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseHTTPSIngress", reflect.TypeOf((*MockConfigurator)(nil).UseHTTPSIngress))
}
