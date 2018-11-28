package main

import (
	"flag"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/botchro/qiita"
)

func main() {

	var (
		p = flag.Int("p", 1, "page number (default 1)")
		n = flag.Int("n", 20, "number of items get once (default 20)")
		q = flag.String("q", "", "search query string (default empty string)")
		s = flag.Int("s", 20, "number of items show in one view")
	)
	flag.Parse()

	for {
		// get first items
		posts, err := qiita.GetItems(*p, *n, *q)
		if err != nil {
			fmt.Println(err)
			return
		}

		items := []qiita.Post{
			qiita.Post{Title: fmt.Sprintf("Next Page #%d", *p)},
			qiita.Post{Title: "Bye"},
		}

		prompt := promptui.Select{
			Label: fmt.Sprintf("Choose an item (#%d)", *p),
			Items: append(posts, items...),
			Templates: &promptui.SelectTemplates{
				Label:    "{{ . }}?",
				Active:   "> {{ .Title | cyan }}",
				Inactive: "  {{ .Title | cyan }}",
				Selected: "{{ .Title | red }}",
			},
			Size: *s,
		}

		i, _, err := prompt.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		// next page
		if i == len(posts) {
			*p++
			continue
		}

		// bye
		if i > len(posts) {
			return
		}

		// show post body
		fmt.Println(posts[i].Body)
	}

}
