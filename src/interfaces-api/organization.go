package interfaces_api

type IOrganizationAPI struct {
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Header      string `json:"header"`
	CreatedAt   int64  `json:"created_at"`
	Type        string `json:"type"`
	PostsCount  int    `json:"posts_count"`
	LikedCount  int    `json:"liked_count"`
	SharedCount int    `json:"shared_count"`
	// IOrganizationAPI specific fields
	PreferredName         string               `bson:"preferred_name"`
	IsVerified            bool                 `bson:"is_verified"`
	Opportunities         []string             `bson:"opportunities"`
}
