package models 

import(
	"database/sql"
)


type Models_wrapper struct {

	Questions interface {
		Index() ([]Question, error)
		Get(question_id string) (*Question, error)
		Create(question *Question) (*Question, error)
		Update(question_id, question_title, question_statement string) (*Question, error)
		Delete(question_id string) error
	}

	Answers interface {
		Get(answer_id string) (*Answer, error)
		Update(question_id string, answer *Answer) (*Answer, error)
	}

	Users interface {
		Index() ([]User, error)
		UserQuestions(user_id string) ([]Question, error)
	}
}

func InitializeModelsWrapper(db *sql.DB) Models_wrapper {
	wrapper := Models_wrapper{
		Questions: QuestionModel{DB: db},
		Answers: AnswerModel{DB: db},
		Users: UserModel{DB: db},
	}

	return wrapper
}
