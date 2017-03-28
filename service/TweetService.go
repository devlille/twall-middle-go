package service

import (
	"context"

	"golang.org/x/oauth2/clientcredentials"

	log "github.com/sirupsen/logrus"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/converter"
	"github.com/fgruchala/twall-middle-go/model"
)

// TweetService define a tweet service
type TweetService struct {
	Client *twitter.Client
}

// NewTweetService initialize a connection with the Twitter API
func NewTweetService(consumerKey *string, consumerSecret *string) *TweetService {
	config := &clientcredentials.Config{
		ClientID:     *consumerKey,
		ClientSecret: *consumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token"}
	_, err := config.Token(context.TODO())

	if err != nil {
		log.Fatal("[TweetService] Failed to fetch Twitter API.")
	}

	httpClient := config.Client(context.TODO())
	return &TweetService{Client: twitter.NewClient(httpClient)}
}

// Search over the Twitter API
func (ts *TweetService) Search(params *twitter.SearchTweetParams) ([]model.Tweet, error) {
	log.Info("[TweetService] Searching tweet(s) over Twitter API ...")
	log.WithField("params", params).Debug("[TweetService]")

	search, _, err := ts.Client.Search.Tweets(params)

	if err != nil {
		log.WithField("error", err).Error("[TweetService] Failed to search tweet(s) over Twitter API.")
		return make([]model.Tweet, 0), err
	}

	log.WithField("count", len(search.Statuses)).Info("[TweetService] Tweet(s) returned by Twitter API.")
	return converter.TweetsFromTwitterToTwall(search.Statuses), nil
}
