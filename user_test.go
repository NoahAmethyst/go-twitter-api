package twitter

import "testing"

const (
	TestUsername = "realDonaldTrump"
	TestUserId   = "25073877"
)

func TestGetUser(t *testing.T) {
	//set your consumeKey & consumerSecret
	twitterConsumeKey := ""
	twitterConsumerSecret := ""

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterUserByUserName(TestUsername)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweeters)
}

func TestGetFollowers(t *testing.T) {
	//set your consumeKey & consumerSecret
	twitterConsumeKey := ""
	twitterConsumerSecret := ""

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterFollowers(TestUserId)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("result:%v", tweeters)
}

func TestGetFollowings(t *testing.T) {
	//set your consumeKey & consumerSecret
	twitterConsumeKey := ""
	twitterConsumerSecret := ""

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterFollowing(TestUserId)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("result:%v", tweeters)
}
