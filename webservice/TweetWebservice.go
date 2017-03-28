package webservice

import (
	"encoding/json"
	"net/http"

	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/service"
	"github.com/gorilla/mux"
)

// TweetWebservice define a tweet webservice
type TweetWebservice struct {
	Routes  *mux.Router
	Service *service.TweetService
}

// NewTweetWebservice initialize the twall API for tweet
func NewTweetWebservice(router *mux.Router, tweetService *service.TweetService) *TweetWebservice {
	webservice := TweetWebservice{
		Routes:  router.PathPrefix("/api/tweet").Subrouter(),
		Service: tweetService}

	webservice.Routes.HandleFunc("/", webservice.GetAll)

	return &webservice
}

// GetAll return all tweets
func (tws *TweetWebservice) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Info("[TweetWebservice] Searching tweet(s) ...")

	since := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	params := &twitter.SearchTweetParams{
		Query:      "#devfestlille -filter:retweets since:" + since,
		Count:      30,
		ResultType: "recent"}
	tweets, err := tws.Service.Search(params)

	if err != nil {
		log.WithField("error", err).Error("[TweetWebservice] Failed to get tweet(s)")
		http.Error(w, "Failed to get tweet(s)", http.StatusInternalServerError)
	} else {
		log.WithField("count", len(tweets)).Info("[TweetWebservice] Tweet(s) returned.")
		json.NewEncoder(w).Encode(tweets)
	}
}
