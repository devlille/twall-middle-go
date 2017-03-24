package main

import (
	"net/http"

	"log"

	"github.com/fgruchala/twall-middle-go/webservice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	webservice.NewTweetWebservice(router)

	log.Fatal(http.ListenAndServe(":3002", router))
}
