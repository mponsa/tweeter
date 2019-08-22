package domain_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestTextTweetPrintsUserAndText(t *testing.T){
	user := domain.NewUser("Hola","blablaasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	tweet := domain.NewTextTweet(user,"My tweet")

	text := tweet.PrintableTweet()

	expectedText := "@blablaasd: My tweet"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s",expectedText,text)
	}
}
