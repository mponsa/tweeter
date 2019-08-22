package domain_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	user := domain.NewUser("Hola","grupoesfera","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	anotherUser := domain.NewUser("Hola","nick","dasdawawd","asdasdasd")
	userManager.RegisterUser(anotherUser.Name,anotherUser.Username,anotherUser.Email,anotherUser.Password)
	userManager.LogIn(anotherUser.Username,anotherUser.Password)
	quotedTweet := domain.NewTextTweet(user, "This is my tweet")
	tweet := domain.NewQuoteTweet(anotherUser, "Awesome", quotedTweet)

	text := tweet.PrintableTweet()

	expectedText := `@nick: Awesome "@grupoesfera: This is my tweet"`
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s",expectedText,text)
	}
}
