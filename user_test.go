package twitter

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

const (
	TestUsername = "realDonaldTrump"
	TestUserId   = "25073877"
)

func TestGetUser(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterUserByUserName(TestUsername)
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("result:%v", tweeters)
}

func TestGetFollowers(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterFollowers(TestUserId, "")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("result:%v", tweeters)
}

func TestGetFollowings(t *testing.T) {
	twitterConsumeKey, twitterConsumerSecret := GetTwitterApiConfig()

	cli := NewClient(twitterConsumeKey, twitterConsumerSecret)

	tweeters, err := cli.GetTwitterFollowing(TestUserId, "")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("result:%v", tweeters)
}

func GetTwitterApiConfig() (string, string) {

	// set path of yaml
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config from yml failed:%s", err))
	}

	// read value from yaml
	return viper.GetString("twitter.consumer.key"), viper.GetString("twitter.consumer.secret")
}
