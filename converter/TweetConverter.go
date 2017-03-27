package converter

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/fgruchala/twall-middle-go/model"
	log "github.com/sirupsen/logrus"
)

// TweetsFromTwitterToTwall convert Twitter tweets into Twall tweets
func TweetsFromTwitterToTwall(tweetsToConvert []twitter.Tweet) []model.Tweet {
	log.WithField("count", len(tweetsToConvert)).Info("[TweetConverter] Converting tweet(s) from Twitter API to Twall ...")

	tweetsConverted := make([]model.Tweet, 0)

	for _, tweetToConvert := range tweetsToConvert {
		tweetsConverted = append(tweetsConverted, TweetFromTwitterToTwall(tweetToConvert))
	}

	log.WithField("count", len(tweetsConverted)).Info("[TweetConverter] Tweet(s) converted successfully.")
	return tweetsConverted
}

// TweetFromTwitterToTwall convert a Twitter tweet into a Twall tweet
func TweetFromTwitterToTwall(tweetToConvert twitter.Tweet) model.Tweet {
	log.WithField("tweet", tweetToConvert).Debug("[TweetConverter] Tweet to convert")

	tweetConverted := model.Tweet{
		ID:           tweetToConvert.ID,
		CreatedAt:    tweetToConvert.CreatedAt,
		User:         &model.User{ScreenName: tweetToConvert.User.ScreenName},
		Text:         tweetToConvert.Text,
		RetweetCount: tweetToConvert.RetweetCount}

	log.WithField("tweet", tweetConverted).Debug("[TweetConverter] Tweet converted")
	return tweetConverted
}
