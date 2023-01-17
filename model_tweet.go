package twitter

import "time"

type Tweet struct {
	Id                  string            `json:"id"`
	Text                string            `json:"text"`
	CreatedAt           time.Time         `json:"created_at"`
	EditHistoryTweetIds []string          `json:"edit_history_tweet_ids"`
	PublicMetrics       PublicMetrics     `json:"public_metrics"`
	AuthorId            string            `json:"author_id"`
	Entities            Entities          `json:"entities"`
	Attachments         Attachments       `json:"attachments"`
	ReferencedTweets    []ReferencedTweet `json:"referenced_tweets"`
}

type TweetResult struct {
	Data Tweet `json:"data"`
}

type TweetsResult struct {
	Data     []Tweet  `json:"data"`
	Includes Includes `json:"includes"`
	MetaData MetaData `json:"meta"`
}

type Includes struct {
	Users []TwitterUser `json:"users"`
	Media []Media       `json:"media"`
}

type Media struct {
	MediaKey string    `json:"media_key"`
	Type     MediaType `json:"type"`
	//photo gif video
	Url string `json:"url"`
}

type ReferencedTweet struct {
	//retweeted quoted replied_to
	Type TweetType `json:"type"`
	Id   string    `json:"id"`
}

type Entities struct {
	Mentions []Mention `json:"mentions"`
	HashTags []HashTag `json:"hashtags"`
}

type HashTag struct {
	Tag string `json:"tag"`
}

type Mention struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}

type Attachments struct {
	MediaKeys []string `json:"media_keys"`
}
