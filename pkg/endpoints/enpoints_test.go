package endpoints

import (
	"errors"
	"testing"

	"context"

	"github.com/golang/mock/gomock"
	"gopkg.in/guregu/null.v4"

	"leonel/prototype_b/mocks/mock_services"
	"leonel/prototype_b/pkg/db/models"
	"leonel/prototype_b/pkg/utils"
)

func TestMakeGetQuestionsEndpoint(t *testing.T){

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	testname := "Get questions return values"
	req := GetQuestionsRequest{}

	t.Run(testname, func(t *testing.T){

		srv.EXPECT().GetQuestions().Return([]models.Question{},nil)
		endpoint := MakeGetQuestionsEndpoint(srv)
		_, err := endpoint(context.Background(), req)
		
		if err != nil {
			t.Errorf("Error not nil: %v",err.Error())
		}
	})

	testname = "Get questions returns empty response"

	t.Run(testname, func(t *testing.T) {
		srv.EXPECT().GetQuestions().Return(nil,utils.ErrNotFound)
		endpoint := MakeGetQuestionsEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err == nil {
			t.Errorf("Error not received on nil response")
		}
	})
}

func TestMakeGetQuestionEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	getQuestionRequest := GetQuestionRequest{
		Question_id: "1",
	}

	testname := "Get question by ID"

	t.Run(testname, func(t *testing.T) {
		srv.EXPECT().GetQuestion("1").Return(&models.Question{}, nil)

		endpoint := MakeGetQuestionEndpoint(srv)

		_, err := endpoint(context.Background(), getQuestionRequest)

		if err != nil {
			t.Errorf("Error received on response: %v",err)
		}

	})

	testname = "Get question by id that is not available"

	t.Run(testname, func(t *testing.T) {
		srv.EXPECT().GetQuestion("2").Return(nil, utils.ErrNotFound)

		getQuestionRequest := GetQuestionRequest{
			Question_id: "2",
		}

		endpoint := MakeGetQuestionEndpoint(srv)
		_, err := endpoint(context.Background(), getQuestionRequest)

		if err != utils.ErrNotFound {
			t.Errorf("Error NotFound not received in response")
		}
	})
}

func TestMakeCreateQustionEndpoint(t *testing.T) {
	
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	req := CreateQuestionRequest{
		Title : "This a title",
		Statement : "This is a statement",
		User_id : "1",
	}

	testname := "Create question"

	t.Run(testname, func(t *testing.T) {

		srv.EXPECT().CreateQuestion(req.User_id, req.Title, req.Statement).Return(&models.Question{
			ID: "1",
			Title: req.Title,
			Statement: req.Statement,
			User : models.User{
				ID: null.StringFrom("1"),
				Username: null.StringFrom("Leonel"),
			},
		}, nil)

		srv.EXPECT().Validate(gomock.Any()).Return(nil)

		endpoint := MakeCreateQuestionsEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != nil {
			t.Errorf("Error received while creating a question: %v",err)
		}
	})

	testname = "Create question validation fail test"

	t.Run(testname, func(t *testing.T) {
		srv.EXPECT().Validate(gomock.Any()).Return(utils.CustomValidationErrors{})
	
		endpoint := MakeCreateQuestionsEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		_, ok := err.(utils.CustomValidationErrors)

		if !ok {
			t.Errorf("Error is not a validation error type: %v", err)
		}
	})

	testname = "Create question failed"

	dummy_error := errors.New("dummy_error")

	t.Run(testname, func(t *testing.T) {

		srv.EXPECT().Validate(gomock.Any()).Return(nil)
		srv.EXPECT().CreateQuestion(req.User_id, req.Title, req.Statement).Return(nil, dummy_error)

		endpoint := MakeCreateQuestionsEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != dummy_error {
			t.Errorf("Didnt returned expected error")
		}
	})
}

func TestMakeGetQuestionsByUserEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	testname := "Get questions by user"

	req := GetQuestionsByUserRequest{
		User_id: "1",
	}

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().GetQuestionsByUser(req.User_id).Return([]models.Question{},nil)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeGetQuestionsByUserEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != nil {
			t.Errorf("Get questios by user returned error: %v",err)
		}

	})

	testname = "Get questions by user validation fail"

	t.Run(testname, func(t *testing.T){

		srv.EXPECT().Validate(req).Return(utils.CustomValidationErrors{})

		endpoint := MakeGetQuestionsByUserEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		_, ok := err.(utils.CustomValidationErrors)

		if !ok {
			t.Errorf("Error is not a validation error type: %v", err)
		}
	})

	testname = "Get questions by user query failed"

	dummy_error := errors.New("dummy_error")

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().GetQuestionsByUser(req.User_id).Return(nil,dummy_error)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeGetQuestionsByUserEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != dummy_error {
			t.Errorf("Get questions by user didn't returned expected error")
		}
	})
}

func TestMakeUpdateQuestionEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	req := UpdateQuestionRequest{
		Question_id: "1",
		Title: "dummy title",
		Statement: "dummy statement",
	}

	testname := "Update question"

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().UpdateQuestion(req.Question_id, req.Title, req.Statement).Return(&models.Question{}, nil)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeUpdateQuestionEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != nil{
			t.Errorf("Update questions returned error: %v",err)
		}
	})

	testname = "Update question failed creation"

	dummy_error := errors.New("dummy_error")

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().UpdateQuestion(req.Question_id, req.Title, req.Statement).Return(nil, dummy_error)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeUpdateQuestionEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != dummy_error {
			t.Errorf("Update questions didnt return error")
		}
	})

	testname = "Update question failed validation"

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().Validate(req).Return(utils.CustomValidationErrors{})

		endpoint := MakeUpdateQuestionEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		_, ok := err.(utils.CustomValidationErrors)

		if !ok {
			t.Errorf("Unexpected error type")
		}
	})
}

func TestMakeUpdateAnswerEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	req := UpdateAnswerRequest{
		Question_id: "1",
		Statement: "dummy statement",
		User_id: "1",
	}

	testname := "Update answer"

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().UpdateAnswer(req.Question_id, req.Statement, req.User_id).Return(&models.Answer{}, nil)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeUpdateAnswerEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != nil{
			t.Errorf("Update answer returned error: %v",err)
		}
	})

	testname = "Update answer failed creation"

	dummy_error := errors.New("dummy_error")

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().UpdateAnswer(req.Question_id, req.Statement, req.User_id).Return(nil, dummy_error)
		srv.EXPECT().Validate(req).Return(nil)

		endpoint := MakeUpdateAnswerEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err != dummy_error {
			t.Errorf("Update answer didnt return error")
		}
	})

	testname = "Update answer failed validation"

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().Validate(req).Return(utils.CustomValidationErrors{})

		endpoint := MakeUpdateAnswerEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		_, ok := err.(utils.CustomValidationErrors)

		if !ok {
			t.Errorf("Unexpected error type")
		}
	})
}

func TestMakeDeleteAnswerEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	req := DeleteQuestionRequest{
		Question_id: "1",
	}

	//dummy_error := errors.New("dummy_error")

	testname := "delete question"

	t.Run(testname, func(t *testing.T){
		srv.EXPECT().DeleteQuestion(req.Question_id).Return(nil)

		endpoint := MakeDeleteQuestionEndpoint(srv)
		_, err := endpoint(context.Background(),req)

		if err != nil {
			t.Errorf("Delete question failed, received error: %v", err)
		}
	})
}

func TestMakeGetUsersEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	srv := mock_services.NewMockQuestionsAndAnswers(mockCtrl)

	req := GetUsersRequest{}

	testname := "Get users"

	t.Run(testname, func(t *testing.T){

		srv.EXPECT().GetUsers().Return([]models.User{},nil)
		endpoint := MakeGetUsersEndpoint(srv)
		_, err := endpoint(context.Background(), req)
		
		if err != nil {
			t.Errorf("Error not nil: %v",err.Error())
		}
	})

	testname = "Get users returns empty response"

	t.Run(testname, func(t *testing.T) {
		srv.EXPECT().GetUsers().Return(nil,utils.ErrNotFound)
		endpoint := MakeGetUsersEndpoint(srv)
		_, err := endpoint(context.Background(), req)

		if err == nil {
			t.Errorf("Error not received on nil response")
		}
	})
}
