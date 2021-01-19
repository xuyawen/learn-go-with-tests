package integers

import "testing"

func TestPerimeter(t *testing.T) {
	got := Rectangle.Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		got := Rectangle.Area(Rectangle{12.0, 6.0})
		want := 72.0
		perimeterAssertCorrectMessage(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		got := Circle.Area(Circle{10.0})
		want := 314.1592653589793
		perimeterAssertCorrectMessage(t, got, want)
	})
}

func perimeterAssertCorrectMessage(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}
