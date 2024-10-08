// Code generated by define_gene; DO NOT EDIT.
package larkehr

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/ehr/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_ehr_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceEhr1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceEhr1Declared = map[string]any{
		"newWorkLocationBuilder":     larkehr.NewWorkLocationBuilder,
		"New":                        larkehr.New,
		"newAttachmentBuilder":       larkehr.NewAttachmentBuilder,
		"newCustomFieldsBuilder":     larkehr.NewCustomFieldsBuilder,
		"newEmergencyContactBuilder": larkehr.NewEmergencyContactBuilder,
		"newJobBuilder":              larkehr.NewJobBuilder,
		"newListEmployeeReqBuilder":  larkehr.NewListEmployeeReqBuilder,
		"newWorkExperienceBuilder":   larkehr.NewWorkExperienceBuilder,
		"UserIdTypeUnionId":          larkehr.UserIdTypeUnionId,
		"ViewFull":                   larkehr.ViewFull,
		"newContractCompanyBuilder":  larkehr.NewContractCompanyBuilder,
		"newDepartmentIdBuilder":     larkehr.NewDepartmentIdBuilder,
		"newGetAttachmentReqBuilder": larkehr.NewGetAttachmentReqBuilder,
		"UserIdTypeOpenId":           larkehr.UserIdTypeOpenId,
		"UserIdTypeUserId":           larkehr.UserIdTypeUserId,
		"newEmployeeBuilder":         larkehr.NewEmployeeBuilder,
		"newManagerBuilder":          larkehr.NewManagerBuilder,
		"ViewBasic":                  larkehr.ViewBasic,
		"newEducationBuilder":        larkehr.NewEducationBuilder,
		"newJobLevelBuilder":         larkehr.NewJobLevelBuilder,
		"newNativeRegionBuilder":     larkehr.NewNativeRegionBuilder,
		"newSystemFieldsBuilder":     larkehr.NewSystemFieldsBuilder,

		"emptyDepartmentId":            engine.Empty[larkehr.DepartmentId],
		"emptyRefDepartmentId":         engine.EmptyRefer[larkehr.DepartmentId],
		"refOfDepartmentId":            engine.ReferOf[larkehr.DepartmentId],
		"unRefDepartmentId":            engine.UnRefer[larkehr.DepartmentId],
		"emptyJob":                     engine.Empty[larkehr.Job],
		"emptyRefJob":                  engine.EmptyRefer[larkehr.Job],
		"refOfJob":                     engine.ReferOf[larkehr.Job],
		"unRefJob":                     engine.UnRefer[larkehr.Job],
		"emptyAttachment":              engine.Empty[larkehr.Attachment],
		"emptyRefAttachment":           engine.EmptyRefer[larkehr.Attachment],
		"refOfAttachment":              engine.ReferOf[larkehr.Attachment],
		"unRefAttachment":              engine.UnRefer[larkehr.Attachment],
		"emptyContractCompany":         engine.Empty[larkehr.ContractCompany],
		"emptyRefContractCompany":      engine.EmptyRefer[larkehr.ContractCompany],
		"refOfContractCompany":         engine.ReferOf[larkehr.ContractCompany],
		"unRefContractCompany":         engine.UnRefer[larkehr.ContractCompany],
		"emptyGetAttachmentReq":        engine.Empty[larkehr.GetAttachmentReq],
		"emptyRefGetAttachmentReq":     engine.EmptyRefer[larkehr.GetAttachmentReq],
		"refOfGetAttachmentReq":        engine.ReferOf[larkehr.GetAttachmentReq],
		"unRefGetAttachmentReq":        engine.UnRefer[larkehr.GetAttachmentReq],
		"emptyV1":                      engine.Empty[larkehr.V1],
		"emptyRefV1":                   engine.EmptyRefer[larkehr.V1],
		"refOfV1":                      engine.ReferOf[larkehr.V1],
		"unRefV1":                      engine.UnRefer[larkehr.V1],
		"emptySystemFields":            engine.Empty[larkehr.SystemFields],
		"emptyRefSystemFields":         engine.EmptyRefer[larkehr.SystemFields],
		"refOfSystemFields":            engine.ReferOf[larkehr.SystemFields],
		"unRefSystemFields":            engine.UnRefer[larkehr.SystemFields],
		"emptyWorkExperience":          engine.Empty[larkehr.WorkExperience],
		"emptyRefWorkExperience":       engine.EmptyRefer[larkehr.WorkExperience],
		"refOfWorkExperience":          engine.ReferOf[larkehr.WorkExperience],
		"unRefWorkExperience":          engine.UnRefer[larkehr.WorkExperience],
		"emptyEmployee":                engine.Empty[larkehr.Employee],
		"emptyRefEmployee":             engine.EmptyRefer[larkehr.Employee],
		"refOfEmployee":                engine.ReferOf[larkehr.Employee],
		"unRefEmployee":                engine.UnRefer[larkehr.Employee],
		"emptyNativeRegion":            engine.Empty[larkehr.NativeRegion],
		"emptyRefNativeRegion":         engine.EmptyRefer[larkehr.NativeRegion],
		"refOfNativeRegion":            engine.ReferOf[larkehr.NativeRegion],
		"unRefNativeRegion":            engine.UnRefer[larkehr.NativeRegion],
		"emptyWorkLocation":            engine.Empty[larkehr.WorkLocation],
		"emptyRefWorkLocation":         engine.EmptyRefer[larkehr.WorkLocation],
		"refOfWorkLocation":            engine.ReferOf[larkehr.WorkLocation],
		"unRefWorkLocation":            engine.UnRefer[larkehr.WorkLocation],
		"emptyEducation":               engine.Empty[larkehr.Education],
		"emptyRefEducation":            engine.EmptyRefer[larkehr.Education],
		"refOfEducation":               engine.ReferOf[larkehr.Education],
		"unRefEducation":               engine.UnRefer[larkehr.Education],
		"emptyEmergencyContact":        engine.Empty[larkehr.EmergencyContact],
		"emptyRefEmergencyContact":     engine.EmptyRefer[larkehr.EmergencyContact],
		"refOfEmergencyContact":        engine.ReferOf[larkehr.EmergencyContact],
		"unRefEmergencyContact":        engine.UnRefer[larkehr.EmergencyContact],
		"emptyListEmployeeReq":         engine.Empty[larkehr.ListEmployeeReq],
		"emptyRefListEmployeeReq":      engine.EmptyRefer[larkehr.ListEmployeeReq],
		"refOfListEmployeeReq":         engine.ReferOf[larkehr.ListEmployeeReq],
		"unRefListEmployeeReq":         engine.UnRefer[larkehr.ListEmployeeReq],
		"emptyListEmployeeResp":        engine.Empty[larkehr.ListEmployeeResp],
		"emptyRefListEmployeeResp":     engine.EmptyRefer[larkehr.ListEmployeeResp],
		"refOfListEmployeeResp":        engine.ReferOf[larkehr.ListEmployeeResp],
		"unRefListEmployeeResp":        engine.UnRefer[larkehr.ListEmployeeResp],
		"emptyListEmployeeRespData":    engine.Empty[larkehr.ListEmployeeRespData],
		"emptyRefListEmployeeRespData": engine.EmptyRefer[larkehr.ListEmployeeRespData],
		"refOfListEmployeeRespData":    engine.ReferOf[larkehr.ListEmployeeRespData],
		"unRefListEmployeeRespData":    engine.UnRefer[larkehr.ListEmployeeRespData],
		"emptyCustomFields":            engine.Empty[larkehr.CustomFields],
		"emptyRefCustomFields":         engine.EmptyRefer[larkehr.CustomFields],
		"refOfCustomFields":            engine.ReferOf[larkehr.CustomFields],
		"unRefCustomFields":            engine.UnRefer[larkehr.CustomFields],
		"emptyListEmployeeIterator":    engine.Empty[larkehr.ListEmployeeIterator],
		"emptyRefListEmployeeIterator": engine.EmptyRefer[larkehr.ListEmployeeIterator],
		"refOfListEmployeeIterator":    engine.ReferOf[larkehr.ListEmployeeIterator],
		"unRefListEmployeeIterator":    engine.UnRefer[larkehr.ListEmployeeIterator],
		"emptyJobLevel":                engine.Empty[larkehr.JobLevel],
		"emptyRefJobLevel":             engine.EmptyRefer[larkehr.JobLevel],
		"refOfJobLevel":                engine.ReferOf[larkehr.JobLevel],
		"unRefJobLevel":                engine.UnRefer[larkehr.JobLevel],
		"emptyGetAttachmentResp":       engine.Empty[larkehr.GetAttachmentResp],
		"emptyRefGetAttachmentResp":    engine.EmptyRefer[larkehr.GetAttachmentResp],
		"refOfGetAttachmentResp":       engine.ReferOf[larkehr.GetAttachmentResp],
		"unRefGetAttachmentResp":       engine.UnRefer[larkehr.GetAttachmentResp],
		"emptyManager":                 engine.Empty[larkehr.Manager],
		"emptyRefManager":              engine.EmptyRefer[larkehr.Manager],
		"refOfManager":                 engine.ReferOf[larkehr.Manager],
		"unRefManager":                 engine.UnRefer[larkehr.Manager]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceEhr1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceEhr1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceEhr1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/ehr/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceEhr1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceEhr1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceEhr1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceEhr1Declared
}
