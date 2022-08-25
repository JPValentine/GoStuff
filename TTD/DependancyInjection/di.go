package main

import(
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w,"Hello, %s", name)
}//Greet

func MyGreeterHandler(w http.ResponseWriter, r *http.Request){
	Greet(w, "World")
}

func main(){
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}//main
