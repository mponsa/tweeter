package domain

import "time"

type TextTweet struct {
	ID   int64
	User *User
	Text string
	Date *time.Time
}

func NewTextTweet(user *User, text string) *TextTweet {
	date := time.Now()

	tweet := TextTweet{
		1,
		user,
		text,
		&date,
	}

	return &tweet
}

func (tweet *TextTweet) PrintableTweet() string {
	return "@" + tweet.User.Username + ": " + tweet.Text
}

func (tweet *TextTweet) Print() string {
	return tweet.PrintableTweet()
}

func (tweet *TextTweet) GetUser() *User {
	return tweet.User
}

func (tweet *TextTweet) GetText() string {
	return tweet.Text
}

func (tweet *TextTweet) GetID() int64 {
	return tweet.ID
}

func (tweet *TextTweet) SetID(id int64) {
	tweet.ID = id
}

func (tweet *TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (tweet *TextTweet) SetText(text string) {
	tweet.Text = text
}
