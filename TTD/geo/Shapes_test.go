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
	checkArea:= func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got:= shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("wow rectangles", func(t *testing.T){
		r := Rectangle{5.0,5.0}
		checkArea(t, r, 25.0)
	})
	t.Run("Circle Area", func(t *testing.T){
		c := Circle{10}
		checkArea(t,c,314.1592653589793)
	})
}

func TestArea (t *testing.T){
	areaTests:= []struct{
		shape Shape
		want float64
	}{
		{Rectangle{12,6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _,tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
