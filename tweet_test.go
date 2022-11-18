package twitter

import "testing"

const (
	TestTweetId = 1567778103438766080
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
