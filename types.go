package gotweet
// Any shared types will be in this file.
// These are fragment types that are shared across collections.
//

// type Mentions []Mention
type Mention struct {
	Start    int16   `json:"start"`
	End      int16   `json:"end"`
	Username *string `json:"username"`
}

// type Hashtags []Hashtag
type Hashtag struct {
	Start   int16   `json:"start"`
	End     int16   `json:"end"`
	Hashtag *string `json:"tag"`
}

// type Cashtags []Cashtag
type Cashtag struct {
	Start   int16   `json:"start"`
	End     int16   `json:"end"`
	Cashtag *string `json:"tag"`
}

type Page struct {
	ResultCount   int     `json:"result_count"`
	PreviousToken *string `json:"previous_token"`
	NextToken     *string `json:"next_token"`
}

type Errors []map[string]string

type TweetLite struct {
	Id       *string `json:"id"`
	Text     *string `json:"text"`
	AuthorId *string `json:"author_id"`
}
