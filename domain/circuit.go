package domain

import (
	"encoding/json"
	"log"
)

type Circuit struct {
	RootArrangement Arrangement
}

func NewCircuit(rootArrangement Arrangement) *Circuit {
	if rootArrangement == nil {
		panic("Root arrangement cannot be nil")
	}

	return &Circuit{rootArrangement}
}

func (circuit Circuit) CalculateTotalResistance() float64 {
	return circuit.RootArrangement.CalculateTotalResistance()
}

func (circuit Circuit) PrettyPrint() {
	empJSON, err := json.MarshalIndent(circuit, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Printf("Circuit: %s\n", string(empJSON))
}
