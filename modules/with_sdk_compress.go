//go:build (sdk_compress || sdk || all) && !no_sdk && !no_sdk_compress

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/bzip2"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/flate"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/gzip"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/lzw"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/zlib"
)
