package main

import (
	"backend/internal/db"
	"backend/internal/handler/userHandler"
	"backend/internal/middleware"
	"backend/internal/repository/user"
	"log"
	"net/http"
)

func main(){
	db,err := db.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	userQueries := user.NewQueries(db)
	userHandler := userhandler.NewUserHandler(userQueries)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /reg",userHandler.Register)
	mux.HandleFunc("POST /login",userHandler.Login)

	serve := http.Server{
		Addr: ":8080",
		Handler: middleware.Logging(mux),
	}


	log.Println("Server listening on port 8080")
	serve.ListenAndServe()
}
