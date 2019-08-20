package service_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T){
	//Inicialización
	var tweet *domain.Tweet //Variable tweet va a teenr un puntero a domain.Tweet.
	user := "mponsa"
	text := "First tweet"
	tweet = domain.NewTweet(user,text) //Devuelve el puntero a tweet.

	//Operacion
	service.PublishTweet(tweet)

	//Validacion
	var publishedTweet = service.GetTweet()
	if publishedTweet.User != user && publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \n but received: %s:%s",user,text,publishedTweet.User,publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Errorf("Expected date cannot be nil")
	}
}

