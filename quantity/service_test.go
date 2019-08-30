package quantity

import "testing"

func TestAddSuccessful(t *testing.T) {
	q1 := &Quantity{
		Unit:  "kg",
		Value: 2,
	}

	q2 := &Quantity{
		Unit:  "g",
		Value: 500,
	}

	s := NewService()

	qTotal, _ := s.Add(q1, q2)

	if qTotal.Value != 2.5 {
		t.Error("Value should be 2.5")
	}

	if qTotal.Unit != "kg" {
		t.Error("Unit should be kg")
	}
}
