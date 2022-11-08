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

// normalizePosition considera que o rover não pode extrapolar o limite da borda
func (r *rover) normalizePosition(xy int) {
	if r.coordinates[xy] > r.plateau.upperRight[xy] {
		r.coordinates[xy] = r.plateau.upperRight[xy]
	} else if r.coordinates[xy] < r.plateau.lowerLeft[xy] {
		r.coordinates[xy] = r.plateau.lowerLeft[xy]
	}
}

func (r *rover) toNorth() {
	r.coordinates[coordinates.Y]++
	r.normalizePosition(coordinates.Y)
}

func (r *rover) toSouth() {
	r.coordinates[coordinates.Y]--
	r.normalizePosition(coordinates.Y)

}

func (r *rover) toEast() {
	r.coordinates[coordinates.X]++
	r.normalizePosition(coordinates.X)
}

func (r *rover) toWest() {
	r.coordinates[coordinates.X]--
	r.normalizePosition(coordinates.X)
}

func (r *rover) setNewPosition(d direction.Direction) {
	r.position = r.position.ToDirection(d)
}

func (r *rover) move(d direction.Direction) {
	if d.Operation() != 0 {
		return
	}

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
}

func (r *rover) toDirection(d direction.Direction) error {
	if err := d.IsValid(); err != nil {
		return err
	}

	r.setNewPosition(d)
	r.move(d)

	return nil
}

func (r *rover) Commands(commands string) error {
	for _, command := range commands {
		if err := r.toDirection(
			direction.Direction(command),
		); err != nil {
			return err
		}
	}

	return nil
}
