package rest

import (
	"fmt"
	"github.com/mponsa/tweeter/src/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mponsa/tweeter/src/service"
)

const port = ":8080"
var userManager = service.NewUserManager()
var tweetWrite = service.NewMemoryTweetWriter()
var tweetManager = service.NewTweetManager(tweetWrite)

func Run() {
	router := gin.Default()

	router.GET("/", home)
	router.POST("/user", registerUser)
	router.POST("/login", loginUser)
	router.POST("/tweet", publishTweet)
	err := router.Run(port)
	if err != nil {

		fmt.Printf(err.Error())
	}
	fmt.Printf("Application listening at %s", port)
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, "Application Program interface working.")
}

func registerUser(c *gin.Context) {
	userDTO := new(UserDTO)

	if err := c.ShouldBindJSON(userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userCreated := userManager.RegisterUser(userDTO.Name, userDTO.Username, userDTO.Email, userDTO.Password)
	if userCreated != nil {
		c.JSON(http.StatusOK, gin.H{"message": "user registered"})
		return
	}
}

func loginUser(c *gin.Context) {
	userDTO := new(UserDTO)

	if err := c.ShouldBindJSON(userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userManager.LogIn(userDTO.Username, userDTO.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in"})
	return

}

func publishTweet(c *gin.Context){
	tweetDTO := new(TweetDTO)
	if err := c.ShouldBindJSON(tweetDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet := composeTweet(tweetDTO, c)

	id, err := tweetManager.PublishTweet(tweet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tweet published","id":id})
	return
}

func composeTweet(dto *TweetDTO, c *gin.Context) domain.Tweet{
	user, err := userManager.FindUser(dto.Username)

	if(err != nil){
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil
	}

	if dto.ImageURL != "" {
		return domain.NewImageTweet(user,dto.Text,dto.ImageURL)
	}
	if dto.TweetQuoted != "" {
		var tweetQuoted domain.Tweet
		id, err := strconv.ParseInt(dto.TweetQuoted,10, 64)
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return nil
		}

		tweetQuoted, err = tweetManager.GetTweet(int64(id))
		if err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return nil
		}

		return domain.NewQuoteTweet(user,dto.Text,tweetQuoted)
	}
	
	return domain.NewTextTweet(user,dto.Text)
}
