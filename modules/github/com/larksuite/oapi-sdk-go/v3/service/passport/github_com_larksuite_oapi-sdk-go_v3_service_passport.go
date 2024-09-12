// Code generated by define_gene; DO NOT EDIT.
package passport

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/passport/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/passport"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_passport.d.ts
	GithubComLarksuiteOapiSdkGo3ServicePassportDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServicePassportDeclared = map[string]any{
		"newService": passport.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServicePassportModule{})
}

type GithubComLarksuiteOapiSdkGo3ServicePassportModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServicePassportModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/passport"
}
func (S GithubComLarksuiteOapiSdkGo3ServicePassportModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServicePassportDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServicePassportModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServicePassportDeclared
}
