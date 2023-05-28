package scheduler

import (
	"fmt"
	"proyecto_final_goland/vehicle"
	"sync"
	"time"
)

// La struct Scheduler guarda una lista de vehículos y una lista de mantenimientos programados.
// También incluye un Mutex para gestionar el acceso concurrente a los datos.
type Scheduler struct {
	Vehicles    []*vehicle.Vehicle
	Maintenance []*vehicle.Maintenance
	Mu          sync.Mutex
}

// NewScheduler es una función para crear un nuevo Scheduler. Recibe una lista de vehículos.
func NewScheduler(vehicles []*vehicle.Vehicle) *Scheduler {
	return &Scheduler{
		Vehicles:    vehicles,
		Maintenance: []*vehicle.Maintenance{},
	}
}

// ScheduleMaintenance es una función que itera sobre todos los vehículos y programa un mantenimiento
// si el vehículo necesita servicio. Utiliza goroutines para hacerlo en paralelo y un WaitGroup para sincronizar.
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

// ScheduleMaintenanceAt es una función que programa un mantenimiento para un vehículo en un tiempo específico.
func (s *Scheduler) ScheduleMaintenanceAt(v *vehicle.Vehicle, t time.Time) {
	maintenance := vehicle.NewMaintenance(v.ID, 100.0, t, 1*time.Hour)
	s.Maintenance = append(s.Maintenance, maintenance)
}

// FinishMaintenance es una función que marca un mantenimiento como finalizado.

func (s *Scheduler) FinishMaintenance(m *vehicle.Maintenance) {
	m.FinishMaintenance()
}

// TotalMaintenanceCost es una función que calcula y devuelve el costo total de todos los mantenimientos programados.
func (s *Scheduler) TotalMaintenanceCost() float64 {
	totalCost := 0.0
	for _, m := range s.Maintenance {
		totalCost += m.Cost
	}
	return totalCost
}
