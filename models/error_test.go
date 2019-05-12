package models

import "testing"

func TestError_Error(t *testing.T) {
	type fields struct {
		Errors string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				Errors: "ping",
			},
			want: "ping",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			er := &Error{
				Errors: tt.fields.Errors,
			}
			if got := er.Error(); got != tt.want {
				t.Errorf("Error.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
