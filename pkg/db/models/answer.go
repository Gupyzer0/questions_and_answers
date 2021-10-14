package models

import (
	"database/sql"
	"leonel/prototype_b/pkg/utils"
	//"log"

	"gopkg.in/guregu/null.v4"
)

type Answer struct {
	ID        null.String `json:"id"`
	Statement null.String `json:"statement" validate:"max=255"`
	User      User   `json:"user"`
}

type AnswerModel struct{
	DB *sql.DB
}

func (m AnswerModel) Get(answer_id string) (*Answer, error) {

	query := `SELECT answers.id, answers.statement, users.id, users.username 
				FROM answers
				INNER JOIN users ON answers.user_id = users.id
				WHERE answers.id = $1`

	row := m.DB.QueryRow(query, answer_id)

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

func (m AnswerModel) Update(question_id string, answer *Answer) (*Answer, error) {

	var answ_id string

	err := m.DB.QueryRow("SELECT answer_id FROM questions WHERE id = $1",question_id).Scan(&answ_id)

	if err != nil {
		return nil, err
	}

	var updated_answer *Answer

	row := m.DB.QueryRow("UPDATE answers SET(statement, user_id) = ($2, $3) WHERE id = $1", answ_id, answer.Statement, answer.User.ID)

	if err= row.Err(); err != nil{
		return nil, err
	}

	updated_answer, err = m.Get(answ_id);

	if err != nil {
		return nil, err
	}

	return updated_answer, nil
}
