// gotweet Oauth implements authentication with the twitter api.
//
// SessionUser contains the id and credentials of a twitter user.
// This establishes a user context for user specific data requests.
//
// App contains the application consumer key and secret.
//
// GetTwitterToken currently utilizes OAuth 2.0 as a confidential client.
// https://developer.twitter.com/en/docs/authentication/oauth-2-0/user-access-token
package gotweet

import (
	// "bufio"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"errors"
)

type App struct {
	ConsumerKey    string
	ConsumerSecret string
}

type SessionUser struct {
	Id           string
	AccessToken  string
	AccessSecret string
}

type OAuth struct {
	App
	User            SessionUser
	Time            string
	Header          string
	Nonce           string
	Signature       string
	SignatureMethod string // "HMAC-SHA1"
	Version         string // "1.0"
}

type AccessToken struct {
	Token string `json:"access_token"`
}

// Init registers an oauth user and the app that will access Twitter on their behalf.
func (auth *OAuth) Init(u SessionUser) error {
	now := fmt.Sprintf("%v", time.Now().Unix())
	nonce := randStr(16)

	auth.Version = "1.0"
	auth.SignatureMethod = "HMAC-SHA1"
	auth.Time = now
	auth.Nonce = nonce

	// App access data is stored in the OAuth struct.
	var exists bool
	auth.ConsumerKey, exists = os.LookupEnv("TWITTER_API_KEY")
	if !exists {
		return errors.New("you need to provide a twitter api_key in an environment variable\n \"API_KEY\"")
	}
	auth.ConsumerSecret, exists = os.LookupEnv("TWITTER_API_SECRET_KEY")
	if !exists {
		return errors.New("you need to provide a twitter api_secret_key in an environment variable\n \"API_SECRET_KEY\"")
	}

	// The SessionUser is embeded in the OAuth struct.
	auth.User = u

	return error(nil)
}

func urlencode(s string) string {
	s = url.QueryEscape(s)
	s = strings.ReplaceAll(s, "+", "%20")
	return s
}

func randStr(len int) string {
	buff := make([]byte, len)
	if _, err := rand.Read(buff); err != nil {
		fmt.Println(err)
	}
	str := base64.StdEncoding.EncodeToString(buff)
	return base64.StdEncoding.EncodeToString([]byte(str))[:len]
}

func hmacSha1(base, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(base))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func (auth *OAuth) genSignature(method string, endpoint *url.URL) {
	params := endpoint.Query()
	vals := url.Values{}
	for k, v := range params {
		vals.Add(k, v[0])
	}
	vals.Add("oauth_nonce", auth.Nonce)
	vals.Add("oauth_consumer_key", auth.ConsumerKey)
	vals.Add("oauth_signature_method", auth.SignatureMethod)
	vals.Add("oauth_timestamp", auth.Time)
	vals.Add("oauth_token", auth.User.AccessToken)
	vals.Add("oauth_version", auth.Version)

	m := strings.ToUpper(method)
	ep := strings.Split(endpoint.String(), "?")[0]
	u := urlencode(ep)
	p := urlencode(vals.Encode())

	base := fmt.Sprintf("%s&%s&%s", m, u, p)
	key := fmt.Sprintf("%s&%s", auth.ConsumerSecret, auth.User.AccessSecret)

	auth.Signature = urlencode(hmacSha1(base, key))
}

func (auth *OAuth) setHeaderParams(k, v string) {
	if auth.Header == "" {
		auth.Header = fmt.Sprintf("OAuth %s=\"%s\"", k, v)
	} else {
		auth.Header += fmt.Sprintf(", %s=\"%s\"", k, v)
	}
}

func (auth OAuth) GenOAuthHeader(method string, endpoint *url.URL) string {
	auth.genSignature(method, endpoint)

	auth.setHeaderParams("oauth_consumer_key", auth.ConsumerKey)
	auth.setHeaderParams("oauth_nonce", auth.Nonce)
	auth.setHeaderParams("oauth_signature", auth.Signature)
	auth.setHeaderParams("oauth_signature_method", auth.SignatureMethod)
	auth.setHeaderParams("oauth_timestamp", auth.Time)
	auth.setHeaderParams("oauth_token", auth.User.AccessToken)
	auth.setHeaderParams("oauth_version", auth.Version)

	return auth.Header
}


// [TODO] update to api version 2
// https://developer.twitter.com/en/docs/authentication/oauth-2-0/user-access-token
func (auth OAuth) GetTwitterToken() (string, error) {
	creds := fmt.Sprintf("%s:%s", auth.App.ConsumerKey, auth.App.ConsumerSecret)
	token := base64.StdEncoding.EncodeToString([]byte(creds))
	base := "https://api.twitter.com/oauth2/token"
	method := "POST"

	req, err := http.NewRequest(method, base, nil)
	if err != nil {
		fmt.Println("GetTwitterToken failed to form request: ", err)
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	q := req.URL.Query()
	q.Add("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp.Status)
		return "", errors.New(resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
		return "", err
	}

	var res AccessToken
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Println(err)
	}
	// [TODO] The token should be embedded into the auth header.
	return res.Token, error(nil)
}
