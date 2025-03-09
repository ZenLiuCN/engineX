//go:build (sdk_image || sdk || all) && !no_sdk && !no_sdk_image

package modules

import (
	_ "github.com/ZenLiuCN/engineX/modules/golang/image"
	_ "github.com/ZenLiuCN/engineX/modules/golang/image/color"
	_ "github.com/ZenLiuCN/engineX/modules/golang/image/draw"
	_ "github.com/ZenLiuCN/engineX/modules/golang/image/gif"
	_ "github.com/ZenLiuCN/engineX/modules/golang/image/jpeg"
	_ "github.com/ZenLiuCN/engineX/modules/golang/image/png"
)
