package main

import (
	"github.com/armr-dev/opaque-go/internal/app/opaque"
)

func main() {
	client := opaque.DefaultOpaqueConfig.Client()
	server := opaque.DefaultOpaqueConfig.Server()

}
