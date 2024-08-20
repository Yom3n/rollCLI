package dice

import (
	"errors"
	"math/rand/v2"
)

type Dice struct {
	sides uint
}

func (d *Dice) SetSides(sides uint) error {
	if sides == 0 {
		return errors.New("sides can't be 0")
	}
	d.sides = sides
	return nil
}

func (d *Dice) Roll() (uint){
	min := 1
	max := int(d.sides + 1)
	return uint(rand.IntN(max - min) + min)
}