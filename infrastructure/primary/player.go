package main

import (
	"bufio"
	"equivalentresistance/usecases"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	circuit, err := NewInputParser().ParseInputs(scanner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	circuit.PrettyPrint()

	resistance := usecases.CalculateCircuitResistanceUsecase(*circuit)
	fmt.Printf("%.1f\n", resistance)
}
