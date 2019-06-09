package main

import (
	"database/sql"
	"fmt"
	"github.com/gojektech/proctor/proctord/logger"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	r.HandleFunc("/bird", updateBirdHandler).Methods("PUT", "PATCH")

	staticFileDirectory := http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	fmt.Println("Starting server...")
	connString := "dbname=bird_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	logger.Info("Initialize database connection")

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	logger.Info("Serving on port 8080")

	http.ListenAndServe(":8080", r)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
