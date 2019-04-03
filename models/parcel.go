package models

// The Parcel object defines the dimensions and the weight of the parcels or document you are sending.
// A shipment can contain more than just one parcel.
// If all parcels you are sending do have the same dimensions and weight,
// you can use the field quantity so specify how many parcels you are sending.
// If dimension or weight differ, you can send several parcel objects in your request.
type Parcel struct {
	Quantity      string        `json:"quantity"`       // The quantity of parcels with the same dimensions and weight.
	Weight        string        `json:"weight"`         // The weight of the parcel.
	WeightUnit    WeightUnit    `json:"weight_unit"`    // The weight unit used. (kg, lbs)
	Height        string        `json:"height"`         //The height of the parcel.
	Length        string        `json:"length"`         // The length of the parcel.
	Width         string        `json:"width"`          // The width of the parcel.
	DimensionUnit DimensionUnit `json:"dimension_unit"` // The dimension unit used.
}

const (
	// WeightUnitKG weight unit kg
	WeightUnitKG WeightUnit = "kg"
	// WeightUnitLBS weight unit lbs
	WeightUnitLBS WeightUnit = "lbs"
	// DimensionUnitCM centimeters
	DimensionUnitCM DimensionUnit = "cm"
	// DimensionUnitIN inches
	DimensionUnitIN DimensionUnit = "in"
)

// WeightUnit The weight unit used. (kg, lbs)
type WeightUnit string

// DimensionUnit The dimension unit used.
type DimensionUnit string
