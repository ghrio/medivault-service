package services

import (
	"context"
	"errors"
	models "medivault-service/internal/db/generated"
	repository "medivault-service/internal/repository"
)

type patientFileCollectionService struct {
	patientFileCollectionRepo repository.PatientFileCollectionRepository
}

func NewPatientFileCollectionService(repo repository.PatientFileCollectionRepository) PatientFileCollectionService {
	return &patientFileCollectionService{
		patientFileCollectionRepo: repo,
	}
}

func (pfcs *patientFileCollectionService) GetPatientCollectionByID(ctx context.Context, id int) (*models.PatientCollection, error) {
	return nil, errors.New("Not implemented")
}
func (pfcs *patientFileCollectionService) GetAllPatientCollections(ctx context.Context) ([]*models.PatientCollection, error) {
	return nil, errors.New("Not implemented")
}
func (pfcs *patientFileCollectionService) CreatePatientCollection(ctx context.Context, patientCollection *models.PatientCollection) error {
	return errors.New("Not implemented")
}
func (pfcs *patientFileCollectionService) UpdatePatientCollection(ctx context.Context, patientCollection *models.PatientCollection) error {
	return errors.New("Not implemented")
}
func (pfcs *patientFileCollectionService) DeletePatientCollection(ctx context.Context, id int) error {
	return errors.New("Not implemented")
}
