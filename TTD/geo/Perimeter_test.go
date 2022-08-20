package main

import(
	"testing"
)

func TestPerimeter(t *testing.T){
	r := Rectangle{5.0,5.0}
	got := Perimeter(r)
	want:= 20.0
	if want != got{
		t.Errorf("got %.2f want %.2f", got,want)
	}//if	
}//TestPerimeter

func TestArea(t *testing.T){
	r := Rectangle{5.0,5.0}
	got := Area(r)
	want:= 25.0
	if got != want{
		t.Errorf("got %2f want %2f",got,want)
	}
}
