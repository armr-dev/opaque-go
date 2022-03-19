package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"io/ioutil"
	"log"
	"net/http"
)

func registrationReq() {
	request := opaque.Client.RegistrationInit([]byte("senha"))
	serializedRequest := request.Serialize()

	postBody, _ := json.Marshal(serializedRequest)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/registration", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	registrationResponse := string(body)

	fmt.Println(registrationResponse)

	//record, exportKey := opaque.Client.RegistrationFinalize("senha", _, registrationResponse)
}

func main() {
	registrationReq()
}
