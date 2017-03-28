package httpservice

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const POST string = "POST"
const GET string = "GET"

// HTTPRequest ...
func HTTPRequest(HTTPMethod string, URL string, form url.Values) (string, error) {

	req, err := http.NewRequest(HTTPMethod, URL, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(req.Body)
	s := buf.String()

	return s, nil

}
