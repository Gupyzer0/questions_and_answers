package models

import(
	"database/sql"

	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID null.String `json:"id"`
	Username null.String `json:"username" validate:"max=30"`
}

func UserQuestions(db *sql.DB, user_id string) ([]Question, error){

	questions := make([]Question,0)

	rows, err := db.Query(`SELECT questions.id, questions.title, questions.statement, questions.user_id, users.username, 
	answers.id AS answer_id, answers.statement AS answer_statement, ua.id AS answer_user_id , ua.username AS answer_user_username
	
	FROM questions 
		LEFT OUTER JOIN users ON users.id = questions.user_id 
		LEFT OUTER JOIN answers ON answer_id = answers.id
		LEFT OUTER JOIN users ua ON ua.id = answers.user_id

	WHERE questions.user_id = $1`, user_id)

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
