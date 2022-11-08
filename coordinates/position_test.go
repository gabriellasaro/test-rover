package coordinates

import (
	"errors"
	"testing"

	"github.com/gabriellasaro/test-rover/direction"
)

func TestOperation(t *testing.T) {
	tests := []struct {
		name  string
		value Position
		arg   direction.Direction
		want  Position
	}{
		{
			name:  "North to Left",
			value: North,
			arg:   direction.Left,
			want:  West,
		},
		{
			name:  "West to Right",
			value: West,
			arg:   direction.Right,
			want:  North,
		},
		{
			name:  "South to Left",
			value: South,
			arg:   direction.Left,
			want:  East,
		},
		{
			name:  "East to Right",
			value: East,
			arg:   direction.Right,
			want:  South,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.ToDirection(tt.arg); got != tt.want {
				t.Errorf("ToDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePosition(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		if _, err := ParsePosition(commandNorth); err != nil {
			t.Errorf("ParsePosition() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid (non-numeric characters)", func(t *testing.T) {
		if _, err := ParsePosition("L"); !errors.Is(err, ErrInvalidPosition) {
			t.Errorf("ParsePosition() = %v, want %v", err, ErrInvalidPosition)
		}
	})
}
