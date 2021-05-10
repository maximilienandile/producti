// Code generated by MockGen. DO NOT EDIT.
// Source: internal/dynamo/clientInterface.go

// Package dynamo is a generated GoMock package.
package dynamo

import (
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAllByPK mocks base method.
func (m *MockClient) GetAllByPK(pkValue PartitionKey) ([]map[string]*dynamodb.AttributeValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByPK", pkValue)
	ret0, _ := ret[0].([]map[string]*dynamodb.AttributeValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByPK indicates an expected call of GetAllByPK.
func (mr *MockClientMockRecorder) GetAllByPK(pkValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByPK", reflect.TypeOf((*MockClient)(nil).GetAllByPK), pkValue)
}

// GetByKey mocks base method.
func (m *MockClient) GetByKey(partitionKeyValue PartitionKey, sortKeyValue string) (map[string]*dynamodb.AttributeValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", partitionKeyValue, sortKeyValue)
	ret0, _ := ret[0].(map[string]*dynamodb.AttributeValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey.
func (mr *MockClientMockRecorder) GetByKey(partitionKeyValue, sortKeyValue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockClient)(nil).GetByKey), partitionKeyValue, sortKeyValue)
}

// Put mocks base method.
func (m *MockClient) Put(req PutRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockClientMockRecorder) Put(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockClient)(nil).Put), req)
}