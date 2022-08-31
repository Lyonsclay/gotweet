package gotweet

import (
	"fmt"
	"time"
)


// The Tweet struct models the full potential data object when using the /tweet
// endpoint of the Twitter api.


type MediaFields Options
type PollFields Options
type PlaceFields Options

func (opts MediaFields) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	return j[:len(j)-1]
}

func (opts PlaceFields) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	return j[:len(j)-1]
}

func (opts PollFields) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	return j[:len(j)-1]
}

var MaxTweetsExpansions = Expansions{"attachments.poll_ids", "attachments.media_keys", "author_id", "entities.mentions.username", "geo.place_id", "in_reply_to_user_id", "referenced_tweets.id", "referenced_tweets.id.author_id"}
var MaxTweetsMediaFields = MediaFields{"duration_ms", "height", "media_key", "preview_image_url", "type", "url", "width", "public_metrics", "non_public_metrics", "organic_metrics", "promoted_metrics", "alt_text"}
var MaxTweetsPlaceFields = PlaceFields{"contained_within", "country", "country_code", "full_name", "geo", "id", "name", "place_type"}
var MaxTweetsPollFields = PollFields{"duration_minutes", "end_datetime", "id", "options", "voting_status"}
var MaxTweetsTweetFields = TweetFields{"attachments", "author_id", "context_annotations", "conversation_id", "created_at", "entities", "geo", "id", "in_reply_to_user_id", "lang", "non_public_metrics", "public_metrics", "organic_metrics", "promoted_metrics", "possibly_sensitive", "referenced_tweets", "reply_settings", "source", "text", "withheld"}
var MaxTweetsUserFields = UserFields{"created_at", "description", "entities", "id", "location", "name", "pinned_tweet_id", "profile_image_url", "protected", "public_metrics", "url", "username", "verified", "withheld"}

type TweetsRequestParams struct {
	Expansions
	MediaFields
	PlaceFields
	PollFields
	TweetFields
	UserFields
	MaxResults      int
	PaginationToken string
}

type Tweet struct {
	Id                string                  `json:"id"`         // DEFAULT
	Text              string                  `json:"text"`       // DEFAULT
	CreatedAt         string                  `json:"created_at"` // date (ISO 8601)
	AuthorId          string                  `json:"author_id"`
	ConversationId    string                  `json:"conversation_id"`
	InReplyToUserId   string                  `json:"in_reply_to_user_id"`
	ReferencedTweets  TweetReferencedTweets   `json:"referenced_tweets"`
	Attachments       TweetAttachments        `json:"attachments"` // object - Specifies the type of attachments (if any) present in this Tweet.
	Geo               TweetGeo                `json:"geo"`
	Annotations       TweetContextAnnotations `json:"context_annotations"`
	Entities          TweetEntities           `json:"entities"`
	Withheld          TweetWithheld           `json:"withheld"`
	PublicMetrics     TweetPublicMetrics      `json:"public_metrics"`
	PossiblySensitive bool                    `json:"possibly_sensitive"`
	Lang              string                  `json:"lang"`
	ReplySettings     TweetReplySettings      `json:"reply_settings"`
	Source            string                  `json:"source"`

	Location       string `json:"location"`
	Url            string `json:"url"`
	Description    string `json:"description"` // aka bio
	Verified       bool   `json:"verified"`
	PublicImageUrl string `json:"public_image_url"`
	PinnedTweetId  string `json:"pinned_tweet_id"`
}

type TweetReferencedTweet struct {
	Type TweetType `json:"type"`
	Id   string    `json:"id"`
}

type TweetReferencedTweets []TweetReferencedTweet

type TweetType Option

var TweetTypeOptions = []TweetType{"retweeted", "quoted", "replied_to"}

type TweetAttachments struct {
	MediaKeys []string `json:"media_keys"`
	PollIds   []string `json:"poll_ids"`
}

type TweetGeo struct {
	Coordinates
}

type Coordinates struct {
	Type        Option       `json:"type"` // only value is "Point"
	Coordinates []Coordinate `json:"coordinates"`
	PlaceId     string       `json:"place_id"`
}

type Coordinate [2]float64

type TweetContextAnnotations []TweetContextAnnotation

type TweetContextAnnotation struct {
	Domain TweetContextAnnotationsDomain `json:"domain"`
	Entity TweetContextAnnotationsEntity `json:"entity"`
}

type TweetContextAnnotationsEntity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetContextAnnotationsDomain struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TweetEntities struct {
	Annotations []TweetEntitiesAnnotation `json:"annotations"`
	Urls        []TweetEntitiesUrl        `json:"urls"`
	Hashtags    []Hashtag                 `json:"hashtags"`
	Mentions    []Mention                 `json:"mentions"`
	Cashtags    []Cashtag                 `json:"cashtags"`
}

type TweetEntitiesAnnotation struct {
	Start          int     `json:"start"`
	End            int     `json:"end"`
	Probability    float32 `json:"probability"`
	Type           string  `json:"type"`
	NormalizedText string  `json:"normalized_text"`
}

type TweetEntitiesUrl struct {
	Start       int    `json:"start"`
	End         int    `json:"end"`
	Url         string `json:"url"`
	ExpandedUrl string `json:"expanded_url"`
	DisplayUrl  string `json:"display_url"`
	UnwoundUrl  string `json:"unwound_url"`
}

type TweetWithheld struct {
	Copyright    bool       `json:"copyright"`
	CountryCodes []string   `json:"country_codes"`
	Scope        TweetScope `json:"scope"`
}

type TweetScope Option // (tweet, user)

type TweetPublicMetrics struct {
	RetweetCount int `json:"retweet_count"`
	ReplyCount   int `json:"reply_count"`
	LikeCount    int `json:"like_count"`
	QuoteCount   int `json:"quote_count"`
}

type TweetReplySettings Option // (everyone, mentionedUsers, following)

type Place struct {
	FullName        string      `json:"full_name"`
	Id              string      `json:"id"`
	ContainedWithin []string    `json:"contained_within"` // array of 'id'
	Country         string      `json:"country"`
	CountryCode     string      `json:"country_code"` // ISO Alpha-2 country code
	Geo             interface{} `json:"geo"`
	Name            string      `json:"name"`
	PlaceType       string      `json:"place_type"`
}

type TweetIncludes struct {
	Tweets []Tweet `json:"tweets"`
	Users  []User  `json:"users"`
	Places []Place `json:"places"`
	Media  []Media `json:"media"`
	Polls  []Poll  `json:"polls"`
}

type Media struct {
	MediaKey        string             `json:"media_key"` // Default
	Type            string             `json:"type"`      // Default
	DurationMs      int                `json:"duration_ms"`
	Height          int                `json:"height"`
	PreviewImageUrl string             `json:"preview_image_url"`
	PublicMetrics   TweetPublicMetrics `json:"public_metrics"`
	Width           int                `json:"width"`
	AltText         string             `json:"alt_text"`
}

type Poll struct {
	Id              string        `json:"id"`      // Default
	Options         []interface{} `json:"options"` // Default
	DurationMinutes int           `json:"duration_minutes"`
	EndDatetime     time.Time     `json:"end_datetime"`
	VotingStatus    string        `json:"voting_status"`
}
