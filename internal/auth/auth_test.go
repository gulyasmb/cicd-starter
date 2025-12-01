package auth

import (
    "reflect"
    "testing"
	"errors"
	"net/http"
)

func TestAuth(t *testing.T) {
	tests := map[string]struct {
        header http.Header
        want   string
        wanterr	   error
    }{
        "EmptyHeader": {header: http.Header{}, 
						want: "", 
						wanterr: errors.New("no authorization header included")},
        "Malformed1":  {header: http.Header{"Authorization": []string{"NotApiKey fake-token"}}, 
						want: "", 
						wanterr: errors.New("malformed authorization header")},
		"Malformed2":  {header: http.Header{"Authorization": []string{"malformed-token"}}, 
						want: "", 
						wanterr: errors.New("malformed authorization header")},
		"Normal1":     {header: http.Header{"Authorization": []string{"ApiKey normal-token"}}, 
						want: "normal-token", 
						wanterr: nil},
		"Normal2":     {header: http.Header{"Authorization": []string{"ApiKey normal-token too"}}, 
						want: "normal-token", 
						wanterr: nil},
    }

	for name, tc := range tests {
        got, goterr := GetAPIKey(tc.header)
        if !reflect.DeepEqual(tc.want, got) {
            t.Fatalf("%s: expected: %v, got: %v", name, tc.want, got)
        }
		if !reflect.DeepEqual(tc.wanterr, goterr) {
            t.Fatalf("%s: expected error: %v, got: %v", name, tc.wanterr, goterr)
        }
    }
}