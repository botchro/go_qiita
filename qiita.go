package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {

	var (
		p = flag.Int("p", 1, "page number (default 1)")
		n = flag.Int("n", 20, "number of items get once (default 20)")
		q = flag.String("q", "", "search query string (default empty string)")
	)
	flag.Parse()

	posts, err := GetItems(*p, *n, *q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}
}

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
		fmt.Println(err)
		return nil, err
	}

	return posts, nil
}
