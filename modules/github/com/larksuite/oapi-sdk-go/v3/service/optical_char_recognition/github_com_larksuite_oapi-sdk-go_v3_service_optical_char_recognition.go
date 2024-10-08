// Code generated by define_gene; DO NOT EDIT.
package optical_char_recognition

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_optical_char_recognition.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionDeclared = map[string]any{
		"newService": optical_char_recognition.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceOptical_char_recognitionDeclared
}
