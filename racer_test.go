package integers

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://www.facebook.com"
	fastURl := "http://www.quil.co.uk"

	want := fastURl
	got := Racer(slowURL, fastURl)

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
