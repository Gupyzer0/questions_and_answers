// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/services/services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "leonel/prototype_b/pkg/db/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuestionsAndAnswers is a mock of QuestionsAndAnswers interface.
type MockQuestionsAndAnswers struct {
	ctrl     *gomock.Controller
	recorder *MockQuestionsAndAnswersMockRecorder
}

// MockQuestionsAndAnswersMockRecorder is the mock recorder for MockQuestionsAndAnswers.
type MockQuestionsAndAnswersMockRecorder struct {
	mock *MockQuestionsAndAnswers
}

// NewMockQuestionsAndAnswers creates a new mock instance.
func NewMockQuestionsAndAnswers(ctrl *gomock.Controller) *MockQuestionsAndAnswers {
	mock := &MockQuestionsAndAnswers{ctrl: ctrl}
	mock.recorder = &MockQuestionsAndAnswersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuestionsAndAnswers) EXPECT() *MockQuestionsAndAnswersMockRecorder {
	return m.recorder
}

// CreateQuestion mocks base method.
func (m *MockQuestionsAndAnswers) CreateQuestion(arg0, arg1, arg2 string) (*models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateQuestion", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateQuestion indicates an expected call of CreateQuestion.
func (mr *MockQuestionsAndAnswersMockRecorder) CreateQuestion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateQuestion", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).CreateQuestion), arg0, arg1, arg2)
}

// DeleteQuestion mocks base method.
func (m *MockQuestionsAndAnswers) DeleteQuestion(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteQuestion", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteQuestion indicates an expected call of DeleteQuestion.
func (mr *MockQuestionsAndAnswersMockRecorder) DeleteQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteQuestion", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).DeleteQuestion), arg0)
}

// GetQuestion mocks base method.
func (m *MockQuestionsAndAnswers) GetQuestion(arg0 string) (*models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestion", arg0)
	ret0, _ := ret[0].(*models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestion indicates an expected call of GetQuestion.
func (mr *MockQuestionsAndAnswersMockRecorder) GetQuestion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestion", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).GetQuestion), arg0)
}

// GetQuestions mocks base method.
func (m *MockQuestionsAndAnswers) GetQuestions() ([]models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestions")
	ret0, _ := ret[0].([]models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestions indicates an expected call of GetQuestions.
func (mr *MockQuestionsAndAnswersMockRecorder) GetQuestions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestions", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).GetQuestions))
}

// GetQuestionsByUser mocks base method.
func (m *MockQuestionsAndAnswers) GetQuestionsByUser(arg0 string) ([]models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQuestionsByUser", arg0)
	ret0, _ := ret[0].([]models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQuestionsByUser indicates an expected call of GetQuestionsByUser.
func (mr *MockQuestionsAndAnswersMockRecorder) GetQuestionsByUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQuestionsByUser", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).GetQuestionsByUser), arg0)
}

// GetUsers mocks base method.
func (m *MockQuestionsAndAnswers) GetUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockQuestionsAndAnswersMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).GetUsers))
}

// UpdateAnswer mocks base method.
func (m *MockQuestionsAndAnswers) UpdateAnswer(arg0, arg1, arg2 string) (*models.Answer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAnswer", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Answer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAnswer indicates an expected call of UpdateAnswer.
func (mr *MockQuestionsAndAnswersMockRecorder) UpdateAnswer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAnswer", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).UpdateAnswer), arg0, arg1, arg2)
}

// UpdateQuestion mocks base method.
func (m *MockQuestionsAndAnswers) UpdateQuestion(arg0, arg1, arg2 string) (*models.Question, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateQuestion", arg0, arg1, arg2)
	ret0, _ := ret[0].(*models.Question)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateQuestion indicates an expected call of UpdateQuestion.
func (mr *MockQuestionsAndAnswersMockRecorder) UpdateQuestion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateQuestion", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).UpdateQuestion), arg0, arg1, arg2)
}

// Validate mocks base method.
func (m *MockQuestionsAndAnswers) Validate(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockQuestionsAndAnswersMockRecorder) Validate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockQuestionsAndAnswers)(nil).Validate), arg0)
}
