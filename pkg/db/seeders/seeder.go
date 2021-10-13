package seeders

import(
	"database/sql"
	_"github.com/lib/pq"
)

type seeder interface{
	seed(*sql.DB)
}

func DatabaseSeed(db *sql.DB) {

	seeders := []seeder{
		users_seeder{name: "users"},
		answers_seeder{name: "answers"},
		questions_seeder{name: "questions"},
	}

	for _, seeder := range seeders{
		seeder.seed(db)
	}

}
