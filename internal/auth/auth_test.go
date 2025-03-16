package auth

import (
	"testing"
)

func TestGetAPIKeySucces(t *testing.T) {
	testcase := make(map[string][]string)
	testcase["Authorization"] = []string{"ApiKey asdofiajesf"}
	var wanterr error
	want, wanterr := "asdofiajesf", nil
	got, goterr := GetAPIKey(testcase)

	if want != got || wanterr != goterr {
		t.Fatalf("Wanted %s, %s\nGot %s, %s", want, wanterr, got, goterr)
	}
}

func TestGetAPIKeyFailNoAuth(t *testing.T) {
	testcase := make(map[string][]string)
	testcase["NotAuth"] = []string{"ApiKey asdofiajesf"}
	var wanterr error
	want, wanterr := "", ErrNoAuthHeaderIncluded
	got, goterr := GetAPIKey(testcase)

	if want != got || wanterr != goterr {
		t.Fatalf("Wanted %s, %s\nGot %s, %s", want, wanterr, got, goterr)
	}
}

func TestGetAPIKeyFailMalformedAuthNoKey(t *testing.T) {
	testcase := make(map[string][]string)
	testcase["Authorization"] = []string{"ApiKey"}
	want := ""
	got, goterr := GetAPIKey(testcase)

	if want != got || goterr == nil {
		t.Fatalf("Wanted %s, error\nGot %s, %s", want, got, goterr)
	}
}

func TestGetAPIKeyFailMalformedAuthWrongPrefix(t *testing.T) {
	testcase := make(map[string][]string)
	testcase["Authorization"] = []string{"Bearer key"}
	want := "broken"
	got, goterr := GetAPIKey(testcase)

	if want != got || goterr == nil {
		t.Fatalf("Wanted %s, error\nGot %s, %s", want, got, goterr)
	}
}
