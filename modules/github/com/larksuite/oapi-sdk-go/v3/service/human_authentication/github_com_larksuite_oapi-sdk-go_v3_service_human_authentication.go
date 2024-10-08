// Code generated by define_gene; DO NOT EDIT.
package human_authentication

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/human_authentication/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/human_authentication"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_human_authentication.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationDeclared = map[string]any{
		"newService": human_authentication.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/human_authentication"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceHuman_authenticationDeclared
}
