package gotweet

import (
	"time"
)

// This describes the full possibility of a response type.
type User struct {
	Id       *string `json:"id"`       // DEFAULT
	Name     *string `json:"name"`     // DEFAULT
	Username *string `json:"username"` // DEFAULT
	// To include the following values in the "user.fields" enum list.
	CreatedAt      *time.Time        `json:"created_at"` // DEFAULT date (ISO 8601)
	Protected      *bool             `json:"protected"`
	Withheld       UserWithheld      `json:"withheld"`
	Location       *string           `json:"location"`
	Url            *string           `json:"url"`
	Description    *string           `json:"description"` // aka bio
	Verified       *bool             `json:"verified"`
	Entities       UserEntities      `json:"entities"`
	PublicImageUrl *string           `json:"public_image_url"`
	PublicMetrics  UserPublicMetrics `json:"public_metrics"`
	PinnedTweetId  *string           `json:"pinned_tweet_id"`
}

type UserWithheld struct {
	CountryCodes []string `json:"country_codes"`
	Scope        *string  `json:"scope"`
}

type UserIncludes struct {
	Tweets []Tweet `json:"tweets"`
}

type UserEntities struct {
	Url         UserEntitiesUrl `json:"url"`
	Description UserDescription `json:"description"`
}

type UserDescription struct {
	Urls     []UserUrlExpansion `json:"urls"`
	Hashtags []Hashtag          `json:"hashtags"`
	Mentions []Mention          `json:"mentions"`
	Cashtags []Cashtag          `json:"cashtags"`
}

type UserEntitiesUrl struct {
	Urls []UserUrlExpansion `json:"urls"`
}
type UserUrlExpansion struct {
	Start       int16   `json:"start"`
	End         int16   `json:"end"`
	Url         *string `json:"url"`
	ExpandedUrl *string `json:"expanded_url"`
	DisplayUrl  *string `json:"display_url"`
}
type UserPublicMetrics struct {
	FollowersCount *int `json:"followers_count"`
	FollowingCount *int `json:"following_count"`
	TweetCount     *int `json:"tweet_count"`
	ListedCount    *int `json:"listed_count"`
}
