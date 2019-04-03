package models

type Surcharge struct {
	Name        string  `json:"name"`
	NetAmount   float64 `json:"net_amount"`
	VatAmount   float64 `json:"vat_amount"`
	VatRate     float64 `json:"vat_rate"`
	TotalAmount float64 `json:"total_amount"`
}

type SurchargeBreakDown struct {
	CarrierSurchargeName string  `json:"carrier_surcharge_name"`
	SurchargeNetAmount   float64 `json:"net_amount"`
	SurchargeVat         float64 `json:"surcharge_vat"`
	SurchargeTotal       float64 `json:"surcharge_total"`
}
