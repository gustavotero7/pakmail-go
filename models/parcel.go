package models

// The Parcel object defines the dimensions and the weight of the parcels or document you are sending.
// A shipment can contain more than just one parcel.
// If all parcels you are sending do have the same dimensions and weight,
// you can use the field quantity so specify how many parcels you are sending.
// If dimension or weight differ, you can send several parcel objects in your request.
type Parcel struct {
	Quantity      float64       `json:"quantity,omitempty"`       // The quantity of parcels with the same dimensions and weight.
	Weight        float64       `json:"weight,omitempty"`         // The weight of the parcel.
	WeightUnit    WeightUnit    `json:"weight_unit,omitempty"`    // The weight unit used. (kg, lbs)
	Height        float64       `json:"height,omitempty"`         //The height of the parcel.
	Length        float64       `json:"length,omitempty"`         // The length of the parcel.
	Width         float64       `json:"width,omitempty"`          // The width of the parcel.
	DimensionUnit DimensionUnit `json:"dimension_unit,omitempty"` // The dimension unit used.
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
