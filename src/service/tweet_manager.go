package service

import (
	"fmt"

	"github.com/mponsa/tweeter/src/domain"
	"github.com/mponsa/tweeter/src/errors"
)

var userManager *UserManager = NewUserManager()

type TweetManager struct {
	tweets      []domain.Tweet
	tweetsUser  map[string][]domain.Tweet
	tweetWriter TweetWriter
}

func NewTweetManager(tweetWriter TweetWriter) *TweetManager {
	tweetManager := new(TweetManager)
	tweetManager.tweets = make([]domain.Tweet, 0)
	tweetManager.tweetsUser = make(map[string][]domain.Tweet)
	tweetManager.tweetWriter = tweetWriter

	return tweetManager
}

func (tweetManager *TweetManager) PublishTweet(tweet domain.Tweet) (int64, error) {
	err := isValidTweet(tweet)
	if err != nil {
		return 0, err
	}

	tweetID := tweetManager.getLastTweetID() + 1
	tweet.SetID(tweetID)

	tweetManager.tweets = append(tweetManager.tweets, tweet)
	tweetManager.tweetsUser[tweet.GetUser().Username] = append(tweetManager.tweetsUser[tweet.GetUser().Username], tweet)
	tweetManager.tweetWriter.SaveTweet(tweet)
	return tweet.GetID(), nil
}

func isValidTweet(tweet domain.Tweet) error {
	if tweet.GetUser() == nil {
		return fmt.Errorf(errors.ERROR_USER_REQUIRED)
	}
	if !userManager.IsLoggedIn(tweet.GetUser().Username) {
		return fmt.Errorf(errors.ERROR_USER_NOT_LOGGED_IN)
	}
	if tweet.GetText() == "" {
		return fmt.Errorf(errors.ERROR_TEXT_REQUIRED)
	}
	if len(tweet.GetText()) > 140 {
		return fmt.Errorf(errors.ERROR_TEXT_LIMIT_EXCEEDED)
	}
	return nil
}

func (tweetManager *TweetManager) getLastTweetID() int64 {
	totalTweets := len(tweetManager.tweets)
	if totalTweets == 0 {
		return 0
	}
	return tweetManager.tweets[totalTweets-1].GetID()
}

func (tweetManager *TweetManager) GetTweets() ([]domain.Tweet, error) {
	if len(tweetManager.tweets) == 0 {
		return nil, fmt.Errorf(errors.ERROR_NO_TWEETS_FOUND)
	}
	return tweetManager.tweets, nil
}

func (tweetManager *TweetManager) GetTweetsByUser(username string) ([]domain.Tweet, error) {
	_, err := userManager.FindUser(username)
	if err != nil {
		return nil, fmt.Errorf(errors.ERROR_USER_NOT_REGISTERED)
	}
	return tweetManager.tweetsUser[username], nil
}

func (tweetManager *TweetManager) GetTweet(id int64) (domain.Tweet, error) {
	for _, tweet := range tweetManager.tweets {
		if tweet != nil {
			if tweet.GetID() == id {
				return tweet, nil
			}
		}
	}
	return nil, fmt.Errorf(errors.ERROR_TWEET_NOT_FOUND)
}

func (tweetManager *TweetManager) DeleteTweet(id int64) error {
	tweet, err := tweetManager.GetTweet(id)
	if err != nil {
		return err
	}
	indexToDeleteFromTweetList := getTweetIndexInList(id, tweetManager.tweets)
	indexToDeleteFromUserTweets := getTweetIndexInList(id, tweetManager.tweetsUser[tweet.GetUser().Username])
	tweetManager.tweets[indexToDeleteFromTweetList] = nil
	tweetManager.tweetsUser[tweet.GetUser().Username][indexToDeleteFromUserTweets] = nil

	return nil
}

func getTweetIndexInList(id int64, tweetList []domain.Tweet) int {
	for index, tweet := range tweetList {
		if tweet.GetID() == id {
			return index
		}
	}
	return -1
}

func (tweetManager *TweetManager) SearchTweetsContaining(query string, searchResult chan domain.Tweet) {
	go tweetManager.tweetWriter.SearchTweetsWithQuery(query, searchResult)
}

func (tweetManager *TweetManager) EditTweet(id int64, newText string) error {
	tweet, err := tweetManager.GetTweet(id)
	if err != nil {
		return err
	}

	tweet.SetText(newText)

	return nil
}
