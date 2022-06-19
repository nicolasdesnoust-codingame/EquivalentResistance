package usecases

import (
	"testing"
)

func TestCalculateCircuitResistance_ShouldReturnValidResultForSimpleCircuit(t *testing.T) {

	codingameCircuit := *NewCodingameCircuit(
		"( ( A B ) ( C A ) )",
		[]CodingameResistor{
			*NewCodingameResistor("A", 24),
			*NewCodingameResistor("B", 8),
			*NewCodingameResistor("C", 48)})

	actualResistance := CalculateCircuitResistanceUsecase(codingameCircuit)

	expectedResistance := 104.0
	if actualResistance != expectedResistance {
		t.Fatalf(`Actual resistance "%f" does not match expected resistance "%f"`, actualResistance, expectedResistance)
	}
}
