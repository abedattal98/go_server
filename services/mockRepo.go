// Code generated by MockGen. DO NOT EDIT.
// Source: rgb/interfaces (interfaces: PostRepository)

// Package mock_interfaces is a generated GoMock package.
package services

import (
	reflect "reflect"
	models "rgb/models"

	gomock "github.com/golang/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// AddPost mocks base method.
func (m *MockPostRepository) AddPost(arg0 int, arg1 models.Post) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPost", arg0, arg1)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPost indicates an expected call of AddPost.
func (mr *MockPostRepositoryMockRecorder) AddPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPost", reflect.TypeOf((*MockPostRepository)(nil).AddPost), arg0, arg1)
}

// DeletePost mocks base method.
func (m *MockPostRepository) DeletePost(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRepositoryMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRepository)(nil).DeletePost), arg0)
}

// GetPostByID mocks base method.
func (m *MockPostRepository) GetPostByID(arg0 int) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", arg0)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockPostRepositoryMockRecorder) GetPostByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockPostRepository)(nil).GetPostByID), arg0)
}

// GetPostsByUserID mocks base method.
func (m *MockPostRepository) GetPostsByUserID(arg0 int) []models.Post {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUserID", arg0)
	ret0, _ := ret[0].([]models.Post)
	return ret0
}

// GetPostsByUserID indicates an expected call of GetPostsByUserID.
func (mr *MockPostRepositoryMockRecorder) GetPostsByUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUserID", reflect.TypeOf((*MockPostRepository)(nil).GetPostsByUserID), arg0)
}

// UpdatePost mocks base method.
func (m *MockPostRepository) UpdatePost(arg0 int, arg1 models.Post) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0, arg1)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockPostRepositoryMockRecorder) UpdatePost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockPostRepository)(nil).UpdatePost), arg0, arg1)
}