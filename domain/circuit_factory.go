package domain

import (
	"log"
	"regexp"
	"strings"
)

type CircuitFactory struct {
}

func NewCircuitFactory() *CircuitFactory {
	return &CircuitFactory{}
}

func (factory CircuitFactory) CreateCircuit(definition string, resistors map[string]Resistor) (*Circuit, error) {

	var rootArrangement Arrangement
	symbols := strings.Split(definition, " ")
	stack := NewStack()

	for _, symbol := range symbols {
		log.Printf("Handling symbol '%v' :\n", string(symbol))
		log.Printf("\tStack size: %v\n", stack.Size())

		if stack.Size() == 1 {
			rootArrangement = stack.Peek()
		}

		if isResistorIdentifier(symbol) {
			log.Println("\tis resistor identifier")
			factory.handleResistorIdentifier(resistors, symbol, stack)
		} else if isArrangementBeginning(symbol) {
			log.Println("\tis arrangement beginning")
			factory.handleArrangementBeginning(symbol, stack)
		} else if isArrangementEnd(symbol) {
			log.Println("\tis arrangement end")
			factory.handleArrangementEnd(stack)
		}
	}

	return NewCircuit(rootArrangement), nil
}

func (CircuitFactory) handleResistorIdentifier(resistors map[string]Resistor, symbol string, stack *Stack) {
	codingameResistor := resistors[symbol]
	arrangement := *NewMonoResistorArrangement(*NewResistor(codingameResistor.Id, codingameResistor.Resistance))

	parent := stack.Peek()
	parent.AddChild(arrangement)
}

func (CircuitFactory) handleArrangementBeginning(symbol string, stack *Stack) {
	var arrangement ArrangementContainer

	if symbol == "[" {
		arrangement = NewInParallelArrangement([]Arrangement{})
	} else if symbol == "(" {
		arrangement = NewInSeriesArrangement([]Arrangement{})
	}

	if !stack.IsEmpty() {
		parent := stack.Peek()
		parent.AddChild(arrangement)
	}
	stack.Push(arrangement)
}

func (CircuitFactory) handleArrangementEnd(stack *Stack) {
	stack.Pop()
}

func isResistorIdentifier(symbol string) bool {
	match, _ := regexp.MatchString("^[a-zA-Z]*$", symbol)
	return match
}

func isArrangementBeginning(symbol string) bool {
	return symbol == "(" || symbol == "["
}

func isArrangementEnd(symbol string) bool {
	return symbol == ")" || symbol == "]"
}
