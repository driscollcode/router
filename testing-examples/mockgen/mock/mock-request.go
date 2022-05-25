// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/driscollcode/router/example_unit_tests/mockgen/mock (interfaces: Request)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	router "github.com/driscollcode/router"
	gomock "github.com/golang/mock/gomock"
)

// MockRequest is a mock of Request interface.
type MockRequest struct {
	ctrl     *gomock.Controller
	recorder *MockRequestMockRecorder
}

// MockRequestMockRecorder is the mock recorder for MockRequest.
type MockRequestMockRecorder struct {
	mock *MockRequest
}

// NewMockRequest creates a new mock instance.
func NewMockRequest(ctrl *gomock.Controller) *MockRequest {
	mock := &MockRequest{ctrl: ctrl}
	mock.recorder = &MockRequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRequest) EXPECT() *MockRequestMockRecorder {
	return m.recorder
}

// ArgExists mocks base method.
func (m *MockRequest) ArgExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArgExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ArgExists indicates an expected call of ArgExists.
func (mr *MockRequestMockRecorder) ArgExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArgExists", reflect.TypeOf((*MockRequest)(nil).ArgExists), arg0)
}

// Body mocks base method.
func (m *MockRequest) Body() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockRequestMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockRequest)(nil).Body))
}

// BodyError mocks base method.
func (m *MockRequest) BodyError() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BodyError")
	ret0, _ := ret[0].(error)
	return ret0
}

// BodyError indicates an expected call of BodyError.
func (mr *MockRequestMockRecorder) BodyError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BodyError", reflect.TypeOf((*MockRequest)(nil).BodyError))
}

// Error mocks base method.
func (m *MockRequest) Error(arg0 ...interface{}) router.Response {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Error", varargs...)
	ret0, _ := ret[0].(router.Response)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockRequestMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockRequest)(nil).Error), arg0...)
}

// GetArg mocks base method.
func (m *MockRequest) GetArg(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArg", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetArg indicates an expected call of GetArg.
func (mr *MockRequestMockRecorder) GetArg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArg", reflect.TypeOf((*MockRequest)(nil).GetArg), arg0)
}

// GetHeader mocks base method.
func (m *MockRequest) GetHeader(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockRequestMockRecorder) GetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockRequest)(nil).GetHeader), arg0)
}

// GetHeaders mocks base method.
func (m *MockRequest) GetHeaders() map[string][]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeaders")
	ret0, _ := ret[0].(map[string][]string)
	return ret0
}

// GetHeaders indicates an expected call of GetHeaders.
func (mr *MockRequestMockRecorder) GetHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeaders", reflect.TypeOf((*MockRequest)(nil).GetHeaders))
}

// GetHost mocks base method.
func (m *MockRequest) GetHost() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHost indicates an expected call of GetHost.
func (mr *MockRequestMockRecorder) GetHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHost", reflect.TypeOf((*MockRequest)(nil).GetHost))
}

// GetIP mocks base method.
func (m *MockRequest) GetIP() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIP")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetIP indicates an expected call of GetIP.
func (mr *MockRequestMockRecorder) GetIP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIP", reflect.TypeOf((*MockRequest)(nil).GetIP))
}

// GetPostVariable mocks base method.
func (m *MockRequest) GetPostVariable(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostVariable", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetPostVariable indicates an expected call of GetPostVariable.
func (mr *MockRequestMockRecorder) GetPostVariable(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostVariable", reflect.TypeOf((*MockRequest)(nil).GetPostVariable), arg0)
}

// GetReferer mocks base method.
func (m *MockRequest) GetReferer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReferer")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetReferer indicates an expected call of GetReferer.
func (mr *MockRequestMockRecorder) GetReferer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReferer", reflect.TypeOf((*MockRequest)(nil).GetReferer))
}

// GetResponseContent mocks base method.
func (m *MockRequest) GetResponseContent() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseContent")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetResponseContent indicates an expected call of GetResponseContent.
func (mr *MockRequestMockRecorder) GetResponseContent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseContent", reflect.TypeOf((*MockRequest)(nil).GetResponseContent))
}

