package server

import (
	"github.com/armr-dev/opaque-go/internal/app/db"
	"net/http"
)

var clientRecords = db.Record{}

func InitServer() {
	registrationService := RegistrationService{credentialId: []byte{}}
	authentiationService := AuthenticationService{}

	http.HandleFunc("/registration-init", registrationService.registrationInit)
	http.HandleFunc("/registration-finalize", registrationService.registrationFinalize)
	http.HandleFunc("/auth-init", authentiationService.authenticationInit)
	http.HandleFunc("/auth-finalize", authentiationService.authenticationFinalize)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
