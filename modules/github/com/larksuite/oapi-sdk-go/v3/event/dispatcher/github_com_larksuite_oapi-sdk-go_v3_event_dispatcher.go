// Code generated by define_gene; DO NOT EDIT.
package dispatcher

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/encoding/json"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/event"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/acs/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/application/v6"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/approval/v4"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/corehr/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/corehr/v2"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/helpdesk/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/hire/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/im/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/meeting_room/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/task/v1"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/service/vc/v1"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_event_dispatcher.d.ts
	GithubComLarksuiteOapiSdkGo3EventDispatcherDefine   []byte
	GithubComLarksuiteOapiSdkGo3EventDispatcherDeclared = map[string]any{
		"newEventDispatcher": dispatcher.NewEventDispatcher,

		"emptyAppTicketEvent":                 engine.Empty[dispatcher.AppTicketEvent],
		"emptyRefAppTicketEvent":              engine.EmptyRefer[dispatcher.AppTicketEvent],
		"refOfAppTicketEvent":                 engine.ReferOf[dispatcher.AppTicketEvent],
		"unRefAppTicketEvent":                 engine.UnRefer[dispatcher.AppTicketEvent],
		"emptyCustomAppTicketEventHandler":    engine.Empty[dispatcher.CustomAppTicketEventHandler],
		"emptyRefCustomAppTicketEventHandler": engine.EmptyRefer[dispatcher.CustomAppTicketEventHandler],
		"refOfCustomAppTicketEventHandler":    engine.ReferOf[dispatcher.CustomAppTicketEventHandler],
		"unRefCustomAppTicketEventHandler":    engine.UnRefer[dispatcher.CustomAppTicketEventHandler]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3EventDispatcherModule{})
}

type GithubComLarksuiteOapiSdkGo3EventDispatcherModule struct{}

func (S GithubComLarksuiteOapiSdkGo3EventDispatcherModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/event/dispatcher"
}
func (S GithubComLarksuiteOapiSdkGo3EventDispatcherModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3EventDispatcherDefine
}
func (S GithubComLarksuiteOapiSdkGo3EventDispatcherModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3EventDispatcherDeclared
}
