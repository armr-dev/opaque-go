package server

import (
	"encoding/json"
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

	deserializedReq, err := opaque.Server.DeserializeRegistrationRequest(registrationReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var _, serverPublicKey = opaque.Server.KeyGen()
	r.credentialId = make([]byte, 64)
	pks, err := opaque.Server.Group.NewElement().Decode(serverPublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serverRegistrationResponse, err := opaque.Server.RegistrationResponse(deserializedReq, pks.Bytes(), r.credentialId, opaque.OPRFSeed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(serverRegistrationResponse)
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

	deserializedRecord, err := opaque.Server.DeserializeRegistrationRecord(registrationRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	clientRecords.Clients = append(clientRecords.Clients, opaqueLib.ClientRecord{
		CredentialIdentifier: r.credentialId,
		ClientIdentity:       opaque.ClientId,
		RegistrationRecord:   deserializedRecord,
	})
}
