package main

import (
	"flag"
	"fmt"

	"github.com/botchro/qiita"
)

func main() {

	var (
		p = flag.Int("p", 1, "page number (default 1)")
		n = flag.Int("n", 20, "number of items get once (default 20)")
		q = flag.String("q", "", "search query string (default empty string)")
	)
	flag.Parse()

	posts, err := qiita.GetItems(*p, *n, *q)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}
}
