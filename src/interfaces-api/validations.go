package interfaces_api

import "interfaces-internal"

type IValidationsAPI struct {
	UserID string `json:"user_id"`
	ExperienceID string `json:"experience_id"`
}

func ConvertToIValidationsAPI(v interfaces_internal.IValidations) IValidationsAPI {
	return IValidationsAPI{
		UserID: v.UserID,
		ExperienceID: v.ExperienceID,
	}
}