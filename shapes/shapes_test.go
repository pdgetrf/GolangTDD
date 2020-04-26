package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{
		Height: 10,
		Width:  10,
	}

	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("Rectangle Area", func(t *testing.T) {
		rect := Rectangle{
			Height: 12,
			Width:  6,
		}
		want := 72.0

		checkArea(t, rect, want)
	})

	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)

	})
}

func TestAreaTable(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{
			Rectangle{
				Height: 12,
				Width:  6},
			72.0,
		},
		{
			Circle{
				Radius: 10,
			},
			314.1592653589793,
		},
		{
			Triangle{
				Base:   12,
				Height: 6,
			},
			36.0,
		},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
