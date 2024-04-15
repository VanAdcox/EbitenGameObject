package main

import (
	"./sprite"
)

type GameObject struct {
	sprite    Sprite
	transform Transform
}
type Transform struct {
	x, y                     int
	rotation, xScale, yScale float64
}

func main() {

}
