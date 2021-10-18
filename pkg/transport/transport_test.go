package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	//"reflect"

	//"log"
	"net/http"
	"testing"

	endpoints "leonel/prototype_b/pkg/endpoints"
	"leonel/prototype_b/pkg/utils"

	"github.com/gorilla/mux"
)

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
	req, _ = http.NewRequest("POST","/questions",bytes.NewReader(jsonBytes))

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

		req, _ := http.NewRequest("GET","questions/"+test.Url_parameter,nil)

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


