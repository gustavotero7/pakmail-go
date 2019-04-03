package models

type Amounts struct {
	NetShippingAmount   string  `json:"net_shipping_amount"`
	NetSurchargesAmount string  `json:"net_surcharges_amount"`
	NetTotalAmount      float64 `json:"net_total_amount"`
	VatAmount           float64 `json:"vat_amount"`
	VatRate             float64 `json:"vat_rate"`
	TotalAmount         float64 `json:"total_amount"`
	Currency            string  `json:""`
}
