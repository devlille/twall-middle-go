package service

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/converter"
	"github.com/fgruchala/twall-middle-go/model"
)

var twitterClient *twitter.Client

// NewTweetService initialize a connection with the Twitter API
func NewTweetService(key, secret *string) {
	/*config := oauth1.NewConfig(key, secret)
	token := oauth1.NewToken("accessToken", "accessSecret")
	httpClient := config.Client(oauth1.NoContext, token)*/

	twitterClient = twitter.NewClient(httpClient)
}

// Search over the Twitter API
func Search(params *twitter.SearchTweetParams) ([]model.Tweet, error) {
	search, resp, err := twitterClient.Search.Tweets(params)

	if err != nil {
		return make([]model.Tweet, 0), err
	}

	return converter.TweetsFromTwitterToTwall(search.Statuses), nil
}
