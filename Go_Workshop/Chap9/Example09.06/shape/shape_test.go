package shape

import "testing"

func TestGetName(t *testing.T) {
	triangle := Triangle{Base: 10, Height: 10}
	rectangle := Rectangle{Length: 10, Width: 10}
	square := Square{Side: 10}

	if name := GetName(triangle); name != "Triangle" {
		t.Errorf("GetName Fail: %T", triangle)
	}

	if name := GetName(rectangle); name != "Rectangle" {
		t.Errorf("GetName Fail: %T", rectangle)
	}

	if name := GetName(square); name != "Square" {
		t.Errorf("GetName Fail: %T", square)
	}
}

func TestGetArea(t *testing.T) {
	triangle := Triangle{Base: 10, Height: 10}
	rectangle := Rectangle{Length: 10, Width: 10}
	square := Square{Side: 10}

	if area := GetArea(triangle); area != float64(50) {
		t.Errorf("GetArea Fail: %T", triangle)
	}

	if area := GetArea(rectangle); area != float64(100) {
		t.Errorf("GetArea Fail: %T", rectangle)
	}

	if area := GetArea(square); area != float64(100) {
		t.Errorf("GetArea Fail: %T", square)
	}

}
