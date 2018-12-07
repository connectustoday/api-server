package interfaces_api

import "interfaces-internal"

type IPointAPI struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

func ConvertToIPointAPI(point interfaces_internal.IPoint) IPointAPI {
	return IPointAPI{
		Type: point.Type,
		Coordinates: point.Coordinates,
	}
}