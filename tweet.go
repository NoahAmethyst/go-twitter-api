package twitter

import (
	"fmt"
	"strconv"
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

	tweetUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%d?tweet.fields=created_at,public_metrics", id)

	err := cli.getApi(tweetUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/tweets/lookup/api-reference/get-tweets#tab2
func (cli *Client) GetTweets(ids []int64) (TweetsResult, error) {
	result := TweetsResult{}

	var idsParam string

	for i, _id := range ids {
		if i < len(ids)-1 {
			idsParam += fmt.Sprintf("%s,", strconv.FormatInt(_id, 10))
		} else {
			idsParam += strconv.FormatInt(_id, 10)
		}
	}

	tweetsUrl := fmt.Sprintf("https://api.twitter.com/2/tweets?ids=%s&tweet.fields=created_at,public_metrics", idsParam)

	err := cli.getApi(tweetsUrl, &result)

	return result, err

}

func (cli *Client) GetQuoteTweets(id int64, needRetweets, needReplies bool, page string) (TweetsResult, error) {
	result := TweetsResult{}
	var exclude string

	if !needRetweets && !needReplies {
		exclude = "retweets,replies"
	} else {
		if needRetweets && !needReplies {
			exclude = "replies"
		} else if !needRetweets && needReplies {
			exclude = "retweets"
		}
	}

	tweetsUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%d/quote_tweets?max_results=100"+
		"&expansions=attachments.media_keys,author_id,entities.mentions.username"+
		"&tweet.fields=public_metrics,entities,referenced_tweets"+
		"&user.fields=username"+
		"&media.fields=url,media_key,type", id)

	if len(exclude) > 0 {
		tweetsUrl += fmt.Sprintf("&exclude=%s", exclude)
	}

	if len(page) > 0 {
		tweetsUrl += fmt.Sprintf("&pagination_token=%s", page)
	}
	err := cli.getApi(tweetsUrl, &result)

	return result, err
}
