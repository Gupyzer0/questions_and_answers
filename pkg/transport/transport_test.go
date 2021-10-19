package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"net/http"
	"net/http/httptest"
	"testing"

	endpoints "leonel/prototype_b/pkg/endpoints"
	"leonel/prototype_b/pkg/utils"

	"github.com/gorilla/mux"
)

func TestDecodeGetQuestionsRequest(t *testing.T) {

	req, _ := http.NewRequest("GET","/questions",nil)

	t.Run("Test Decode get questions request", func(t * testing.T){
		deq_req, err := decodeGetQuestionsRequest(context.Background(), req)

		if deq_req != (endpoints.GetQuestionsRequest{}) || err != nil {
			t.Errorf("Unexpected error in the get questions request decoder")
		}
	})
}

func TestDecodeCreateQuestionRequest(t *testing.T) {
	//TODO change to table testing scheme
	req_struct := endpoints.CreateQuestionRequest{
		Title: "Test Title",
		Statement: "Test statement",
		User_id: "1",
	}

	jsonBytes, _ := json.Marshal(req_struct)
	req, _ := http.NewRequest("POST","/questions",bytes.NewReader(jsonBytes))

	t.Run("Decoding create question request", func (t *testing.T){
		dec_req, err := decodeCreateQuestionRequest(context.Background(), req)

		if err != nil {
			t.Errorf("Unexpected error result: %v",err)
		}

		if dec_req != req_struct {
			t.Errorf("Error decoding json")
		} else {
			fmt.Println("test passed")
		}
	})

	req_bad_data := "dummy string"
	jsonBytes, _ = json.Marshal(req_bad_data)
	req = httptest.NewRequest("POST","/questions",bytes.NewReader(jsonBytes))

	t.Run("Decoding create question with bad data", func(t *testing.T) {
		_, err := decodeCreateQuestionRequest(context.Background(), req)

		if err != utils.ErrBadData {
			t.Errorf("No expected error result")
		}
	})
}

