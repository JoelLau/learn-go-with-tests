package main

import "strings"

const (
	english = "EN"
	spanish = "ES"
	french  = "FR"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bounjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return getGreetingPrefix(language) + name + "!"
}

// named result parameters
// https://go.dev/doc/effective_go#named-results
func getGreetingPrefix(language string) (prefix string) {
	switch strings.ToUpper(language) {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case english:
		fallthrough
	default:
		prefix = englishHelloPrefix
	}
	return
}
