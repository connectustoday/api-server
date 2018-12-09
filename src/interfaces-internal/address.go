package interfaces_internal

type IAddress struct {
	SchemaVersion uint32 `bson:"schema_version" json:"schema_version"`
	Street        string `bson:"street" json:"street"`
	City          string `bson:"city" json:"city"`
	Province      string `bson:"province" json:"province"`
	Country       string `bson:"country" json:"country"`
	PostalCode    string `bson:"postal_code" json:"postal_code"`
	AptNumber     string `bson:"apt_number" json:"apt_number"`
	GeoJSON       IPoint `bson:"geojson" json:"geojson"`
}