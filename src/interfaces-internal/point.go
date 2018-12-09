package interfaces_internal

import "interfaces-api"

type IPoint struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

func ConvertToIPointInternal(api interfaces_api.IPointAPI) IPoint {
	return IPoint{
		Type: api.Type,
		Coordinates: api.Coordinates,
	}
}
