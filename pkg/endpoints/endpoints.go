package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	models "leonel/prototype_b/pkg/db/models"
	services "leonel/prototype_b/pkg/services"
)

//todo test what is inside, create endpoint and pass the data, use encode and decode

func MakeGetQuestionsEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){

		questions, err := srv.GetQuestions()

		if err != nil{
			return nil, err
		}

		return GetQuestionsResponse{Questions: questions}, nil
	}
}

func MakeGetQuestionEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(GetQuestionRequest)
			
		question, err := srv.GetQuestion(req.Question_id)

		if err != nil{
			return nil, err
		}

		return GetQuestionResponse{ Question: question },nil
	}
}

func MakeCreateQuestionsEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){

		req := request.(CreateQuestionRequest)

		err := srv.Validate(req)

		if err != nil{
			return nil, err
		}

		question, err := srv.CreateQuestion(req.User_id, req.Title, req.Statement)
		
		if err != nil{
			return nil, err
		}
		
		return CreateQuestionResponse{Question: *question}, nil
	}
}

func MakeGetQuestionsByUserEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){

		req := request.(GetQuestionsByUserRequest)

		err := srv.Validate(req)

		if err!= nil {
			return nil, err
		}

		questions, err := srv.GetQuestionsByUser(req.User_id)

		if err != nil{
			return nil,err
		}

		return GetQuestionsByUserResponse{Questions: questions }, nil
	}
}

func MakeUpdateQuestionEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)

		err := srv.Validate(req)

		if err != nil{
			return nil, err
		}

		question, err := srv.UpdateQuestion(req.Question_id, req.Title, req.Statement)

		if err != nil{
			return nil, err
		}

		return UpdateQuestionResponse{Question: question}, nil
	}
}

func MakeUpdateAnswerEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAnswerRequest)

		err := srv.Validate(req)

		if err != nil{
			return nil, err
		}

		answer, err := srv.UpdateAnswer(req.Question_id, req.Statement, req.User_id)
		
		if err != nil{
			return err, err
		}

		return UpdateAnswerResponse{Answer: answer}, nil

	}
}

func MakeDeleteQuestionEndpoint(srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteQuestionRequest)

		err := srv.DeleteQuestion(req.Question_id)

		return DeleteQuestionResponse{ Error: err }, err

	}
}

func MakeGetUsersEndpoint( srv services.QuestionsAndAnswers) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		users, err := srv.GetUsers()

		if err != nil {
			return nil, err
		}

		return GetUsersReponse{ Users: users }, nil
	}
}

//Requests

type GetQuestionsRequest struct{}

type GetQuestionRequest struct{
	Question_id string `validate:"required,max=10"`
}

type GetQuestionsByUserRequest struct{
	User_id string `validate:"required,max=10"`
}

type CreateQuestionRequest struct{
	Title     string `validate:"required,max=100"`
	Statement string `validate:"required,max=255"`

	User_id string `validate:"required,max=10"`
}

type UpdateQuestionRequest struct{
	Question_id string `validate:"required,max=10"`
	Title string `validate:"required,max=100"`
	Statement string `validate:"required,max=255"`
}

type UpdateAnswerRequest struct{
	Question_id string `validate:"required,max=10"`
	Statement string `validate:"required,max=255"`
	User_id string `validate:"required,max=10"`
}

type DeleteQuestionRequest struct{
	Question_id string `validate:"required,max=10"`
}

type GetUsersRequest struct {}

//Responses
type GetQuestionResponse struct{
	Question *models.Question `json:"question"`
	Err error `json:"error,omitempty"`
}

type GetQuestionsResponse struct{
	Questions []models.Question `json:"questions"`
}

type GetQuestionsByUserResponse struct{
	Questions []models.Question `json:"questions"`
	Err error `json:"error,omitempty"`
}

type CreateQuestionResponse struct{
	Question models.Question `json:"question"`
}

type UpdateQuestionResponse struct{
	Question *models.Question `json:"question"`
	Err error `json:"error,omitempty"`
}

type UpdateAnswerResponse struct{
	Answer *models.Answer `json:"answer"`
	Err error `json:"error,omitempty"`
}

type DeleteQuestionResponse struct{
	Error error `json:"error,omitempty"`
}

type GetUsersReponse struct {
	Users []models.User `json:"users"`
}
