package transport

import (
	"context"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	endpoints "leonel/prototype_b/pkg/endpoints"
	services "leonel/prototype_b/pkg/services"
	utils "leonel/prototype_b/pkg/utils"
)

//test definir new create question  request, crear contrusctor para cada uno

func MakeHttpHandler(srv services.QuestionsAndAnswers) *mux.Router{

	router := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	router.Methods("GET").Path("/questions").Handler(httptransport.NewServer(
		endpoints.MakeGetQuestionsEndpoint(srv),
		decodeGetQuestionsRequest,
		encodeResponse,
		options...,
	))

	router.Methods("POST").Path("/question").Handler(httptransport.NewServer(
		endpoints.MakeCreateQuestionsEndpoint(srv),
		decodeCreateQuestionRequest,
		encodeResponse,
		options...,
	))

	router.Methods("GET").Path("/question/{id}").Handler(httptransport.NewServer(
		endpoints.MakeGetQuestionEndpoint(srv),
		decodeGetQuestionRequest,
		encodeResponse,
		options...,
	))

	router.Methods("GET").Path("/user/{id}/questions").Handler(httptransport.NewServer(
		endpoints.MakeGetQuestionsByUserEndpoint(srv),
		decodeGetQuestionsByUserRequest,
		encodeResponse,
		options...,
	))

	router.Methods("PUT").Path("/question/{id}").Handler(httptransport.NewServer(
		endpoints.MakeUpdateQuestionEndpoint(srv),
		decodeUpdateQuestionRequest,
		encodeResponse,
		options...,
	))

	router.Methods("POST").Path("/question/{id}/answer").Handler(httptransport.NewServer(
		endpoints.MakeUpdateAnswerEndpoint(srv),
		decodeUpdateAnswerRequest,
		encodeResponse,
		options...,
	))

	router.Methods("DELETE").Path("/question/{id}").Handler(httptransport.NewServer(
		endpoints.MakeDeleteQuestionEndpoint(srv),
		decodeDeleteQuestionRequest,
		encodeResponse,
		options...,
	))

	router.Methods("GET").Path("/users").Handler(httptransport.NewServer(
		endpoints.MakeGetUsersEndpoint(srv),
		decodeGetUsersRequest,
		encodeResponse,
		options...,
	))

	return router
}

func decodeGetQuestionsRequest(_ context.Context, r *http.Request) (interface{}, error){
	return endpoints.GetQuestionsRequest{}, nil
}

func decodeCreateQuestionRequest(_ context.Context, r *http.Request) (interface{},error){

	var req endpoints.CreateQuestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, utils.ErrBadData
	}

	return req, nil
}

func decodeGetQuestionRequest(_ context.Context, r *http.Request) (interface{}, error){
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok || id == "" {
		return nil, utils.ErrBadData
	}
	
	return endpoints.GetQuestionRequest{ Question_id: id, }, nil
}

func decodeGetQuestionsByUserRequest(_ context.Context, r *http.Request) (interface{}, error){
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok  || id == "" {
		return nil, utils.ErrBadData
	}

	return endpoints.GetQuestionsByUserRequest{ User_id: id,}, nil
}

func decodeUpdateQuestionRequest(_ context.Context, r *http.Request)(interface{}, error){
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok || id == "" {
		return nil, utils.ErrBadData
	}

	var req endpoints.UpdateQuestionRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, utils.ErrBadData
	}

	return endpoints.UpdateQuestionRequest{ Question_id:id, Title:req.Title, Statement:req.Statement }, nil
}

func decodeUpdateAnswerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok || id == "" {
		return nil, utils.ErrBadData
	}

	var req endpoints.UpdateAnswerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, utils.ErrBadData
	}

	return endpoints.UpdateAnswerRequest{ Question_id: id, Statement: req.Statement, User_id: req.User_id }, nil
}

func decodeDeleteQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	id, ok := vars["id"]

	if !ok || id == "" {
		return nil, utils.ErrBadData
	}

	return endpoints.DeleteQuestionRequest{ Question_id: id }, nil
}

func decodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.GetUsersRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error{

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	
	if err == nil{
		panic("Error: nil error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	validationError, ok := err.(*utils.CustomValidationErrors)

	if ok {

		w.WriteHeader(http.StatusBadRequest)

		errorStruct :=  struct {
			Errors []string
		}{
			Errors: validationError.GetErrors(),
		}

		json.NewEncoder(w).Encode(errorStruct)
	
		return

	} else {
		switch err{
		case utils.ErrNotFound:
			w.WriteHeader(utils.ErrNotFound.GetCode())
		case utils.ErrBadData:
			w.WriteHeader(utils.ErrBadData.GetCode())
		case utils.ServerError:
			w.WriteHeader(utils.ServerError.GetCode())
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}	

	json.NewEncoder(w).Encode(err.Error())
}
