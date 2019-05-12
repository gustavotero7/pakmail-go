package pakmail

import (
	"reflect"
	"testing"

	"github.com/gustavotero7/pakmail-go/models"
)

func TestLookUp(t *testing.T) {
	type args struct {
		enviaya_id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.ShipmentLookupResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				enviaya_id: "8KKBK9H2",
			},
			want:    nil,
			wantErr: false,
		},
	}
	APIKey = ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LookUp(tt.args.enviaya_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("LookUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBooking(t *testing.T) {
	type args struct {
		req models.ShipmentBookingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *models.ShipmentBookingResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				req: models.ShipmentBookingRequest{
					Carrier:            "ups",
					RateID:             "41624964",
					CarrierServiceCode: "07",
					OriginDirection: models.Direction{
						FullName:    "Leopold stoch",
						City:        "Ciudad Mexiko",
						Phone:       "23232323",
						StateCode:   "DF",
						PostalCode:  "11550",
						CountryCode: "MX",
						Direction1:  "Av siempre viva",
					},
					DestinationDirection: models.Direction{
						FullName:    "Stan Marsh",
						City:        "Ciudad Mexiko",
						Phone:       "23232323",
						StateCode:   "DF",
						PostalCode:  "01210",
						CountryCode: "MX",
						Direction1:  "Av siempre viva",
					},
					Shipment: models.Shipment{
						ShipmentType:         "Package",
						InsuredValue:         5000,
						InsuredValueCurrency: "MXN",
						Content:              "Description",
						Parcels: []models.Parcel{
							models.Parcel{
								Quantity:      1,
								Weight:        3,
								WeightUnit:    models.WeightUnitKG,
								Length:        10,
								Height:        20,
								Width:         30,
								DimensionUnit: models.DimensionUnitCM,
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	APIKey = ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Booking(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Booking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("Booking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancellation(t *testing.T) {
	type args struct {
		req models.ShipmentCancellationRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *models.ShipmentCancellationResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				req: models.ShipmentCancellationRequest{
					ShipmentID: 4656974,
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	APIKey = ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Cancellation(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cancellation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("Cancellation() = %v, want %v", got, tt.want)
			}
		})
	}
}
