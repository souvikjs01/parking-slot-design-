package services

import (
	"errors"
	"fmt"
	"parking-slot/internal/model"
	"time"
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

func (p *ParkingLot) GenerateTicket(vehicle *model.Vehicle, entrance *model.Entrance) (*model.Ticket, error) {
	slot, err := p.FindAvailableSlots(vehicle.Type, entrance)
	if err != nil {
		return nil, err
	}

	err = slot.Park(vehicle)
	if err != nil {
		return nil, err
	}
	p.ticketSeq++
	return &model.Ticket{
		ID:        p.ticketSeq,
		Vehicle:   vehicle,
		Slot:      slot,
		EntryTime: time.Now(),
	}, nil
}

func (p *ParkingLot) Exit(ticket *model.Ticket, payment model.PaymentMethod) error {
	cost := p.CostStrategy.CalculateCost(ticket.EntryTime, time.Now(), ticket.Vehicle.Type)

	ticket.Cost = cost
	ticket.Paid = true
	ticket.Payment = payment
	ticket.Slot.Unpark()

	fmt.Printf("Vehicle %s exited. Paid $%.2f\n", ticket.Vehicle.Number, cost)
	return nil
}
