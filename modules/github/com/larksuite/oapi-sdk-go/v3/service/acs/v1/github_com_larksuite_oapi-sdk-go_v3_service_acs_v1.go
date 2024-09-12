// Code generated by define_gene; DO NOT EDIT.
package larkacs

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_acs_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceAcs1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceAcs1Declared = map[string]any{
		"newGetUserFaceReqBuilder":                    larkacs.NewGetUserFaceReqBuilder,
		"UserIdTypeDeleteVisitorUnionId":              larkacs.UserIdTypeDeleteVisitorUnionId,
		"UserIdTypeGetUserFaceOpenId":                 larkacs.UserIdTypeGetUserFaceOpenId,
		"newCreateVisitorReqBodyBuilder":              larkacs.NewCreateVisitorReqBodyBuilder,
		"newDepartmentIdBuilder":                      larkacs.NewDepartmentIdBuilder,
		"newGetAccessRecordAccessPhotoReqBuilder":     larkacs.NewGetAccessRecordAccessPhotoReqBuilder,
		"newDeviceBindRuleExternalPathReqBodyBuilder": larkacs.NewDeviceBindRuleExternalPathReqBodyBuilder,
		"UserIdTypeGetUserUserId":                     larkacs.UserIdTypeGetUserUserId,
		"UserIdTypeListUserUserId":                    larkacs.UserIdTypeListUserUserId,
		"newP2AccessRecordCreatedV1Handler":           larkacs.NewP2AccessRecordCreatedV1Handler,
		"newP2UserUpdatedV1Handler":                   larkacs.NewP2UserUpdatedV1Handler,
		"UserIdTypeGetRuleExternalOpenId":             larkacs.UserIdTypeGetRuleExternalOpenId,
		"newCreateRuleExternalReqBodyBuilder":         larkacs.NewCreateRuleExternalReqBodyBuilder,
		"newDeviceExternalBuilder":                    larkacs.NewDeviceExternalBuilder,
		"newFileBuilder":                              larkacs.NewFileBuilder,
		"newUpdateUserFaceReqBuilder":                 larkacs.NewUpdateUserFaceReqBuilder,
		"UserIdTypeDeleteVisitorUserId":               larkacs.UserIdTypeDeleteVisitorUserId,
		"UserIdTypeOpenId":                            larkacs.UserIdTypeOpenId,
		"newCreateRuleExternalPathReqBodyBuilder":     larkacs.NewCreateRuleExternalPathReqBodyBuilder,
		"newPropertyBuilder":                          larkacs.NewPropertyBuilder,
		"newRuleBuilder":                              larkacs.NewRuleBuilder,
		"New":                                         larkacs.New,
		"newDeviceBindRuleExternalReqBuilder":         larkacs.NewDeviceBindRuleExternalReqBuilder,
		"newListAccessRecordReqBuilder":               larkacs.NewListAccessRecordReqBuilder,
		"UserIdTypeCreateVisitorUnionId":              larkacs.UserIdTypeCreateVisitorUnionId,
		"UserIdTypeGetUserFaceUserId":                 larkacs.UserIdTypeGetUserFaceUserId,
		"UserIdTypeGetUserOpenId":                     larkacs.UserIdTypeGetUserOpenId,
		"newAccessRecordBuilder":                      larkacs.NewAccessRecordBuilder,
		"newDeleteRuleExternalReqBuilder":             larkacs.NewDeleteRuleExternalReqBuilder,
		"newGetRuleExternalReqBuilder":                larkacs.NewGetRuleExternalReqBuilder,
		"newGetUserReqBuilder":                        larkacs.NewGetUserReqBuilder,
		"UserIdTypeGetRuleExternalUserId":             larkacs.UserIdTypeGetRuleExternalUserId,
		"UserIdTypePatchUserUnionId":                  larkacs.UserIdTypePatchUserUnionId,
		"UserIdTypeUpdateUserFaceOpenId":              larkacs.UserIdTypeUpdateUserFaceOpenId,
		"newDeviceBindRuleExternalReqBodyBuilder":     larkacs.NewDeviceBindRuleExternalReqBodyBuilder,
		"newOpeningTimeExternalBuilder":               larkacs.NewOpeningTimeExternalBuilder,
		"newPatchUserReqBuilder":                      larkacs.NewPatchUserReqBuilder,
		"UserIdTypeGetRuleExternalUnionId":            larkacs.UserIdTypeGetRuleExternalUnionId,
		"UserIdTypeUpdateUserFaceUnionId":             larkacs.UserIdTypeUpdateUserFaceUnionId,
		"newCreateRuleExternalReqBuilder":             larkacs.NewCreateRuleExternalReqBuilder,
		"newFeatureBuilder":                           larkacs.NewFeatureBuilder,
		"newUserBuilder":                              larkacs.NewUserBuilder,
		"UserIdTypeDeleteVisitorOpenId":               larkacs.UserIdTypeDeleteVisitorOpenId,
		"UserIdTypeListUserUnionId":                   larkacs.UserIdTypeListUserUnionId,
		"UserIdTypeUnionId":                           larkacs.UserIdTypeUnionId,
		"newCreateVisitorReqBuilder":                  larkacs.NewCreateVisitorReqBuilder,
		"UserIdTypeCreateRuleExternalUnionId":         larkacs.UserIdTypeCreateRuleExternalUnionId,
		"UserIdTypeCreateVisitorUserId":               larkacs.UserIdTypeCreateVisitorUserId,
		"newUserIdBuilder":                            larkacs.NewUserIdBuilder,
		"UserIdTypeCreateRuleExternalOpenId":          larkacs.UserIdTypeCreateRuleExternalOpenId,
		"UserIdTypeCreateVisitorOpenId":               larkacs.UserIdTypeCreateVisitorOpenId,
		"newCreateVisitorPathReqBodyBuilder":          larkacs.NewCreateVisitorPathReqBodyBuilder,
		"newDeviceBuilder":                            larkacs.NewDeviceBuilder,
		"newOpeningTimePeriodExternalBuilder":         larkacs.NewOpeningTimePeriodExternalBuilder,
		"UserIdTypeGetUserUnionId":                    larkacs.UserIdTypeGetUserUnionId,
		"UserIdTypeUpdateUserFaceUserId":              larkacs.UserIdTypeUpdateUserFaceUserId,
		"UserIdTypeUserId":                            larkacs.UserIdTypeUserId,
		"newOpeningTimeValidDayExternalBuilder":       larkacs.NewOpeningTimeValidDayExternalBuilder,
		"UserIdTypeCreateRuleExternalUserId":          larkacs.UserIdTypeCreateRuleExternalUserId,
		"UserIdTypeGetUserFaceUnionId":                larkacs.UserIdTypeGetUserFaceUnionId,
		"UserIdTypePatchUserOpenId":                   larkacs.UserIdTypePatchUserOpenId,
		"newDeleteVisitorReqBuilder":                  larkacs.NewDeleteVisitorReqBuilder,
		"newListUserReqBuilder":                       larkacs.NewListUserReqBuilder,
		"UserIdTypeListUserOpenId":                    larkacs.UserIdTypeListUserOpenId,
		"newUserExternalBuilder":                      larkacs.NewUserExternalBuilder,
		"UserIdTypePatchUserUserId":                   larkacs.UserIdTypePatchUserUserId,

		"emptyGetAccessRecordAccessPhotoResp":    engine.Empty[larkacs.GetAccessRecordAccessPhotoResp],
		"emptyRefGetAccessRecordAccessPhotoResp": engine.EmptyRefer[larkacs.GetAccessRecordAccessPhotoResp],
		"refOfGetAccessRecordAccessPhotoResp":    engine.ReferOf[larkacs.GetAccessRecordAccessPhotoResp],
		"unRefGetAccessRecordAccessPhotoResp":    engine.UnRefer[larkacs.GetAccessRecordAccessPhotoResp],
		"emptyP2UserUpdatedV1Data":               engine.Empty[larkacs.P2UserUpdatedV1Data],
		"emptyRefP2UserUpdatedV1Data":            engine.EmptyRefer[larkacs.P2UserUpdatedV1Data],
		"refOfP2UserUpdatedV1Data":               engine.ReferOf[larkacs.P2UserUpdatedV1Data],
		"unRefP2UserUpdatedV1Data":               engine.UnRefer[larkacs.P2UserUpdatedV1Data],
		"emptyPatchUserReq":                      engine.Empty[larkacs.PatchUserReq],
		"emptyRefPatchUserReq":                   engine.EmptyRefer[larkacs.PatchUserReq],
		"refOfPatchUserReq":                      engine.ReferOf[larkacs.PatchUserReq],
		"unRefPatchUserReq":                      engine.UnRefer[larkacs.PatchUserReq],
		"emptyCreateVisitorReq":                  engine.Empty[larkacs.CreateVisitorReq],
		"emptyRefCreateVisitorReq":               engine.EmptyRefer[larkacs.CreateVisitorReq],
		"refOfCreateVisitorReq":                  engine.ReferOf[larkacs.CreateVisitorReq],
		"unRefCreateVisitorReq":                  engine.UnRefer[larkacs.CreateVisitorReq],
		"emptyDeviceBindRuleExternalReqBody":     engine.Empty[larkacs.DeviceBindRuleExternalReqBody],
		"emptyRefDeviceBindRuleExternalReqBody":  engine.EmptyRefer[larkacs.DeviceBindRuleExternalReqBody],
		"refOfDeviceBindRuleExternalReqBody":     engine.ReferOf[larkacs.DeviceBindRuleExternalReqBody],
		"unRefDeviceBindRuleExternalReqBody":     engine.UnRefer[larkacs.DeviceBindRuleExternalReqBody],
		"emptyGetUserResp":                       engine.Empty[larkacs.GetUserResp],
		"emptyRefGetUserResp":                    engine.EmptyRefer[larkacs.GetUserResp],
		"refOfGetUserResp":                       engine.ReferOf[larkacs.GetUserResp],
		"unRefGetUserResp":                       engine.UnRefer[larkacs.GetUserResp],
		"emptyListAccessRecordResp":              engine.Empty[larkacs.ListAccessRecordResp],
		"emptyRefListAccessRecordResp":           engine.EmptyRefer[larkacs.ListAccessRecordResp],
		"refOfListAccessRecordResp":              engine.ReferOf[larkacs.ListAccessRecordResp],
		"unRefListAccessRecordResp":              engine.UnRefer[larkacs.ListAccessRecordResp],
		"emptyUpdateUserFaceResp":                engine.Empty[larkacs.UpdateUserFaceResp],
		"emptyRefUpdateUserFaceResp":             engine.EmptyRefer[larkacs.UpdateUserFaceResp],
		"refOfUpdateUserFaceResp":                engine.ReferOf[larkacs.UpdateUserFaceResp],
		"unRefUpdateUserFaceResp":                engine.UnRefer[larkacs.UpdateUserFaceResp],
		"emptyCreateVisitorRespData":             engine.Empty[larkacs.CreateVisitorRespData],
		"emptyRefCreateVisitorRespData":          engine.EmptyRefer[larkacs.CreateVisitorRespData],
		"refOfCreateVisitorRespData":             engine.ReferOf[larkacs.CreateVisitorRespData],
		"unRefCreateVisitorRespData":             engine.UnRefer[larkacs.CreateVisitorRespData],
		"emptyDeleteVisitorReq":                  engine.Empty[larkacs.DeleteVisitorReq],
		"emptyRefDeleteVisitorReq":               engine.EmptyRefer[larkacs.DeleteVisitorReq],
		"refOfDeleteVisitorReq":                  engine.ReferOf[larkacs.DeleteVisitorReq],
		"unRefDeleteVisitorReq":                  engine.UnRefer[larkacs.DeleteVisitorReq],
		"emptyDevice":                            engine.Empty[larkacs.Device],
		"emptyRefDevice":                         engine.EmptyRefer[larkacs.Device],
		"refOfDevice":                            engine.ReferOf[larkacs.Device],
		"unRefDevice":                            engine.UnRefer[larkacs.Device],
		"emptyDeviceExternal":                    engine.Empty[larkacs.DeviceExternal],
		"emptyRefDeviceExternal":                 engine.EmptyRefer[larkacs.DeviceExternal],
		"refOfDeviceExternal":                    engine.ReferOf[larkacs.DeviceExternal],
		"unRefDeviceExternal":                    engine.UnRefer[larkacs.DeviceExternal],
		"emptyListUserRespData":                  engine.Empty[larkacs.ListUserRespData],
		"emptyRefListUserRespData":               engine.EmptyRefer[larkacs.ListUserRespData],
		"refOfListUserRespData":                  engine.ReferOf[larkacs.ListUserRespData],
		"unRefListUserRespData":                  engine.UnRefer[larkacs.ListUserRespData],
		"emptyP2UserUpdatedV1":                   engine.Empty[larkacs.P2UserUpdatedV1],
		"emptyRefP2UserUpdatedV1":                engine.EmptyRefer[larkacs.P2UserUpdatedV1],
		"refOfP2UserUpdatedV1":                   engine.ReferOf[larkacs.P2UserUpdatedV1],
		"unRefP2UserUpdatedV1":                   engine.UnRefer[larkacs.P2UserUpdatedV1],
		"emptyPatchUserResp":                     engine.Empty[larkacs.PatchUserResp],
		"emptyRefPatchUserResp":                  engine.EmptyRefer[larkacs.PatchUserResp],
		"refOfPatchUserResp":                     engine.ReferOf[larkacs.PatchUserResp],
		"unRefPatchUserResp":                     engine.UnRefer[larkacs.PatchUserResp],
		"emptyUser":                              engine.Empty[larkacs.User],
		"emptyRefUser":                           engine.EmptyRefer[larkacs.User],
		"refOfUser":                              engine.ReferOf[larkacs.User],
		"unRefUser":                              engine.UnRefer[larkacs.User],
		"emptyCreateRuleExternalReqBody":         engine.Empty[larkacs.CreateRuleExternalReqBody],
		"emptyRefCreateRuleExternalReqBody":      engine.EmptyRefer[larkacs.CreateRuleExternalReqBody],
		"refOfCreateRuleExternalReqBody":         engine.ReferOf[larkacs.CreateRuleExternalReqBody],
		"unRefCreateRuleExternalReqBody":         engine.UnRefer[larkacs.CreateRuleExternalReqBody],
		"emptyCreateVisitorResp":                 engine.Empty[larkacs.CreateVisitorResp],
		"emptyRefCreateVisitorResp":              engine.EmptyRefer[larkacs.CreateVisitorResp],
		"refOfCreateVisitorResp":                 engine.ReferOf[larkacs.CreateVisitorResp],
		"unRefCreateVisitorResp":                 engine.UnRefer[larkacs.CreateVisitorResp],
		"emptyDeviceBindRuleExternalReq":         engine.Empty[larkacs.DeviceBindRuleExternalReq],
		"emptyRefDeviceBindRuleExternalReq":      engine.EmptyRefer[larkacs.DeviceBindRuleExternalReq],
		"refOfDeviceBindRuleExternalReq":         engine.ReferOf[larkacs.DeviceBindRuleExternalReq],
		"unRefDeviceBindRuleExternalReq":         engine.UnRefer[larkacs.DeviceBindRuleExternalReq],
		"emptyListAccessRecordRespData":          engine.Empty[larkacs.ListAccessRecordRespData],
		"emptyRefListAccessRecordRespData":       engine.EmptyRefer[larkacs.ListAccessRecordRespData],
		"refOfListAccessRecordRespData":          engine.ReferOf[larkacs.ListAccessRecordRespData],
		"unRefListAccessRecordRespData":          engine.UnRefer[larkacs.ListAccessRecordRespData],
		"emptyUpdateUserFaceReq":                 engine.Empty[larkacs.UpdateUserFaceReq],
		"emptyRefUpdateUserFaceReq":              engine.EmptyRefer[larkacs.UpdateUserFaceReq],
		"refOfUpdateUserFaceReq":                 engine.ReferOf[larkacs.UpdateUserFaceReq],
		"unRefUpdateUserFaceReq":                 engine.UnRefer[larkacs.UpdateUserFaceReq],
		"emptyUserId":                            engine.Empty[larkacs.UserId],
		"emptyRefUserId":                         engine.EmptyRefer[larkacs.UserId],
		"refOfUserId":                            engine.ReferOf[larkacs.UserId],
		"unRefUserId":                            engine.UnRefer[larkacs.UserId],
		"emptyDepartmentId":                      engine.Empty[larkacs.DepartmentId],
		"emptyRefDepartmentId":                   engine.EmptyRefer[larkacs.DepartmentId],
		"refOfDepartmentId":                      engine.ReferOf[larkacs.DepartmentId],
		"unRefDepartmentId":                      engine.UnRefer[larkacs.DepartmentId],
		"emptyFeature":                           engine.Empty[larkacs.Feature],
		"emptyRefFeature":                        engine.EmptyRefer[larkacs.Feature],
		"refOfFeature":                           engine.ReferOf[larkacs.Feature],
		"unRefFeature":                           engine.UnRefer[larkacs.Feature],
		"emptyListDeviceResp":                    engine.Empty[larkacs.ListDeviceResp],
		"emptyRefListDeviceResp":                 engine.EmptyRefer[larkacs.ListDeviceResp],
		"refOfListDeviceResp":                    engine.ReferOf[larkacs.ListDeviceResp],
		"unRefListDeviceResp":                    engine.UnRefer[larkacs.ListDeviceResp],
		"emptyListUserIterator":                  engine.Empty[larkacs.ListUserIterator],
		"emptyRefListUserIterator":               engine.EmptyRefer[larkacs.ListUserIterator],
		"refOfListUserIterator":                  engine.ReferOf[larkacs.ListUserIterator],
		"unRefListUserIterator":                  engine.UnRefer[larkacs.ListUserIterator],
		"emptyOpeningTimePeriodExternal":         engine.Empty[larkacs.OpeningTimePeriodExternal],
		"emptyRefOpeningTimePeriodExternal":      engine.EmptyRefer[larkacs.OpeningTimePeriodExternal],
		"refOfOpeningTimePeriodExternal":         engine.ReferOf[larkacs.OpeningTimePeriodExternal],
		"unRefOpeningTimePeriodExternal":         engine.UnRefer[larkacs.OpeningTimePeriodExternal],
		"emptyDeleteVisitorResp":                 engine.Empty[larkacs.DeleteVisitorResp],
		"emptyRefDeleteVisitorResp":              engine.EmptyRefer[larkacs.DeleteVisitorResp],
		"refOfDeleteVisitorResp":                 engine.ReferOf[larkacs.DeleteVisitorResp],
		"unRefDeleteVisitorResp":                 engine.UnRefer[larkacs.DeleteVisitorResp],
		"emptyFile":                              engine.Empty[larkacs.File],
		"emptyRefFile":                           engine.EmptyRefer[larkacs.File],
		"refOfFile":                              engine.ReferOf[larkacs.File],
		"unRefFile":                              engine.UnRefer[larkacs.File],
		"emptyGetRuleExternalReq":                engine.Empty[larkacs.GetRuleExternalReq],
		"emptyRefGetRuleExternalReq":             engine.EmptyRefer[larkacs.GetRuleExternalReq],
		"refOfGetRuleExternalReq":                engine.ReferOf[larkacs.GetRuleExternalReq],
		"unRefGetRuleExternalReq":                engine.UnRefer[larkacs.GetRuleExternalReq],
		"emptyP2AccessRecordCreatedV1":           engine.Empty[larkacs.P2AccessRecordCreatedV1],
		"emptyRefP2AccessRecordCreatedV1":        engine.EmptyRefer[larkacs.P2AccessRecordCreatedV1],
		"refOfP2AccessRecordCreatedV1":           engine.ReferOf[larkacs.P2AccessRecordCreatedV1],
		"unRefP2AccessRecordCreatedV1":           engine.UnRefer[larkacs.P2AccessRecordCreatedV1],
		"emptyP2AccessRecordCreatedV1Data":       engine.Empty[larkacs.P2AccessRecordCreatedV1Data],
		"emptyRefP2AccessRecordCreatedV1Data":    engine.EmptyRefer[larkacs.P2AccessRecordCreatedV1Data],
		"refOfP2AccessRecordCreatedV1Data":       engine.ReferOf[larkacs.P2AccessRecordCreatedV1Data],
		"unRefP2AccessRecordCreatedV1Data":       engine.UnRefer[larkacs.P2AccessRecordCreatedV1Data],
		"emptyDeviceBindRuleExternalResp":        engine.Empty[larkacs.DeviceBindRuleExternalResp],
		"emptyRefDeviceBindRuleExternalResp":     engine.EmptyRefer[larkacs.DeviceBindRuleExternalResp],
		"refOfDeviceBindRuleExternalResp":        engine.ReferOf[larkacs.DeviceBindRuleExternalResp],
		"unRefDeviceBindRuleExternalResp":        engine.UnRefer[larkacs.DeviceBindRuleExternalResp],
		"emptyListDeviceRespData":                engine.Empty[larkacs.ListDeviceRespData],
		"emptyRefListDeviceRespData":             engine.EmptyRefer[larkacs.ListDeviceRespData],
		"refOfListDeviceRespData":                engine.ReferOf[larkacs.ListDeviceRespData],
		"unRefListDeviceRespData":                engine.UnRefer[larkacs.ListDeviceRespData],
		"emptyCreateVisitorReqBody":              engine.Empty[larkacs.CreateVisitorReqBody],
		"emptyRefCreateVisitorReqBody":           engine.EmptyRefer[larkacs.CreateVisitorReqBody],
		"refOfCreateVisitorReqBody":              engine.ReferOf[larkacs.CreateVisitorReqBody],
		"unRefCreateVisitorReqBody":              engine.UnRefer[larkacs.CreateVisitorReqBody],
		"emptyGetRuleExternalRespData":           engine.Empty[larkacs.GetRuleExternalRespData],
		"emptyRefGetRuleExternalRespData":        engine.EmptyRefer[larkacs.GetRuleExternalRespData],
		"refOfGetRuleExternalRespData":           engine.ReferOf[larkacs.GetRuleExternalRespData],
		"unRefGetRuleExternalRespData":           engine.UnRefer[larkacs.GetRuleExternalRespData],
		"emptyGetUserFaceReq":                    engine.Empty[larkacs.GetUserFaceReq],
		"emptyRefGetUserFaceReq":                 engine.EmptyRefer[larkacs.GetUserFaceReq],
		"refOfGetUserFaceReq":                    engine.ReferOf[larkacs.GetUserFaceReq],
		"unRefGetUserFaceReq":                    engine.UnRefer[larkacs.GetUserFaceReq],
		"emptyListUserResp":                      engine.Empty[larkacs.ListUserResp],
		"emptyRefListUserResp":                   engine.EmptyRefer[larkacs.ListUserResp],
		"refOfListUserResp":                      engine.ReferOf[larkacs.ListUserResp],
		"unRefListUserResp":                      engine.UnRefer[larkacs.ListUserResp],
		"emptyProperty":                          engine.Empty[larkacs.Property],
		"emptyRefProperty":                       engine.EmptyRefer[larkacs.Property],
		"refOfProperty":                          engine.ReferOf[larkacs.Property],
		"unRefProperty":                          engine.UnRefer[larkacs.Property],
		"emptyCreateRuleExternalRespData":        engine.Empty[larkacs.CreateRuleExternalRespData],
		"emptyRefCreateRuleExternalRespData":     engine.EmptyRefer[larkacs.CreateRuleExternalRespData],
		"refOfCreateRuleExternalRespData":        engine.ReferOf[larkacs.CreateRuleExternalRespData],
		"unRefCreateRuleExternalRespData":        engine.UnRefer[larkacs.CreateRuleExternalRespData],
		"emptyGetAccessRecordAccessPhotoReq":     engine.Empty[larkacs.GetAccessRecordAccessPhotoReq],
		"emptyRefGetAccessRecordAccessPhotoReq":  engine.EmptyRefer[larkacs.GetAccessRecordAccessPhotoReq],
		"refOfGetAccessRecordAccessPhotoReq":     engine.ReferOf[larkacs.GetAccessRecordAccessPhotoReq],
		"unRefGetAccessRecordAccessPhotoReq":     engine.UnRefer[larkacs.GetAccessRecordAccessPhotoReq],
		"emptyGetUserRespData":                   engine.Empty[larkacs.GetUserRespData],
		"emptyRefGetUserRespData":                engine.EmptyRefer[larkacs.GetUserRespData],
		"refOfGetUserRespData":                   engine.ReferOf[larkacs.GetUserRespData],
		"unRefGetUserRespData":                   engine.UnRefer[larkacs.GetUserRespData],
		"emptyRule":                              engine.Empty[larkacs.Rule],
		"emptyRefRule":                           engine.EmptyRefer[larkacs.Rule],
		"refOfRule":                              engine.ReferOf[larkacs.Rule],
		"unRefRule":                              engine.UnRefer[larkacs.Rule],
		"emptyCreateRuleExternalReq":             engine.Empty[larkacs.CreateRuleExternalReq],
		"emptyRefCreateRuleExternalReq":          engine.EmptyRefer[larkacs.CreateRuleExternalReq],
		"refOfCreateRuleExternalReq":             engine.ReferOf[larkacs.CreateRuleExternalReq],
		"unRefCreateRuleExternalReq":             engine.UnRefer[larkacs.CreateRuleExternalReq],
		"emptyCreateRuleExternalResp":            engine.Empty[larkacs.CreateRuleExternalResp],
		"emptyRefCreateRuleExternalResp":         engine.EmptyRefer[larkacs.CreateRuleExternalResp],
		"refOfCreateRuleExternalResp":            engine.ReferOf[larkacs.CreateRuleExternalResp],
		"unRefCreateRuleExternalResp":            engine.UnRefer[larkacs.CreateRuleExternalResp],
		"emptyGetRuleExternalResp":               engine.Empty[larkacs.GetRuleExternalResp],
		"emptyRefGetRuleExternalResp":            engine.EmptyRefer[larkacs.GetRuleExternalResp],
		"refOfGetRuleExternalResp":               engine.ReferOf[larkacs.GetRuleExternalResp],
		"unRefGetRuleExternalResp":               engine.UnRefer[larkacs.GetRuleExternalResp],
		"emptyGetUserReq":                        engine.Empty[larkacs.GetUserReq],
		"emptyRefGetUserReq":                     engine.EmptyRefer[larkacs.GetUserReq],
		"refOfGetUserReq":                        engine.ReferOf[larkacs.GetUserReq],
		"unRefGetUserReq":                        engine.UnRefer[larkacs.GetUserReq],
		"emptyListUserReq":                       engine.Empty[larkacs.ListUserReq],
		"emptyRefListUserReq":                    engine.EmptyRefer[larkacs.ListUserReq],
		"refOfListUserReq":                       engine.ReferOf[larkacs.ListUserReq],
		"unRefListUserReq":                       engine.UnRefer[larkacs.ListUserReq],
		"emptyDeleteRuleExternalReq":             engine.Empty[larkacs.DeleteRuleExternalReq],
		"emptyRefDeleteRuleExternalReq":          engine.EmptyRefer[larkacs.DeleteRuleExternalReq],
		"refOfDeleteRuleExternalReq":             engine.ReferOf[larkacs.DeleteRuleExternalReq],
		"unRefDeleteRuleExternalReq":             engine.UnRefer[larkacs.DeleteRuleExternalReq],
		"emptyDeleteRuleExternalResp":            engine.Empty[larkacs.DeleteRuleExternalResp],
		"emptyRefDeleteRuleExternalResp":         engine.EmptyRefer[larkacs.DeleteRuleExternalResp],
		"refOfDeleteRuleExternalResp":            engine.ReferOf[larkacs.DeleteRuleExternalResp],
		"unRefDeleteRuleExternalResp":            engine.UnRefer[larkacs.DeleteRuleExternalResp],
		"emptyUserExternal":                      engine.Empty[larkacs.UserExternal],
		"emptyRefUserExternal":                   engine.EmptyRefer[larkacs.UserExternal],
		"refOfUserExternal":                      engine.ReferOf[larkacs.UserExternal],
		"unRefUserExternal":                      engine.UnRefer[larkacs.UserExternal],
		"emptyAccessRecord":                      engine.Empty[larkacs.AccessRecord],
		"emptyRefAccessRecord":                   engine.EmptyRefer[larkacs.AccessRecord],
		"refOfAccessRecord":                      engine.ReferOf[larkacs.AccessRecord],
		"unRefAccessRecord":                      engine.UnRefer[larkacs.AccessRecord],
		"emptyGetUserFaceResp":                   engine.Empty[larkacs.GetUserFaceResp],
		"emptyRefGetUserFaceResp":                engine.EmptyRefer[larkacs.GetUserFaceResp],
		"refOfGetUserFaceResp":                   engine.ReferOf[larkacs.GetUserFaceResp],
		"unRefGetUserFaceResp":                   engine.UnRefer[larkacs.GetUserFaceResp],
		"emptyListAccessRecordIterator":          engine.Empty[larkacs.ListAccessRecordIterator],
		"emptyRefListAccessRecordIterator":       engine.EmptyRefer[larkacs.ListAccessRecordIterator],
		"refOfListAccessRecordIterator":          engine.ReferOf[larkacs.ListAccessRecordIterator],
		"unRefListAccessRecordIterator":          engine.UnRefer[larkacs.ListAccessRecordIterator],
		"emptyListAccessRecordReq":               engine.Empty[larkacs.ListAccessRecordReq],
		"emptyRefListAccessRecordReq":            engine.EmptyRefer[larkacs.ListAccessRecordReq],
		"refOfListAccessRecordReq":               engine.ReferOf[larkacs.ListAccessRecordReq],
		"unRefListAccessRecordReq":               engine.UnRefer[larkacs.ListAccessRecordReq],
		"emptyOpeningTimeExternal":               engine.Empty[larkacs.OpeningTimeExternal],
		"emptyRefOpeningTimeExternal":            engine.EmptyRefer[larkacs.OpeningTimeExternal],
		"refOfOpeningTimeExternal":               engine.ReferOf[larkacs.OpeningTimeExternal],
		"unRefOpeningTimeExternal":               engine.UnRefer[larkacs.OpeningTimeExternal],
		"emptyOpeningTimeValidDayExternal":       engine.Empty[larkacs.OpeningTimeValidDayExternal],
		"emptyRefOpeningTimeValidDayExternal":    engine.EmptyRefer[larkacs.OpeningTimeValidDayExternal],
		"refOfOpeningTimeValidDayExternal":       engine.ReferOf[larkacs.OpeningTimeValidDayExternal],
		"unRefOpeningTimeValidDayExternal":       engine.UnRefer[larkacs.OpeningTimeValidDayExternal],
		"emptyV1":                                engine.Empty[larkacs.V1],
		"emptyRefV1":                             engine.EmptyRefer[larkacs.V1],
		"refOfV1":                                engine.ReferOf[larkacs.V1],
		"unRefV1":                                engine.UnRefer[larkacs.V1]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceAcs1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceAcs1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceAcs1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceAcs1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceAcs1Declared
}
