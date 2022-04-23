package server

import (
	"bytes"
	"encoding/json"
	"github.com/armr-dev/opaque-go/internal/app/client"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"github.com/armr-dev/opaque-go/internal/utils"
	"log"
	"net/http"
)

type AuthenticationService struct {
	state     []byte
	serverKey []byte
}

func (a *AuthenticationService) AuthenticationInit(w http.ResponseWriter, req *http.Request) {
	var authReq client.AuthInit

	var err = json.NewDecoder(req.Body).Decode(&authReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deserializedReq, err := opaque.Server.Deserialize.KE1(authReq.Auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := utils.FindClient(clientRecords.Clients, authReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ke2, err := opaque.Server.LoginInit(deserializedReq, opaque.ServerId, opaque.ServerSecretKey, opaque.ServerPublicKey, opaque.OPRFSeed, &client)

	state := opaque.Server.SerializeState()
	ke2s := ke2.Serialize()

	a.state = state

	_, err = w.Write(ke2s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *AuthenticationService) AuthenticationFinalize(w http.ResponseWriter, req *http.Request) {
	var authReq = client.AuthFinish{}

	var err = json.NewDecoder(req.Body).Decode(&authReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	m6s := authReq.M6s
	clientSessionKey := authReq.SessionKey

	m6, err := opaque.Server.Deserialize.KE3(m6s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := opaque.Server.LoginFinish(m6); err != nil {
		log.Fatalln(err)
	}

	var serverSessionKey = opaque.Server.SessionKey()

	if !bytes.Equal(serverSessionKey, clientSessionKey) {
		err = json.NewEncoder(w).Encode("User not Authenticated")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	err = json.NewEncoder(w).Encode("User Authenticated")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
