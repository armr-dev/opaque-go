package db

import "github.com/bytemare/opaque"

type Record struct {
	Clients []opaque.ClientRecord `default:"[]"`
}
