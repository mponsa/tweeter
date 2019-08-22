package service_test

import (
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
	"strings"
	"testing"
)



func TestPublishedTweetIsSaved(t *testing.T){
	//Inicialización
	var tweet domain.Tweet //Variable tweet va a teenr un puntero a domain.TextTweet.
	user := domain.NewUser("Hola","blabla","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "First tweet"
	tweet = domain.NewTextTweet(user,text) //Devuelve el puntero a tweet.
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)

	//Operacion
	tweetManager.PublishTweet(tweet)

	//Validacion
	tweets, _ := tweetManager.GetTweets()
	var publishedTweet = tweets[len(tweets) - 1]
	if publishedTweet.GetUser() != user && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \n but received: %s:%s",user,text,publishedTweet.GetUser(),publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Errorf("Expected date cannot be nil")
	}
}

func TestTweetWithoutUserIsNotPublished(t *testing.T){
	var tweet *domain.TextTweet;
	user := domain.NewUser("Hola","","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "This is my first tweet"
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweet = domain.NewTextTweet(user,text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T){
	var tweet *domain.TextTweet;
	user := domain.NewUser("Hola","blabla","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	var text string;
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweet = domain.NewTextTweet(user,text)

	//Operation
	var err error
	_,err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T){
	var tweet *domain.TextTweet;
	user := domain.NewUser("Hola","blabla","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	var text string = "aaaaaaaaakakakakakakakakakakakjsahfdkjadshkjsdhfkjaslhdfljkashdfkajshdfkjashdfkljashdfaskjdfhaskdjfhaskjdfhasksadjfhaskjdhfakshdfaksjhdfkasjhdfkjashdfkjhasdkfjhaskdjhfaksjhdfkajshdflkajshdflkjahsdfkajsdflkjhsalfkjhaskdjfhaskdhfkasdjhfkasjhdfksfkshfd";
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweet = domain.NewTextTweet(user,text)

	//Operation
	var err error
	_,err = tweetManager.PublishTweet(tweet)

	if err != nil && err.Error() != "text shouldn't be longer than 140 characters" {
		t.Error("Expected error is text shouldn't be longer than 140 characters")
	}
}

func TestPublishMultipleTweets(t *testing.T){
	user := domain.NewUser("Holaasdasd","blablaasdsad","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	first_tweet := domain.NewTextTweet(user,"hola")
	second_tweet := domain.NewTextTweet(user, "hola k ace")
	third_tweet := domain.NewTextTweet(user, "hola k acessss")
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweetManager.PublishTweet(first_tweet)
	tweetManager.PublishTweet(second_tweet)
	tweetManager.PublishTweet(third_tweet)

	tweets_fetched, _ := tweetManager.GetTweetsByUser(user.Username)
	if(len(tweets_fetched) != 3){
		t.Errorf("Should fetch %d tweets but received %d",3,len(tweets_fetched))
	}
}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T){
	//Initialization
	var tweet, second_tweet, third_tweet *domain.TextTweet
	user := domain.NewUser("Hola","blablaasdasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	another_user := domain.NewUser("Holaadasd","blablaasdawds","dasdawawd","asdasdasd")
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user,secondText)
	second_tweet = domain.NewTextTweet(user,text)
	third_tweet = domain.NewTextTweet(another_user,secondText)
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(second_tweet)
	tweetManager.PublishTweet(third_tweet)

	//Operation
	tweets, _ := tweetManager.GetTweetsByUser(user.Username)

	//Validation
	if len(tweets) != 2 {
		t.Errorf("Lenght should be 2 but got %d",len(tweets))
	}
	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]
	if firstPublishedTweet.GetText() != secondText || secondPublishedTweet.GetText() != text {
		t.Errorf("Tweets not valid")
	}
}

func TestTweetWithUserNotLoggedInIsNotPublished(t *testing.T){
	//Initialization
	var tweet *domain.TextTweet;
	var user *domain.User;
	user = domain.NewUser("Hola","blabla","dasdawawd","asdasdasd")
	text := "My first tweet"
	tweet = domain.NewTextTweet(user,text);
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation.
	if err != nil {
		t.Errorf("Expected user not logged in but got %s",err.Error())
	}
}

func TestCanRetrieveTweetById(t *testing.T){
	//Inicialización
	var tweet *domain.TextTweet //Variable tweet va a teenr un puntero a domain.TextTweet.
	user := domain.NewUser("Hola","blablaasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "First tweet"
	tweet = domain.NewTextTweet(user,text) //Devuelve el puntero a tweet.
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	//Operacion
	id, _ := tweetManager.PublishTweet(tweet)

	//Validacion
	var publishedTweet, err = tweetManager.GetTweet(id)
	if(err != nil){
		t.Errorf(err.Error())
	}
	if publishedTweet.GetUser() != user && publishedTweet.GetText() != text {
		t.Errorf("Expected tweet is %s: %s \n but received: %s:%s",user,text,publishedTweet.GetUser(),publishedTweet.GetText())
	}
	if publishedTweet.GetDate() == nil {
		t.Errorf("Expected date cannot be nil")
	}

}

func TestCanDeleteTweet(t *testing.T){
	var tweet *domain.TextTweet //Variable tweet va a teenr un puntero a domain.TextTweet.
	user := domain.NewUser("Hola","blablaasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "First tweet"
	tweet = domain.NewTextTweet(user,text) //Devuelve el puntero a tweet.
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)
	id, _ := tweetManager.PublishTweet(tweet)

	tweetManager.DeleteTweet(id)

	expectedError := "tweet not found"
	_, err := tweetManager.GetTweet(id)
	if err != nil && err.Error() != expectedError{
		t.Errorf("The expected error is %s but was %s",expectedError,err.Error())
	}

}

func TestPublishedTweetIsSavedToExternalResource(t *testing.T){
	user := domain.NewUser("Hola","blablaasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "First tweet"
	tweet := domain.NewTextTweet(user,text) //Devuelve el puntero a tweet.
	tweetWriter := service.NewMemoryTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)

	id, _ := tweetManager.PublishTweet(tweet)

	memoryWriter := (tweetWriter).(*service.MemoryTweetWriter)
	savedTweet := memoryWriter.GetLastSavedTweet()

	if savedTweet == nil {

	}
	if savedTweet.GetID() != id {

	}
}

func TestWriterCanWriteAFile(t *testing.T){
	user := domain.NewUser("Hola","blablaasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "First tweet"
	tweet := domain.NewTextTweet(user,text) //Devuelve el puntero a tweet.
	tweetWriter := service.NewFileTweetWriter() //MockImplementation
	tweetManager := service.NewTweetManager(tweetWriter)

	tweetManager.PublishTweet(tweet)
}

func TestCanSearchForTweetContainingText(t *testing.T){
	var tweetWriter service.TweetWriter
	tweetWriter = service.NewMemoryTweetWriter()
	tweetManager := service.NewTweetManager(tweetWriter)
	// Create and publish a tweet
	var tweet, second_tweet, third_tweet *domain.TextTweet
	user := domain.NewUser("Hola","blablaasdasd","dasdawawd","asdasdasd")
	userManager := service.NewUserManager()
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	another_user := domain.NewUser("Holaadasd","blablaasdawds","dasdawawd","asdasdasd")
	userManager.RegisterUser(user.Name,user.Username,user.Email,user.Password)
	userManager.LogIn(user.Username,user.Password)
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user,secondText)
	second_tweet = domain.NewTextTweet(user,text)
	third_tweet = domain.NewTextTweet(another_user,text)
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(second_tweet)
	tweetManager.PublishTweet(third_tweet)

	// Operation
	searchResult := make(chan domain.Tweet)
	query := "first"
	tweetManager.SearchTweetsContaining(query, searchResult)

	// Validation
	foundTweet := <-searchResult

	if foundTweet == nil {
		t.Error("Haven't fetched any tweets")
	}
	if !strings.Contains(foundTweet.GetText(), query) {
		t.Errorf("Expected to find a tweet with %s but got %s",query,foundTweet.GetText())
	}
}

