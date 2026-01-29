package main

import "fmt"

const (
	chinese = "Chinese"
	french  = "French"
	spanish = "Spanish"

	chineseHelloPrefix = "Hi, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case chinese:
		prefix = chineseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

// if language == spanish {
// 	return spanishHelloPrefix + name
// }
// if language == french {
// 	return frenchHelloPrefix + name
// }
// return englishHelloPrefix + name
