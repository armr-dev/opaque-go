package opaque

import opaqueLib "github.com/bytemare/opaque"

var DefaultOpaqueConfig = opaqueLib.DefaultConfiguration()
var Client, _ = DefaultOpaqueConfig.Client()
var Server, _ = DefaultOpaqueConfig.Server()

var serverSecretKey, pks = DefaultOpaqueConfig.KeyGen()

var ServerSecretKey = serverSecretKey
var ServerPublicKey = pks

var ServerId = []byte("server")
var ClientId = []byte("client")

var OPRFSeed = DefaultOpaqueConfig.GenerateOPRFSeed()
