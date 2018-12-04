package interfaces_internal

type IOrganizationSettings struct {
	*IAccountSettings
	IsNonprofit bool `bson:"is_nonprofit"`
}
