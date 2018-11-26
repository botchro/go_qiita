package qiita

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GetItems get an array of post
func GetItems(page int, perPage int, query string) ([]Post, error) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://qiita.com/api/v2/items?page=%d&per_page=%d", page, perPage)

	if len(query) > 0 {
		url += fmt.Sprintf("&query=%s", query)
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var posts []Post
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
