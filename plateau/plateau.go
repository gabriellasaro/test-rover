package plateau

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gabriellasaro/test-rover/coordinates"
)

var ErrRoverCommandIsInvalid = errors.New("rover command is invalid")

type Plateau interface {
	// AddRoverByParse adiciona um novo rover no plateau. Ex.: 1 2 N
	AddRoverByParse(command string) (Rover, error)

	AddRover(coordinates coordinates.Coordinates, position coordinates.Position) (Rover, error)
}

type plateau struct {
	upperRight coordinates.Coordinates
	lowerLeft  coordinates.Coordinates
}

func NewPlateauByParse(upperRight string) (Plateau, error) {
	c, err := coordinates.ParseCoordinates(upperRight)
	if err != nil {
		return nil, err
	}

	return NewPlateau(c), nil
}

func NewPlateau(upperRight coordinates.Coordinates) Plateau {
	return &plateau{
		upperRight: upperRight,
		lowerLeft:  coordinates.Coordinates{0, 0},
	}
}

func (p *plateau) AddRover(coordinates coordinates.Coordinates, position coordinates.Position) (Rover, error) {
	rover := &rover{
		plateau:     p,
		coordinates: coordinates,
		position:    position,
	}

	if err := rover.isValid(); err != nil {
		return nil, err
	}

	return rover, nil
}

func (p *plateau) AddRoverByParse(command string) (Rover, error) {
	values := strings.Split(strings.TrimSpace(command), " ")

	if len(values) != 3 {
		return nil, fmt.Errorf("%w: %s", ErrRoverCommandIsInvalid, command)
	}

	xy, err := coordinates.ParseCoordinatesByXY(values[coordinates.X], values[coordinates.Y])
	if err != nil {
		return nil, err
	}

	position, err := coordinates.ParsePosition(values[2])
	if err != nil {
		return nil, err
	}

	return p.AddRover(xy, position)
}
