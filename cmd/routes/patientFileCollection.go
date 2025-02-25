package routes

import (
	"medivault-service/internal/handler"
	"net/http"
)

func PatientFileCollectionRouter(patientFileCollectionHandler *handler.PatientFileCollectionHandler) *http.ServeMux {
	patientFileCollectionMux := http.NewServeMux()
	patientFileCollectionMux.HandleFunc("POST /create", patientFileCollectionHandler.AddPatientFileToPatientFileCollection)
	patientFileCollectionMux.HandleFunc("GET /patientCollections", patientFileCollectionHandler.CreatePatientFileCollection)
	patientFileCollectionMux.HandleFunc("GET /patientCollections/patient/{id}", nil)
	return patientFileCollectionMux
}
