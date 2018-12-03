package interfaces_internal

type IAddress struct {
	SchemaVersion uint32 `bson:"schema_version"`
	Street string `bson:"street"`
	City string `bson:"city"`
	Province string `bson:"province"`
	Country string `bson:"country"`
	PostalCode string `bson:"postal_code"`
	AptNumber string `bson:"apt_number"`
	GeoJSON IPoint `bson:"geojson"`
}