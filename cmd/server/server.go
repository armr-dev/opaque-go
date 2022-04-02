package main

import (
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"github.com/armr-dev/opaque-go/internal/app/server"
)

func main() {
	opaque.DefaultOpaqueConfig.Context = []byte("TCC")

	server.InitServer()
}
