package twitter

import "time"

type Retweeter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type Tweet struct {
	Id                  string        `json:"id"`
	Text                string        `json:"text"`
	CreatedAt           time.Time     `json:"created_at"`
	EditHistoryTweetIds []string      `json:"edit_history_tweet_ids"`
	PublicMetrics       PublicMetrics `json:"public_metrics"`
}

type TweetResult struct {
	Data Tweet `json:"data"`
}

type TweetsResult struct {
	Data []Tweet `json:"data"`
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
	//user
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
	//tweet
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type TwitterErrDetail struct {
	Detail  string `json:"detail"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type TwitterError struct {
	ErrorDetails []TwitterErrDetail `json:"errors"`
}
