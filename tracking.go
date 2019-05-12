package pakmail

import (
	"encoding/json"

	"github.com/gustavotero7/pakmail-go/models"
)

const (
	pathTrakings = "/trackings" // POST /api/v1/trackings
)

// ShipmentTracking POST /api/v1/trackings
func ShipmentTracking(req models.TrackingRequest) (*models.TrackingResponse, error) {
	// do request
	res, err := Post(pathTrakings, req)
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.TrackingResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}
