package repository

import (
	"context"
	db "medivault-service/internal/db/generated"
	models "medivault-service/internal/db/generated"
	"medivault-service/internal/repository"

	"github.com/jackc/pgx/v5"
)

type PatientFileRepoSqlc struct {
	Queries *db.Queries
	conn    *pgx.Conn
}

func NewPatientFileRepositorySqlc(conn *pgx.Conn) repository.PatientFileRepository {
	return &PatientFileRepoSqlc{
		Queries: db.New(conn),
		conn:    conn,
	}
}

func (r *PatientFileRepoSqlc) GetAllPatientFiles(ctx context.Context) ([]models.PatientFile, error) {
	return nil, nil
}
func (r *PatientFileRepoSqlc) GetPatientFileByID(ctx context.Context, id int) (models.PatientFile, error) {
	return models.PatientFile{}, nil
}
func (r *PatientFileRepoSqlc) CreatePatientFile(ctx context.Context, patientID int, fileID int, description string) (models.PatientFile, error) {
	return models.PatientFile{}, nil
}
func (r *PatientFileRepoSqlc) UpdatePatientFile(ctx context.Context, id int, patientID int, fileID int, description string) (models.PatientFile, error) {
	return models.PatientFile{}, nil
}
func (r *PatientFileRepoSqlc) DeletePatientFile(ctx context.Context, id int) error {
	return nil
}
