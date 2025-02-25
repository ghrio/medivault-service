package routes

import (
	"medivault-service/internal/handler"
	"net/http"
)

func PatientFileRouter(patientFileHandler *handler.PatientFileHandler) *http.ServeMux {
	patientFileMux := http.NewServeMux()
	patientFileMux.HandleFunc("POST /create", patientFileHandler.CreatePatientFile)
	patientFileMux.HandleFunc("GET /patientFiles/patient/{id}", patientFileHandler.GetAllPatientFilesByPatientID)
	patientFileMux.HandleFunc("UPDATE /{id}", patientFileHandler.UpdatePatientFile)
	patientFileMux.HandleFunc("DELETE /{id}", patientFileHandler.DeletePatientFile)
	return patientFileMux
}
