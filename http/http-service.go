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

	client := &http.Client{}

	req, err := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	return string(body), nil

}
