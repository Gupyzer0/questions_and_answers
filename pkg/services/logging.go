package services

import (
	"fmt"
	"leonel/prototype_b/pkg/db/models"

	"time"

	"github.com/go-kit/kit/log"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   QuestionsAndAnswers
}

func (mw LoggingMiddleware) GetQuestions() (output []models.Question, err error) {

	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetQuestions",
			"input", nil,
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetQuestions()

	return
}

func (mw LoggingMiddleware) GetQuestion(question_id string) (output *models.Question, err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetQuestion",
			"input", question_id,
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetQuestion(question_id)

	return
}

func (mw LoggingMiddleware) CreateQuestion(title,statement,user_id string) (output *models.Question, err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "CreateQuestion",
			"input", fmt.Sprintln(title,statement,user_id),
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.CreateQuestion(title,statement,user_id)

	return
}

func (mw LoggingMiddleware) GetQuestionsByUser(user_id string) (output []models.Question, err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetQuestionsByUser",
			"input", user_id,
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetQuestionsByUser(user_id)

	return
}

func (mw LoggingMiddleware) UpdateQuestion(question_id, question_title, question_statement string) (output *models.Question, err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "UpdateQuestion",
			"input", fmt.Sprintln(question_id, question_title, question_statement),
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.UpdateQuestion(question_id, question_title, question_statement)

	return
}

func (mw LoggingMiddleware) DeleteQuestion(question_id string) (err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "DeleteQuestion",
			"input",question_id,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.DeleteQuestion(question_id)

	return
}

func (mw LoggingMiddleware) UpdateAnswer(question_id, statement, user_id string) (output *models.Answer, err error){
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "UpdateAnswer",
			"input",fmt.Sprintln(question_id, statement, user_id),
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.UpdateAnswer(question_id, statement, user_id)

	return
}

func (mw LoggingMiddleware) GetUsers() (output []models.User, err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetUsers",
			"input",nil,
			"output", fmt.Sprintln(output),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.GetUsers()

	return
}

func (mw LoggingMiddleware) Validate(req interface{}) (err error) {
	defer func(begin time.Time) {
		mw.Logger.Log(
			"method", "GetQuestion",
			"input",req,
			"output", nil,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	err = mw.Next.Validate(req)

	return
}
