# Config
Defines the configuration for running this protocol. 

This configuration is based on the configuration defined in the [draft](https://www.ietf.org/archive/id/draft-irtf-cfrg-opaque-07.html#name-opaque-3dh-real-test-vector), to be used with the test vector to verify everything
is working as expected.


```
OPRF: 0001
Hash: SHA512
MHF: Identity
KDF: HKDF-SHA512
MAC: HMAC-SHA512
Group: ristretto255
Context: 4f50415155452d504f43
Nh: 64
Npk: 32
Nsk: 32
Nm: 64
Nx: 64
Nok: 32
```