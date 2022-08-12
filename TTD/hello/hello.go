package main

import "fmt"

//main func
func main() {
	fmt.Println(hello("world"))
}

//simple example of a unit test
func hello(name string) string {
	return "hello " + name
}
