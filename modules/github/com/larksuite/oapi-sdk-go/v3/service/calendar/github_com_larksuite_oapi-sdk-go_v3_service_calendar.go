// Code generated by define_gene; DO NOT EDIT.
package calendar

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_calendar.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceCalendarDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceCalendarDeclared = map[string]any{
		"newService": calendar.NewService,
	}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceCalendarModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceCalendarModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceCalendarModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/calendar"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceCalendarModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceCalendarDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceCalendarModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceCalendarDeclared
}
