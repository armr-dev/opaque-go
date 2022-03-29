package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/registration-init", registrationInit)
	http.HandleFunc("/registration-finalize", registrationFinalize)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		panic(err)
	}
}
