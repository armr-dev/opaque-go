package crypto

import "crypto/rand"

func GenRandomByte(len int) ([]byte, error) {
	r := make([]byte, len)
	if _, err := rand.Read(r); err != nil {
		return nil, err
	}

	return r, nil
}
