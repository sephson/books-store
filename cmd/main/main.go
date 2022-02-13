package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sephson/books-store/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}