// GetResponseHeaders mocks base method.
func (m *MockRequest) GetResponseHeaders() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseHeaders")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetResponseHeaders indicates an expected call of GetResponseHeaders.
func (mr *MockRequestMockRecorder) GetResponseHeaders() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseHeaders", reflect.TypeOf((*MockRequest)(nil).GetResponseHeaders))
}

// GetResponseRedirect mocks base method.
func (m *MockRequest) GetResponseRedirect() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseRedirect")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetResponseRedirect indicates an expected call of GetResponseRedirect.
func (mr *MockRequestMockRecorder) GetResponseRedirect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseRedirect", reflect.TypeOf((*MockRequest)(nil).GetResponseRedirect))
}

// GetResponseStatusCode mocks base method.
func (m *MockRequest) GetResponseStatusCode() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseStatusCode")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetResponseStatusCode indicates an expected call of GetResponseStatusCode.
func (mr *MockRequestMockRecorder) GetResponseStatusCode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseStatusCode", reflect.TypeOf((*MockRequest)(nil).GetResponseStatusCode))
}

// GetURL mocks base method.
func (m *MockRequest) GetURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetURL indicates an expected call of GetURL.
func (mr *MockRequestMockRecorder) GetURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockRequest)(nil).GetURL))
}

// GetUserAgent mocks base method.
func (m *MockRequest) GetUserAgent() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAgent")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUserAgent indicates an expected call of GetUserAgent.
func (mr *MockRequestMockRecorder) GetUserAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAgent", reflect.TypeOf((*MockRequest)(nil).GetUserAgent))
}

// HasBody mocks base method.
func (m *MockRequest) HasBody() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasBody")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasBody indicates an expected call of HasBody.
func (mr *MockRequestMockRecorder) HasBody() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasBody", reflect.TypeOf((*MockRequest)(nil).HasBody))
}

// HeaderExists mocks base method.
func (m *MockRequest) HeaderExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeaderExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HeaderExists indicates an expected call of HeaderExists.
func (mr *MockRequestMockRecorder) HeaderExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeaderExists", reflect.TypeOf((*MockRequest)(nil).HeaderExists), arg0)
}

// PermanentRedirect mocks base method.
func (m *MockRequest) PermanentRedirect(arg0 string) router.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PermanentRedirect", arg0)
	ret0, _ := ret[0].(router.Response)
	return ret0
}

// PermanentRedirect indicates an expected call of PermanentRedirect.
func (mr *MockRequestMockRecorder) PermanentRedirect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PermanentRedirect", reflect.TypeOf((*MockRequest)(nil).PermanentRedirect), arg0)
}

// PostVariableExists mocks base method.
func (m *MockRequest) PostVariableExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostVariableExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// PostVariableExists indicates an expected call of PostVariableExists.
func (mr *MockRequestMockRecorder) PostVariableExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostVariableExists", reflect.TypeOf((*MockRequest)(nil).PostVariableExists), arg0)
}

// Redirect mocks base method.
func (m *MockRequest) Redirect(arg0 string) router.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Redirect", arg0)
	ret0, _ := ret[0].(router.Response)
	return ret0
}

// Redirect indicates an expected call of Redirect.
func (mr *MockRequestMockRecorder) Redirect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Redirect", reflect.TypeOf((*MockRequest)(nil).Redirect), arg0)
}

// Response mocks base method.
func (m *MockRequest) Response(arg0 ...interface{}) router.Response {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Response", varargs...)
	ret0, _ := ret[0].(router.Response)
	return ret0
}

// Response indicates an expected call of Response.
func (mr *MockRequestMockRecorder) Response(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Response", reflect.TypeOf((*MockRequest)(nil).Response), arg0...)
}

// SetHeader mocks base method.
func (m *MockRequest) SetHeader(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHeader", arg0, arg1)
}

// SetHeader indicates an expected call of SetHeader.
func (mr *MockRequestMockRecorder) SetHeader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockRequest)(nil).SetHeader), arg0, arg1)
}

// Success mocks base method.
func (m *MockRequest) Success(arg0 ...interface{}) router.Response {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Success", varargs...)
	ret0, _ := ret[0].(router.Response)
	return ret0
}

// Success indicates an expected call of Success.
func (mr *MockRequestMockRecorder) Success(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Success", reflect.TypeOf((*MockRequest)(nil).Success), arg0...)
}