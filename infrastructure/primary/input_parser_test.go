package main

import (
	"bufio"
	"equivalentresistance/usecases"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInputs_ShouldParseSimpleCircuit(t *testing.T) {
	fileHandle, _ := os.Open("../../resources/test_case_01.txt")
	defer fileHandle.Close()
	scanner := bufio.NewScanner(fileHandle)

	circuit, err := NewInputParser().ParseInputs(scanner)

	if err != nil {
		t.Fatalf(`%v`, err)
	}

	expectedDefinition := "( A B )"
	actualDefinition := circuit.Definition
	if actualDefinition != expectedDefinition {
		t.Fatalf(`Actual definition "%q" does not match expected definition "%q"`, actualDefinition, expectedDefinition)
	}

	expectedResistor := *usecases.NewCodingameResistor("A", 20)
	actualResistor := circuit.Resistors[0]
	if !cmp.Equal(expectedResistor, actualResistor) {
		t.Fatalf(`Actual resitor %+v does not match expected resitor %+v`, actualResistor, expectedResistor)
	}

	expectedResistor = *usecases.NewCodingameResistor("B", 10)
	actualResistor = circuit.Resistors[1]
	if !cmp.Equal(expectedResistor, actualResistor) {
		t.Fatalf(`Actual resitor %+v does not match expected resitor %+v`, actualResistor, expectedResistor)
	}

}
