package pakmail

import (
	"encoding/json"

	"github.com/gustavotero7/pakmail-go/models"
)

const (
	pathPickupBooking = "/pickups" // POST /api/v1/pickups
)

// PickupBooking POST /api/v1/pickups
func PickupBooking(req models.PickupBookingRequest) (*models.PickupBookingResponse, error) {
	// do request
	res, err := Post(pathShipments, req)
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.PickupBookingResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}
