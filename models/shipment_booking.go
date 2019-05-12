package models

type ShipmentBookingRequest struct {
	Accounts
	APIKey               string    `json:"api_key,omitempty"`
	RateID               string    `json:"rate_id,omitempty"`
	Carrier              string    `json:"carrier,omitempty"` // DHL, FedEX, UPS
	CarrierServiceCode   string    `json:"carrier_service_code,omitempty"`
	OriginDirection      Direction `json:"origin_direction,omitempty"`
	DestinationDirection Direction `json:"destination_direction,omitempty"`
	Shipment             Shipment  `json:"shipment,omitempty"`
	LabelFormat          string    `json:"label_format,omitempty"`
	Customs              Customs   `json:"customs,omitempty"` // Mandatory for international packages; See Customs object.
}

type ShipmentBookingResponse struct {
	Shipment
	Error    bool        `json:"bool,omitempty"`
	Errors   interface{} `json:"errors,omitempty"`
	Warnings interface{} `json:"warnings,omitempty"`
}

type Customs struct {
	DutiesPayment      string      `json:"duties_payment,omitempty"`
	TotalDeclaredValue string      `json:"total_declared_value,omitempty"`
	Commodities        []Commodity `json:"commodities,omitempty"`
}

type Commodity struct {
	Quantity             string `json:"quantity,omitempty"`
	Description          string `json:"description,omitempty"`
	CountryOfManufacture string `json:"country_of_manufacture,omitempty"`
	UnitPrice            int    `json:"unit_price,omitempty"`
	TotalCustomsValue    int    `json:"total_customs_value,omitempty"`
}
