package httpservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// HTTPRequest ...
func HTTPRequest(URL string, form url.Values) (string, error) {

	req, err := http.Post(URL, "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer req.Body.Close()

	body, _ := ioutil.ReadAll(req.Body)

	return string(body), nil

}
