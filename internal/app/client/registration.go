package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	opaque2 "github.com/bytemare/opaque"
	"github.com/bytemare/opaque/message"
	"io"
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var registrationResponse message.RegistrationResponse
	err = json.Unmarshal(body, &registrationResponse)
	if err != nil {
		log.Fatalln(err)
	}

	clientCredentials := &opaque2.Credentials{
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(uploadResponse.Body)
	uploadBody, err := ioutil.ReadAll(uploadResponse.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(uploadBody))
}
