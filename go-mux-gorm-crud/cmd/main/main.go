package main

import (
	"LHXHL/go-mux-gorm-crud/internal/model"
	"LHXHL/go-mux-gorm-crud/internal/router"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func readEnv() {
	err := godotenv.Load("../../deploy/.env")
	if err != nil {
		log.Println("File loading failed...")
	} else {
		log.Println("ReadEnv success!")
	}
}

func main() {
	readEnv()
	model.InitDB()
	r := mux.NewRouter()
	router.RegisterRoutes(r)
	http.Handle("/", r)
	log.Println("listening on port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
