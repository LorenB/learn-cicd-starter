package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "")
	got, gotErr := GetAPIKey(headers)
	want := ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected %v, got: %v", want, got)
	}

	// Compare error messages, not error instances
	if gotErr == nil {
		t.Fatal("expected error, got nil")
	}
	wantErrMsg := "no authorization header included"
	if gotErr.Error() != wantErrMsg {
		t.Fatalf("expected error: %v, got: %v", wantErrMsg, gotErr.Error())
	}
}
