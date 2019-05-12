package models

type Pickup struct {
	Accounts
	Status                          string    `json:"status"`
	PickupDate                      string    `json:"pickup_date"`
	ScheduleFrom                    string    `json:"schedule_from"`
	ScheduleTo                      string    `json:"schedule_to"`
	Carrier                         string    `json:"carrier"`
	ConfirmationNumber              string    `json:"confirmation_number"`
	Direction                       Direction `json:"direction"`
	Shipment                        Shipment  `json:"shipment"`
	PickupUUID                      string    `json:"pickup_uuid"`
	UUID                            string    `json:"uuid"`
	EnviayaPickupUUID               string    `json:"enviaya_pickup_uuid"`
	PickupStatus                    string    `json:"pickup_status"`
	CarrierPickupConfirmationNumber string    `json:"carrier_pickup_confirmation_number"`
}

type Messages struct {
	Message string `json:"message"`
}
