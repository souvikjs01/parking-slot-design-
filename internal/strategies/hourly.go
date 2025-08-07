package strategies

import (
	"parking-slot/internal/model"
	"time"
)

type HourlyCostStrategy struct {
	Rates map[model.VehicleType]float64
}

func (h *HourlyCostStrategy) CalculateCost(entry, exit time.Time, vType model.VehicleType) float64 {
	duration := exit.Sub(entry).Hours()
	return h.Rates[vType] * duration
}
