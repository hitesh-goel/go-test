package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if want != got {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaV1(t *testing.T) {
	calculateArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("calucate area for rectangle", func(t *testing.T) {
		shape := Rectangle{10.0, 10.0}
		want := 100.0
		calculateArea(t, shape, want)
	})
	t.Run("calculate area of circle", func(t *testing.T) {
		shape := Circle{10.0}
		want := 628.0
		calculateArea(t, shape, want)
	})
}

func TestArea(t *testing.T) {
	arrTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{2.0, 3.0}, 6.0},
		{Circle{10.0}, 628.0},
	}

	for _, tt := range arrTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
