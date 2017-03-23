package model

import "fmt"

// Tweet define a tweet
type Tweet struct {
	ID           int64  `json:id`
	CreatedAt    string `json:created_at`
	User         *User  `json:user`
	Text         string `json:text`
	RetweetCount int    `json:retweet_count`
}

// String stringify the TwallTweet struct
func (t *Tweet) String() string {
	return fmt.Sprintf("TwallTweet : Id %v, User %v, CreatedAt %v", t.ID, t.User, t.CreatedAt)
}
