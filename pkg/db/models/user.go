package models

import(
	"database/sql"

	"gopkg.in/guregu/null.v4"
	utils "leonel/prototype_b/pkg/utils"
)

type User struct {
	ID null.String `json:"id"`
	Username null.String `json:"username" validate:"max=30"`
}

type UserModel struct{
	DB *sql.DB
}

//Index all users, not needed on the test, just for fun
func (m UserModel) Index() ([]User, error) {
	
	rows, err := m.DB.Query("SELECT * FROM users")
	
	defer rows.Close()

	if err != nil{

		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, err
		}
	}

	var users []User

	for rows.Next() {

		var user User

		err := rows.Scan(&user.ID ,&user.Username)

		if err!= nil {
			return nil, err
		}

		users = append(users, user)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}


func (m UserModel) UserQuestions(user_id string) ([]Question, error){

	questions := make([]Question,0)

	query := `SELECT questions.id, questions.title, questions.statement, questions.user_id, users.username, 
				answers.id AS answer_id, answers.statement AS answer_statement, ua.id AS answer_user_id , ua.username AS answer_user_username
				
				FROM questions 
					LEFT OUTER JOIN users ON users.id = questions.user_id 
					LEFT OUTER JOIN answers ON answer_id = answers.id
					LEFT OUTER JOIN users ua ON ua.id = answers.user_id

				WHERE questions.user_id = $1`

	rows, err := m.DB.Query(query, user_id)

	defer rows.Close()

	if err != nil{

		switch err {
		case sql.ErrNoRows:
			return nil, utils.ErrNotFound
		default:
			return nil, err
		}
	}

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
