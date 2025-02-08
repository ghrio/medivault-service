package repository

import (
	"context"
	models "medivault-service/internal/db/generated"
)

type PatientFileRepository interface {
	GetAllPatientFiles(ctx context.Context) ([]models.PatientFile, error)
	GetPatientFileByID(ctx context.Context, id int) (models.PatientFile, error)
	CreatePatientFile(ctx context.Context, patientID int, fileID int, description string) (models.PatientFile, error)
	UpdatePatientFile(ctx context.Context, id int, patientID int, fileID int, description string) (models.PatientFile, error)
	DeletePatientFile(ctx context.Context, id int) error
}
