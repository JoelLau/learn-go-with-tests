package hello

import "testing"

// file needs to end with "test.go"
// function must start with "Test"
// function must take exactly one argument `t *testing.T`
// ^ imported from "testing"
func TestHello(t *testing.T) {
	t.Run("when language is 'en' and name is provided", func(t *testing.T)  {
		have := Hello("Joel", "en")
		want := "Hello, Joel"

		assertCorrectMessage(t, have, want)
	})

	t.Run("when language is 'en' and no name is provided", func(t *testing.T)  {
		have := Hello("", "en")
		want := "Hello, World"

		assertCorrectMessage(t, have, want)
	})

	t.Run("when language is 'es' and name is provided", func(t *testing.T)  {
		have := Hello("Joel", "es")
		want := "Hola, Joel"

		assertCorrectMessage(t, have, want)
	})

	t.Run("when language is 'es' and no name is provided", func(t *testing.T)  {
		have := Hello("", "es")
		want := "Hola, World"

		assertCorrectMessage(t, have, want)
	})

	t.Run("when language is 'fr' and name is provided", func(t *testing.T)  {
		have := Hello("Joel", "fr")
		want := "Bonjour, Joel"

		assertCorrectMessage(t, have, want)
	})

	t.Run("when language is 'fr' and no name is provided", func(t *testing.T)  {
		have := Hello("", "fr")
		want := "Bonjour, World"

		assertCorrectMessage(t, have, want)
	})
}

// testing.TB is asdf
func assertCorrectMessage(t testing.TB, have, want string) {
	// tells Go that this is a helper function
	// this ensures that the stack trace shows the caller
	t.Helper()

	if have != want {
		t.Errorf("have %q want %q", have, want)
	}
}
