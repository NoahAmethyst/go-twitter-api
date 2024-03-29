package twitter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	usernameRule = "^[A-Za-z0-9_]{1,15}$"
)

// give a tweet url then return the tweet id
func ParseTweetIdByUrl(tweetUrl string) (int64, error) {
	twitterUrl, err := url.Parse(tweetUrl)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("invalid tweet tweetUrl:%s", tweetUrl))
	}

	v := strings.Split(twitterUrl.Path, "/")

	if len(v) == 0 {
		return 0, errors.New(fmt.Sprintf("invalid tweet tweetUrl:%s", tweetUrl))
	}
	tweetId, err := strconv.ParseInt(v[len(v)-1], 10, 64)
	if err != nil {
		err = errors.New(fmt.Sprintf("invalid tweet tweetUrl:%s", tweetUrl))
	}
	return tweetId, err
}

// check whether user name match the twitter username rule
func CheckUsername(username string) bool {
	match, _ := regexp.MatchString(usernameRule, username)
	return match
}

func (cli *Client) getApi(apiUrl string, data interface{}) error {

	get, err := cli.Get(apiUrl)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(get.Body)

	if get.StatusCode == http.StatusNotFound {
		return errors.New(get.Status)
	}

	body, err := io.ReadAll(get.Body)

	if err != nil {

		return err
	}

	twitterErr := TwitterError{}
	if get.StatusCode != http.StatusOK {
		if err := json.Unmarshal(body, &twitterErr); err != nil {
			return err
		}
		if len(twitterErr.ErrorDetails) > 0 {
			twitterErrDetail := twitterErr.ErrorDetails[0]
			if len(twitterErrDetail.Title) > 0 {
				err = errors.New(twitterErrDetail.Title)
			}
			if len(twitterErrDetail.Message) > 0 {
				err = errors.New(twitterErrDetail.Message)
			}
			if len(twitterErrDetail.Detail) > 0 {
				err = errors.New(twitterErrDetail.Detail)
			}
		} else {
			singleErr := TwitterErrDetail{}
			if err := json.Unmarshal(body, &twitterErr); err != nil {
				return err
			}
			if len(singleErr.Title) > 0 {
				err = errors.New(singleErr.Title)
			}
			if len(singleErr.Detail) > 0 {
				err = errors.New(singleErr.Detail)
			}
		}
		return err
	}

	if strings.Contains(string(body), "errors") {
		if err := json.Unmarshal(body, &twitterErr); err != nil {
			return err
		}
		if len(twitterErr.ErrorDetails) > 0 {
			twitterErrDetail := twitterErr.ErrorDetails[0]
			if len(twitterErrDetail.Title) > 0 {
				err = errors.New(twitterErrDetail.Title)
			}
			if len(twitterErrDetail.Message) > 0 {
				err = errors.New(twitterErrDetail.Message)
			}
			if len(twitterErrDetail.Detail) > 0 {
				err = errors.New(twitterErrDetail.Detail)
			}
		}
	}

	err = json.Unmarshal(body, &data)

	return nil
}
