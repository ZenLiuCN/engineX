// Code generated by define_gene; DO NOT EDIT.
package contact

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_contact.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceContactDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceContactDeclared = map[string]any{
		"newService": contact.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceContactModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceContactModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceContactModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/contact"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceContactModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceContactDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceContactModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceContactDeclared
}
