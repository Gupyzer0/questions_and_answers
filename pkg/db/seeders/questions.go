package seeders

import(
	"database/sql"
	"log"

	_"github.com/lib/pq"
)

type questions_seeder struct{
	name string
}

func (s questions_seeder) seed(db *sql.DB) {

	sql := `
		INSERT INTO questions (user_id,answer_id,title,statement) VALUES('1','1','First Question','This is the first question?');
		INSERT INTO questions (user_id,answer_id,title,statement) VALUES('2','2','Second Question','This is the second question?');
		INSERT INTO questions (user_id,answer_id,title,statement) VALUES('2',NULL,'Third Question','Is this really the last question?');
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",s.name,"seeded")
}

