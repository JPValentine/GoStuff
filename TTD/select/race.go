package main

import(
	"time"
	"net/http"
//	"net/http/httptest"
)

func measureResponse(url string) time.Duration{
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(a, b string)(winner string){
	startA := time.Now()
	http.Get(a)
	aDur := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDur := time.Since(startB)

	if aDur < bDur{
		return a
	}
	return b
}
