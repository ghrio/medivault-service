package services

import (
	"context"
	models "medivault-service/internal/db/generated"
	"medivault-service/internal/repository"
)

type patientFileService struct {
	patientFileRepo repository.PatientFileRepository
}

func NewPatientFileService(repo repository.PatientFileRepository) PatientFileService {
	return &patientFileService{
		patientFileRepo: repo,
	}
}

func (s *patientFileService) CreateFile(ctx context.Context, file *models.PatientFile) error {
	panic("not implemented") // TODO: Implement
}

func (s *patientFileService) GetFileByID(ctx context.Context, id int) (*models.PatientFile, error) {
	panic("not implemented") // TODO: Implement
}

func (s *patientFileService) GetAllFiles(ctx context.Context) ([]*models.PatientFile, error) {
	panic("not implemented") // TODO: Implement
}

func (s *patientFileService) UpdateFile(ctx context.Context, file *models.PatientFile) error {
	panic("not implemented") // TODO: Implement
}
