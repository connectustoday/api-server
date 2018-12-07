package interfaces_api

import "interfaces-internal"

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

func ConvertToIExperienceAPI(experience interfaces_internal.IExperience) IExperienceAPI {
	return IExperienceAPI{
		Location:     ConvertToIAddressAPI(experience.Location),
		ID:           experience.ID.String(),
		Name:         experience.Name,
		Organization: experience.Organization,
		Opportunity:  experience.Opportunity,
		When:         When{Begin: experience.When.Begin, End: experience.When.End},
		IsVerified:   experience.IsVerified,
		EmailVerify:  experience.EmailVerify,
		CreatedAt:    experience.CreatedAt,
		Hours:        experience.Hours,
	}
}
