package api

import "testing"

var shortURLTests = []struct {
	n string
}{
	{"http://example.com"},
	{"http://google.com"},
	{"ASDFGHJKL"},
	{"1234567!@#$%^ASDFGH"},
}

func TestCreateShortURL(t *testing.T) {
	for _, tc := range shortURLTests {
		url := tc.n
		s1 := createShortURL(url)
		s2 := createShortURL(url)
		if s1 != s2 {
			t.Fail()
		}
		if s1 == url {
			t.Fail()
		}
	}
}
