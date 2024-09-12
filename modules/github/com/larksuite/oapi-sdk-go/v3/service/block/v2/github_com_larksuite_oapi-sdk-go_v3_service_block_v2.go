// Code generated by define_gene; DO NOT EDIT.
package larkblock

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/block/v2"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_block_v2.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceBlock2Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceBlock2Declared = map[string]any{
		"newCreateEntityReqBuilder":         larkblock.NewCreateEntityReqBuilder,
		"newDepartmentIdBuilder":            larkblock.NewDepartmentIdBuilder,
		"newMessageBuilder":                 larkblock.NewMessageBuilder,
		"newEntityBuilder":                  larkblock.NewEntityBuilder,
		"newUpdateEntityReqBuilder":         larkblock.NewUpdateEntityReqBuilder,
		"New":                               larkblock.New,
		"newCreateEntityPathReqBodyBuilder": larkblock.NewCreateEntityPathReqBodyBuilder,
		"newCreateEntityReqBodyBuilder":     larkblock.NewCreateEntityReqBodyBuilder,
		"newCreateMessageReqBuilder":        larkblock.NewCreateMessageReqBuilder,

		"emptyUpdateEntityReq":         engine.Empty[larkblock.UpdateEntityReq],
		"emptyRefUpdateEntityReq":      engine.EmptyRefer[larkblock.UpdateEntityReq],
		"refOfUpdateEntityReq":         engine.ReferOf[larkblock.UpdateEntityReq],
		"unRefUpdateEntityReq":         engine.UnRefer[larkblock.UpdateEntityReq],
		"emptyUpdateEntityResp":        engine.Empty[larkblock.UpdateEntityResp],
		"emptyRefUpdateEntityResp":     engine.EmptyRefer[larkblock.UpdateEntityResp],
		"refOfUpdateEntityResp":        engine.ReferOf[larkblock.UpdateEntityResp],
		"unRefUpdateEntityResp":        engine.UnRefer[larkblock.UpdateEntityResp],
		"emptyEntity":                  engine.Empty[larkblock.Entity],
		"emptyRefEntity":               engine.EmptyRefer[larkblock.Entity],
		"refOfEntity":                  engine.ReferOf[larkblock.Entity],
		"unRefEntity":                  engine.UnRefer[larkblock.Entity],
		"emptyCreateEntityResp":        engine.Empty[larkblock.CreateEntityResp],
		"emptyRefCreateEntityResp":     engine.EmptyRefer[larkblock.CreateEntityResp],
		"refOfCreateEntityResp":        engine.ReferOf[larkblock.CreateEntityResp],
		"unRefCreateEntityResp":        engine.UnRefer[larkblock.CreateEntityResp],
		"emptyCreateMessageReq":        engine.Empty[larkblock.CreateMessageReq],
		"emptyRefCreateMessageReq":     engine.EmptyRefer[larkblock.CreateMessageReq],
		"refOfCreateMessageReq":        engine.ReferOf[larkblock.CreateMessageReq],
		"unRefCreateMessageReq":        engine.UnRefer[larkblock.CreateMessageReq],
		"emptyCreateMessageResp":       engine.Empty[larkblock.CreateMessageResp],
		"emptyRefCreateMessageResp":    engine.EmptyRefer[larkblock.CreateMessageResp],
		"refOfCreateMessageResp":       engine.ReferOf[larkblock.CreateMessageResp],
		"unRefCreateMessageResp":       engine.UnRefer[larkblock.CreateMessageResp],
		"emptyDepartmentId":            engine.Empty[larkblock.DepartmentId],
		"emptyRefDepartmentId":         engine.EmptyRefer[larkblock.DepartmentId],
		"refOfDepartmentId":            engine.ReferOf[larkblock.DepartmentId],
		"unRefDepartmentId":            engine.UnRefer[larkblock.DepartmentId],
		"emptyV2":                      engine.Empty[larkblock.V2],
		"emptyRefV2":                   engine.EmptyRefer[larkblock.V2],
		"refOfV2":                      engine.ReferOf[larkblock.V2],
		"unRefV2":                      engine.UnRefer[larkblock.V2],
		"emptyCreateEntityReq":         engine.Empty[larkblock.CreateEntityReq],
		"emptyRefCreateEntityReq":      engine.EmptyRefer[larkblock.CreateEntityReq],
		"refOfCreateEntityReq":         engine.ReferOf[larkblock.CreateEntityReq],
		"unRefCreateEntityReq":         engine.UnRefer[larkblock.CreateEntityReq],
		"emptyCreateEntityReqBody":     engine.Empty[larkblock.CreateEntityReqBody],
		"emptyRefCreateEntityReqBody":  engine.EmptyRefer[larkblock.CreateEntityReqBody],
		"refOfCreateEntityReqBody":     engine.ReferOf[larkblock.CreateEntityReqBody],
		"unRefCreateEntityReqBody":     engine.UnRefer[larkblock.CreateEntityReqBody],
		"emptyCreateEntityRespData":    engine.Empty[larkblock.CreateEntityRespData],
		"emptyRefCreateEntityRespData": engine.EmptyRefer[larkblock.CreateEntityRespData],
		"refOfCreateEntityRespData":    engine.ReferOf[larkblock.CreateEntityRespData],
		"unRefCreateEntityRespData":    engine.UnRefer[larkblock.CreateEntityRespData],
		"emptyMessage":                 engine.Empty[larkblock.Message],
		"emptyRefMessage":              engine.EmptyRefer[larkblock.Message],
		"refOfMessage":                 engine.ReferOf[larkblock.Message],
		"unRefMessage":                 engine.UnRefer[larkblock.Message]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceBlock2Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceBlock2Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceBlock2Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/block/v2"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBlock2Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceBlock2Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBlock2Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceBlock2Declared
}
