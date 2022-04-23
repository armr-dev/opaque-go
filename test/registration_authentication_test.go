package test

import (
	"github.com/armr-dev/opaque-go/internal/app/client"
	"github.com/armr-dev/opaque-go/internal/app/opaque"
	"testing"
)

func BenchmarkRegistrationAuthentication(b *testing.B) {
	client.Registration(opaque.DefaultUsername, opaque.DefaultPassword)
	client.Authentication(opaque.DefaultUsername, opaque.DefaultPassword)
}
