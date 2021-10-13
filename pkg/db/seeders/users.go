package seeders

import(
	"database/sql"
	"log"

	_"github.com/lib/pq"
)

type users_seeder struct{
	name string
}

func (s users_seeder) seed(db *sql.DB) {

	sql := `
		INSERT INTO users (username) VALUES('Leonel');
		INSERT INTO users (username) VALUES('Eduardo');
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}

	log.Println("table",s.name,"seeded")
}
