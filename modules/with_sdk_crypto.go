//go:build (sdk_crypto || sdk || all) && !no_sdk && !no_sdk_crypto

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/aes"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/cipher"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/des"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/dsa"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/ecdh"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/ecdsa"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/ed25519"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/elliptic"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/hmac"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/md5"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/rand"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/rc4"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/rsa"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/sha1"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/sha256"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/sha512"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/subtle"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/tls"
	_ "github.com/ZenLiuCN/engineX/modules/golang/crypto/x509"
)
