package interfaces_api

import "interfaces-internal"

type IAddressAPI struct {
	Street        string    `json:"street"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	AptNumber     string    `json:"apt_number"`
	GeoJSON       IPointAPI `json:"geojson"`
}

func ConvertToIAddressAPI(address interfaces_internal.IAddress) IAddressAPI {
	return IAddressAPI{
		Street: address.Street,
		City: address.City,
		Province: address.Province,
		Country: address.Country,
		PostalCode: address.PostalCode,
		AptNumber: address.AptNumber,
		GeoJSON: ConvertToIPointAPI(address.GeoJSON),
	}
}