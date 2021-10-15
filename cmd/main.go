package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	db "leonel/prototype_b/pkg/db"
	//models "leonel/prototype_b/pkg/db/models"
	migrations "leonel/prototype_b/pkg/db/migrations"
	"leonel/prototype_b/pkg/db/models"
	seeders "leonel/prototype_b/pkg/db/seeders"
	services "leonel/prototype_b/pkg/services"
	transport "leonel/prototype_b/pkg/transport"
)

var migrate = flag.Bool("migrate",false,"Run migrations")
var seed = flag.Bool("seed",false,"Seed the database")

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
	
	models_wrapper := models.InitializeModelsWrapper(Db_conn)

	srv := services.NewQuestionsAndAnswersService(&models_wrapper)
	
	router := transport.MakeHttpHandler(srv)

	server := &http.Server{
		Handler: router,
		Addr: ":4000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
