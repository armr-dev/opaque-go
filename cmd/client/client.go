package main

import (
	"github.com/armr-dev/opaque-go/internal/app/client"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
)

func main() {
	opaque.DefaultOpaqueConfig.Context = []byte("TCC")

	client.InitClient()
}
