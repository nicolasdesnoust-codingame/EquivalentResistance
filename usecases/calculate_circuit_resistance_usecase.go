package usecases

import (
	"equivalentresistance/domain"
	"log"
)

func CalculateCircuitResistanceUsecase(codingameCircuit CodingameCircuit) float64 {
	resistors := buildResistorMap(codingameCircuit.Resistors)

	circuit, err := domain.NewCircuitFactory().CreateCircuit(codingameCircuit.Definition, resistors)

	if err != nil {
		log.Fatalf(err.Error())
	}
	circuit.PrettyPrint()

	return circuit.CalculateTotalResistance()
}

func buildResistorMap(resistors []CodingameResistor) map[string]domain.Resistor {
	var resistorMap = make(map[string]domain.Resistor)

	for _, resistor := range resistors {
		resistorMap[resistor.Id] = *domain.NewResistor(resistor.Id, resistor.Resistance)
	}

	return resistorMap
}
