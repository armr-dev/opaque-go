package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthFinish struct {
	M6s        []byte
	SessionKey []byte
}

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

	ke3, exportKeyLogin, err := opaque.Client.LoginFinish(opaque.ClientId, opaque.ServerId, m5)
	if err != nil {
		log.Fatalln(err)
	}
	if ke3 == nil {
		log.Fatalln("User not found")
	}

	sessionKey := opaque.Client.SessionKey()
	m6s := ke3.Serialize()

	body := AuthFinish{m6s, sessionKey}
	marshalledBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	buffer := bytes.NewBuffer(marshalledBody)

	resp2, err := http.Post("http://localhost:8090/auth-finalize", "application/json", buffer)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	var response string
	err = json.NewDecoder(resp2.Body).Decode(&response)

	fmt.Println(response)
	fmt.Println(exportKeyLogin)
}
