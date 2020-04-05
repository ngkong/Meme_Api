package reddit

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// EncodeCredentials : Return base64 Encoded client ID and Secret required for authentication
func (r *Reddit) EncodeCredentials() (encodedCredentials string) {
	data := r.ClientID + ":" + r.ClientSecret
	encodedCredentials = base64.StdEncoding.EncodeToString([]byte(data))
	return
}

// MakeGetRequest : Makes a GET Request to Reddit API with Access Token
func (r *Reddit) MakeGetRequest(url string) (responseBody []byte, errorCode int) {
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "Bearer "+r.AccessToken)
	req.Header.Add("User-Agent", r.UserAgent)
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Host", "oauth.reddit.com")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("Error while making request", err)
		return nil, res.StatusCode
	}
	// Close the response body
	defer res.Body.Close()

	// Read the response
	body, _ := ioutil.ReadAll(res.Body)

	return body, res.StatusCode
}

// GetSubredditAPIURL : Returns API Reddit URL with Limit
func GetSubredditAPIURL(subreddit string, limit int) (url string) {
	url = "https://oauth.reddit.com/r/" + subreddit + "/hot?limit=" + strconv.Itoa(limit)
	return
}
