package interfaces_internal

type IValidations struct {
	UserID       string `bson:"user_id"`
	ExperienceID string `bson:"experience_id"`
}

type IOrganization struct {
	*IAccount
	PreferredName         string               `bson:"preferred_name"`
	IsVerified            bool                 `bson:"is_verified"`
	Opportunities         []string             `bson:"opportunities"`
	OrgInfo               IOrganizationProfile `bson:"org_info"`
	ExperienceValidations []IValidations       `bson:"experience_validations"`
}
