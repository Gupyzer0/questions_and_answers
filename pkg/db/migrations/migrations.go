package migrations

import(
	"database/sql"
	_"github.com/lib/pq"
)

type migrations interface{
	migrate_up(*sql.DB)
}

func DatabaseMigrate(db *sql.DB) {

	migrations := []migrations{
		users_migration{name: "users"},
		answers_migration{name: "answers"},
		questions_migration{name: "questions"},
	}

	for _, migration := range migrations{
		migration.migrate_up(db)
	}

}
