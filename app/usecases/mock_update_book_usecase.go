// Code generated by MockGen. DO NOT EDIT.
// Source: update_book_usecase.go

// Package usecases is a generated GoMock package.
package usecases

import (
	entity "library/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUpdateBookUseCase is a mock of UpdateBookUseCase interface.
type MockUpdateBookUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateBookUseCaseMockRecorder
}

// MockUpdateBookUseCaseMockRecorder is the mock recorder for MockUpdateBookUseCase.
type MockUpdateBookUseCaseMockRecorder struct {
	mock *MockUpdateBookUseCase
}

// NewMockUpdateBookUseCase creates a new mock instance.
func NewMockUpdateBookUseCase(ctrl *gomock.Controller) *MockUpdateBookUseCase {
	mock := &MockUpdateBookUseCase{ctrl: ctrl}
	mock.recorder = &MockUpdateBookUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateBookUseCase) EXPECT() *MockUpdateBookUseCaseMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockUpdateBookUseCase) Execute(arg0 *UpdateBookUseCaseInputDTO) (*UpdateBookUseCaseOutputDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(*UpdateBookUseCaseOutputDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockUpdateBookUseCaseMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockUpdateBookUseCase)(nil).Execute), arg0)
}

// MockUpdateBookUseCaseRepository is a mock of UpdateBookUseCaseRepository interface.
type MockUpdateBookUseCaseRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateBookUseCaseRepositoryMockRecorder
}

// MockUpdateBookUseCaseRepositoryMockRecorder is the mock recorder for MockUpdateBookUseCaseRepository.
type MockUpdateBookUseCaseRepositoryMockRecorder struct {
	mock *MockUpdateBookUseCaseRepository
}

// NewMockUpdateBookUseCaseRepository creates a new mock instance.
func NewMockUpdateBookUseCaseRepository(ctrl *gomock.Controller) *MockUpdateBookUseCaseRepository {
	mock := &MockUpdateBookUseCaseRepository{ctrl: ctrl}
	mock.recorder = &MockUpdateBookUseCaseRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateBookUseCaseRepository) EXPECT() *MockUpdateBookUseCaseRepositoryMockRecorder {
	return m.recorder
}

// Update mocks base method.
func (m *MockUpdateBookUseCaseRepository) Update(arg0 *entity.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUpdateBookUseCaseRepositoryMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUpdateBookUseCaseRepository)(nil).Update), arg0)
}
