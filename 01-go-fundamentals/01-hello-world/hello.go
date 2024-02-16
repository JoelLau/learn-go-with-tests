package main

import "fmt"

// language codes
const (
	english = "en"
	spanish = "es"
	french = "fr"
)

func Hello(name, languageCode string) string {
	if name == "" {
		name = "World"
	}

	switch languageCode {
		case spanish:
			return fmt.Sprintf("Hola, %s", name)
		case french:
			return fmt.Sprintf("Bonjour, %s", name)
	}

	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(Hello("Joel", "en"))
}
