package db

import(
	"fmt"
	"database/sql"
	_"github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "docker"
	password = "123456"
	dbname = "questions_and_answers"
)

func Db_connect() *sql.DB{
	
	conn_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres",conn_info)

	if err != nil{
		panic(err)
	}

	return db

}
