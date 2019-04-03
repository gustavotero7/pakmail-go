package pakmail

import "github.com/gustavotero7/pakmail-go/models"

type Shipments interface{}

const (
	PATH_SHIPMENTS = "/shipments"
)

// Booking Post https://api224.pakmail.com.mx/api/v1/shipments
func Booking(s models.ShipmentBooking) error {
	if _, err := Post(PATH_SHIPMENTS, s); err != nil {
		return err
	}
	return nil
}
func LookUp()       {}
func Cancellation() {}
