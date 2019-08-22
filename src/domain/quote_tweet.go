package domain

import "time"

type QuoteTweet struct {
	TextTweet
	quote *TextTweet
}

func NewQuoteTweet(user *User, text string, quotedText *TextTweet) *QuoteTweet{
	date := time.Now()

	tweet := QuoteTweet{
		TextTweet{
			1,
			user,
			text,
			&date,
		},
		quotedText,
	}

	return &tweet
}

func (tweet *QuoteTweet) PrintableTweet() string{
	return "@"+tweet.GetUser().Username+": "+tweet.GetText()+` "`+tweet.quote.PrintableTweet()+`"`
}