package model

import "errors"

type Slot struct {
	ID         int
	Vehicle    *Vehicle
	IsOccupied bool
	Distance   int
}

// two methods:

func (s *Slot) Park(v *Vehicle) error {
	if s.IsOccupied {
		return errors.New("Slot already occupied")
	}

	s.Vehicle = v
	s.IsOccupied = true

	return nil
}

func (s *Slot) Unpark() {
	s.Vehicle = nil
	s.IsOccupied = false
}
