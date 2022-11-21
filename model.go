package twitter

import "time"

type Retweeter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type MetaData struct {
	ResultCount   int    `json:"result_count"`
	NextToken     string `json:"next_token"`
	PreviousToken string `json:"previous_token"`
}

type ReTweeterResult struct {
	Data []Retweeter `json:"data"`
	Meta MetaData    `json:"meta"`
}

type TwitterUserResult struct {
	Data TwitterUser `json:"data"`
}

type TwitterFollowResult struct {
	Data []TwitterUser `json:"data"`
	Meta MetaData      `json:"meta"`
}

type TwitterUser struct {
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	UserName        string        `json:"username"`
	Url             string        `json:"url"`
	ProfileImageUrl string        `json:"profile_image_url"`
	CreatedAt       time.Time     `json:"created_at"`
	PublicMetrics   PublicMetrics `json:"public_metrics"`
}

type PublicMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type TwitterErrDetail struct {
	Detail  string `json:"detail"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type TwitterError struct {
	ErrorDetails []TwitterErrDetail `json:"errors"`
}
