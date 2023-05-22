package main

import (
	"koifer"
	"koifer/db/memory"
	"log"
	"net/http"
)

var db koifer.UserRepository

func init() {
	db = memory.NewDB()
}

func main() {
	authService := koifer.NewAuthService(db)
	http.Handle("/api/auth", authService)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
