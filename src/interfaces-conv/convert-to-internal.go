package interfaces_conv

import (
	"interfaces-api"
	"interfaces-internal"
)

func ConvertToIAddressInternal(api interfaces_api.IAddressAPI) interfaces_internal.IAddress {
	return interfaces_internal.IAddress{
		SchemaVersion: 0,
		Street:        api.Street,
		City:          api.City,
		Province:      api.Province,
		Country:       api.Country,
		PostalCode:    api.PostalCode,
		AptNumber:     api.AptNumber,
		GeoJSON:       ConvertToIPointInternal(api.GeoJSON),
	}
}

func ConvertToIPointInternal(api interfaces_api.IPointAPI) interfaces_internal.IPoint {
	return interfaces_internal.IPoint{
		Type:        api.Type,
		Coordinates: api.Coordinates,
	}
}
