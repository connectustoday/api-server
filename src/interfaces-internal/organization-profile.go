package interfaces_internal

type IOrganizationProfile struct {
	SchemaVersion  string   `bson:"schema_version"`
	Mission        string   `bson:"mission"`
	Quote          string   `bson:"quote"`
	Address        IAddress `bson:"address"`
	AffiliatedOrgs []string `bson:"affiliated_orgs"`
	Interests      []string `bson:"interests"`
}
