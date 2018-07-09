package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestHandler(t *testing.T) {
	tests := []struct {
		name    string
		want    Response
		wantErr bool
	}{
		{
			name: "success",
			want: Response{
				Message: "Okay so your other function also executed successfully!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Handler()
			if (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Handler() = %v, want %v", got, tt.want)
			}
		})
	}
}
