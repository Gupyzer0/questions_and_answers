package services

import (
	"database/sql"
	"log"
	//"sync"

	models "leonel/prototype_b/pkg/db/models"
	//utils "leonel/prototype_b/pkg/utils"
)

//service
type Service interface{
	GetQuestion(string) (*models.Question,error)

	GetQuestions() []models.Question
	
	GetQuestionsByUser(string) ([]models.Question, error)
	
	CreateQuestion(*models.Question) (*models.Question, error)
	
	UpdateQuestion(string,string,string) (*models.Question, error)
	
	UpdateAnswer(string, *models.Answer) (*models.Answer, error)
	
	DeleteQuestion(string) error
}

type QuestionsAndAnswersService struct{
	db *sql.DB
	questions map[string]*models.Question
}

//constructor
func NewQuestionsAndAnswersService(db *sql.DB) *QuestionsAndAnswersService{
	return &QuestionsAndAnswersService{
		db: db,
	}
}

func (srv *QuestionsAndAnswersService) GetQuestion(question_id string) (*models.Question, error){

	question, err := models.GetQuestion(srv.db, question_id)

	if err != nil{
		log.Println("services error")
		return nil, err
	}

	return question, nil
}

func (srv *QuestionsAndAnswersService) GetQuestions() ([]models.Question) {
	
	questions, err := models.IndexQuestions(srv.db)

	if err != nil{
		log.Println(err)
		return nil
	}

	return questions
}

func (srv *QuestionsAndAnswersService) GetQuestionsByUser(user_id string) ([]models.Question, error){

	questions, err := models.UserQuestions(srv.db, user_id)

	if err != nil{
		return nil, err
	}

	return questions, nil
}

func (srv *QuestionsAndAnswersService) CreateQuestion(question *models.Question) (*models.Question, error){

	question, err := models.CreateQuestion(srv.db, question) 

	if err != nil {
		return nil, err
	}

	return question, nil
}

func (srv *QuestionsAndAnswersService) UpdateQuestion(question_id string, question_title string, question_statement string) (*models.Question, error){

	updated_question, err := models.UpdateQuestion(srv.db, question_id, question_title, question_statement)

	if err != nil{
		return nil, err
	}

	return updated_question, nil
}

func (srv *QuestionsAndAnswersService) UpdateAnswer(question_id string, answer *models.Answer) (*models.Answer, error){

	answer, err := models.UpdateAnswer(srv.db, question_id, answer)

	if err != nil{
		return nil, err
	}

	return answer, nil	
}

func (srv *QuestionsAndAnswersService) DeleteQuestion(question_id string) error{

	err := models.DeleteQuestion(srv.db, question_id)

	if err != nil {
		return err
	}

	return nil
}
