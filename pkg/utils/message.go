package utils

import "github.com/jackc/pgx/pgtype"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Timestamp = pgtype.Timestamp
