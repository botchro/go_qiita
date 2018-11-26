package qiita

// Post a qiita post
type Post struct {
	RenderedBody   string `json:"rendered_body"`
	Body           string `json:"body"`
	Coediting      bool   `json:"coediting"`
	CommentsCount  int    `json:"comments_count"`
	CreatedAt      string `json:"created_at"`
	Group          Group  `json:"group"`
	ID             string `json:"id"`
	LikesCount     int    `json:"likes_count"`
	Private        bool   `json:"private"`
	ReactionsCount int    `json:"reactions_count"`
	Tags           Tags   `json:"tags"`
	Title          string `json:"title"`
	UpdatedAt      string `json:"updated_at"`
	URL            string `json:"url"`
	User           User   `json:"user"`
	PageViewsCount int    `json:"page_views_count"`
}
