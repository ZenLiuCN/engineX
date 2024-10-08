// Code generated by define_gene; DO NOT EDIT.
package larkext

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/ext"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_ext.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceExtDefine   []byte
	GithubComLarksuiteOapiSdkGo3ServiceExtDeclared = map[string]any{
		"newAuthenAccessTokenReqBodyBuilder":        larkext.NewAuthenAccessTokenReqBodyBuilder,
		"newAuthenAccessTokenReqBuilder":            larkext.NewAuthenAccessTokenReqBuilder,
		"newCreateFileReqBodyBuilder":               larkext.NewCreateFileReqBodyBuilder,
		"newRefreshAuthenAccessTokenReqBuilder":     larkext.NewRefreshAuthenAccessTokenReqBuilder,
		"newService":                                larkext.NewService,
		"FileTypeDoc":                               larkext.FileTypeDoc,
		"FileTypeSheet":                             larkext.FileTypeSheet,
		"GrantTypeAuthorizationCode":                larkext.GrantTypeAuthorizationCode,
		"newRefreshAuthenAccessTokenReqBodyBuilder": larkext.NewRefreshAuthenAccessTokenReqBodyBuilder,
		"FileTypeBitable":                           larkext.FileTypeBitable,
		"GrantTypeRefreshCode":                      larkext.GrantTypeRefreshCode,
		"newCreateFileReqBuilder":                   larkext.NewCreateFileReqBuilder,

		"emptyAuthenUserInfoResp":                  engine.Empty[larkext.AuthenUserInfoResp],
		"emptyRefAuthenUserInfoResp":               engine.EmptyRefer[larkext.AuthenUserInfoResp],
		"refOfAuthenUserInfoResp":                  engine.ReferOf[larkext.AuthenUserInfoResp],
		"unRefAuthenUserInfoResp":                  engine.UnRefer[larkext.AuthenUserInfoResp],
		"emptyCreateFileResp":                      engine.Empty[larkext.CreateFileResp],
		"emptyRefCreateFileResp":                   engine.EmptyRefer[larkext.CreateFileResp],
		"refOfCreateFileResp":                      engine.ReferOf[larkext.CreateFileResp],
		"unRefCreateFileResp":                      engine.UnRefer[larkext.CreateFileResp],
		"emptyRefreshAuthenAccessTokenRespBody":    engine.Empty[larkext.RefreshAuthenAccessTokenRespBody],
		"emptyRefRefreshAuthenAccessTokenRespBody": engine.EmptyRefer[larkext.RefreshAuthenAccessTokenRespBody],
		"refOfRefreshAuthenAccessTokenRespBody":    engine.ReferOf[larkext.RefreshAuthenAccessTokenRespBody],
		"unRefRefreshAuthenAccessTokenRespBody":    engine.UnRefer[larkext.RefreshAuthenAccessTokenRespBody],
		"emptyAuthenAccessTokenReq":                engine.Empty[larkext.AuthenAccessTokenReq],
		"emptyRefAuthenAccessTokenReq":             engine.EmptyRefer[larkext.AuthenAccessTokenReq],
		"refOfAuthenAccessTokenReq":                engine.ReferOf[larkext.AuthenAccessTokenReq],
		"unRefAuthenAccessTokenReq":                engine.UnRefer[larkext.AuthenAccessTokenReq],
		"emptyAuthenAccessTokenResp":               engine.Empty[larkext.AuthenAccessTokenResp],
		"emptyRefAuthenAccessTokenResp":            engine.EmptyRefer[larkext.AuthenAccessTokenResp],
		"refOfAuthenAccessTokenResp":               engine.ReferOf[larkext.AuthenAccessTokenResp],
		"unRefAuthenAccessTokenResp":               engine.UnRefer[larkext.AuthenAccessTokenResp],
		"emptyCreateFileReq":                       engine.Empty[larkext.CreateFileReq],
		"emptyRefCreateFileReq":                    engine.EmptyRefer[larkext.CreateFileReq],
		"refOfCreateFileReq":                       engine.ReferOf[larkext.CreateFileReq],
		"unRefCreateFileReq":                       engine.UnRefer[larkext.CreateFileReq],
		"emptyRefreshAuthenAccessTokenResp":        engine.Empty[larkext.RefreshAuthenAccessTokenResp],
		"emptyRefRefreshAuthenAccessTokenResp":     engine.EmptyRefer[larkext.RefreshAuthenAccessTokenResp],
		"refOfRefreshAuthenAccessTokenResp":        engine.ReferOf[larkext.RefreshAuthenAccessTokenResp],
		"unRefRefreshAuthenAccessTokenResp":        engine.UnRefer[larkext.RefreshAuthenAccessTokenResp],
		"emptyAuthenAccessTokenReqBody":            engine.Empty[larkext.AuthenAccessTokenReqBody],
		"emptyRefAuthenAccessTokenReqBody":         engine.EmptyRefer[larkext.AuthenAccessTokenReqBody],
		"refOfAuthenAccessTokenReqBody":            engine.ReferOf[larkext.AuthenAccessTokenReqBody],
		"unRefAuthenAccessTokenReqBody":            engine.UnRefer[larkext.AuthenAccessTokenReqBody],
		"emptyAuthenAccessTokenRespBody":           engine.Empty[larkext.AuthenAccessTokenRespBody],
		"emptyRefAuthenAccessTokenRespBody":        engine.EmptyRefer[larkext.AuthenAccessTokenRespBody],
		"refOfAuthenAccessTokenRespBody":           engine.ReferOf[larkext.AuthenAccessTokenRespBody],
		"unRefAuthenAccessTokenRespBody":           engine.UnRefer[larkext.AuthenAccessTokenRespBody],
		"emptyCreateFileReqBody":                   engine.Empty[larkext.CreateFileReqBody],
		"emptyRefCreateFileReqBody":                engine.EmptyRefer[larkext.CreateFileReqBody],
		"refOfCreateFileReqBody":                   engine.ReferOf[larkext.CreateFileReqBody],
		"unRefCreateFileReqBody":                   engine.UnRefer[larkext.CreateFileReqBody],
		"emptyRefreshAuthenAccessTokenReq":         engine.Empty[larkext.RefreshAuthenAccessTokenReq],
		"emptyRefRefreshAuthenAccessTokenReq":      engine.EmptyRefer[larkext.RefreshAuthenAccessTokenReq],
		"refOfRefreshAuthenAccessTokenReq":         engine.ReferOf[larkext.RefreshAuthenAccessTokenReq],
		"unRefRefreshAuthenAccessTokenReq":         engine.UnRefer[larkext.RefreshAuthenAccessTokenReq],
		"emptyRefreshAuthenAccessTokenReqBody":     engine.Empty[larkext.RefreshAuthenAccessTokenReqBody],
		"emptyRefRefreshAuthenAccessTokenReqBody":  engine.EmptyRefer[larkext.RefreshAuthenAccessTokenReqBody],
		"refOfRefreshAuthenAccessTokenReqBody":     engine.ReferOf[larkext.RefreshAuthenAccessTokenReqBody],
		"unRefRefreshAuthenAccessTokenReqBody":     engine.UnRefer[larkext.RefreshAuthenAccessTokenReqBody],
		"emptyAuthenUserInfoRespBody":              engine.Empty[larkext.AuthenUserInfoRespBody],
		"emptyRefAuthenUserInfoRespBody":           engine.EmptyRefer[larkext.AuthenUserInfoRespBody],
		"refOfAuthenUserInfoRespBody":              engine.ReferOf[larkext.AuthenUserInfoRespBody],
		"unRefAuthenUserInfoRespBody":              engine.UnRefer[larkext.AuthenUserInfoRespBody],
		"emptyCreateFileRespData":                  engine.Empty[larkext.CreateFileRespData],
		"emptyRefCreateFileRespData":               engine.EmptyRefer[larkext.CreateFileRespData],
		"refOfCreateFileRespData":                  engine.ReferOf[larkext.CreateFileRespData],
		"unRefCreateFileRespData":                  engine.UnRefer[larkext.CreateFileRespData],
		"emptyExtService":                          engine.Empty[larkext.ExtService],
		"emptyRefExtService":                       engine.EmptyRefer[larkext.ExtService],
		"refOfExtService":                          engine.ReferOf[larkext.ExtService],
		"unRefExtService":                          engine.UnRefer[larkext.ExtService]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceExtModule{})
}

type GithubComLarksuiteOapiSdkGo3ServiceExtModule struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceExtModule) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/ext"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceExtModule) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceExtDefine
}
func (S GithubComLarksuiteOapiSdkGo3ServiceExtModule) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceExtDeclared
}
