package main

import "testing"

func testhello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("chris", "")
		want := "hello chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := hello("", "")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := hello("elodie", "spanish")
		want := "hola elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := hello("name", "french")
		want := "bonjour name"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
