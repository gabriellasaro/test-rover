package coordinates

import (
	"errors"
	"fmt"

	"github.com/gabriellasaro/test-rover/direction"
)

var ErrInvalidPosition = errors.New("invalid position")

type Position int8

const (
	North Position = iota
	East
	South
	West

	commandNorth = "N"
	commandEast  = "E"
	commandSouth = "S"
	commandWest  = "W"
)

func (p Position) String() string {
	switch p {
	case North:
		return commandNorth
	case East:
		return commandEast
	case South:
		return commandSouth
	case West:
		return commandWest
	}

	return ""
}

func (p Position) ToDirection(d direction.Direction) Position {
	p = p + Position(d.Operation())

	if p == -1 {
		return West
	} else if p == 4 {
		return North
	}

	return p
}

func ParsePosition(position string) (Position, error) {
	switch position {
	case commandNorth:
		return North, nil
	case commandEast:
		return East, nil
	case commandSouth:
		return South, nil
	case commandWest:
		return West, nil
	}

	return 0, fmt.Errorf("%w: %s", ErrInvalidPosition, position)
}
