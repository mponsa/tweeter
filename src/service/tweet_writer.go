package service

import "github.com/mponsa/tweeter/src/domain"

type TweetWriter interface {
	SaveTweet(tweet domain.Tweet)
}

