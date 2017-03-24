package webservice

import (
	"encoding/json"
	"net/http"

	"time"

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
func NewTweetWebservice(router *mux.Router) *TweetWebservice {
	webservice := TweetWebservice{
		Routes:  router.PathPrefix("/api/tweet").Subrouter(),
		Service: service.NewTweetService()}

	webservice.Routes.HandleFunc("/", webservice.GetAll)

	return &webservice
}

// GetAll return all tweets
func (tws *TweetWebservice) GetAll(w http.ResponseWriter, r *http.Request) {
	since := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	params := &twitter.SearchTweetParams{
		Query:      "#devfestlille -filter:retweets since:" + since,
		Count:      30,
		ResultType: "recent"}
	tweets, err := tws.Service.Search(params)

	if err != nil {
		http.Error(w, "Failed to get tweets", http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(tweets)
	}
}
