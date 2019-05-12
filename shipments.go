package pakmail

import (
	"encoding/json"
	"fmt"

	"github.com/gustavotero7/pakmail-go/models"
)

const (
	pathShipments    = "/shipments"            // POST https://api224.pakmail.com.mx/api/v1/shipments
	pathLookUp       = "/shipments/%s"         // GET https://api224.pakmail.com.mx/api/v1/shipments/{enviaya_id}?api_key={api_key}
	pathCancellation = "/request_cancellation" // POST https://api224.pakmail.com.mx/api/v1/request_cancellation
)

// Booking Post https://api224.pakmail.com.mx/api/v1/shipments
func Booking(req models.ShipmentBookingRequest) (*models.ShipmentBookingResponse, error) {
	// do request
	res, err := Post(pathShipments, req)
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.ShipmentBookingResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}

// LookUp GET /api/v1/shipments/enviaya_id?api_key=YOUR_API_KEY
func LookUp(enviaya_id string) (*models.ShipmentLookupResponse, error) {
	// do request
	res, err := Get(fmt.Sprintf(pathLookUp, enviaya_id))
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.ShipmentLookupResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}

// Cancellation POST https://api224.pakmail.com.mx/api/v1/request_cancellation
func Cancellation(req models.ShipmentCancellationRequest) (*models.ShipmentCancellationResponse, error) {
	// do request
	res, err := Post(pathCancellation, req)
	if err != nil {
		return nil, err
	}
	// decode response
	r := &models.ShipmentCancellationResponse{}
	if err := json.Unmarshal(res, r); err != nil {
		return nil, err
	}
	return r, nil
}
