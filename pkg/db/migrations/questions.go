package migrations

import(
	"log"
	"database/sql"
	_"github.com/lib/pq"
)

type questions_migration struct{
	name string
}

func (m questions_migration) migrate_up(db *sql.DB){

	sql := `DROP TABLE IF EXISTS questions;	
	
		CREATE TABLE public.questions (
			id SERIAL PRIMARY KEY,
			user_id integer,
			answer_id integer,
			title character varying,
			statement character varying
	)`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",m.name,"created")
}
