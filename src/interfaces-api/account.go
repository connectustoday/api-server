package interfaces_api

type IAccountAPI struct {
	ID string `json:"id"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Header string `json:"header"`
	CreatedAt int64 `json:"created_at"`
	Type string `json:"type"`
	PostsCount int `json:"posts_count"`
	LikedCount int `json:"liked_count"`
	SharedCount int `json:"shared_count"`
}