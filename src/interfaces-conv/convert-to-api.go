package interfaces_conv

import (
	"interfaces-api"
	"interfaces-internal"
)

func ConvertToIAccountAPI(acc interfaces_internal.IAccount) interfaces_api.IAccountAPI {
	return interfaces_api.IAccountAPI{
		UserName: acc.UserName,
		Email: acc.Email,
		Avatar: acc.Avatar,
		Header: acc.Header,
		CreatedAt: acc.CreatedAt,
		Type: acc.Type,
		PostsCount: len(acc.Posts),
		LikedCount: len(acc.Liked),
		SharedCount: len(acc.Shared),
	}
}

func ConvertToIAddressAPI(address interfaces_internal.IAddress) interfaces_api.IAddressAPI {
	return interfaces_api.IAddressAPI{
		Street: address.Street,
		City: address.City,
		Province: address.Province,
		Country: address.Country,
		PostalCode: address.PostalCode,
		AptNumber: address.AptNumber,
		GeoJSON: ConvertToIPointAPI(address.GeoJSON),
	}
}

func ConvertToIExperienceAPI(experience interfaces_internal.IExperience) interfaces_api.IExperienceAPI {
	return interfaces_api.IExperienceAPI{
		Location:     ConvertToIAddressAPI(experience.Location),
		ID:           experience.ID.Hex(),
		Name:         experience.Name,
		Organization: experience.Organization,
		Opportunity:  experience.Opportunity,
		When:         interfaces_api.When{Begin: experience.When.Begin, End: experience.When.End},
		IsVerified:   experience.IsVerified,
		EmailVerify:  experience.EmailVerify,
		CreatedAt:    experience.CreatedAt,
		Hours:        experience.Hours,
	}
}

func ConvertToIPointAPI(point interfaces_internal.IPoint) interfaces_api.IPointAPI {
	return interfaces_api.IPointAPI{
		Type: point.Type,
		Coordinates: point.Coordinates,
	}
}

func ConvertToIValidationsAPI(v interfaces_internal.IValidations) interfaces_api.IValidationsAPI {
	return interfaces_api.IValidationsAPI{
		UserID: v.UserID,
		ExperienceID: v.ExperienceID,
	}
}