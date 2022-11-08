package coordinates

import (
	"errors"
	"testing"
)

func TestCoordinates(t *testing.T) {
	var (
		x int32 = 1
		y int32 = 2
	)

	coordinates := Coordinates{x, y}

	t.Run("Get X", func(t *testing.T) {
		if got := coordinates.X(); got != x {
			t.Errorf("X() = %d, want %d", got, x)
		}
	})

	t.Run("Get Y", func(t *testing.T) {
		if got := coordinates.Y(); got != y {
			t.Errorf("Y() = %d, want %d", got, y)
		}
	})
}

func TestParseCoordinates(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		if _, err := ParseCoordinates("1 2"); err != nil {
			t.Errorf("ParseCoordinates() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid (non-numeric characters)", func(t *testing.T) {
		if _, err := ParseCoordinates("1 d2"); !errors.Is(err, ErrInvalidCoordinates) {
			t.Errorf("ParseCoordinates() = %v, want %v", err, ErrInvalidCoordinates)
		}
	})

	t.Run("Operation (many values)", func(t *testing.T) {
		if _, err := ParseCoordinates("1 2 0"); !errors.Is(err, ErrInvalidCoordinates) {
			t.Errorf("ParseCoordinates() = %v, want %v", err, ErrInvalidCoordinates)
		}
	})
}

func TestParseCoordinatesByXY(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		if _, err := ParseCoordinatesByXY("1", "2"); err != nil {
			t.Errorf("ParseCoordinatesByXY() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid (non-numeric characters)", func(t *testing.T) {
		if _, err := ParseCoordinatesByXY("1", "d2"); !errors.Is(err, ErrInvalidCoordinates) {
			t.Errorf("ParseCoordinatesByXY() = %v, want %v", err, ErrInvalidCoordinates)
		}
	})
}
