package endpoints

import (
	"context"
	//"log"

	"github.com/go-kit/kit/endpoint"

	models "leonel/prototype_b/pkg/db/models"
	services "leonel/prototype_b/pkg/services"

)


func MakeGetQuestionsEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		////req := request.(getQuestionsRequest) -> not used
		questions := srv.GetQuestions()
		return GetQuestionsResponse{Questions: questions}, nil
	}
}

func MakeGetQuestionEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(GetQuestionRequest)
			
		question, err := srv.GetQuestion(req.Question_id)

		if err != nil{
			return nil, err
		}

		//return question, nil
		return GetQuestionResponse{ Question: question },nil
	}
}

func MakeCreateQuestionsEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){

		req := request.(CreateQuestionRequest)

		err := srv.Validate(req.Question)

		if err != nil{
			return nil, err
		}

		question, err := srv.CreateQuestion(req.Question)
		
		if err != nil{
			return nil, err
		}
		
		return CreateQuestionResponse{Question: *question}, nil
	}
}

func MakeGetQuestionsByUserEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error){

		req := request.(GetQuestionsByUserRequest)
		questions, err := srv.GetQuestionsByUser(req.User_id)

		if err != nil{
			return nil,err
		}

		return GetQuestionsByUserResponse{Questions: questions }, nil
	}
}

func MakeUpdateQuestionEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateQuestionRequest)

		question, err := srv.UpdateQuestion(req.Question_id, req.Title, req.Statement)

		if err != nil{
			return nil, err
		}

		return UpdateQuestionResponse{Question: question}, nil
	}
}

func MakeUpdateAnswerEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAnswerRequest)

		answer, err := srv.UpdateAnswer(req.Question_id, &req.Answer)
		
		if err != nil{
			return err, err
		}

		return UpdateAnswerResponse{Answer: answer}, nil

	}
}

func MakeDeleteAnswerEndpoint(srv services.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteQuestionRequest)

		err := srv.DeleteQuestion(req.Question_id)

		return DeleteQuestionResponse{ Error: err }, err

	}
}

//Requests

type GetQuestionRequest struct{
	Question_id string `validate:"required,max=10"`
}

type GetQuestionsRequest struct{}

type GetQuestionsByUserRequest struct{
	User_id string
}

type CreateQuestionRequest struct{
	Question *models.Question
}

type UpdateQuestionRequest struct{
	Question_id string 
	Title string
	Statement string
}

type UpdateAnswerRequest struct{
	Question_id string
	Answer models.Answer
}

type DeleteQuestionRequest struct{
	Question_id string
}

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
