package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const randomUrl = "http://www.thecocktaildb.com/api/json/v1/1/random.php"
const baseSearchUrl = "http://www.thecocktaildb.com/api/json/v1/1/search.php?s="

type Client struct {
	httpClient http.Client
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) Query(drinkName string) (Cocktails, error) {
	var url string
	if drinkName == "" {
		url = randomUrl
	} else {
		url = baseSearchUrl + drinkName
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Cocktails{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Cocktails{}, err
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		return Cocktails{}, fmt.Errorf("status code err: %v", response.StatusCode)
	}

	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return Cocktails{}, err
	}

	var drinkResponse Cocktails
	err = json.Unmarshal(resBytes, &drinkResponse)
	if err != nil {
		return Cocktails{}, err
	}

	return drinkResponse, nil

}
