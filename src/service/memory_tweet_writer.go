package service

import (
	"github.com/mponsa/tweeter/src/domain"
	"strings"
)

type MemoryTweetWriter struct{
	tweets     []domain.Tweet
	tweetsUser map[string][]domain.Tweet
}

func NewMemoryTweetWriter() TweetWriter{
	tweetWriter := new(MemoryTweetWriter)
	tweetWriter.tweets = make([]domain.Tweet,0)
	tweetWriter.tweetsUser = make(map[string][]domain.Tweet)

	return tweetWriter
}

func (tweetWriter *MemoryTweetWriter) SaveTweet(tweet domain.Tweet){
	tweetWriter.tweets = append(tweetWriter.tweets,tweet)
	tweetWriter.tweetsUser[tweet.GetUser().Username] = append(tweetWriter.tweetsUser[tweet.GetUser().Username],tweet)
}

func (tweetWriter *MemoryTweetWriter) GetLastSavedTweet() domain.Tweet{
	if(len(tweetWriter.tweets) > 0){
		return(tweetWriter.tweets[len(tweetWriter.tweets) -1])
	}

	return nil
}

func (tweetWriter *MemoryTweetWriter) SearchTweetsWithQuery(query string, searchResult chan domain.Tweet){
	for _, tweet := range tweetWriter.tweets{
		if(strings.Contains(tweet.GetText(),query)){
			searchResult <- tweet
		}
	}
}
