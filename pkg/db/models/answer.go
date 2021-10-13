package models

import (
	"database/sql"
	"leonel/prototype_b/pkg/utils"
	//"log"

	"gopkg.in/guregu/null.v4"
)

type Answer struct {
	ID        null.String `json:"id"`
	Statement null.String `json:"statement"`
	User      User   `json:"user"`
}

func GetAnswer(db *sql.DB, answer_id string) (*Answer, error){

	row := db.QueryRow(`SELECT answers.id, answers.statement, users.id, users.username 
		FROM answers

		INNER JOIN users ON answers.user_id = users.id

		WHERE answers.id = $1`, answer_id)

	answer := new(Answer)

	err := row.Scan(&answer.ID, &answer.Statement, &answer.User.ID, &answer.User.Username)

	if err != nil{

		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, err
		}
	}

	return answer, nil

}

func UpdateAnswer(db *sql.DB, question_id string, answer *Answer) (*Answer, error) {

	var answ_id string

	err := db.QueryRow("SELECT answer_id FROM questions WHERE id = $1",question_id).Scan(&answ_id)

	if err != nil {
		return nil, err
	}

	var updated_answer *Answer

	row := db.QueryRow("UPDATE answers SET(statement, user_id) = ($2, $3) WHERE id = $1", answ_id, answer.Statement, answer.User.ID)

	if err= row.Err(); err != nil{
		return nil, err
	}

	updated_answer, err = GetAnswer(db, answ_id);

	if err != nil {
		return nil, err
	}

	return updated_answer, nil
}
