package services

import (
	mock_models "leonel/prototype_b/mocks/mock_db"
	"leonel/prototype_b/pkg/db/models"
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"

	//"golang.org/x/text/search"
	"gopkg.in/guregu/null.v4"
)

type testConfig struct {
	mockModels models.Models_wrapper
	service    QuestionsAndAnswers
}

func setupTest(t *testing.T) *testConfig {

	var suite testConfig

	ctrl := gomock.NewController(t)

	suite.mockModels = models.Models_wrapper{
		Questions: mock_models.NewMockQuestionsInterface(ctrl),
		Answers:   mock_models.NewMockAnswersInterface(ctrl),
		Users:     mock_models.NewMockUsersInterface(ctrl),
	}

	suite.service = NewQuestionsAndAnswersService(&suite.mockModels)

	return &suite
}

func (s *testConfig) GetQuestionModel() *mock_models.MockQuestionsInterfaceMockRecorder {
	return s.mockModels.Questions.(*mock_models.MockQuestionsInterface).EXPECT()
}

func (s *testConfig) GetAnswersModel() *mock_models.MockAnswersInterfaceMockRecorder {
	return s.mockModels.Answers.(*mock_models.MockAnswersInterface).EXPECT()
}

func (s *testConfig) GetUsersModel() *mock_models.MockUsersInterfaceMockRecorder {
	return s.mockModels.Users.(*mock_models.MockUsersInterface).EXPECT()
}

