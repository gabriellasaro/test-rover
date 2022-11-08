package coordinates

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidCoordinates = errors.New("invalid coordinates")

	EmptyCoordinates Coordinates
)

const (
	X = 0
	Y = 1
)

type Coordinates [2]int32

func (c Coordinates) X() int32 {
	return c[X]
}

func (c Coordinates) Y() int32 {
	return c[Y]
}

func (c Coordinates) String() string {
	return fmt.Sprintf("%d %d", c.X(), c.Y())
}

func ParseCoordinatesByXY(x, y string) (Coordinates, error) {
	xc, err := strconv.ParseInt(x, 10, 32)
	if err != nil {
		return EmptyCoordinates, fmt.Errorf("%w: %s", ErrInvalidCoordinates, err)
	}

	yc, err := strconv.ParseInt(y, 10, 32)
	if err != nil {
		return EmptyCoordinates, fmt.Errorf("%w: %s", ErrInvalidCoordinates, err)
	}

	return Coordinates{int32(xc), int32(yc)}, nil
}

func ParseCoordinates(coordinates string) (Coordinates, error) {
	values := strings.Split(strings.TrimSpace(coordinates), " ")

	if len(values) != 2 {
		return EmptyCoordinates, fmt.Errorf("%w: %s", ErrInvalidCoordinates, coordinates)
	}

	return ParseCoordinatesByXY(values[X], values[Y])
}
