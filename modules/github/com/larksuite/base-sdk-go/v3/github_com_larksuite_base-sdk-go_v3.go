// Code generated by define_gene; DO NOT EDIT.
package lark

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/net/http"
	_ "github.com/ZenLiuCN/engine/modules/golang/time"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/base-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/base-sdk-go/v3/service/base/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/base-sdk-go/v3/service/drive/v1"
	"github.com/larksuite/base-sdk-go/v3"
)

var (
	//go:embed github_com_larksuite_base-sdk-go_v3.d.ts
	GithubComLarksuiteBaseSdkGo3Define   []byte
	GithubComLarksuiteBaseSdkGo3Declared = map[string]any{
		"withHeaders":       lark.WithHeaders,
		"withLogLevel":      lark.WithLogLevel,
		"withLogReqAtDebug": lark.WithLogReqAtDebug,
		"withReqTimeout":    lark.WithReqTimeout,
		"withSerialization": lark.WithSerialization,
		"FeishuBaseUrl":     lark.FeishuBaseUrl,
		"LarkBaseUrl":       lark.LarkBaseUrl,
		"newClient":         lark.NewClient,
		"withHttpClient":    lark.WithHttpClient,
		"withLogger":        lark.WithLogger,
		"withOpenBaseUrl":   lark.WithOpenBaseUrl,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteBaseSdkGo3Module{})
}

type GithubComLarksuiteBaseSdkGo3Module struct{}

func (S GithubComLarksuiteBaseSdkGo3Module) Identity() string {
	return "github.com/larksuite/base-sdk-go/v3"
}
func (S GithubComLarksuiteBaseSdkGo3Module) TypeDefine() []byte {
	return GithubComLarksuiteBaseSdkGo3Define
}
func (S GithubComLarksuiteBaseSdkGo3Module) Exports() map[string]any {
	return GithubComLarksuiteBaseSdkGo3Declared
}
