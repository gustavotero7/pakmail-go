package models

type ShipmentBooking struct {
	ApiKey               string      `json:"api_key"`
	RateId               int         `json:"rate_id"`
	Carrier              string      `json:"carrier"` // DHL, FedEX, UPS
	CarrierServiceCode   string      `json:"carrier_service_code"`
	CarrierAccount       interface{} `json:"carrier_account"`
	EnviayaAccount       string      `json:"enviaya_account"`
	OriginDirection      Direction   `json:"origin_direction"`
	DestinationDirection Direction   `json:"destination_direction"`
	Shipment             Shipment    `json:"shipment"`
	LabelFormat          string      `json:"label_format"`
	Customs              Customs     `json:"customs"` // Mandatory for international packages; See Customs object.
}

type Customs struct {
	DutiesPayment      string      `json:"duties_payment"`
	TotalDeclaredValue string      `json:"total_declared_value"`
	Commodities        []Commodity `json:"commodities"`
}

type Commodity struct {
	Quantity             string `json:"quantity"`
	Description          string `json:"description"`
	CountryOfManufacture string `json:"country_of_manufacture"`
	UnitPrice            int    `json:"unit_price"`
	TotalCustomsValue    int    `json:"total_customs_value"`
}
