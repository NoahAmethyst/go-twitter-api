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
