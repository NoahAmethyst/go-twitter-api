package twitter

import (
	"fmt"
)

func (cli *Client) GetReTweeters(id int64, nextPageToken string) (ReTweeterResult, error) {
	result := ReTweeterResult{}

	retweetedByUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%d/retweeted_by?max_results=100", id)

	if len(nextPageToken) > 0 {
		retweetedByUrl += fmt.Sprintf("&pagination_token=%s", nextPageToken)
	}
	err := cli.callApi(retweetedByUrl, &result)

	return result, err
}
