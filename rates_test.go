package pakmail

import (
	"reflect"
	"testing"

	"github.com/gustavotero7/pakmail-go/models"
)

func TestRate(t *testing.T) {
	type args struct {
		req models.RateRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *models.RateResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				req: models.RateRequest{
					OrderTotalAmount: 60,
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
			got, err := Rate(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Rate() = %v, want %v", got, tt.want)
			}
		})
	}
}
