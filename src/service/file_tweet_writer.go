package service

import (
	"github.com/mponsa/tweeter/src/domain"
	"os"
)

type FileTweetWriter struct{
	file *os.File
}

func NewFileTweetWriter() TweetWriter{
	file , _ := os.OpenFile(
		"tweets.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)

	tweetWriter := new (FileTweetWriter)
	tweetWriter.file = file

	return tweetWriter
}

func (tweetWriter *FileTweetWriter) SaveTweet(tweet domain.Tweet){
	go writeTweetInFile(tweet,tweetWriter.file)
}

func writeTweetInFile(tweet domain.Tweet,file *os.File){
	if(file != nil){
		byteSlice := []byte(tweet.Print() + "\n")
		file.Write(byteSlice)
	}
}

func (tweetWriter *FileTweetWriter) SearchTweetsWithQuery(query string, searchResult chan domain.Tweet){

}