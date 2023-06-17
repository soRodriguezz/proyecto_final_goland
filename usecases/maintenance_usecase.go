package usecases

import (
	"proyecto_final_goland/maintenance"
	"proyecto_final_goland/models"
)

func CreateMaintenanceUseCase() {
	newmaintenance := models.NewMaintenance("SC-45-12", 3, "Jhon", "9234234")
	maintenance.MaintenanceArr = append(maintenance.MaintenanceArr, newmaintenance)
}
