// Code generated by MockGen. DO NOT EDIT.
// Source: entites/student.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	entites "github.com/piovani/aula/entites"
)

// MockStudentRepository is a mock of StudentRepository interface.
type MockStudentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStudentRepositoryMockRecorder
}

// MockStudentRepositoryMockRecorder is the mock recorder for MockStudentRepository.
type MockStudentRepositoryMockRecorder struct {
	mock *MockStudentRepository
}

// NewMockStudentRepository creates a new mock instance.
func NewMockStudentRepository(ctrl *gomock.Controller) *MockStudentRepository {
	mock := &MockStudentRepository{ctrl: ctrl}
	mock.recorder = &MockStudentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStudentRepository) EXPECT() *MockStudentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStudentRepository) Create(student *entites.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStudentRepositoryMockRecorder) Create(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStudentRepository)(nil).Create), student)
}

// Delete mocks base method.
func (m *MockStudentRepository) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStudentRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStudentRepository)(nil).Delete), id)
}

// FindByID mocks base method.
func (m *MockStudentRepository) FindByID(id uuid.UUID) (entites.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(entites.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockStudentRepositoryMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockStudentRepository)(nil).FindByID), id)
}

// List mocks base method.
func (m *MockStudentRepository) List() ([]entites.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]entites.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStudentRepositoryMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStudentRepository)(nil).List))
}

// Update mocks base method.
func (m *MockStudentRepository) Update(student *entites.Student) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", student)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStudentRepositoryMockRecorder) Update(student interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStudentRepository)(nil).Update), student)
}
