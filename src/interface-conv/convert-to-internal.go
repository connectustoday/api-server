package interface_conv

import (
	"interfaces-api"
	"interfaces-internal"
)

func ConvertToIAddressInternal(api interfaces_api.IAddressAPI) interfaces_internal.IAddress {
	return interfaces_internal.IAddress{
		SchemaVersion: 0,
		Street: api.Street,
		City: api.City,
		Province: api.Province,
		Country: api.Country,
		PostalCode: api.PostalCode,
		AptNumber: api.AptNumber,
		GeoJSON: ConvertToIPointInternal(api.GeoJSON),
	}
}
