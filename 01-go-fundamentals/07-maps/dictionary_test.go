package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dict := Dictionary{
		"test": "this is just a test",
		"asdf": "asdf",
		"qwer": "qwer",
	}

	for idx, tt := range []struct {
		dict   Dictionary
		key    string
		expect string
		err    error
	}{
		{
			dict:   dict,
			key:    "test",
			expect: "this is just a test",
			err:    nil,
		},
		{
			dict:   dict,
			key:    "asdf",
			expect: "asdf",
			err:    nil,
		},
		{
			dict:   dict,
			key:    "some key that doesn't exisdt 80nlkjsd0v",
			expect: "",
			err:    ErrSearchKeyNotFound,
		},
	} {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			have, err := tt.dict.Search(tt.key)
			want := tt.expect

			assertError(t, err, tt.err)
			assertStrings(t, have, want)
		})
	}
}

func TestAdd(t *testing.T) {
	for idx, tt := range []struct {
		dict   Dictionary
		key    string
		expect string
		err    error
	}{
		{
			dict:   Dictionary{},
			key:    "test",
			expect: "this is just a test",
			err:    nil,
		},
		{
			dict:   Dictionary{"asdf": "asdf"},
			key:    "asdf",
			expect: "asdf",
			err:    ErrAddKeyCollision,
		},
	} {
		t.Run(fmt.Sprintf("Test %d", idx), func(t *testing.T) {
			err := tt.dict.Add(tt.key, tt.expect)
			assertError(t, err, tt.err)

			have, err := tt.dict.Search(tt.key)
			want := tt.expect

			assertError(t, err, nil)
			assertStrings(t, have, want)
		})
	}
}

func TestUpdate(t *testing.T) {
	t.Run("replace old value", func(t *testing.T) {
		const (
			key      = "key"
			oldValue = "old value"
			newValue = "new value"
		)

		dict := Dictionary{key: oldValue}
		err := dict.Update(key, newValue)
		assertError(t, err, nil)

		found, err := dict.Search(key)
		assertError(t, err, nil)
		assertStrings(t, found, newValue)
	})

	t.Run("non-existant key", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("non-existant key", "dummy val")
		assertError(t, err, ErrUpdateKeyNotFound)
	})
}

func assertStrings(t testing.TB, have, want string) {
	t.Helper()

	if have != want {
		t.Errorf("have %q want %q", have, want)
	}
}

func assertError(t testing.TB, have, want error) {
	t.Helper()

	if !errors.Is(have, want) {
		t.Errorf("have error '%s' want error '%s'", have, want)
	}
}
