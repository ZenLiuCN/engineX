//go:build (sdk || all) && !no_sdk

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang"
	_ "github.com/ZenLiuCN/engineX/modules/golang/archive/tar"
	_ "github.com/ZenLiuCN/engineX/modules/golang/archive/zip"
	_ "github.com/ZenLiuCN/engineX/modules/golang/bufio"
	_ "github.com/ZenLiuCN/engineX/modules/golang/bytes"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/bzip2"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/flate"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/gzip"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/lzw"
	_ "github.com/ZenLiuCN/engineX/modules/golang/compress/zlib"
)
