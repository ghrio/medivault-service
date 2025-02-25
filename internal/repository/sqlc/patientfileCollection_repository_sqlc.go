package repository

import (
	"context"
	"medivault-service/internal/db/generated"
	"medivault-service/internal/repository"

	"github.com/jackc/pgx/v5"
)

type PatientFileCollectionSqlc struct {
	queries *generated.Queries
	conn    *pgx.Conn
}

func NewPatientFileCollectionRepository(conn *pgx.Conn) repository.PatientFileCollectionRepository {
	return &PatientFileCollectionSqlc{queries: generated.New(conn), conn: conn}
}

func (p *PatientFileCollectionSqlc) GetPatientCollectionByUser(ctx context.Context, userID int32) (generated.PatientCollection, error) {
	return generated.PatientCollection{}, nil
}

func (p *PatientFileCollectionSqlc) CreatePatientCollection(ctx context.Context, userID int32) (generated.PatientCollection, error) {
	return generated.PatientCollection{}, nil
}
