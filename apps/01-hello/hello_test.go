// wtf is this
package main

import "testing"

// Table-Driven Tests
// https://dev.to/boncheff/table-driven-unit-tests-in-go-407b
type helloTestCase struct {
	// person to say hello to
	name string

	// person to say hello to
	language string

	// expected output given the above args
	expect string
}

func TestHello(t *testing.T) {
	for _, scenario := range []helloTestCase{
		{name: "", language: "", expect: "Hello, World!"},
		{name: "Bobby", language: "", expect: "Hello, Bobby!"},
		{name: "John", language: "en", expect: "Hello, John!"},
		{name: "", language: "es", expect: "Hola, World!"},
		{name: "Brian", language: "FR", expect: "Bounjour, Brian!"},
	} {
		t.Run("say hello to '"+scenario.name+"' in language '"+scenario.language+"'", func(t *testing.T) {

			assertCorrectMessage(t, Hello(scenario.name, scenario.language), scenario.expect)
		})
	}
}

func assertCorrectMessage(t testing.TB, have, want string) {
	t.Helper()

	if have != want {
		t.Errorf("have %q want %q", have, want)
	}
}
