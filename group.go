package main

// Group a qiita Group
type Group struct {
	CreatedAt string `json:"created_at"`
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Private   bool   `json:"private"`
	UpdatedAt string `json:"updated_at"`
	URLName   string `json:"url_name"`
}
