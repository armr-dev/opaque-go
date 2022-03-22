package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	opaqueLib "github.com/bytemare/opaque"
	"github.com/bytemare/opaque/message"
	"io/ioutil"
	"log"
	"net/http"
)

func registrationReq() {
	request := opaque.Client.RegistrationInit([]byte("senha"))
	serializedRequest := request.Serialize()

	postBody, _ := json.Marshal(serializedRequest)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/registration-init", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var registrationResponse message.RegistrationResponse
	json.Unmarshal(body, &registrationResponse)

	clientCredentials := &opaqueLib.Credentials{
		Client: opaque.ClientId,
		Server: opaque.ServerId,
	}

	record, _, err := opaque.Client.RegistrationFinalize([]byte("senha"), clientCredentials, &registrationResponse)
	if err != nil {
		log.Fatalln(err)
	}

	serializedUpload := record.Serialize()
	postUploadBody, _ := json.Marshal(serializedUpload)
	responseUploadBody := bytes.NewBuffer(postUploadBody)

	uploadResponse, err := http.Post("http://localhost:8090/registration-finalize", "application/json", responseUploadBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer uploadResponse.Body.Close()
	uploadBody, err := ioutil.ReadAll(uploadResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(uploadBody))
}

func main() {
	registrationReq()
}
