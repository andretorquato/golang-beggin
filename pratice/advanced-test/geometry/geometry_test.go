package geometry

import "testing"

func TestPoints(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		point := Point{X: 1, Y: 2}
		got := point.Sum()
		want := 3.0
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
	t.Run("Sub", func(t *testing.T) {
		point := Point{X: 1, Y: 2}
		got := point.Sub()
		want := -1.0
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
