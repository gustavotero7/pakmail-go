package models

type Shipment struct {
	ID                    float64  `json:"id,omitempty"`                     // The shipment date. If not provided, we will use the current date in your users time zone.
	ShipmentDate          string   `json:"shipment_date,omitempty"`          // The shipment date. If not provided, we will use the current date in your users time zone.
	ShipmentType          string   `json:"shipment_type,omitempty"`          // Defines if the shipment is a document or a package. If no value is provided, Package is assumed as shipment type.
	InsuredValue          float64  `json:"insured_value,omitempty"`          // Defines the amount you want to insure. Please note that shipment insurance usually generates an additional cost.
	InsuredValueCurrency  string   `json:"insured_value_currency,omitempty"` //The 3 letter currency code in ISO 4217 alphabetic standard.
	Content               string   `json:"content,omitempty"`                // The shipment content.Only required for the labelling service.
	Carrier               string   `json:"carrier,omitempty"`
	CarrierServiceCode    string   `json:"carrier_service_code,omitempty"` // The service code used by the carrier. Only mandatory if you provide at least one parcel.
	EnviayaServiceCode    string   `json:"enviaya_service_code,omitempty"` // A unique service code assigned by us and used to identify this specific service. Can be used for label creation and pickup service.
	EnviayaShipmentNumber string   `json:"enviaya_shipment_number,omitempty"`
	CarrierShipmentNumber string   `json:"carrier_shipment_number,omitempty"`
	Parcels               []Parcel `json:"parcels,omitempty"`
	Label                 string   `json:"label,omitempty"`           // The label in DBASE code
	LabelFormat           string   `json:"label_format,omitempty"`    // The printer formar of the label.
	LabelFileType         string   `json:"label_file_type,omitempty"` // The fily type of the label the dbase code has to be converted to.
	LabelURL              string   `json:"label_url,omitempty"`       // The URL where the label can be downloaded.
	Rate                  Rate     `json:"rate,omitempty"`            // The rate and charges break down of the shipment cost.
	ShipmentStatus        string   `json:"shipment_status"`
	Accounts
}
