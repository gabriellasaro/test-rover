package plateau

import (
	"errors"
	"testing"

	"github.com/gabriellasaro/test-rover/coordinates"
)

func TestNewPlateauByParse(t *testing.T) {
	t.Run("Valid upper right coordinates", func(t *testing.T) {
		if _, err := NewPlateauByParse("5 5"); err != nil {
			t.Errorf("NewPlateauByParse() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid upper right coordinates", func(t *testing.T) {
		if _, err := NewPlateauByParse("5"); !errors.Is(err, coordinates.ErrInvalidCoordinates) {
			t.Errorf("NewPlateauByParse() = %v, want %v", err, coordinates.ErrInvalidCoordinates)
		}
	})
}

func TestAddRoverByParse(t *testing.T) {
	plateau := NewPlateau(coordinates.Coordinates{5, 5})

	t.Run("Valid rover", func(t *testing.T) {
		if _, err := plateau.AddRoverByParse("1 2 N"); err != nil {
			t.Errorf("AddRoverByParse() = %v, want %v", err, nil)
		}
	})

	t.Run("Rover with invalid position", func(t *testing.T) {
		if _, err := plateau.AddRoverByParse("1 2 L"); !errors.Is(err, coordinates.ErrInvalidPosition) {
			t.Errorf("AddRoverByParse() = %v, want %v", err, coordinates.ErrInvalidPosition)
		}
	})

	t.Run("Rover with invalid coordinates", func(t *testing.T) {
		if _, err := plateau.AddRoverByParse("1 L N"); !errors.Is(err, coordinates.ErrInvalidCoordinates) {
			t.Errorf("AddRoverByParse() = %v, want %v", err, coordinates.ErrInvalidCoordinates)
		}
	})

	t.Run("Rover with invalid command", func(t *testing.T) {
		if _, err := plateau.AddRoverByParse("1 5N"); !errors.Is(err, ErrRoverCommandIsInvalid) {
			t.Errorf("AddRoverByParse() = %v, want %v", err, ErrRoverCommandIsInvalid)
		}
	})
}
