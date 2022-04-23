package client

import "github.com/armr-dev/opaque-go/internal/app/opaque"

func InitClient() {
	Registration(opaque.DefaultUsername, opaque.DefaultPassword)
	Authentication(opaque.DefaultUsername, opaque.DefaultPassword)
}
