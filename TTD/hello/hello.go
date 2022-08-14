package main

import "fmt"

//main func
func main() {
	fmt.Println(Hello("world", ""))
}

const englishHelloPrefix = "hello "
const spanishHelloPrefix = "hola "
const frenchHelloPrefix = "bonjour "
const spanish = "spanish"
const french = "french"

//simple example of a unit test
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return prefix(lang) + name
}

func prefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
