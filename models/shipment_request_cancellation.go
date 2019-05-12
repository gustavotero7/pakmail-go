package models

type ShipmentCancellationRequest struct {
	APIKey         string `json:"api_key"`         // YOUR_API_KEY
	EnviayaAccount string `json:"enviaya_account"` // The Billing account.
	EnviayaID      string `json:"enviaya_id"`      // The EnviaYa ID of your shipment.
	ShipmentID     int64  `json:"shipment_id"`     // The shipment ID.
}

type ShipmentCancellationResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
}
