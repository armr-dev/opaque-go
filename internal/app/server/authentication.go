package server

import (
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"github.com/armr-dev/opaque-go/internal/utils"
	"io/ioutil"
	"net/http"
)

func authenticationInit(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	authReq, err := ioutil.ReadAll(req.Body)

	deserializedReq, err := opaque.Server.Deserialize.KE1(authReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ke2, err := opaque.Server.LoginInit(deserializedReq, opaque.ServerId, opaque.ServerSecretKey, opaque.ServerPublicKey, opaque.OPRFSeed, &clientRecords.Clients[0])

	state := opaque.Server.SerializeState()
	ke2s := ke2.Serialize()

	utils.Use(state)

	_, err = w.Write(ke2s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
