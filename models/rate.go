package models

type Rate struct {
	ID                        string               `json:"id"`                           // The ID of the rate.
	RateID                    string               `json:"rate_id"`                      // The ID of the rate (alternative name).
	Date                      string               `json:"date"`                         // Date and Time of the rate. This can be important for time definte services, because they might not be delivered on time if shipment is realized until a later hour of the day.
	Carrier                   string               `json:"carrier"`                      // Carrier providing the service.
	CarrierServiceName        string               `json:"carrier_service_name"`         // The service name used by the carrier.
	CarrierServiceCode        string               `json:"carrier_service_code"`         // The service code used by the carrier.
	CarrierLogoURL            string               `json:"carrier_logo_url"`             // The URL to the carrier logo. Can be used if you want to display the carrier logo to your users.
	CarrierLogoURLSvg         string               `json:"carrier_logo_url_svg"`         // The URL to the carrier logo. Can be used if you want to display the carrier logo to your users.
	EstimatedDelivery         string               `json:"estimated_delivery"`           // The estimated delivery date and time. Time is only provided for a time definte service. For all other services, only an estimated delivery date is provided.
	NetShippingAmount         float64              `json:"net_shipping_amount"`          // The net amount of the shipping service. This amount does not include additional surcharges that migth apply.
	NetSurchargesAmount       float64              `json:"net_surcharges_amount"`        // The net amount of all surcharges that apply additional to the net shipping amount.
	NetTotalAmount            float64              `json:"net_total_amount"`             // The total net amount of the shipment, including all charges that apply (shipping amound and surcharges) You will probably want to use this field if you want to know the total net shipping amount of a service.
	VATAmount                 float64              `json:"vat_amount"`                   // The vat amount applying.
	VATRate                   float64              `json:"vat_rate"`                     // The vat rate.
	TotalAmount               float64              `json:"total_amount"`                 // The total amount / final price of the shipment. This amount includes all charges that apply (shipping and surcharges) and VAT. You will probably want to use this field if you want to know the final total price of a service.
	PublicTotalAmount         float64              `json:"public_total_amount"`          // The public total amount.
	PublicTotalAmountCurrency string               `json:"public_total_amount_currency"` // The 3 letter currency code in ISO 4217 alphabetic standard.
	Currency                  string               `json:"currency"`                     // The 3 letter currency code in ISO 4217 alphabetic standard.
	ServiceTerms              string               `json:"service_terms"`                // If special terms apply for a service, these are mentioned here.
	SurchargesBreakDown       []SurchargeBreakDown `json:"surcharges_break_down"`        // The breakdown of surcharges that apply to your shipment.
	EnviayaServiceName        string               `json:"enviaya_service_name"`         // A standardized service name assigned by us. This can be used for easier understanding and comparision of the carrier services. Example: "2 day shipping".
	EnviayaServiceCode        string               `json:"enviaya_service_code"`         // A unique service code assigned by us and used to identify this specific service. Can be used for label creation and pickup service.
}
