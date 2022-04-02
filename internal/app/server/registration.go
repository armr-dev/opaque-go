package server

import (
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	opaqueLib "github.com/bytemare/opaque"
	"net/http"
)

type RegistrationService struct {
	credentialId []byte
}

func (r *RegistrationService) registrationInit(w http.ResponseWriter, req *http.Request) {
	var registrationReq []byte

	var err = json.NewDecoder(req.Body).Decode(&registrationReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deserializedReq, err := opaque.Server.Deserialize.RegistrationRequest(registrationReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	r.credentialId = make([]byte, 32)
	pks, err := opaque.Server.Deserialize.DecodeAkePublicKey(opaque.ServerPublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serverRegistrationResponse := opaque.Server.RegistrationResponse(deserializedReq, pks, r.credentialId, opaque.OPRFSeed)
	serializedResponse := serverRegistrationResponse.Serialize()

	err = json.NewEncoder(w).Encode(serializedResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (r *RegistrationService) registrationFinalize(w http.ResponseWriter, req *http.Request) {
	var registrationRecord []byte

	var err = json.NewDecoder(req.Body).Decode(&registrationRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deserializedRecord, err := opaque.Server.Deserialize.RegistrationRecord(registrationRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clientRecords.Clients = append(clientRecords.Clients, opaqueLib.ClientRecord{
		CredentialIdentifier: r.credentialId,
		ClientIdentity:       opaque.ClientId,
		RegistrationRecord:   deserializedRecord,
	})

	fmt.Println("User registered successfully!")
}
