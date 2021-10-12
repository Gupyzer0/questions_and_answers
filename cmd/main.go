package main

import (
	"log"
	"net/http"
	"time"

	db "leonel/prototype_b/pkg/db"
	services "leonel/prototype_b/pkg/services"
	transport "leonel/prototype_b/pkg/transport"
)

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
