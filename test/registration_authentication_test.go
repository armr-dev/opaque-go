package test

import (
	"github.com/armr-dev/opaque-go/internal/app/client"
	"testing"
)

func BenchmarkRegistrationAuthentication(b *testing.B) {
	client.Registration()
	client.AuthenticationInit()
}
