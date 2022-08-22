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

func TestArea (t *testing.T){
	areaTests:= []struct{
		name string
		shape Shape
		hasArea float64
	}{
		{name:"Rectagle",shape:Rectangle{Width:12,Height:6},hasArea: 72.0},
		{name:"Circle", shape: Circle{Radius:10}, hasArea: 314.1592653589793},
		{name:"Triangle", shape: Triangle{Base: 12,Height: 6}, hasArea:36.0},
	}
	for _,tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.hasArea{
			t.Errorf("shape %v, got %g, want %g",tt.shape, got, tt.hasArea)
		}
	}
}
