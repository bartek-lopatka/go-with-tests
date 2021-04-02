package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 4.0}, 40.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{5.0, 8.0}, 20.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
	//checkArea := func(tb testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g, want %g", got, want)
	//	}
	//}
	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{6.0, 4.0}
	//	checkArea(t, rectangle, 24.0)
	//})
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkArea(t, circle, 314.1592653589793)
	//})
}
