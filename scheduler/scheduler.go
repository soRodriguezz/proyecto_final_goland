package scheduler

import (
	"fmt"
	"proyecto_final_goland/vehicle"
	"sync"
	"time"
)

type Scheduler struct {
	Vehicles    []*vehicle.Vehicle
	Maintenance []*vehicle.Maintenance
	Mu          sync.Mutex
}

func NewScheduler(vehicles []*vehicle.Vehicle) *Scheduler {
	return &Scheduler{
		Vehicles:    vehicles,
		Maintenance: []*vehicle.Maintenance{},
	}
}

func (s *Scheduler) ScheduleMaintenance() {
	for _, v := range s.Vehicles {
		go func(v *vehicle.Vehicle) {
			if v.NeedsService(time.Now()) {
				s.Mu.Lock()
				// Crea un nuevo mantenimiento y añádelo a la lista de mantenimientos
				maintenance := &vehicle.Maintenance{
					ID:        v.ID,
					Cost:      100.0, // Puedes reemplazar esto por un valor real
					StartTime: time.Now(),
					Duration:  1 * time.Hour, // Puedes reemplazar esto por un valor real
				}
				s.Maintenance = append(s.Maintenance, maintenance)
				fmt.Printf("El vehículo con ID: %s necesita mantenimiento\n", v.ID)
				s.Mu.Unlock()
			}
		}(v)
	}
}
