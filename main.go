package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/fgruchala/twall-middle-go/webservice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	webservice.NewTweetWebservice(router)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	log.WithFields(log.Fields{"port": "3002"}).Info("Starting server ...")
	log.Debug(http.ListenAndServe(":3002", router))
}
