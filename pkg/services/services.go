package services

import (
	"fmt"
	"leonel/prototype_b/pkg/db/models"
	"leonel/prototype_b/pkg/utils"

	"github.com/go-playground/validator/v10"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)
//mock models for testing services

//service
type QuestionsAndAnswers interface{
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

type qaService struct{
	mdl *models.Models_wrapper
	validator *validator.Validate
	Validator_err_translatator *ut.Translator
}

//Questions and answer service constructor
func NewQuestionsAndAnswersService(wrapper *models.Models_wrapper) QuestionsAndAnswers{
	
	validator := validator.New()

	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	_ = en_translations.RegisterDefaultTranslations(validator, trans)

	service := qaService{
		mdl: wrapper,
		validator: validator,
		Validator_err_translatator: &trans,
	}

	return &service
}

func translateError(err error, trans ut.Translator) (errs *utils.CustomValidationErrors) {

	var cstm_val_errors utils.CustomValidationErrors

	validator_errors := err.(validator.ValidationErrors)
	
	for _, e := range validator_errors {
		translatedErr := fmt.Errorf(e.Translate(trans))
		cstm_val_errors.AddError(translatedErr)
	}

	return &cstm_val_errors
}

func (srv *qaService) Validate(req interface{}) error {
	err := srv.validator.Struct(req)

	if err == nil{
		return nil
	}

	return translateError(err, *srv.Validator_err_translatator)
}

func (srv *qaService) GetQuestion(question_id string) (*models.Question, error){
	
	question, err := srv.mdl.Questions.Get(question_id)

	if err != nil{
		return nil, err
	}

	return question, nil
}

func (srv *qaService) GetQuestions() ([]models.Question, error) {
	
	questions, err := srv.mdl.Questions.Index()

	if err != nil{
		return nil, err
	}

	return questions, err
}

func (srv *qaService) GetQuestionsByUser(user_id string) ([]models.Question, error){

	questions, err := srv.mdl.Users.UserQuestions(user_id)

	if err != nil{
		return nil, err
	}

	return questions, nil
}

func (srv *qaService) CreateQuestion(title,statement,user_id string) (*models.Question, error){

	question, err := srv.mdl.Questions.Create(title,statement,user_id)

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (srv *qaService) UpdateQuestion(question_id, question_title, question_statement string) (*models.Question, error){

	updated_question, err := srv.mdl.Questions.Update(question_id, question_title, question_statement)

	if err != nil{
		return nil, err
	}

	return updated_question, nil
}

func (srv *qaService) UpdateAnswer(question_id, statement, user_id string) (*models.Answer, error){

	answer, err := srv.mdl.Answers.Update(question_id, statement, user_id)

	if err != nil{
		return nil, err
	}

	return answer, nil	
}

func (srv *qaService) DeleteQuestion(question_id string) error{

	err := srv.mdl.Questions.Delete(question_id)

	if err != nil {
		return err
	}

	return nil
}

func (srv *qaService) GetUsers() ([]models.User, error) {
	
	users, err := srv.mdl.Users.Index()

	if err != nil {
		return nil, err
	}

	return users, nil

}
