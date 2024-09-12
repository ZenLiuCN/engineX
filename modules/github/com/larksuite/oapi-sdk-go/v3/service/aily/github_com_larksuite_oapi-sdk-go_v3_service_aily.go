// Code generated by define_gene; DO NOT EDIT.
package aily

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/aily/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/aily"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_aily.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceAilyDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceAilyDeclared = map[string]any{
		"newService": aily.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceAilyModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceAilyModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceAilyModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/aily"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAilyModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceAilyDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAilyModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceAilyDeclared
}
