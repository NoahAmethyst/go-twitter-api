package twitter

type TweetType string
type MediaType string

const (
	//media type
	PHOTO MediaType = "photo"
	Gif   MediaType = "gif"
	Video MediaType = "video"

	//tweet type
	Retweeted TweetType = "retweeted"
	Quoted    TweetType = "quoted"
	Reply     TweetType = "replied_to"
)
