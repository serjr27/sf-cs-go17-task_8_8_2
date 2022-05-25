package auto

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		// сконвертировать value в заданный в параметре UnitType
		if t == CM {
			return value * 2.54
		}

		if t == Inch {
			if value != 0 {
				return value / 2.54
			}
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type dimensions struct {
	lenght Unit
	width  Unit
	height Unit
}

func newDimensions(lenght, width, height float64, unit UnitType) *dimensions {
	return &dimensions{Unit{lenght, unit}, Unit{width, unit}, Unit{height, unit}}
}

func (d dimensions) Length() Unit { return d.lenght }

func (d dimensions) Width() Unit { return d.width }

func (d dimensions) Height() Unit { return d.height }

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type auto struct {
	brand       string
	model       string
	dimensions  dimensions
	maxSpeed    int
	enginePower int
}

func NewAuto(brand, model string, lenght, width, height float64, unit UnitType, speed, power int) *auto {
	return &auto{brand, model, *newDimensions(lenght, width, height, unit), speed, power}
}

func (a auto) Brand() string { return a.brand }

func (a auto) Model() string { return a.model }

func (a auto) Dimensions() Dimensions { return a.dimensions }

func (a auto) MaxSpeed() int { return a.maxSpeed }

func (a auto) EnginePower() int { return a.enginePower }
