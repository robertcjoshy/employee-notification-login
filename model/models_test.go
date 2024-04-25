package model

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAuthenticate(t *testing.T) {
	type args struct {
		user string
		c    *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Authenticate(tt.args.user, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
