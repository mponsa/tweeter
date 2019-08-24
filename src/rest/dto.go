package rest

type UserDTO struct {
	Name     string `json: "name"`
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

type TweetDTO struct {
	Username    string `json: "name"`
	Text        string `json: "text"`
	ImageURL    string `json: "imageUrl"`
	TweetQuoted string `json: "quoteID"`
}
