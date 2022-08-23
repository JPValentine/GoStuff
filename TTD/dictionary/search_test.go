package main

import(
	"testing"
)

func assertStrings(t testing.TB, got , want string){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}//assertStrings

func assertError(t testing.TB, got , want error){
	t.Helper()
	if got != want{
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}//assertError

func TestSearch(t *testing.T){
	d := Dictionary{"test":"this is a test"} 

	t.Run("known word", func(t *testing.T){
		got,_ := d.Search("test")
		want:= "this is a test"

		assertError(t,got,want)
	})

	t.Run("unknown word", func(t *testing.T){
		_, err := d.Search("unknown")

		assertStrings(t,err.Error(),ErrNotFound)
	})
}//TestSearch

func TestAdd(t *testing.T){
	d := Dictionary{}
	d.Add("test","this is a test")
	want:="this is a test"
	got, err := d.Search("test")
	if err != nil{
		t.Fatal("should find added word: ",err)
	}
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}//TestAdd

