package direction

import (
	"errors"
	"fmt"
)

var ErrInvalidDirection = errors.New("invalid direction")

type Direction rune

const (
	Left  Direction = 'L'
	Right Direction = 'R'
	Move  Direction = 'M'
)

func (d Direction) Operation() int8 {
	if d == Left {
		return -1
	} else if d == Right {
		return 1
	}

	return 0
}

func (d Direction) IsValid() error {
	switch d {
	case Left, Right, Move:
		return nil
	}

	return fmt.Errorf("%w: %s", ErrInvalidDirection, string(d))
}
