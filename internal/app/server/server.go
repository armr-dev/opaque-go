package server

import (
	"github.com/armr-dev/opaque-go/internal/app/db"
	"net/http"
)

var clientRecords = db.Record{}

func InitServer() {
	registrationService := RegistrationService{CredentialId: []byte{}}
	authenticationService := AuthenticationService{}

	http.HandleFunc("/registration-init", registrationService.RegistrationInit)
	http.HandleFunc("/registration-finalize", registrationService.RegistrationFinalize)
	http.HandleFunc("/auth-init", authenticationService.AuthenticationInit)
	http.HandleFunc("/auth-finalize", authenticationService.AuthenticationFinalize)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
