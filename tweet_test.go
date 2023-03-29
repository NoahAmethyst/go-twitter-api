package twitter

import "testing"

const (
	TestTweetId = 1347569870578266115
)

func TestGetReTweeters(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetReTweeters(TestTweetId, "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweeters)
}

func TestGetTweet(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweet, err := cli.GetTweet(TestTweetId)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweet)
}

func TestGetTweets(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweets, err := cli.GetTweets([]int64{1604578908124450816, 1604862158843645952, 1604862398032056320})
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, _tweet := range tweets.Data {
		t.Logf("%v", _tweet)
	}
}

func TestGetQuoteTweets(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweets, err := cli.GetQuoteTweets(1604578908124450816, false, true, "")
	if err != nil {
		t.Error(err.Error())
		return
	}
	for _, _tweet := range tweets.Data {
		t.Logf("%v", _tweet)
	}
}
