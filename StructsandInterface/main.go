package main

import (
	"fmt"
)

type engine interface {
	milesleft() uint8
}
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type owner struct {
	name string
}

func (e gasEngine) milesleft() uint8 {
	return e.gallons * e.mpg
}
type electricEngine struct {
	mpkwh uint8 // miles per kWh
	kwh   uint8 // battery charge in kWh
	model string
}

func (e electricEngine) milesleft() uint8 {
	return e.kwh * e.mpkwh
}

func main() {
	gas := gasEngine{25, 15, owner{"Alex"}, 10}
	electric := electricEngine{4, 20, "Tesla Model 3"}
	var engines []engine = []engine{gas, electric}

	for _, eng := range engines {
		fmt.Printf("Miles left: %d\n", eng.milesleft())
	}
}
