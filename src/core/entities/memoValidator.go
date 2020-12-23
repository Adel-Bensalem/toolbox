package entities

func IsMemoValid(title string, body string) bool {
	return len(title) > 0 && len(body) > 0
}
