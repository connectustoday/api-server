package interfaces_api

type When struct {
	Begin string `json:"begin"`
	End   string `json:"end"`
}

type IExperienceAPI struct {
	Location     IAddressAPI `json:"location"`
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Organization string      `json:"organization"`
	Opportunity  string      `json:"opportunity"`
	When         When        `json:"when"`
	IsVerified   bool        `json:"is_verified"`
	EmailVerify  bool        `json:"email_verify"`
	CreatedAt    int64       `json:"created_at"`
	Hours        int64       `json:"hours"`
}
