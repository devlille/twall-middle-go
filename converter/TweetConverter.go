package converter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/model"
)

// TweetsFromTwitterToTwall convert Twitter tweets into Twall tweets
func TweetsFromTwitterToTwall(tweetsToConvert []twitter.Tweet) []model.Tweet {
	var tweetsConverted []model.Tweet

	for _, tweetToConvert := range tweetsToConvert {
		tweetsConverted = append(tweetsConverted, TweetFromTwitterToTwall(tweetToConvert))
	}

	return tweetsConverted
}

// TweetFromTwitterToTwall convert a Twitter tweet into a Twall tweet
func TweetFromTwitterToTwall(tweetToConvert twitter.Tweet) model.Tweet {
	return model.Tweet{
		ID:           tweetToConvert.ID,
		CreatedAt:    tweetToConvert.CreatedAt,
		User:         &model.User{ScreenName: tweetToConvert.User.ScreenName},
		Text:         tweetToConvert.Text,
		RetweetCount: tweetToConvert.RetweetCount}
}
