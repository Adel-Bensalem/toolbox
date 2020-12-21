package usecases

import (
	"core/adapters"
	"core/entities"
	"errors"
)

type requestModel struct {
	Data   string
	Url    string
	Method string
}

type responseModel struct {
	Code string
	Body string
}

type RequestSendInteractor func(model requestModel) (responseModel, error)

func CreateRequestSendInteractor(requestClient adapters.RequestClient) RequestSendInteractor {
	return func(model requestModel) (responseModel, error) {
		if !entities.IsRequestValid(model.Method, model.Url) {
			return responseModel{Body: "", Code: ""}, errors.New("invalid request provided")
		}

		if res, err := requestClient.SendRequest(adapters.RequestData{
			Body:   model.Data,
			Url:    model.Url,
			Method: model.Method,
		}); err != nil {
			return responseModel{Body: "", Code: "500"}, err
		} else {
			return responseModel{Body: res.Body, Code: res.Code}, nil
		}
	}
}
