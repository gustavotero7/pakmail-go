package models

type Shipment struct {
	ShipmentDate          string      `json:"shipment_date"`          // The shipment date. If not provided, we will use the current date in your users time zone.
	ShipmentType          string      `json:"shipment_type"`          // Defines if the shipment is a document or a package. If no value is provided, Package is assumed as shipment type.
	InsuredValue          int         `json:"insured_value"`          // Defines the amount you want to insure. Please note that shipment insurance usually generates an additional cost.
	InsuredValueCurrency  string      `json:"insured_value_currency"` //The 3 letter currency code in ISO 4217 alphabetic standard.
	Content               string      `json:"content"`                // The shipment content.Only required for the labelling service.
	Carrier               string      `json:"carrier"`
	CarrierServiceCode    string      `json:"carrier_service_code"` // The service code used by the carrier. Only mandatory if you provide at least one parcel.
	EnviayaServiceCode    string      `json:"enviaya_service_code"` // A unique service code assigned by us and used to identify this specific service. Can be used for label creation and pickup service.
	EnviayaShipmentNumber string      `json:"enviaya_shipment_number"`
	CarrierShipmentNumber string      `json:"carrier_shipment_number"`
	Parcels               []Parcel    `json:"parcels"`
	Label                 string      `json:"label"`           // The label in DBASE code
	LabelFormat           string      `json:"label_format"`    // The printer formar of the label.
	LabelFileType         string      `json:"label_file_type"` // The fily type of the label the dbase code has to be converted to.
	LabelURL              string      `json:"label_url"`       // The URL where the label can be downloaded.
	Rate                  []Rate      `json:"rate"`            // The rate and charges break down of the shipment cost.
	Accounts              interface{} `json:"accounts"`
}
