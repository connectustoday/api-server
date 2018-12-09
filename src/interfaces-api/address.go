package interfaces_api

import (
	"encoding/json"
)

type IAddressAPI struct {
	Street        string    `json:"street"`
	City          string    `json:"city"`
	Province      string    `json:"province"`
	Country       string    `json:"country"`
	PostalCode    string    `json:"postal_code"`
	AptNumber     string    `json:"apt_number"`
	GeoJSON       IPointAPI `json:"geojson"`
}

func (address *IAddressAPI) UnmarshalJSON(data []byte) error {
	println(string(data))
	var v []string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	} // TODO in general
	address.Street = v[1]
	address.City = v[2]
	address.Province = v[3]
	address.Country = v[4]
	address.PostalCode = v[5]
	address.AptNumber = v[6]

	return nil
}