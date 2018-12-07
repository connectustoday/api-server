package interfaces_api

type IAddressAPI struct {
	SchemaVersion uint32    `json:"schema_version"`
	Street        string    `json:"street"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	AptNumber     string    `json:"apt_number"`
	GeoJSON       IPointAPI `json:"geojson"`
}
