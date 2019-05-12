package models

type PickupLookupRequest struct {
	Pickup
	Parcels     []Parcel      `json:"parcels"`
	ShipmentIDs []interface{} `json:"shipment_ids"`
}
