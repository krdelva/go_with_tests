package helloworld

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Bob", ""))
	fmt.Println(Hello("", ""))
}
