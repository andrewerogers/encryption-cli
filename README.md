# encryption-cli

Simple encryption cli using AES-GCM. This package uses a nil padding nonce so passwords
are reproducible. In production-grade applications, nonce should be randomly generated. 