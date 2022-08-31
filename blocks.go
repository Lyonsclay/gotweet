// This file contains types and data used to construct twitter api requests and responses.
//

package gotweet

import (
	"fmt"
	// "time"
)

type Option string
type Options []Option
type Expansions Options
type TweetFields Options
type UserFields Options

// [TODO] Seems like something a url method should handle.
func (opts Expansions) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	// Don't return last comma.
	return j[:len(j)-1]
}
func (opts TweetFields) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	return j[:len(j)-1]
}
func (opts UserFields) Join() string {
	var j string
	for _, o := range opts {
		j += fmt.Sprintf("%s,", o)
	}
	return j[:len(j)-1]
}

var MaxExpansions = Expansions{"pinned_tweet_id"}

var ExcludedTweetFields = []string{
	"non_public_metrics",
	"organic_metrics",
	"promoted_metrics",
}

var IncludedTweetFields = TweetFields{
	"attachments",
	"author_id",
	"context_annotations",
	"conversation_id",
	"created_at",
	"entities",
	"geo",
	"id",
	"in_reply_to_user_id",
	"lang",
	"public_metrics",
	"possibly_sensitive",
	"referenced_tweets",
	"reply_settings",
	"source",
	"text",
	"withheld",
}

var MaxTweetFields = TweetFields{
	"attachments",
	"author_id",
	"context_annotations",
	"conversation_id",
	"created_at",
	"entities",
	"geo",
	"id",
	"in_reply_to_user_id",
	"lang",
	"public_metrics",
	"possibly_sensitive",
	"referenced_tweets",
	"reply_settings",
	"source",
	"text",
	"withheld",
}

var MaxUserFields = UserFields{
	"created_at",
	"description",
	"entities",
	"id",
	"location",
	"name",
	"pinned_tweet_id",
	"profile_image_url",
	"protected",
	"public_metrics",
	"url",
	"username",
	"verified",
	"withheld",
}

type Blocks struct {
	UserId       string
	Data         []User `json:"data"`
	UserIncludes `json:"includes"`
	Meta         Page    `json:"meta"`
	Errors       []Error `json:"errors"`
}

type Error struct {
	Value        string `json:"value"`
	Detail       string `json:"detail"`
	Title        string `json:"title"`
	ResourceType string `json:"resource_type"`
	Parameter    string `json:"parameter"`
	ResourceId   string `json:"resource_id"`
	Type         string `json:"type"`
}
