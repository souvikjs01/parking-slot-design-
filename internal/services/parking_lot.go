package services

import (
	"errors"
	"parking-slot/internal/model"
)

type ParkingLot struct {
	Slots        []*model.Slot
	Entrances    []*model.Entrance
	CostStrategy model.CostStrategy
	ticketSeq    int
}

func NewParkingLot(slots []*model.Slot, entrance []*model.Entrance, strategy model.CostStrategy) *ParkingLot {
	return &ParkingLot{
		Slots:        slots,
		Entrances:    entrance,
		CostStrategy: strategy,
		ticketSeq:    0,
	}
}

func (p *ParkingLot) FindAvailableSlots(vType model.VehicleType, entrance *model.Entrance) (*model.Slot, error) {
	var nearest *model.Slot
	for _, slot := range p.Slots {
		if !slot.IsOccupied && (nearest == nil || slot.Distance < nearest.Distance) {
			nearest = slot
		}
	}

	if nearest == nil {
		return nil, errors.New("no slot available")
	}
	return nearest, nil
}
