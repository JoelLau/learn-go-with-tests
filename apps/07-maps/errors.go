package main

type DictionaryError string

const (
	NotFoundError   = DictionaryError("Not found")
	TermExistsError = DictionaryError("Term exists")
)

func (e DictionaryError) Error() string {
	return string(e)
}
