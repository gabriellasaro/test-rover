package plateau

import (
	"fmt"

	"github.com/gabriellasaro/test-rover/coordinates"
	"github.com/gabriellasaro/test-rover/direction"
)

type Rover interface {
	// Commands operações com o rover.
	// L: vire à esquerda; R: vire à direira; M: avançar um ponto na grade. Ex.: LMLMLMLMM
	Commands(commands string) error

	// CurrentPosition obter a posição atual do rover
	CurrentPosition() string

	String() string
}

type rover struct {
	plateau     *plateau
	coordinates coordinates.Coordinates
	position    coordinates.Position
}

func (r *rover) CurrentPosition() string {
	return fmt.Sprintf("%s %s", r.coordinates, r.position)
}

func (r *rover) String() string {
	return r.CurrentPosition()
}

func (r *rover) isValid() error {
	if r.coordinates.X() > r.plateau.upperRight.X() ||
		r.coordinates.Y() > r.plateau.upperRight.Y() {
		return fmt.Errorf("coordinates %s is invalid: limit %s", r.coordinates, r.plateau.upperRight)
	} else if r.coordinates.X() < r.plateau.lowerLeft.X() ||
		r.coordinates.Y() < r.plateau.lowerLeft.Y() {
		return fmt.Errorf("coordinates %s is invalid: minimum %s", r.coordinates, r.plateau.lowerLeft)
	}

	return nil
}

func (r *rover) toNorth() {
	r.coordinates[coordinates.Y]++
}

func (r *rover) toSouth() {
	r.coordinates[coordinates.Y]--
}

func (r *rover) toEast() {
	r.coordinates[coordinates.X]++
}

func (r *rover) toWest() {
	r.coordinates[coordinates.X]--
}

func (r *rover) setNewPosition(d direction.Direction) {
	r.position = r.position.ToDirection(d)
}

func (r *rover) toDirection(d direction.Direction) error {
	r.setNewPosition(d)

	switch r.position {
	case coordinates.North:
		r.toNorth()
	case coordinates.East:
		r.toEast()
	case coordinates.South:
		r.toSouth()
	case coordinates.West:
		r.toWest()
	}

	return nil
}

func (r *rover) Commands(commands string) error {
	for _, command := range commands {
		direction := direction.Direction(command)
		if err := direction.IsValid(); err != nil {
			return err
		}

		if err := r.toDirection(direction); err != nil {
			return err
		}
	}

	return nil
}
