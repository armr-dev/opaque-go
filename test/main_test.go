package test

import (
	"context"
	"github.com/armr-dev/opaque-go/internal/app/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"testing"
	"time"
)

var registrationService = server.RegistrationService{CredentialId: []byte{}}
var authenticationService = server.AuthenticationService{}

func setup(ctx context.Context) error {
	var err error
	mux := http.NewServeMux()
	mux.Handle("/registration-init", http.HandlerFunc(registrationService.RegistrationInit))
	mux.Handle("/registration-finalize", http.HandlerFunc(registrationService.RegistrationFinalize))
	mux.Handle("/auth-init", http.HandlerFunc(authenticationService.AuthenticationInit))
	mux.Handle("/auth-finalize", http.HandlerFunc(authenticationService.AuthenticationFinalize))

	srv := &http.Server{
		Addr:    ":8090",
		Handler: mux,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	<-ctx.Done()

	log.Printf("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return nil
}

func TestMain(m *testing.M) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		m.Run()
		cancel()
	}()

	if err := setup(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
