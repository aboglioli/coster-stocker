package quantity

type Quantity struct {
	Unit  string
	Value float32
}

func (q *Quantity) Multiply(s float32) *Quantity {
	return &Quantity{
		Unit:  q.Unit,
		Value: s * q.Value,
	}
}
