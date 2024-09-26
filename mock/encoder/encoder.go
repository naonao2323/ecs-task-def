// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/encoder/encoder.go
//
// Generated by this command:
//
//	mockgen -source=pkg/encoder/encoder.go -destination=mock/encoder/encoder.go
//

// Package mock_encoder is a generated GoMock package.
package mock_encoder

import (
	encoder "ecs-task-def-action/pkg/encoder"
	ecs "ecs-task-def-action/pkg/plovider/ecs"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockEncoder is a mock of Encoder interface.
type MockEncoder[P ecs.EcsTarget] struct {
	ctrl     *gomock.Controller
	recorder *MockEncoderMockRecorder[P]
}

// MockEncoderMockRecorder is the mock recorder for MockEncoder.
type MockEncoderMockRecorder[P ecs.EcsTarget] struct {
	mock *MockEncoder[P]
}

// NewMockEncoder creates a new mock instance.
func NewMockEncoder[P ecs.EcsTarget](ctrl *gomock.Controller) *MockEncoder[P] {
	mock := &MockEncoder[P]{ctrl: ctrl}
	mock.recorder = &MockEncoderMockRecorder[P]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEncoder[P]) EXPECT() *MockEncoderMockRecorder[P] {
	return m.recorder
}

// Encode mocks base method.
func (m *MockEncoder[P]) Encode(in []byte, format encoder.Format) (*P, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", in, format)
	ret0, _ := ret[0].(*P)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockEncoderMockRecorder[P]) Encode(in, format any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockEncoder[P])(nil).Encode), in, format)
}