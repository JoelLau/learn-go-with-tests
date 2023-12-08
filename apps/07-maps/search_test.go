package main

import (
	"fmt"
	"testing"
)

type SearchTest struct {
	SearchSpace      Dictionary
	SearchTerm       string
	SearchDefinition string
	IsExpectingError bool
}

// `const` cannot be used here because arrays and slices are mutable
// `[...]` "locks in" the size of the array at compile time
var SearchTests = [...]SearchTest{
	{
		SearchSpace:      Dictionary{"test": "this is just a test"},
		SearchTerm:       "test",
		SearchDefinition: "this is just a test",
		IsExpectingError: false,
	},
	{
		SearchSpace:      Dictionary{"123": "456"},
		SearchTerm:       "123",
		SearchDefinition: "456",
		IsExpectingError: false,
	},
	{
		SearchSpace:      Dictionary{"123": "456"},
		SearchTerm:       "abc",
		SearchDefinition: "",
		IsExpectingError: true,
	},
}

type AddTest struct {
	ExistingDictionary Dictionary
	AddTerm            string
	NewDefinition      string
	ExpectedDefinition string
	IsExpectingError   bool
}

var AddTests = [...]AddTest{
	{
		ExistingDictionary: Dictionary{},
		AddTerm:            "test",
		NewDefinition:      "this is just a test",
		ExpectedDefinition: "this is just a test",
		IsExpectingError:   false,
	},
	{
		ExistingDictionary: Dictionary{"test": "this is just a test"},
		AddTerm:            "test",
		NewDefinition:      "this is not just a test",
		ExpectedDefinition: "this is just a test",
		IsExpectingError:   true,
	},
	{
		ExistingDictionary: Dictionary{"test": "this is just a test"},
		AddTerm:            "test2",
		NewDefinition:      "this is just a second test",
		ExpectedDefinition: "this is just a second test",
		IsExpectingError:   false,
	},
}

func TestSearch(t *testing.T) {
	t.Run("Search", func(t *testing.T) {
		for _, test := range SearchTests {
			title := fmt.Sprintf(
				"given '%v' when searching '%v', return '%v'",
				test.SearchSpace,
				test.SearchTerm,
				test.SearchDefinition,
			)

			t.Run(title, func(t *testing.T) {
				result, err := test.SearchSpace.Search(test.SearchTerm)

				assertString(t, result, test.SearchDefinition)
				if test.IsExpectingError {
					assertHasError(t, err)
				} else {
					assertNoError(t, err)
				}
			})
		}
	})

	t.Run("Add", func(t *testing.T) {
		for _, test := range AddTests {
			title := fmt.Sprintf("given '%v' when adding '%v'", test.ExistingDictionary, test.AddTerm)

			t.Run(title, func(t *testing.T) {
				addResult, err := test.ExistingDictionary.Add(test.AddTerm, test.NewDefinition)

				assertString(t, addResult, test.ExpectedDefinition)
				if test.IsExpectingError {
					assertHasError(t, err)
				} else {
					assertNoError(t, err)
				}

				searchResult, err := test.ExistingDictionary.Search(test.AddTerm)
				assertNoError(t, err)
				assertString(t, searchResult, test.ExpectedDefinition)
			})
		}
	})
}

func assertHasError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Errorf("was expecting error, but received %v instead", err)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("was not expecting error, but received %v instead", err)
	}
}

func assertString(t testing.TB, have, want string) {
	t.Helper()

	if have != want {
		t.Errorf("expected %v and %v to match", have, want)
	}
}
