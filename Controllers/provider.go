package controllers

import (
	"C214-teoria-GO/database"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(dest interface{}) database.DBInterface {
	args := m.Called(dest)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) database.DBInterface {
	args := m.Called(dest)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) Create(value interface{}) database.DBInterface {
	args := m.Called(value)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) database.DBInterface {
	args := m.Called(value)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) Model(value interface{}) database.DBInterface {
	args := m.Called(value)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) UpdateColumns(value interface{}) database.DBInterface {
	args := m.Called(value)
	return args.Get(0).(database.DBInterface)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) database.DBInterface {
	m.Called(append([]interface{}{query}, args...)...)
	return m
}

type ErrorMockDatabase struct {
	err error
}

func (e *ErrorMockDatabase) Find(dest interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) First(dest interface{}, conds ...interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) Create(value interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) Delete(value interface{}, conds ...interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) Model(value interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) UpdateColumns(value interface{}) database.DBInterface {
	return e
}

func (e *ErrorMockDatabase) Where(value interface{}, args ...interface{}) database.DBInterface {
	return e
}
