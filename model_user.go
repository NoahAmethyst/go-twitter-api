package twitter

import "time"

type Retweeter struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

type ReTweeterResult struct {
	Data []Retweeter `json:"data"`
	Meta MetaData    `json:"meta"`
}

type TwitterUserResult struct {
	Data TwitterUser `json:"data"`
}

type TwitterFollowResult struct {
	Data []TwitterUser `json:"data"`
	Meta MetaData      `json:"meta"`
}

type TwitterUser struct {
	Id              string        `json:"id"`
	Name            string        `json:"name"`
	UserName        string        `json:"username"`
	Url             string        `json:"url"`
	ProfileImageUrl string        `json:"profile_image_url"`
	CreatedAt       time.Time     `json:"created_at"`
	PublicMetrics   PublicMetrics `json:"public_metrics"`
}
