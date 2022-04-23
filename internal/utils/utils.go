package utils

import (
	"errors"
	"github.com/bytemare/opaque"
)

// Used to ignore declared and not used annoying message
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}

func FindClient(clients []opaque.ClientRecord, userName string) (opaque.ClientRecord, error) {
	for i := range clients {
		if string(clients[i].ClientIdentity) == userName {
			return clients[i], nil
		}
	}
	return opaque.ClientRecord{}, errors.New("client not found")
}
