package domain_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"testing"
)

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	user := domain.NewUser("Hola","grupoesfera","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	tweet := domain.NewImageTweet(user, "This is my image",
		"http://www.grupoesfera.com.ar/common/img/grupoesfera.png")
	// Operation
	text := tweet.PrintableTweet()
	// Validation
	expectedText := "@grupoesfera: This is my image" +
	"http://www.grupoesfera.com.ar/common/img/grupoesfera.png"
	if text != expectedText {
		t.Errorf("The expected text is %s but was %s",expectedText,text)
	}

}

