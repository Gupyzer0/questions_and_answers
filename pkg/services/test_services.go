package services

import (
	"leonel/prototype_b/pkg/db/models"
	"leonel/prototype_b/mocks/mock_db"

	//"testing"
)

func makeMockModelsWrapper() models.Models_wrapper {

	MockWrapper := models.Models_wrapper {
		Questions: &mock_models.MockQuestionsInterface{},
		Answers: &mock_models.MockAnswersInterface{},
		Users: &mock_models.MockUsersInterface{},
	}   

	return MockWrapper

}
/*
func TestGetQuestions(t *testing.T) {
	mock_wrapper := makeMockModelsWrapper()

	mock_wrapper.Answers.Get()

}
*/