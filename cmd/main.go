package main

import (
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

	time.Sleep(20 * time.Second) // vehicle parked for 20 second

	payment := &payments.CardPayment{CardNumber: "1234567890"}

}
