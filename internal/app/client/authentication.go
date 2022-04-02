package client

import (
	"bytes"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"io/ioutil"
	"log"
	"net/http"
)

func authenticationInit() {
	password := []byte("senha")

	auth := opaque.Client.LoginInit(password)
	requestBody := bytes.NewReader(auth.Serialize())

	resp, err := http.Post("http://localhost:8090/auth-init", "application/json", requestBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var m5s []byte
	defer resp.Body.Close()
	m5s, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	m5, err := opaque.Client.Deserialize.KE2(m5s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(m5)
	//
	//auth2, exportKeyLogin, err := opaque.Client.LoginFinish(opaque.ClientId, opaque.ServerId, deserializedResponse)
	//serializedRequest2 := auth2.Serialize()
	//
	//postBody2, _ := json.Marshal(serializedRequest2)
	//responseBody2 := bytes.NewBuffer(postBody2)
	//resp2, err := http.Post("http://localhost:8090/auth-init", "application/json", responseBody2)
	//if err != nil {
	//	log.Fatalf("An Error Occured %v", err)
	//}
	//
	//fmt.Println(resp2)
	//fmt.Println(exportKeyLogin)
}
