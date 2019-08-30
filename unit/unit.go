package unit

type Unit struct {
	Type     string
	Name     string
	Modifier float32
}

func NewUnit(t string, n string, m float32) *Unit {
	return &Unit{
		Type:     t,
		Name:     n,
		Modifier: m,
	}
}
