package main

import "fmt"

//main func
func main() {
	fmt.Println(hello("world", ""))
}

const englishHelloPrefix = "hello "
const spanishHelloPrefix = "hola "
const frenchHelloPrefix = "bonjour "
const spanish = "spanish"
const french = "french"

//simple example of a unit test
func hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	if lang == spanish {
		return spanishHelloPrefix + name
	}
	if lang == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
