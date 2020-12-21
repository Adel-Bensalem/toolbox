package libs

import (
	"core/adapters"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RequestClient struct{}

func (requestClient *RequestClient) SendRequest(data adapters.RequestData) (adapters.ResponseData, error) {
	client := http.Client{}
	parsedUrl, _ := url.ParseRequestURI(data.Url)
	requestBody := ioutil.NopCloser(strings.NewReader(data.Body))
	headers := http.Header{}

	headers.Add("Content-Type", "application/json")

	res, err := client.Do(&http.Request{
		Method: strings.ToUpper(data.Method),
		URL:    parsedUrl,
		Body:   requestBody,
		Header: headers,
	})

	if err != nil {
		return adapters.ResponseData{
			Code: "500",
			Body: "",
		}, err
	}

	defer res.Body.Close()

	responseData, _ := ioutil.ReadAll(res.Body)

	return adapters.ResponseData{
		Code: strconv.Itoa(res.StatusCode),
		Body: string(responseData),
	}, nil
}
