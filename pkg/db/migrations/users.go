package migrations

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type users_migration struct{
	name string
}

func (m users_migration) migrate_up(db *sql.DB){
	sql := `DROP TABLE IF EXISTS users;
	
		CREATE TABLE public.users (
			id SERIAL PRIMARY KEY,
			username character varying(255)
	)`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",m.name,"created")
}