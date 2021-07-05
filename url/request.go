package url

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Result represents URL visit result
type Result struct {
	Address string
	Hash    string
	Err     error
}

// VisitURL do http request to url
func VisitURL(url string) Result {
	url = validateURLScheme(url)
	resp, err := http.Get(url)
	if err != nil {
		return Result{Address: url, Err: err}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{Address: url, Err: err}
	}
	return Result{Address: url, Hash: fmt.Sprintf("%x", md5.Sum(body))}

}

// adds https scheme if not found in the url
func validateURLScheme(url string) string {
	if !strings.HasPrefix(url, "http") {
		return "https://" + url
	}
	return url
}
