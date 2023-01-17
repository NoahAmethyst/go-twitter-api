package twitter

type MetaData struct {
	ResultCount   int    `json:"result_count"`
	NextToken     string `json:"next_token"`
	PreviousToken string `json:"previous_token"`
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
