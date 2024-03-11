package main

const (
	ErrSearchKeyNotFound = DictionaryErr("could not find the key you were looking for")
	ErrAddKeyCollision   = DictionaryErr("could not add entry because key already exists")
	ErrUpdateKeyNotFound = DictionaryErr("could not update entry because key doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
