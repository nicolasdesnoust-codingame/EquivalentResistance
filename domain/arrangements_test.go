package domain

import (
	"testing"
)

func TestMonoResistorArrangement_ShouldReturnChildResistance(t *testing.T) {
	arrangement := NewMonoResistorArrangement(*NewResistor("A", 20))

	expectedResistance := 20.0
	actualResistance := arrangement.CalculateTotalResistance()
	if actualResistance != expectedResistance {
		t.Fatalf(`Actual resistance "%f" does not match expected resistance "%f"`, actualResistance, expectedResistance)
	}
}

func TestInSeriesArrangement_ShouldCalculateTotalResistanceOfTwoResistances(t *testing.T) {
	children := []Arrangement{
		NewMonoResistorArrangement(*NewResistor("A", 10)),
		NewMonoResistorArrangement(*NewResistor("B", 20)),
	}
	arrangement := NewInSeriesArrangement(children)

	expectedResistance := 30.0
	actualResistance := arrangement.CalculateTotalResistance()
	if actualResistance != expectedResistance {
		t.Fatalf(`Actual resistance "%f" does not match expected resistance "%f"`, actualResistance, expectedResistance)
	}
}

func TestInParallelArrangementt_ShouldCalculateTotalResistanceOfTwoResistances(t *testing.T) {
	children := []Arrangement{
		NewMonoResistorArrangement(*NewResistor("A", 10)),
		NewMonoResistorArrangement(*NewResistor("B", 20)),
	}
	arrangement := NewInParallelArrangement(children)

	expectedResistance := 1 / 0.15
	actualResistance := arrangement.CalculateTotalResistance()
	epsilon := 0.01
	if actualResistance-expectedResistance > epsilon {
		t.Fatalf(`Actual resistance "%f" does not match expected resistance "%f"`, actualResistance, expectedResistance)
	}
}
