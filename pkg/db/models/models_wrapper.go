package models 

import(
	"database/sql"
)

type QuestionsInterface interface {
	Index() ([]Question, error)
	Get(question_id string) (*Question, error)
	Create(title,statement,user_id string) (*Question, error)
	Update(question_id, question_title, question_statement string) (*Question, error)
	Delete(question_id string) error
}

type AnswersInterface interface {
	Get(answer_id string) (*Answer, error)
	Update(question_id,statement,user_id string,) (*Answer, error)
}

type UsersInterface interface {
	Index() ([]User, error)
	UserQuestions(user_id string) ([]Question, error)
}

type Models_wrapper struct {
	Questions QuestionsInterface
	Answers AnswersInterface
	Users UsersInterface
}

func InitializeModelsWrapper(db *sql.DB) Models_wrapper {
	wrapper := Models_wrapper{
		Questions: QuestionModel{DB: db},
		Answers: AnswerModel{DB: db},
		Users: UserModel{DB: db},
	}

	return wrapper
}
