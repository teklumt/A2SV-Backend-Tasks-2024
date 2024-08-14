// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "clean_architecture_Testing/domain"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: task
func (_m *TaskRepository) CreateTask(task domain.Task) (domain.Task, error) {
	ret := _m.Called(task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Task) (domain.Task, error)); ok {
		return rf(task)
	}
	if rf, ok := ret.Get(0).(func(domain.Task) domain.Task); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(domain.Task) error); ok {
		r1 = rf(task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTask provides a mock function with given fields: id
func (_m *TaskRepository) DeleteTask(id string) (domain.Task, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Task, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyTasks provides a mock function with given fields: usename
func (_m *TaskRepository) GetMyTasks(usename string) ([]domain.Task, error) {
	ret := _m.Called(usename)

	if len(ret) == 0 {
		panic("no return value specified for GetMyTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Task, error)); ok {
		return rf(usename)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Task); ok {
		r0 = rf(usename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(usename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: id, creater, Rol_
func (_m *TaskRepository) GetTaskByID(id string, creater string, Rol_ string) (domain.Task, error) {
	ret := _m.Called(id, creater, Rol_)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (domain.Task, error)); ok {
		return rf(id, creater, Rol_)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) domain.Task); ok {
		r0 = rf(id, creater, Rol_)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(id, creater, Rol_)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields:
func (_m *TaskRepository) GetTasks() ([]domain.Task, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTasks")
	}

	var r0 []domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.Task, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: id, task
func (_m *TaskRepository) UpdateTask(id string, task domain.Task) (domain.Task, error) {
	ret := _m.Called(id, task)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 domain.Task
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.Task) (domain.Task, error)); ok {
		return rf(id, task)
	}
	if rf, ok := ret.Get(0).(func(string, domain.Task) domain.Task); ok {
		r0 = rf(id, task)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(string, domain.Task) error); ok {
		r1 = rf(id, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}