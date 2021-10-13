package migrations

import(
	"log"
	"database/sql"
	_"github.com/lib/pq"
)

type answers_migration struct{
	name string
}

func (m answers_migration) migrate_up(db *sql.DB){
	sql := `DROP TABLE IF EXISTS answers;

			CREATE TABLE public.answers (
				id SERIAL PRIMARY KEY,
				user_id integer,
				statement character varying
	)`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",m.name,"created")
}
