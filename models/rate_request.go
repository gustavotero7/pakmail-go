package models

type RateRequest struct {
	Currency             string    `json:"currency"`
	InsuredValue         float64   `json:"insured_value"`
	InsuredValueCurrency string    `json:"insured_value_currency"`
	OriginDirection      Direction `json:"origin_direction"`
	DestinationDirection Direction `json:"destination_direction"`
	Shipment             Shipment  `json:"shipment"`
	OrderTotalAmount     float64   `json:"order_total_amount"`    // Order total amount for subsidy calculate
	IntelligentFiltering bool      `json:"intelligent_filtering"` // If a carrier offers more than one services with the same delivery date and time, only the cheapest one of those services will be returned if the intelligent filtering option is enabled (true). On the contrary, all services will always be returned if this option is disabled (false).
	Accounts
}

type RateResponse struct {
	Warning []string `json:"warning"`
	Dhl     []Rate   `json:"dhl"`
	Fedex   []Rate   `json:"fedex"`
	Redpack []Rate   `json:"redpack"`
	Ups     []Rate   `json:"ups"`
}
