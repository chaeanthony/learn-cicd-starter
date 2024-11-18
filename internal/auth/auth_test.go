package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   http.Header
		wantRes string
		wantErr error
	}{
		"valid":                      {input: http.Header{"Authorization": []string{"ApiKey api_key"}, "Content-Type": []string{"application/json"}}, wantRes: "api_key", wantErr: nil},
		"invalid auth header format": {input: http.Header{"Authorization": []string{"bad api_key"}}, wantRes: "", wantErr: errors.New("malformed authorization header")},
		"no auth header":             {input: http.Header{"Content-Type": []string{"application/json"}}, wantRes: "", wantErr: errors.New("no authorization header included")},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if !reflect.DeepEqual(tc.wantRes, got) {
				t.Fatalf("expected: %v, got: %v\nexpected error: %v\nerror: %v", tc.wantRes, got, tc.wantErr, err)
			}
		})
	}

}
