package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, expectedArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		//I should also name fields in the test cases, but my IDE is clever and shows me what they are
		//So in the name of laziness, I'll leave it as it is
		{"Rectangle", Rectangle{10.0, 4.0}, 40.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{5.0, 8.0}, 20.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expectedArea {
				t.Errorf("%#v got %g, expectedArea %g", tt.shape, got, tt.expectedArea)
			}
		})
	}
	//checkArea := func(tb testing.TB, shape Shape, expectedArea float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != expectedArea {
	//		t.Errorf("got %g, expectedArea %g", got, expectedArea)
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
