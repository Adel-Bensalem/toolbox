package adapters

type MemoRepository interface {
	SaveMemo(title string, body string) error
}
