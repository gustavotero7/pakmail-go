package models

type TrackingRequest struct {
	Accounts
	APIKey  string `json:"api_key,omitempty"` // YOUR_API_KEY
	Carrier string `json:"carrier"`           // The carrier you want to track. If you create your shipments with us, you recommend to use the Pak Mail Via 02 MTY IDs of your shipments for tracking.
	// shipment_number also known as enviaya_shipment_number (in this context)
	ShipmentNumber string `json:"shipment_number"`     // The shipment number or Pak Mail Via 02 MTY ID of your shipment.
	Reference      string `json:"reference,omitempty"` // If you provide one or more reference(s) when creating your shipment, you can use these for tracking instead of a shipment number.
}

type TrackingResponse struct {
	Status                string `json:"status"`
	EnviayaShipmentNumber string `json:"enviaya_shipment_number"`
	CarrierTrackingNumber string `json:"carrier_tracking_number"`
	EstimatedDeliveryDate string `json:"estimated_delivery_date"`
	ExpectedDeliveryDate  string `json:"expected_delivery_date"`
	// Unlike the estimated delivery date, the expected delivery date considers all existing shipment events and dynamically calculates the probable date of delivery.
	// The main difference between the estimated delivery date and the expected delivery date is, that the estimated delivery is returned during rating and label creation and is never updated, whereas the expected delivery date will be updated over time.
	// If you want to know the realistic date of arrival of a shipment, we recommend to consult the expected_delivery_date.
	DeliveryDate        string      `json:"delivery_date"`
	PickupDate          string      `json:"pickup_date"`
	ShipmentStatus      string      `json:"shipment_status"`
	EventCode           float64     `json:"event_code"`
	EventDescription    string      `json:"event_description"`
	Event               string      `json:"event"`
	StatusCode          float64     `json:"status_code"`
	SubEventCode        interface{} `json:"sub_event_code"`
	SubEvent            interface{} `json:"sub_event"`
	SubEventDescription interface{} `json:"sub_event_description"`
	Checkpoints         interface{} `json:"checkpoints"`
}
