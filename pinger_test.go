package pinger

import (
	"net/http/httptest"
	"testing"
)

func TestPinger(t *testing.T) {
	req := httptest.NewRequest("GET", "/?p=7555", nil)
	rr := httptest.NewRecorder()
	Pinger(rr, req)
	if rr.Result().StatusCode != 200 {
		t.Errorf("Expected 200 response")
	}
}
