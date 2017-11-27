package REST

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	getEndpoints(router)
	//router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	//router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getEndpoints(router *mux.Router) {
	router.HandleFunc("/account/{id}", GetAccountEndpoint).Methods("GET")
}
