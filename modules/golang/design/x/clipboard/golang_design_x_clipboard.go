package clipboard

import (
	_ "embed"
	"fmt"
	"github.com/ZenLiuCN/engine"
	"github.com/dop251/goja"

	_ "github.com/ZenLiuCN/engine/modules/golang/archive/zip"
	_ "github.com/ZenLiuCN/engine/modules/golang/bytes"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/xml"
	_ "github.com/ZenLiuCN/engine/modules/golang/image/color"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engine/modules/golang/sync"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	"golang.design/x/clipboard"
)

var (
	//go:embed golang_design_x_clipboard.d.ts
	GolangDesignXClipboardDefine []byte
)

func init() {
	engine.RegisterModule(&GolangDesignXClipboardModule{})
}

type GolangDesignXClipboardModule struct {
	engine.BaseInitializeModule
}

func (S *GolangDesignXClipboardModule) Identity() string {
	return "golang.design/x/clipboard"
}
func (S *GolangDesignXClipboardModule) TypeDefine() []byte {
	return GolangDesignXClipboardDefine
}

func (S *GolangDesignXClipboardModule) ExportsWithEngine(eng *engine.Engine) map[string]any {
	pr := func(i clipboard.Format, data []byte) *goja.Promise {
		p, r, j := eng.NewPromise()
		ch := clipboard.Write(i, data)
		if ch == nil {
			j(fmt.Errorf("write to clipboard fail"))
		}
		go func() {
			for _ = range ch {
				r(nil)
				break
			}
		}()
		return p
	}
	return map[string]any{
		"init": clipboard.Init,
		"read": func(i int) []byte {
			return clipboard.Read(clipboard.Format(i))
		},
		"readImage": func(i int) []byte {
			return clipboard.Read(clipboard.FmtImage)
		},
		"readText": func(i int) string {
			return string(clipboard.Read(clipboard.FmtText))
		},
		"write": func(i int, data []byte) *goja.Promise {
			return pr(clipboard.Format(i), data)
		},
		"writeText": func(d string) *goja.Promise {
			return pr(clipboard.FmtText, []byte(d))
		},
		"writeImage": func(d []byte) *goja.Promise {
			return pr(clipboard.FmtImage, d)
		},
	}
}
