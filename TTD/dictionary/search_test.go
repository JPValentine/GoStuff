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
		t.Errorf("got %q want %q given", got, want)
	}
}//assertError

func assertDefinition(t testing.TB, d Dictionary, word, definition string){
	t.Helper()
	got, err := d.Search(word)
	if err != nil{
		t.Fatal("should find added word: ", err)
	}//if
	if got != definition{
		t.Errorf("got %q want %q", got, definition)
	}
}//assertDefinition

func TestUpdate(t *testing.T){
	word := "test"
	definition := "this is a test"
	d := Dictionary{word: definition}
	newDefinition := "this is a new definition"

	d.Update(word, newDefinition)

	assertDefinition(t, d, word, newDefinition)

	t.Run("existing word", func(t *testing.T){
		word := "test"
		definition := "this is a test"
		d := Dictionary{word: definition}
		newDefinition := "this is a new definition"

		err := d.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T){
		word := "test"
		definition := "this is a test"
		d := Dictionary{}

		err := d.Update(word,definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}//TestUpdate

func TestDelete(t *testing.T){
	word:= "test"
	d := Dictionary{word: "test definition"}
	d.Delete(word)

	_, err := d.Search(word)

	if err != ErrNotFound{
		t.Errorf("Expected %q to be deleted", word)
	}//if
}//TestDelete
func TestAdd(t *testing.T){
	t.Run("new word", func(t *testing.T){
		d := Dictionary{}
		word:="test"
		definition:="this is a test"

		err:= d.Add(word, definition)

		assertError(t,err,nil)
		assertDefinition(t,d,word,definition)
	})
	t.Run("exsisting word", func(t *testing.T){
		word:="test"
		definition:="this is a test"
		d := Dictionary{word: definition}
		err:=d.Add(word, "new test")

		assertError(t,err,ErrWordExists)
		assertDefinition(t,d,word,definition)
	})
}//testadd

func TestSearch(t *testing.T){
	d := Dictionary{"test":"this is a test"} 

	t.Run("known word", func(t *testing.T){
		got,_ := d.Search("test")
		want:= "this is a test"

		assertStrings(t,got,want)
	})

	t.Run("unknown word", func(t *testing.T){
		_, err := d.Search("unknown")

		assertError(t,err,ErrNotFound)
	})
}//TestSearch


