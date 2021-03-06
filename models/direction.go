package models

// Direction objects are the core of every shipment quote, a shipment label request and a pickup scheduling request. A direction object can be an origin, destination or pickup address.
type Direction struct {
	FullName     string `json:"full_name,omitempty"`    // Full name: First name(s) and last name(s) of the direction.
	Company      string `json:"company,omitempty"`      // Company name
	Direction1   string `json:"direction_1,omitempty"`  // First street line. In most cases, this is the street name and street number.
	Direction2   string `json:"direction_2,omitempty"`  // Second street line. E.g. for internal numbers or further references to encounter the direction.
	PostalCode   string `json:"postal_code,omitempty"`  // Postal Code (not mandatory for countries that do not manage postal codes)
	Neighborhood string `json:"neighborhood,omitempty"` // The neighborhood is only used for directions in Mexico.
	District     string `json:"district,omitempty"`     // The district is only used for directions in Mexico.
	City         string `json:"city,omitempty"`         // The city is optional for quote requests (unless the direction country does not manage postal codes), but mandatory for pickup booking
	StateCode    string `json:"state_code,omitempty"`   // The state code is only needed for directions within the USA, Canada or Mexico. If no state code is provided, we try to determine
	CountryCode  string `json:"country_code,omitempty"` // The 2 digits ISO 3166 country code of the origin.
	Phone        string `json:"phone,omitempty"`        // Carriers use the phone number to call the direction contact in case of any questions or problems. If you do not know the phone number of an direction, you might consider putting your own.
	Email        string `json:"email,omitempty"`        // The e-mail is used for automatic shipment and tracking notification. You can leave it empty if you don’t want such notifications to be sent.
}
