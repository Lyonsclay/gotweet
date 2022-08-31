package gotweet

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	base = "https://api.twitter.com/2"
)

func GetUrl(method string) string {
	var paths = map[string]string{"users_blocking": "/users/%s/blocking"}
	return base + paths[method]
}

// ## Not Included ##
// _ Spaces
// _ Lists
// _ Compliance

type AuthMethod int

const (
	Any AuthMethod = iota
	OAuth1
	OAuth2
)

type TwitterRequest struct {
	Path        string
	Description string
	AuthMethod  string // enum of types
	// Collection  string // enum of endpoint collections
	// [TODO] Should this be part of the description.
	// Action      string // ex. Mangage Likes or Likes Lookup. Describes one or more endpoints.
	BaseUrl     string
	Method      string
	QueryParams url.Values
	PathParams  []interface{}
	Body        string
}

func ConstructUrl() {

}

// OAuth 2.0 scopes required by this endpoint
// tweet.read
// users.read
// block.read
//

var	UsersBlockingRequest = TwitterRequest{Path: "/2/users/blocking"}

// [TODO] OAuth should contain user id
func (auth OAuth) GetUsersBlocking() ([]byte, error) {
	base := fmt.Sprintf(GetUrl("users_blocking"), auth.User.Id)
	endpoint, err := url.Parse(base)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	method := "GET"
	// Query params
	params := url.Values{}
	params.Add("expansions", MaxExpansions.Join())
	params.Add("tweet.fields", MaxTweetFields.Join())
	params.Add("user.fields", MaxUserFields.Join())
	params.Add("max_results", "1000")
	endpoint.RawQuery = params.Encode()
	header := auth.GenOAuthHeader(method, endpoint)
	req, _ := http.NewRequest(method, endpoint.String(), nil)
	req.Header.Add("Authorization", header)
	q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	// fmt.Println(resp.Status)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func PullOpenApiDoc() error {
	url := "https://api.twitter.com/2/openapi.json"
	client := http.Client{Timeout: time.Second * 2}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("../openapi/twitter.openapi.json", jsonPrettyPrint(body), 0644)

	return err
}
