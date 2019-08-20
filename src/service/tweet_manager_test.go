package service_test

import (
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	var tweet string = "My first tweet"

	service.PublishTweet(tweet)

	if service.GetTweet() != tweet {
		t.Error("Expected tweet is",tweet)
	}
}

