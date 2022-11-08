package testrover

import (
	"testing"

	"github.com/gabriellasaro/test-rover/coordinates"
)

func TestRoverCommands(t *testing.T) {
	plateau := NewPlateau(coordinates.Coordinates{5, 5})

	rover, err := plateau.AddRover(coordinates.Coordinates{1, 2}, coordinates.North)
	if err != nil {
		t.Errorf("AddRover() = %v, want %v", err, nil)
	}

	t.Run("Valid rover commands", func(t *testing.T) {
		if err := rover.Commands("LMLMLMLMM"); err != nil {
			t.Errorf("Commands() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid rover commands", func(t *testing.T) {
		if err := rover.Commands("LRN"); err == nil {
			t.Errorf("Commands() = %v, want %v", err, nil)
		}
	})
}

func TestCurrentPosition(t *testing.T) {
	plateau := NewPlateau(coordinates.Coordinates{5, 5})

	tests := []struct {
		name        string
		coordinates coordinates.Coordinates
		position    coordinates.Position
		command     string
		want        string
	}{
		{
			name:        "Rover (1 2 N)",
			coordinates: coordinates.Coordinates{1, 2},
			position:    coordinates.North,
			command:     "LMLMLMLMM",
			want:        "1 3 N",
		},
		{
			name:        "Rover (3 3 E)",
			coordinates: coordinates.Coordinates{3, 3},
			position:    coordinates.East,
			command:     "MMRMMRMRRM",
			want:        "5 1 E",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover, err := plateau.AddRover(tt.coordinates, tt.position)
			if err != nil {
				t.Errorf("AddRover() = %v, want %v", err, nil)
			}

			if err := rover.Commands(tt.command); err != nil {
				t.Errorf("Commands() = %v, want %v", err, nil)
			}

			if got := rover.CurrentPosition(); got != tt.want {
				t.Errorf("CurrentPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
