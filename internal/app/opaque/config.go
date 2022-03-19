package opaque

import opaqueLib "github.com/bytemare/opaque"

var DefaultOpaqueConfig = opaqueLib.DefaultConfiguration()
var Client = DefaultOpaqueConfig.Client()
var Server = DefaultOpaqueConfig.Server()

var OPRFSeed = []byte("380d78c283bf98e26334038293e47865922a3b54d3722d8e9ced1c8729c42f5a")
var CredentialIdentifier = []byte("31323334")

var ServerPublicKey = []byte("583f7bccccbc1907ae1506bac950d08266eb3b33ba452b8df7061a390ffd736e")

var ClientSecretKey = []byte("")
