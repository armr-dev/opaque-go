package kdf

import (
	"fmt"
	"github.com/armr-dev/opaque-go/internals/crypto"
)

const Nseed = 32
const Nn = 32

const DefaultNonceLength = 32
const DefaultMaskingKeyInfo = "MaskingKey"
const DefaultAuthKeyInfo = "AuthKey"
const DefaultExportKeyInfo = "ExportKey"
const DefaultPrivateKeyInfo = "PrivateKey"

type CleartextCredentials struct {
	serverPublicKey []byte
	serverIdentity  []byte
	clientIdentity  []byte
}

type Envelope struct {
	nonce   []byte
	authTag []byte
}

func Store(c *CleartextCredentials, k *crypto.KDF, randomizedPwd []byte) (e *Envelope, maskingKey []byte, clientPublicKey, exportKey []byte) {
	envelopeNonce, err := crypto.GenRandomByte(DefaultNonceLength)

	if err != nil {
		panic(fmt.Errorf("Error generating random bytes: %w", err))
	}

	maskingKey := k.Expand(randomizedPwd, []byte(DefaultMaskingKeyInfo), DefaultNonceLength)
	authKey := k.Expand(randomizedPwd, append(envelopeNonce, []byte(DefaultAuthKeyInfo)...), DefaultNonceLength)
	exportKey := k.Expand(randomizedPwd, append(envelopeNonce, []byte(DefaultExportKeyInfo)...), DefaultNonceLength)

	seed := k.Expand(randomizedPwd, append(envelopeNonce, []byte(DefaultPrivateKeyInfo)...), Nseed)

}

func CreateCleartextCredentials(serverPublicKey []byte, serverIdentity []byte, clientPublicKey []byte, clientIdentity []byte) *CleartextCredentials {
	if serverIdentity == 0 {
		serverIdentity = serverPublicKey
	}
	if clientIdentity == 0 {
		clientIdentity = clientPublicKey
	}

	return &CleartextCredentials{
		serverPublicKey,
		serverIdentity,
		clientIdentity,
	}
}
