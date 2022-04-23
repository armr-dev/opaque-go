package opaque

import opaqueLib "github.com/bytemare/opaque"

var DefaultOpaqueConfig = opaqueLib.DefaultConfiguration()
var Client, _ = DefaultOpaqueConfig.Client()
var Server, _ = DefaultOpaqueConfig.Server()

var serverSecretKey, pks = DefaultOpaqueConfig.KeyGen()

var ServerSecretKey = serverSecretKey
var ServerPublicKey = pks

var ServerId = []byte("server")

var OPRFSeed = DefaultOpaqueConfig.GenerateOPRFSeed()

var DefaultUsername = "Usuario"
var DefaultPassword = "Senha"
