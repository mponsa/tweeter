package domain

import "time"

type ImageTweet struct {
	TextTweet
	url string
}

func NewImageTweet(user *User, text string, url string) *ImageTweet {
	date := time.Now()

	tweet := ImageTweet{
		TextTweet{
			1,
			user,
			text,
			&date,
		},
		url,
	}

	return &tweet
}

func (imageTweet *ImageTweet) PrintableTweet() string {
	return "@" + imageTweet.User.Username + ": " + imageTweet.Text + imageTweet.url
}

