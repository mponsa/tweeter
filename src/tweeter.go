package main

import (
	"github.com/abiosoft/ishell"
	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/service"
)

var tweetWriter = service.NewFileTweetWriter()
var tweetManager = service.NewTweetMananager(tweetWriter)
var userManager = service.NewUserManager()

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")


	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Enter username: ")

			user, _ := userManager.FindUser(c.ReadLine())

			c.Print("Write your tweet: ")

			tweet := c.ReadLine()


			_, err := tweetManager.PublishTweet(domain.NewTextTweet(user,tweet))

			if err != nil{
				c.Print(err.Error())
			}
			c.Print("TextTweet sent\n")

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showLastestTweet",
		Help: "Shows a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			fetched_tweets, err := tweetManager.GetTweets()

			if err != nil {
				c.Printf(err.Error())

				return
			}

			for _, tweet := range fetched_tweets {
				c.Printf(tweet.Print())
			}


			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "Register",
		Help: "Create a new user to use tweeter",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print(" Enter your name: ")
			name := c.ReadLine()
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter your email: ")
			email := c.ReadLine()
			c.Print("Enter your password: ")
			password := c.ReadLine()

			registeredUser := userManager.RegisterUser(name, user, email, password)

			if registeredUser != nil {
				c.Printf(" %s registered! ", registeredUser.Username)
			}
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "Login",
		Help: "Logs in",
		Func: func(c *ishell.Context) {
			defer c.ShowPrompt(true)
			c.Print("Enter your username: ")
			user := c.ReadLine()
			c.Print("Enter your password: ")
			password := c.ReadLine()

			err := userManager.LogIn(user, password)

			if err != nil {
				c.Print(err.Error())
			}
			return
		},
	})

	shell.Run()

}