func TestGetQuestions(t *testing.T) {
	config := setupTest(t)

	expected := []models.Question{
		{
			ID: uuid.NewV4().String(),
			Title:     "Is this a question",
			Statement: "The real question is the answer?",
			User: models.User{
				ID:       null.NewString(uuid.NewV4().String(), true),
				Username: null.NewString("user1", true),
			},
		},
	}

	config.GetQuestionModel().Index().Return(expected, nil)

	got, err := config.service.GetQuestions()
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestGetQuestion(t *testing.T) {
	config := setupTest(t)

	type test struct {
		Search_id string
		ExpectedQuestionResponse *models.Question
		ExpectedErrorResponse interface{}
	}

	tests := []test{
		{
			Search_id : uuid.NewV4().String(),
			ExpectedQuestionResponse: &models.Question {
				ID: uuid.NewV4().String(),
				Title: "Question 1",
				Statement: "Question statement",
				User: models.User{
					ID: null.NewString(uuid.NewV4().String(), true),
					Username: null.StringFrom("user1"),
				},
			},
			ExpectedErrorResponse: nil,
		},
		{
			Search_id: uuid.NewV4().String(),
			ExpectedQuestionResponse: nil,
			ExpectedErrorResponse: assert.AnError,
		},
	}

	for _, test := range tests {

		t.Run("Testing GetQuestion service", func(t *testing.T){

			config.GetQuestionModel().Get(test.Search_id).Return(test.ExpectedQuestionResponse, test.ExpectedErrorResponse)

			question, err := config.service.GetQuestion(test.Search_id)

			assert.Equal(t, test.ExpectedQuestionResponse, question, "Response not equal")
			assert.Equal(t, test.ExpectedErrorResponse, err)
		})
	}
}

func TestGetQuestionsByUser(t *testing.T){
	
	config := setupTest(t)

	type test struct {
		Search_id string
		ExpectedQuestionsResponse []models.Question
		ExpectedErrorResponse interface{}
	}

	tests := []test {
		{
			Search_id: uuid.NewV4().String(),
			ExpectedQuestionsResponse: []models.Question {
				{
					ID: uuid.NewV4().String(),
					Title: "Question 1",
					Statement: "Question statement",
					User: models.User{
						ID: null.NewString(uuid.NewV4().String(), true),
						Username: null.StringFrom("user1"),
					},
				},
				{
					ID: uuid.NewV4().String(),
					Title: "Question 2",
					Statement: "Question statement 2",
					User: models.User{
						ID: null.NewString(uuid.NewV4().String(), true),
						Username: null.StringFrom("user2"),
					},
				},
			},
			ExpectedErrorResponse : nil,
		},
		{
			Search_id: uuid.NewV4().String(),
			ExpectedQuestionsResponse: nil,
			ExpectedErrorResponse : assert.AnError,
		},
	}

	for _, test := range tests {

		t.Run("Testing GetQuestionByUser service", func(t *testing.T){
			
			config.GetUsersModel().UserQuestions(test.Search_id).Return(test.ExpectedQuestionsResponse, test.ExpectedErrorResponse)

			questions, err := config.service.GetQuestionsByUser(test.Search_id)

			assert.Equal(t, test.ExpectedQuestionsResponse, questions, "Response not equal")
			assert.Equal(t, test.ExpectedErrorResponse, err)
		})
	}
}

func TestCreateQuestion(t *testing.T){

	config := setupTest(t)

	type test struct {
		Title,Statement,User_id string
		ExpectedQuestionResponse *models.Question
		ExpectedErrorResponse interface{}
	}

	tests := []test{
		{
			Title: "Question Title",
			Statement: "Question statement",
			User_id: string(uuid.NewV4().String()),
			ExpectedQuestionResponse: &models.Question{
				ID: uuid.NewV4().String(),
				Title: "Question Title",
				Statement: "Question statement",
				User: models.User{
					ID: null.NewString(uuid.NewV4().String(), true),
					Username: null.StringFrom("user1"),
				},
			},
			ExpectedErrorResponse: nil,
		},
		{
			Title: "Question Title",
			Statement: "Question statement",
			User_id: string(uuid.NewV4().String()),
			ExpectedQuestionResponse: nil,
			ExpectedErrorResponse: assert.AnError,
		},
	}

	for _, test := range tests {

		t.Run("Testing CreateQuestion service", func(t *testing.T){

			config.GetQuestionModel().Create(test.Title, test.Statement, test.User_id).Return(test.ExpectedQuestionResponse, test.ExpectedErrorResponse)

			question, err := config.service.CreateQuestion(test.Title, test.Statement, test.User_id)

			assert.Equal(t, test.ExpectedQuestionResponse, question, "Response not equal")
			assert.Equal(t, test.ExpectedErrorResponse, err)
		})
	}
}


func TestUpdateQuestion(t *testing.T) {
	config := setupTest(t)

	type test struct {
		Question_id, Question_title, Question_statement string
		ExpectedQuestionResponse *models.Question
		ExpectedErrorResponse interface{}
	}

	tests := []test {
		{
			Question_id: uuid.NewV4().String(),
			Question_title: "Question 1",
			Question_statement: "Question statement",
			ExpectedQuestionResponse: &models.Question{
				ID: uuid.NewV4().String(),
					Title: "Question 1",
					Statement: "Question statement",
					User: models.User{
						ID: null.NewString(uuid.NewV4().String(), true),
						Username: null.StringFrom("user1"),
				},
			},
			ExpectedErrorResponse: nil,
		},
		{
			Question_id: uuid.NewV4().String(),
			Question_title: "Question 1",
			Question_statement: "Question statement",
			ExpectedQuestionResponse: nil,
			ExpectedErrorResponse: assert.AnError,
		},
	}

	for _, test := range tests {

		t.Run("Testing UpdateQuestionService", func(t *testing.T){

			config.GetQuestionModel().Update(test.Question_id, test.Question_title, test.Question_statement).Return(test.ExpectedQuestionResponse, test.ExpectedErrorResponse)

			question, err := config.service.UpdateQuestion(test.Question_id, test.Question_title, test.Question_statement)

			assert.Equal(t, test.ExpectedQuestionResponse, question, "Response not equal")
			assert.Equal(t, test.ExpectedErrorResponse, err, "Errors aren't equal")
		})
	}
}

func TestUpdateAnswer(t *testing.T) {

	config := setupTest(t)

	type test struct {
		Question_id, Statement, User_id string
		ExpectedQuestionResponse *models.Answer
		ExpectedErrorResponse interface{}
	}

	tests := []test {
		{
			Question_id: uuid.NewV4().String(),
			Statement: "New Statement",
			User_id: uuid.NewV4().String(),
			ExpectedQuestionResponse: &models.Answer {
				User: models.User {
					ID: null.StringFrom(uuid.NewV4().String()),
					Username: null.StringFrom("User2"),
				},
			},
			ExpectedErrorResponse: nil,
		},
		{
			Question_id: uuid.NewV4().String(),
			Statement: "New Statement",
			User_id: uuid.NewV4().String(),
			ExpectedQuestionResponse: nil,
			ExpectedErrorResponse: assert.AnError,
		},
	}

	for _, test := range tests {

		t.Run("Testing UpdateAnswerService", func(t *testing.T){

			config.GetAnswersModel().Update(test.Question_id, test.Statement, test.User_id).Return(test.ExpectedQuestionResponse, test.ExpectedErrorResponse)

			question, err := config.service.UpdateAnswer(test.Question_id, test.Statement, test.User_id)

			assert.Equal(t, test.ExpectedQuestionResponse, question, "Response not equal")
			assert.Equal(t, test.ExpectedErrorResponse, err, "Errors aren't equal")
		})

	}

}

func TestDeleteQuestion( t *testing.T) {

	config := setupTest(t)

	type test struct {
		Question_id string
		ExpectedErrorResponse interface{}
	}

	tests := []test {
		{
			Question_id: uuid.NewV4().String(),
			ExpectedErrorResponse: nil,
		},
		{
			Question_id: uuid.NewV4().String(),
			ExpectedErrorResponse: assert.AnError,
		},
	}
	
	for _, test := range tests {

		t.Run("Testing DeleteQuestionService", func(t *testing.T){

			config.GetQuestionModel().Delete(test.Question_id).Return(test.ExpectedErrorResponse)

			err := config.service.DeleteQuestion(test.Question_id)

			assert.Equal(t, test.ExpectedErrorResponse, err, "Errors aren't equal")

		})
	}
}

func TestGetUsers(t *testing.T) {	

	config := setupTest(t)

	type test struct{
		ExpectedResponse []models.User
		ExpectedErrorResponse interface{}
	}

	tests := []test {
		{
			ExpectedResponse: []models.User{
				{
					ID: null.StringFrom(uuid.NewV4().String()),
					Username: null.StringFrom("Test user"),
				},
			},
			ExpectedErrorResponse: nil,
		},
		{
			ExpectedResponse: nil,
			ExpectedErrorResponse: assert.AnError,
		},
	}

	for _, test := range tests{

		t.Run("Testing GetUsers Service", func(t *testing.T){

			config.GetUsersModel().Index().Return(test.ExpectedResponse, test.ExpectedErrorResponse)

			users, err := config.service.GetUsers()

			assert.Equal(t,test.ExpectedResponse, users)
			assert.Equal(t,test.ExpectedErrorResponse, err)
		})
	}

}
