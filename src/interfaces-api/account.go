package interfaces_api

import "interfaces-internal"

type IAccountAPI struct {
	UserName string `json:"username"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	Header string `json:"header"`
	CreatedAt int64 `json:"created_at"`
	Type string `json:"type"`
	PostsCount int `json:"posts_count"`
	LikedCount int `json:"liked_count"`
	SharedCount int `json:"shared_count"`
}

func ConvertToIAccountAPI(acc interfaces_internal.IAccount) IAccountAPI {
	return IAccountAPI{
		UserName: acc.UserName,
		Email: acc.Email,
		Avatar: acc.Avatar,
		Header: acc.Header,
		CreatedAt: acc.CreatedAt,
		Type: acc.Type,
		PostsCount: len(acc.Posts),
		LikedCount: len(acc.Liked),
		SharedCount: len(acc.Shared),
	}
}