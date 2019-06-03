package pakmail

import (
	"encoding/json"
	"fmt"

	"github.com/gustavotero7/pakmail-go/models"
)

const (
	pathPostalCodeInformation = "/get_postal_code_information?country_code=%s&postal_code=%s"
)

// GetPostalCodeInformation GET /get_postal_code_information?country_code=COUNTRY_CODE&page=1&per_page=500&postal_code=POSTAL_CODE"
func GetPostalCodeInformation(countryCode, postalCode string) (*models.PostalCodeInformationResponse, error) {
	// do request
	res, err := Get(fmt.Sprintf(pathPostalCodeInformation, countryCode, postalCode))
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.PostalCodeInformationResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}
