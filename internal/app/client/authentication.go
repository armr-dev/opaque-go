package client

import (
	"bytes"
	"encoding/json"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthInit struct {
	Auth     []byte
	Username string
}

type AuthFinish struct {
	M6s        []byte
	SessionKey []byte
}

func Authentication(username, password string) {
	auth := opaque.Client.LoginInit([]byte(password))
	sAuth := auth.Serialize()

	var authInit = AuthInit{
		Auth:     sAuth,
		Username: username,
	}

	postBody, _ := json.Marshal(authInit)
	bufferBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:8090/auth-init", "application/json", bufferBody)
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

	ke3, _, err := opaque.Client.LoginFinish([]byte(username), opaque.ServerId, m5)
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

}
