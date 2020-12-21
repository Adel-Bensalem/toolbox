package adapters

type RequestData struct {
	Body   string
	Url    string
	Method string
}

type ResponseData struct {
	Code string
	Body string
}

type RequestClient interface {
	SendRequest(data RequestData) (ResponseData, error)
}
