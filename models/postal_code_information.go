package models

type PostalCodeInformation struct {
	PostalCode   string `json:"postal_code"`
	Neighborhood string `json:"neighborhood"`
	District     string `json:"district"`
	State        string `json:"state"`
	City         string `json:"city"`
	StateCode    string `json:"state_code"`
}

type PostalCodeInformationResponse struct {
	PostalCodeInformation []PostalCodeInformation `json:"postal_code_information"`
}
