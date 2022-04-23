package client

import (
	"bytes"
	"encoding/json"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"github.com/armr-dev/opaque-go/internal/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ClientRegistration struct {
	Upload   []byte
	Username string
}

func Registration(username, password string) {
	request := opaque.Client.RegistrationInit([]byte(password))
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

	var registrationResponse []byte
	err = json.Unmarshal(body, &registrationResponse)

	m2, err := opaque.Client.Deserialize.RegistrationResponse(registrationResponse)

	upload, key := opaque.Client.RegistrationFinalize(m2, []byte(username), opaque.ServerId)
	exportKeyReg := key

	utils.Use(exportKeyReg)

	serializedUpload := upload.Serialize()
	var regFinalize = ClientRegistration{
		Upload:   serializedUpload,
		Username: username,
	}

	postUploadBody, _ := json.Marshal(regFinalize)
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

	utils.Use(uploadBody)
}
