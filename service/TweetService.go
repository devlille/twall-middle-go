package service

import (
	"context"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/converter"
	"github.com/fgruchala/twall-middle-go/model"
	"golang.org/x/oauth2"
)

// TweetService define a tweet service
type TweetService struct {
	Client *twitter.Client
}

// NewTweetService initialize a connection with the Twitter API
func NewTweetService() *TweetService {
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: "715300702298447872-ogVmHS8fX10SyxTeYRk3C9QViecTtar"}
	httpClient := config.Client(context.Background(), token)

	return &TweetService{Client: twitter.NewClient(httpClient)}
}

// Search over the Twitter API
func (ts *TweetService) Search(params *twitter.SearchTweetParams) ([]model.Tweet, error) {
	search, _, err := ts.Client.Search.Tweets(params)

	if err != nil {
		return make([]model.Tweet, 0), err
	}

	return converter.TweetsFromTwitterToTwall(search.Statuses), nil
}
