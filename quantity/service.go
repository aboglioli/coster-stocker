package quantity

import (
	"errors"

	"github.com/aboglioli/coster-stocker/unit"
)

type Service interface {
	Add(*Quantity, *Quantity) (*Quantity, error)
	Substract(*Quantity, *Quantity) (*Quantity, error)
	MultiplyByScalar(float32, *Quantity) (*Quantity, error)
}

type service struct {
	unitRepository unit.Repository
}

func NewService() Service {
	return &service{
		unitRepository: unit.NewRepository(),
	}
}

func (s *service) Add(q1 *Quantity, q2 *Quantity) (*Quantity, error) {
	u1 := s.unitRepository.FindByName(q1.Unit)
	if u1 == nil {
		return nil, errors.New("Unit does not exist")
	}

	u2 := s.unitRepository.FindByName(q2.Unit)
	if u2 == nil {
		return nil, errors.New("Unit does not exist")
	}

	if u1.Type != u2.Type {
		return nil, errors.New("Incompatible units")
	}

	nQ1 := normalizeQuantity(q1, u1)
	nQ2 := normalizeQuantity(q2, u2)

	total := nQ1 + nQ2
	total = total / u1.Modifier

	return &Quantity{
		Unit:  q1.Unit,
		Value: total,
	}, nil
}

func (s *service) Substract(q1 *Quantity, q2 *Quantity) (*Quantity, error) {
	u1 := s.unitRepository.FindByName(q1.Unit)
	if u1 == nil {
		return nil, errors.New("Unit does not exist")
	}

	u2 := s.unitRepository.FindByName(q2.Unit)
	if u2 == nil {
		return nil, errors.New("Unit does not exist")
	}

	if u1.Type != u2.Type {
		return nil, errors.New("Incompatible units")
	}

	nQ1 := normalizeQuantity(q1, u1)
	nQ2 := normalizeQuantity(q2, u2)

	total := nQ1 - nQ2
	total = total / u1.Modifier

	return &Quantity{
		Unit:  q1.Unit,
		Value: total,
	}, nil
}

func (s *service) MultiplyByScalar(s float32, q *Quantity) (*Quantity, error) {
	return q.Multiply(s)
}

func normalizeQuantity(q *Quantity, u *unit.Unit) float32 {
	return q.Value * u.Modifier
}
