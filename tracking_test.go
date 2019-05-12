package pakmail

import (
	"reflect"
	"testing"

	"github.com/gustavotero7/pakmail-go/models"
)

func TestShipmentTracking(t *testing.T) {
	type args struct {
		req models.TrackingRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *models.TrackingResponse
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				req: models.TrackingRequest{
					ShipmentNumber: "U4J8Z8ZX",
					Carrier:        "ups",
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	APIKey = ""
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ShipmentTracking(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShipmentTracking() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				//t.Errorf("ShipmentTracking() = %v, want %v", got, tt.want)
			}
		})
	}
}
