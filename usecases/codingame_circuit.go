package usecases

import (
	"encoding/json"
	"log"
)

type CodingameCircuit struct {
	Definition string
	Resistors  []CodingameResistor
}

type CodingameResistor struct {
	Id         string
	Resistance int
}

func NewCodingameCircuit(definition string, resistors []CodingameResistor) *CodingameCircuit {
	if len(definition) == 0 {
		panic("Circuit definition cannot be empty")
	}
	if len(resistors) == 0 {
		panic("Circuit must have at least one resistor")
	}

	return &CodingameCircuit{definition, resistors}
}

func (circuit CodingameCircuit) GetResistorById(id string) CodingameResistor {
	for _, resistor := range circuit.Resistors {
		if resistor.Id == id {
			return resistor
		}
	}

	return CodingameResistor{}
}

func (circuit CodingameCircuit) PrettyPrint() {
	empJSON, err := json.MarshalIndent(circuit, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Circuit: %s\n", string(empJSON))
}

func NewCodingameResistor(id string, resistance int) *CodingameResistor {
	if resistance < 0 {
		panic("Resistor resistance must be a positive integer")
	}

	return &CodingameResistor{id, resistance}
}
