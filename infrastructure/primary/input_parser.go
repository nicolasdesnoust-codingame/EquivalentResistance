package main

import (
	"bufio"
	"equivalentresistance/usecases"
	"fmt"
)

type InputParser struct {
}

func NewInputParser() *InputParser {
	return &InputParser{}
}

func (inputParser InputParser) ParseInputs(scanner *bufio.Scanner) (*usecases.CodingameCircuit, error) {
	var resistorCount int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &resistorCount)

	var resistors []usecases.CodingameResistor
	for i := 0; i < resistorCount; i++ {
		var id string
		var resistance int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &id, &resistance)
		resistors = append(resistors, usecases.CodingameResistor{id, resistance})
	}

	scanner.Scan()
	definition := scanner.Text()

	return usecases.NewCodingameCircuit(definition, resistors), nil
}
