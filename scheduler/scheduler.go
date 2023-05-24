package scheduler

import (
	"fmt"
	"sync"
	"time"

	"proyecto_final_goland/vehicle"
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
	var wg sync.WaitGroup

	for _, v := range s.Vehicles {
		wg.Add(1)
		go func(v *vehicle.Vehicle) {
			defer wg.Done()
			if v.NeedsService(time.Now()) {
				s.Mu.Lock()
				maintenance := vehicle.NewMaintenance(v.ID, 100.0, time.Now(), 1*time.Hour)
				s.Maintenance = append(s.Maintenance, maintenance)
				fmt.Printf("El vehículo con ID: %s necesita mantenimiento\n", v.ID)
				s.Mu.Unlock()
			}
		}(v)
	}
	wg.Wait()
}

func (s *Scheduler) ScheduleMaintenanceAt(v *vehicle.Vehicle, t time.Time) {
	maintenance := vehicle.NewMaintenance(v.ID, 100.0, t, 1*time.Hour)
	s.Maintenance = append(s.Maintenance, maintenance)
}

func (s *Scheduler) FinishMaintenance(m *vehicle.Maintenance) {
	m.FinishMaintenance()
}

func (s *Scheduler) TotalMaintenanceCost() float64 {
	totalCost := 0.0
	for _, m := range s.Maintenance {
		totalCost += m.Cost
	}
	return totalCost
}
