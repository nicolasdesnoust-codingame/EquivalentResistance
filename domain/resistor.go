package domain

type Resistor struct {
	Id         string
	Resistance int
}

func NewResistor(id string, resistance int) *Resistor {
	if resistance < 0 {
		panic("Resistor resistance must be a positive integer")
	}

	return &Resistor{id, resistance}
}
