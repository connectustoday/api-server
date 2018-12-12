package interfaces_api

type IOrganizationProfileAPI struct {
	Mission        string      `json:"mission"`
	Quote          string      `json:"quote"`
	Address        IAddressAPI `json:"address"`
	AffiliatedOrgs []string    `json:"affiliated_orgs"`
	Interests      []string    `json:"interests"`
	Type           string      `json:"type"`
}
