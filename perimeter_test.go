package integers

import "testing"

func TestPerimeter(t *testing.T) {
	got := Rectangle.Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestAreaOne(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		Rectangle := Rectangle{12.0, 6.0}
		checkArea(t, Rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

// 表格驱动测试
func TestAreaTwo(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{width: 12.0, height: 6.0}, hasArea: 72.0},
		{shape: Circle{radius: 10.0}, hasArea: 314.1592653589793},
		{shape: Triangle{base: 12.0, height: 6.0}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("got '%.2f' want '%.2f'. type is '%T'", got, tt.hasArea, tt.shape)
		}
	}
}
