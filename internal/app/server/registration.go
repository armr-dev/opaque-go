package main

import (
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"net/http"
)

func registrationInit(w http.ResponseWriter, req *http.Request) {
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
	credentialId := make([]byte, 64)
	pks, err := opaque.Server.Group.NewElement().Decode(serverPublicKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	serverRegistrationResponse, err := opaque.Server.RegistrationResponse(deserializedReq, pks.Bytes(), credentialId, opaque.OPRFSeed)
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

func registrationFinalize(w http.ResponseWriter, req *http.Request) {
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

	err = json.NewEncoder(w).Encode(deserializedRecord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(deserializedRecord)
}
