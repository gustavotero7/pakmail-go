package models

type PickupBookingRequest struct {
	Accounts
	APIKey             string    `json:"api_key"`       // YOUR_API_KEY
	Carrier            string    `json:"carrier"`       // The carrier you want to book a pickup with.
	PickupDate         string    `json:"pickup_date"`   // The date you want the pickup to be realized.
	ScheduleFrom       string    `json:"schedule_from"` // The earliest hour you want the carrier to pick up your shipment. Please note that nor do all carriers accept a schedule_from time neither can we guarantee the pickup in the time window provided.
	ScheduleTo         string    `json:"schedule_to"`   // The latest hour you want the carrier to pick up your shipment. Please note that nor do all carriers accept a schedule_to time neither can we guarantee the pickup in the time window provided.
	Direction          Direction `json:"direction"`
	Shipment           Shipment  `json:"shipment"`
	CarrierServiceCode string    `json:"carrier_service_code"`
	EnviayaServiceCode string    `json:"enviaya_service_code"`
	EnviayaID          string    `json:"enviaya_id"`
}

type PickupBookingResponse struct {
	Pickup
	Messages Messages `json:"messages"`
}
