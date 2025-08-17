package main

import (
	"fmt"
	"parking-slot/internal/model"
	"parking-slot/internal/payments"
	"parking-slot/internal/services"
	"parking-slot/internal/strategies"
	"time"
)

func main() {
	entrances := []*model.Entrance{
		{ID: 1, Name: "Main gate"},
	}

	slots := []*model.Slot{
		{ID: 1, Distance: 5},
		{ID: 2, Distance: 10},
	}

	strategy := &strategies.HourlyCostStrategy{
		Rates: map[model.VehicleType]float64{
			model.TwoWheeler:  10,
			model.FourWheeler: 20,
		},
	}

	lot := services.NewParkingLot(slots, entrances, strategy)

	vehicle := &model.Vehicle{Number: "WB124AB123", Type: model.FourWheeler}

	ticket, err := lot.GenerateTicket(vehicle, entrances[0])
	fmt.Println("Ticket generated", ticket.ID, "for vehicle", vehicle.Number)
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second) // vehicle parked for 20 second

	payment := &payments.CardPayment{CardNumber: "1234567890"}

	if err := lot.Exit(ticket, payment); err != nil {
		panic(err)
	}
}
