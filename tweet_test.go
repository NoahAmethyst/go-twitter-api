package twitter

import "testing"

const (
	TestTweetId = 1347569870578266115
)

func TestGetReTweeters(t *testing.T) {
	//set your consumeKey & consumerSecret
	twitterConsumeKey := ""
	twitterConsumerSecret := ""

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetReTweeters(TestTweetId, "")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweeters)
}

func TestGetTweet(t *testing.T) {
	/**
	TWITTER_CONSUMER_KEY=hXPFra6AxNN5cEnuyWo7q5hUb
	TWITTER_CONSUMER_SECRET=S63n5QPU9Y9sIni0zP97d9KhZLkgxR8rC9bIvAi5eVwwUXjlXr
	*/
	twitterConsumeKey := "hXPFra6AxNN5cEnuyWo7q5hUb"
	twitterConsumerSecret := "S63n5QPU9Y9sIni0zP97d9KhZLkgxR8rC9bIvAi5eVwwUXjlXr"

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweet, err := cli.GetTweet(TestTweetId)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweet)
}
