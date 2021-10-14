package services

import (
	"leonel/prototype_b/pkg/db/models"
	//models "leonel/prototype_b/pkg/db/models"
	"github.com/go-playground/validator/v10"

)

//service
type Service interface{
	GetQuestion(string) (*models.Question,error)
	GetQuestions() ([]models.Question, error)	
	GetQuestionsByUser(string) ([]models.Question, error)	
	CreateQuestion(string,string,string) (*models.Question, error)	
	UpdateQuestion(string,string,string) (*models.Question, error)	
	UpdateAnswer(string,string,string) (*models.Answer, error)
	DeleteQuestion(string) error
	GetUsers() ([]models.User, error)

	Validate(interface{}) error
}

type QuestionsAndAnswersService struct{
	mdl *models.Models_wrapper
	validator *validator.Validate
}

//Questions and answer service constructor
func NewQuestionsAndAnswersService(wrapper *models.Models_wrapper) *QuestionsAndAnswersService{
	return &QuestionsAndAnswersService{
		mdl: wrapper,
		validator: validator.New(),
	}
}

func (srv *QuestionsAndAnswersService) Validate(req interface{}) error {
	return srv.validator.Struct(req)
}

func (srv *QuestionsAndAnswersService) GetQuestion(question_id string) (*models.Question, error){
	
	question, err := srv.mdl.Questions.Get(question_id)

	if err != nil{
		return nil, err
	}

	return question, nil
}

func (srv *QuestionsAndAnswersService) GetQuestions() ([]models.Question, error) {
	
	questions, err := srv.mdl.Questions.Index()

	if err != nil{
		return nil, err
	}

	return questions, err
}

func (srv *QuestionsAndAnswersService) GetQuestionsByUser(user_id string) ([]models.Question, error){

	questions, err := srv.mdl.Users.UserQuestions(user_id)

	if err != nil{
		return nil, err
	}

	return questions, nil
}

func (srv *QuestionsAndAnswersService) CreateQuestion(title,statement,user_id string) (*models.Question, error){

	question, err := srv.mdl.Questions.Create(title,statement,user_id)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (srv *QuestionsAndAnswersService) UpdateQuestion(question_id, question_title, question_statement string) (*models.Question, error){

	updated_question, err := srv.mdl.Questions.Update(question_id, question_title, question_statement)

	if err != nil{
		return nil, err
	}

	return updated_question, nil
}

func (srv *QuestionsAndAnswersService) UpdateAnswer(question_id, statement, user_id string) (*models.Answer, error){

	answer, err := srv.mdl.Answers.Update(question_id, statement, user_id)

	if err != nil{
		return nil, err
	}

	return answer, nil	
}

func (srv *QuestionsAndAnswersService) DeleteQuestion(question_id string) error{

	err := srv.mdl.Questions.Delete(question_id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *QuestionsAndAnswersService) GetUsers() ([]models.User, error) {
	
	users, err := srv.mdl.Users.Index()

	if err != nil {
		return nil, err
	}

	return users, nil

}
