package main

import "fmt"

// Metal - mass and volume information
type Metal struct {
	mass   float64
	volume float64
}

// Density - return density of metal
func (m *Metal) Density() float64 {
	// density = mass/volume
	return m.mass / m.volume
}

// IsDenser - compare density of two objects
func IsDenser(a, b Dense) bool {
	return a.Density() > b.Density()
}

type Dense interface {
	Density() float64
}

type Gas struct {
	temperature   float64
	molecularMass float64
}

func (m *Gas) Density() float64 {
	return m.molecularMass / m.temperature
}

func main() {
	gold := Metal{478, 24}
	silver := Metal{100, 10}

	result := IsDenser(&gold, &silver)
	if result {
		fmt.Println("gold has higher density than silver")
	} else {
		fmt.Println("silver has higher density than gold")
	}

	oxygen := Gas{34, 300}
	hydrogen := Gas{14, 300}

	result = IsDenser(&oxygen, &hydrogen)
	if result {
		fmt.Println("oxygen is higher in density than hydrogen")
	} else {
		fmt.Println("hydrogen is higher in density than oxygen")
	}
}
