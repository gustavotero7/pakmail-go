package pakmail

import (
	"encoding/json"

	"github.com/gustavotero7/pakmail-go/models"
)

const (
	pathRates = "/rates" // POST https://api224.pakmail.com.mx/api/v1/rates
)

// Rate POST /api/v1/rates
func Rate(req models.RateRequest) (*models.RateResponse, error) {
	// do request
	res, err := Post(pathRates, req)
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.RateResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}
