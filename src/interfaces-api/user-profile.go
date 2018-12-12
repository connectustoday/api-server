package interfaces_api

type IUserProfileAPI struct {
	Interests        []string `json:"interests"`
	Biography        string   `json:"biography"`
	Education        string   `json:"education"` // TODO
	Quote            string   `json:"quote"`
	CurrentResidence string   `json:"current_residence"`
	Certifications   string   `json:"certifications"` // TODO
	Type             string   `json:"type"`
}
