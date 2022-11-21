package twitter

import (
	"errors"
	"fmt"
)

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-by-username-username
func (cli *Client) GetTwitterUserByUserName(username string) (TwitterUserResult, error) {
	result := TwitterUserResult{}

	if !CheckUsername(username) {
		return result, errors.New("username does not match ^[A-Za-z0-9_]{1,15}$")
	}

	userLookupByUrl := fmt.Sprintf("https://api.twitter.com/2/users/by/username/%s?user.fields=created_at,url,public_metrics,profile_image_url", username)

	err := cli.callApi(userLookupByUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/users/lookup/api-reference/get-users-id
func (cli *Client) GetTwitterUserById(id string) (TwitterUserResult, error) {
	result := TwitterUserResult{}

	userLookupByUrl := fmt.Sprintf("https://api.twitter.com/2/users/%s?user.fields=created_at,url,public_metrics,profile_image_url", id)

	err := cli.callApi(userLookupByUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-followers
func (cli *Client) GetTwitterFollowers(id string) (TwitterFollowResult, error) {
	result := TwitterFollowResult{}

	followersUrl := fmt.Sprintf("https://api.twitter.com/2/users/%s/followers?max_results=1000", id)

	err := cli.callApi(followersUrl, &result)

	return result, err
}

// https://developer.twitter.com/en/docs/twitter-api/users/follows/api-reference/get-users-id-following
func (cli *Client) GetTwitterFollowing(id string) (TwitterFollowResult, error) {
	result := TwitterFollowResult{}

	followingUrl := fmt.Sprintf("https://api.twitter.com/2/users/%s/following?max_results=1000", id)

	err := cli.callApi(followingUrl, &result)

	return result, err
}
