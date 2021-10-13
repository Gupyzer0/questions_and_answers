package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	db "leonel/prototype_b/pkg/db"
	services "leonel/prototype_b/pkg/services"
	transport "leonel/prototype_b/pkg/transport"
	migrations "leonel/prototype_b/pkg/db/migrations"
	seeders "leonel/prototype_b/pkg/db/seeders"
)

var migrate = flag.Bool("migrate",false,"Run migrations")
var seed = flag.Bool("seed",false,"Seed the database")

/*TODO
implement logging middleware
response codes
go kit handlers middlewares
*/

func main(){

	Db_conn := db.Db_connect()

	err := Db_conn.Ping()

	if err != nil{
		panic(err)
	} else {
		log.Println("db connection success")
	}

	flag.Parse()

	if *migrate == true{
		log.Println("Migrating . . .")
		migrations.DatabaseMigrate(Db_conn)
	}

	if *seed == true{
		log.Println("Seeding . . .")
		seeders.DatabaseSeed(Db_conn)
	}
	
	srv := services.NewQuestionsAndAnswersService(Db_conn)
	
	router := transport.MakeHttpHandler(srv)

	server := &http.Server{
		Handler: router,
		Addr: ":4000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