func TestDecodeGetQuestionRequest(t *testing.T) {

	type TestQuestionRequest struct {
		Url_parameter string
		Expected_decode interface{}
		Expected_error interface{}
	}

	tests := []TestQuestionRequest{
		{
			Url_parameter: "1",
			Expected_decode: endpoints.GetQuestionRequest{
				Question_id: "1",
			},
			Expected_error: nil,
		},
		{
			Url_parameter: "",
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
	}

	testname := "Decode GET question request"

	for _, test := range tests {

		req, _ := http.NewRequest("GET","/questions/"+test.Url_parameter,nil)

		vars := map[string]string{
			"id": test.Url_parameter,
		}		

		req = mux.SetURLVars(req, vars)

		t.Run(testname, func(t *testing.T){
			dec_req, err := decodeGetQuestionRequest(context.Background(),req)
			
			if dec_req != test.Expected_decode || err != test.Expected_error {
				t.Errorf("Unexpected result:\n decode result: %v, expected: %v \n decode error: %v, expected: %v",dec_req,test.Expected_decode,err,test.Expected_error)
			}
		})
	}
}

func TestDecodeUpdateQuestionRequest(t *testing.T) {
	type TestUpdateQuestionRequest struct {
		Url_parameter string
		EndpointRequest interface{}
		Expected_decode interface{}
		Expected_error interface{}
	}

	tests := []TestUpdateQuestionRequest {
		{
			Url_parameter: "1",
			EndpointRequest: endpoints.UpdateQuestionRequest {
				Title: "Mock Title",
				Statement: "Mock statement",
			},
			Expected_decode: endpoints.UpdateQuestionRequest{ 
				Question_id: "1",
				Title: "Mock Title", 
				Statement: "Mock statement", 
			},
			Expected_error: nil,
		},
		{
			Url_parameter: "",
			EndpointRequest: endpoints.UpdateQuestionRequest {
				Title: "Mock Title",
				Statement: "Mock statement",
			},
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
		{
			Url_parameter: "1",
			EndpointRequest: "dummy string",
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
	}

	for _, test := range tests {

		jsonBytes, _ := json.Marshal(test.EndpointRequest)

		req, _ := http.NewRequest("PUT","questions/"+test.Url_parameter,bytes.NewReader(jsonBytes))

		vars := map[string]string{
			"id": test.Url_parameter,
		}

		req = mux.SetURLVars(req, vars)

		t.Run("Testing update questions decoder", func(t *testing.T) {

			dec_req, err := decodeUpdateQuestionRequest(context.Background(),req)

			if dec_req != test.Expected_decode || err != test.Expected_error {
				t.Errorf("Unexpected result:\n %v, expected: %v\n error received: %v, expected: %v",dec_req, test.Expected_decode, err, test.Expected_error )
			}

			//decoded_req := dec_req.(endpoints.UpdateQuestionRequest)
			//expected_decode := test.Expected_decode.(endpoints.UpdateQuestionRequest)

			/*fmt.Println("REFLECT DEEP EQUAL",reflect.DeepEqual(decoded_req, expected_decode))

			if decoded_req.Question_id != expected_decode.Question_id {
				t.Errorf("Unexpected result in question_id:\n %v, expected: %v",decoded_req.Question_id, expected_decode.Question_id )
			}

			if decoded_req.Title != expected_decode.Title {
				t.Errorf("Unexpected result in title:\n %v, expected: %v",decoded_req.Title, expected_decode.Title )
				t.Errorf("Types: %T, expected: %T",decoded_req.Title, expected_decode.Title )
			}

			if decoded_req.Statement != expected_decode.Statement {
				t.Errorf("Unexpected result in statement:\n %v, expected: %v",decoded_req.Statement, expected_decode.Statement )
			}

			if err != test.Expected_error {
				t.Errorf("Unexpected error: %v, expected: %v",err, test.Expected_error )
			}*/
		})
	}
}

func TestDecodeUpdateAnswerRequest(t *testing.T) {
	type TestUpdateAnswerRequest struct {
		Url_parameter string
		EndpointRequest interface{}
		Expected_decode interface{}
		Expected_error interface{}
	}

	tests := []TestUpdateAnswerRequest {
		{
			Url_parameter: "1",
			EndpointRequest: endpoints.UpdateAnswerRequest {
				Question_id: "1",
				Statement: "Mock Statement",
				User_id: "1",				
			},
			Expected_decode: endpoints.UpdateAnswerRequest{ 
				Question_id: "1",
				Statement: "Mock Statement", 
				User_id: "1",
			},
			Expected_error: nil,
		},
		{
			Url_parameter: "",
			EndpointRequest: endpoints.UpdateAnswerRequest {
				Question_id: "1",
				Statement: "Mock Statement",
				User_id: "1",				
			},
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
		{
			Url_parameter: "1",
			EndpointRequest: "Dummy string",
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
	}

	for _, test := range tests {

		jsonBytes, _ := json.Marshal(test.EndpointRequest)

		req, _ := http.NewRequest("POST","questions/"+test.Url_parameter+"/answer",bytes.NewReader(jsonBytes))

		vars := map[string]string{
			"id": test.Url_parameter,
		}

		req = mux.SetURLVars(req, vars)

		t.Run("Testing update answers decoder", func(t *testing.T) {

			dec_req, err := decodeUpdateAnswerRequest(context.Background(),req)

			if dec_req != test.Expected_decode || err != test.Expected_error {
				t.Errorf("Unexpected result:\n %v, expected: %v\n error received: %v, expected: %v",dec_req, test.Expected_decode, err, test.Expected_error )
			}
		})
	}
}


func TestDecodeGetQuestionsByUserRequest(t *testing.T) {
	type TestGetQuestionsByUserRequest struct {
		Url_parameter string
		Expected_decode interface{}
		Expected_error interface{}
	}

	tests := []TestGetQuestionsByUserRequest {
		{
			Url_parameter: "1",
			Expected_decode: endpoints.GetQuestionsByUserRequest{
				User_id: "1",
			},
			Expected_error: nil,
		},
		{	
			Url_parameter: "",
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
	}

	for _, test := range tests {

		req, _ := http.NewRequest("GET","users/"+test.Url_parameter+"/questions",nil)

		vars := map[string]string{
			"id": test.Url_parameter,
		}

		req = mux.SetURLVars(req, vars)

		t.Run("Testing get questions by user decoder", func(t *testing.T) {

			dec_req, err := decodeGetQuestionsByUserRequest(context.Background(),req)

			if dec_req != test.Expected_decode || err != test.Expected_error {
				t.Errorf("Unexpected result:\n %v, expected: %v\n error received: %v, expected: %v",dec_req, test.Expected_decode, err, test.Expected_error )
			}
		})
	}
}

func TestDecodeDeleteQuestionRequest(t *testing.T) {
	type TestDeleteQuestion struct {
		Url_parameter string
		Expected_decode interface{}
		Expected_error interface{}
	}

	tests := []TestDeleteQuestion {
		{
			Url_parameter: "1",
			Expected_decode: endpoints.DeleteQuestionRequest{
				Question_id: "1",
			},
			Expected_error: nil,
		},
		{	
			Url_parameter: "",
			Expected_decode: nil,
			Expected_error: utils.ErrBadData,
		},
	}

	for _, test := range tests {

		req, _ := http.NewRequest("DELETE","questions/"+test.Url_parameter,nil)

		vars := map[string]string{
			"id": test.Url_parameter,
		}

		req = mux.SetURLVars(req, vars)

		t.Run("Testing delete questions request decoder", func(t *testing.T) {

			dec_req, err := decodeDeleteQuestionRequest(context.Background(),req)

			if dec_req != test.Expected_decode || err != test.Expected_error {
				t.Errorf("Unexpected result:\n %v, expected: %v\n error received: %v, expected: %v",dec_req, test.Expected_decode, err, test.Expected_error )
			}
		})
	}
}

func TestDecodeGetUsersRequest(t *testing.T) {

	req, _ := http.NewRequest("GET","/users",nil)

	t.Run("Test Decode get users request", func(t * testing.T){
		deq_req, err := decodeGetUsersRequest(context.Background(), req)

		if deq_req != (endpoints.GetUsersRequest{}) || err != nil {
			t.Errorf("Unexpected error in the get users request decoder")
		}
	})
}

//TODO more tests here
func TestEncodeResponse(t *testing.T) {

	t.Run("Testing encoding response", func(t *testing.T){

		handler := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "{ response:{} }")
		}

		req, _ := http.NewRequest("GET","/questions",nil)
		w := httptest.NewRecorder()

		handler(w, req)

		resp := w.Result()

		err := encodeResponse(context.Background(), w, resp)

		if err != nil {
			t.Errorf("Error encoding the response: %v",err)
		}
	})
}

func TestEncodeError(t *testing.T) {

	var cstm_val_errors utils.CustomValidationErrors
	cstm_val_errors.AddError(fmt.Errorf("Error1"))
	cstm_val_errors.AddError(fmt.Errorf("Error2"))

	type EncodeErrorTest struct {
		Err error
		ExpectedResponse string
		ExpectedStatusCode int
	}

	tests := []EncodeErrorTest {
		{
			Err: utils.ErrBadData,
			ExpectedResponse: "Bad Request",
			ExpectedStatusCode: utils.ErrBadData.GetCode(),
		},
		{
			Err: utils.ErrNotFound,
			ExpectedResponse: "Resource Not Found",
			ExpectedStatusCode: utils.ErrNotFound.GetCode(),
		},
		{
			Err: utils.ServerError,
			ExpectedResponse: "Server Error",
			ExpectedStatusCode: utils.ServerError.GetCode(),
		},
		{
			Err: &cstm_val_errors,
			ExpectedResponse: `{Errors:[Error1,Error2]}`,
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Err: fmt.Errorf("dummy unknow error"),
			ExpectedResponse: "",
			ExpectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {

		t.Run("Testing encoding of errors", func(t *testing.T){

			w := httptest.NewRecorder()
			err := test.Err

			encodeError(context.Background(), err, w)

			body_string := w.Body.String()
			body_string = strings.ReplaceAll(body_string,"\"","")
			body_string = strings.TrimSpace(body_string)

			if body_string != test.ExpectedResponse {
				t.Errorf("Wrong response, wanted %v but got %v", test.ExpectedResponse, body_string)
			}

			if test.ExpectedStatusCode != w.Code {
				t.Errorf("Wrong response code, wanted %v but got %v", test.ExpectedStatusCode, w.Code)
			}

			fmt.Println("------>",w.Body.String())

		})
	}
}
