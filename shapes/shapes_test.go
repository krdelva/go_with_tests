package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	}

	t.Run("calculate perimeter of a rectangle given height and width", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v: got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}

}
