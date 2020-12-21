package entities

import (
	"regexp"
	"strings"
)

func hasEndpoint(endpoint string) bool {
	return len(endpoint) > 0
}

func isUrlValid(endpoint string) bool {
	res, _ := regexp.MatchString(
		"^(https?:\\/\\/)?((([a-z\\d]([a-z\\d-]*[a-z\\d])*)\\.)+[a-z]{2,}|((\\d{1,3}\\.){3}\\d{1,3}))(\\:\\d+)?(\\/[-a-z\\d%_.~+]*)*(\\?[;&a-z\\d%_.~+=-]*)?(\\#[-a-z\\d_]*)?$",
		endpoint,
	)

	return res
}

func hasValidMethod(method string) bool {
	methods := []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
		"OPTIONS",
	}

	for _, m := range methods {
		if strings.ToUpper(method) == m {
			return true
		}
	}

	return false
}

func IsRequestValid(method string, endpoint string) bool {
	return hasEndpoint(endpoint) && isUrlValid(endpoint) && hasValidMethod(method)
}
