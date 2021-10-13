package models

import (
	"database/sql"
	//"log"
	utils "leonel/prototype_b/pkg/utils"
)

type Question struct {
	ID        string `json:"id"`
	Title     string `json:"title" validate:"required,max=100"`
	Statement string `json:"statement" validate:"required,max=255"`

	User   User   `json:"user"`
	Answer Answer `json:"answer"`
}

func IndexQuestions(db *sql.DB) ([]Question, error) {

	questions := make([]Question, 0)

	sql := `SELECT questions.id, questions.title, questions.statement, questions.user_id, users.username, 
			answers.id AS answer_id, answers.statement AS answer_statement, ua.id AS answer_user_id , ua.username AS answer_user_username
	
			FROM questions 

				LEFT OUTER JOIN users ON users.id = questions.user_id 
				LEFT OUTER JOIN answers ON answer_id = answers.id
				LEFT OUTER JOIN users ua ON ua.id = answers.user_id`

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var question Question

		err := rows.Scan(&question.ID, &question.Title, &question.Statement, &question.User.ID, &question.User.Username, 
			&question.Answer.ID, &question.Answer.Statement, &question.Answer.User.ID, &question.Answer.User.Username)

		if err != nil {
			return nil, err
		}

		questions = append(questions, question)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func CreateQuestion(db *sql.DB, question *Question) (*Question, error) {

	var new_id string
	
	err := db.QueryRow("INSERT INTO questions(user_id, title, statement) VALUES($1, $2, $3) RETURNING ID", question.User.ID, question.Title, question.Statement).Scan(&new_id)

	if err != nil{
		return nil, err
	}

	question, err = GetQuestion(db, new_id)

	if err != nil{
		return nil, err
	}

	return question, nil
}

func GetQuestion(db *sql.DB, question_id string) (*Question, error){
	
	row := db.QueryRow(`SELECT questions.id, questions.title, questions.statement, questions.user_id, users.username, 
		answers.id AS answer_id, answers.statement AS answer_statement, ua.id AS answer_user_id , ua.username AS answer_user_username
		
		FROM questions 
			LEFT OUTER JOIN users ON users.id = questions.user_id 
			LEFT OUTER JOIN answers ON answer_id = answers.id
			LEFT OUTER JOIN users ua ON ua.id = answers.user_id

		WHERE questions.id = $1`, question_id)

	question := new(Question)

	err := row.Scan(&question.ID, &question.Title, &question.Statement, &question.User.ID, &question.User.Username, 
		&question.Answer.ID, &question.Answer.Statement, &question.Answer.User.ID, &question.Answer.User.Username)

	if err != nil{

		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, err
		}
	}

	return question, nil
}

func UpdateQuestion(db *sql.DB, question_id, question_title, question_statement string ) (*Question, error) {

	_, err := db.Exec("UPDATE questions SET(title, statement) = ($2, $3) WHERE id = $1", question_id, question_title, question_statement)

	if err != nil{
		return nil, err
	}

	return GetQuestion(db, question_id)
}

func DeleteQuestion(db *sql.DB, question_id string) error {

	_, err := db.Exec("DELETE FROM questions WHERE id = $1", question_id)

	if err != nil {
		return err
	}

	return nil
}
