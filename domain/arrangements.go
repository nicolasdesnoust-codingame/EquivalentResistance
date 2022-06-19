package domain

type Arrangement interface {
	CalculateTotalResistance() float64
}

type ArrangementContainer interface {
	Arrangement
	AddChild(child Arrangement)
}

type InSeriesArrangement struct {
	Children []Arrangement
}

func NewInSeriesArrangement(children []Arrangement) *InSeriesArrangement {
	return &InSeriesArrangement{children}
}

func (arrangement InSeriesArrangement) CalculateTotalResistance() float64 {
	totalResistance := 0.0

	for _, child := range arrangement.Children {
		totalResistance += child.CalculateTotalResistance()
	}

	return totalResistance
}

func (container *InSeriesArrangement) AddChild(child Arrangement) {
	container.Children = append(container.Children, child)
}

type InParallelArrangement struct {
	Children []Arrangement
}

func NewInParallelArrangement(children []Arrangement) *InParallelArrangement {
	return &InParallelArrangement{children}
}

func (container *InParallelArrangement) AddChild(child Arrangement) {
	container.Children = append(container.Children, child)
}

func (arrangement InParallelArrangement) CalculateTotalResistance() float64 {
	temp := 0.0

	for _, child := range arrangement.Children {
		temp += 1 / child.CalculateTotalResistance()
	}

	return 1 / temp
}

type MonoResistorArrangement struct {
	Child Resistor
}

func NewMonoResistorArrangement(child Resistor) *MonoResistorArrangement {
	return &MonoResistorArrangement{child}
}

func (arrangement MonoResistorArrangement) CalculateTotalResistance() float64 {
	return float64(arrangement.Child.Resistance)
}
