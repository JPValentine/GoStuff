package main

import(
	"testing"
	"reflect"
	"time"
)

func mockWebChecker(url string) bool{
	if url == "waat://furhurterwe.geds"{
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool{
	time.Sleep(20 * time.Millisecond)
	return true
}//slowStubWebsiteChecker

func benchmarkCheckWebsites(b *testing.B){
	urls := make([]string, 100)
	for i:=0; i<len(urls); i++{
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func TestWebChecker(t *testing.T){
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}
	want := map[string]bool{
		"http://google.com": true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds": false,
	}

	got := CheckWebsites(mockWebChecker, websites)

	if !reflect.DeepEqual(want, got){
		t.Fatalf("wanted %v, got %v",want, got)
	}
}
