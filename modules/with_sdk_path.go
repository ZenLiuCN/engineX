//go:build (sdk_path || sdk || all) && !no_sdk && !no_sdk_path

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/path"
	_ "github.com/ZenLiuCN/engineX/modules/golang/path/filepath"
)
