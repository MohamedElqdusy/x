package url

import (
	"testing"
)

func TestValidateURLScheme(t *testing.T) {
	var cases = []struct {
		Name     string
		URL      string
		Expected string
	}{
		{
			Name:     "url without scheme",
			URL:      "google.com",
			Expected: "https://google.com",
		},
		{
			Name:     "url with http scheme",
			URL:      "http://adjust.com",
			Expected: "http://adjust.com",
		},
		{
			Name:     "url with https scheme",
			URL:      "https://reddit.com",
			Expected: "https://reddit.com",
		},
	}

	for _, input := range cases {
		actual := validateURLScheme(input.URL)
		if input.Expected != actual {
			t.Errorf("Test failed:  %s \n Want to be: %v but we got %v",
				input.Name,
				input.Expected,
				actual,
			)
		}
	}
}
