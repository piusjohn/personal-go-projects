package main

type vehicle struct {
	brand string
	speed float64
}

type car struct {
	vehicle
	doors string
}

type boat struct {
	vehicle
	hasAnchor string
}

type mover interface {
	move() string
}

type describer interface {
	describe() string
}

func (c car) move() string {
	return c.
}