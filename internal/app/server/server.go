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
		panic(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	deserializedReq, err := opaque.Server.DeserializeRegistrationRequest(registrationReq)
	if err != nil {
		//panic(err)
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
		panic(err)
	}

	json.NewEncoder(w).Encode(serverRegistrationResponse)
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/registration", registrationInit)

	http.ListenAndServe(":8090", nil)
}
