package twitter

import (
	"fmt"
)

// https://developer.twitter.com/en/docs/twitter-api/tweets/retweets/api-reference/get-tweets-id-retweeted_by
func (cli *Client) GetReTweeters(id int64, nextPageToken string) (ReTweeterResult, error) {
	result := ReTweeterResult{}

	retweetedByUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%d/retweeted_by?max_results=100", id)

	if len(nextPageToken) > 0 {
		retweetedByUrl += fmt.Sprintf("&pagination_token=%s", nextPageToken)
	}
	err := cli.getApi(retweetedByUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets-id
func (cli *Client) GetTweet(id int64) (TweetResult, error) {
	result := TweetResult{}

	tweetUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%d?tweet.fields=created_at", id)

	err := cli.getApi(tweetUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets#tab2
//func (cli *Client) GetTweets(ids []int64) {
//	//https://api.twitter.com/2/tweets
//	result := TweetsResult{}
//	tweetsUrl := fmt.Sprintf("https://api.twitter.com/2/tweets?tweet.fields=created_at")
//
//}
