package seeders

import(
	"database/sql"
	"log"

	_"github.com/lib/pq"
)

type answers_seeder struct{
	name string
}

func (s answers_seeder) seed(db *sql.DB) {

	sql := `
		INSERT INTO answers (user_id, statement) VALUES('1','First answer statement');
		INSERT INTO answers (user_id, statement) VALUES('2','Second answer statement');
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",s.name,"seeded")

}
