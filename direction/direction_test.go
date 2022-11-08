package direction

import (
	"errors"
	"testing"
)

func TestOperation(t *testing.T) {
	tests := []struct {
		name  string
		value Direction
		want  int8
	}{
		{
			name:  "Operation +1",
			value: 'R',
			want:  1,
		},
		{
			name:  "Operation -1",
			value: 'L',
			want:  -1,
		},
		{
			name:  "Operation 0",
			value: 'M',
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.Operation(); got != tt.want {
				t.Errorf("Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		if err := Direction('R').IsValid(); err != nil {
			t.Errorf("Operation() = %v, want %v", err, nil)
		}
	})

	t.Run("Invalid", func(t *testing.T) {
		if err := Direction('V').IsValid(); !errors.Is(err, ErrInvalidDirection) {
			t.Errorf("Operation() = %v, want %v", err, ErrInvalidDirection)
		}
	})
}
