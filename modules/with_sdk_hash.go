//go:build (sdk_hash || sdk || all) && !no_sdk && !no_sdk_hash

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash"
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash/adler32"
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash/crc32"
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash/crc64"
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash/fnv"
	_ "github.com/ZenLiuCN/engineX/modules/golang/hash/maphash"
)
