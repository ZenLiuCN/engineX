// Code generated by define_gene; DO NOT EDIT.
package larkbaike

import (
	_ "embed"
	"github.com/ZenLiuCN/engine"

	_ "github.com/ZenLiuCN/engine/modules/golang/context"
	_ "github.com/ZenLiuCN/engine/modules/golang/io"
	_ "github.com/ZenLiuCN/engineX/modules/github/com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike/v1"
)

var (
	//go:embed github_com_larksuite_oapi-sdk-go_v3_service_baike_v1.d.ts
	GithubComLarksuiteOapiSdkGo3ServiceBaike1Define   []byte
	GithubComLarksuiteOapiSdkGo3ServiceBaike1Declared = map[string]any{
		"newExtractEntityReqBodyBuilder":       larkbaike.NewExtractEntityReqBodyBuilder,
		"UserIdTypeSearchEntityUnionId":        larkbaike.UserIdTypeSearchEntityUnionId,
		"UserIdTypeUnionId":                    larkbaike.UserIdTypeUnionId,
		"newCreateEntityReqBuilder":            larkbaike.NewCreateEntityReqBuilder,
		"newDepartmentIdBuilder":               larkbaike.NewDepartmentIdBuilder,
		"UserIdTypeOpenId":                     larkbaike.UserIdTypeOpenId,
		"newCorrectPairBuilder":                larkbaike.NewCorrectPairBuilder,
		"newSearchEntityReqBuilder":            larkbaike.NewSearchEntityReqBuilder,
		"newStatisticalReportBuilder":          larkbaike.NewStatisticalReportBuilder,
		"newStatisticsBuilder":                 larkbaike.NewStatisticsBuilder,
		"newEntityBuilder":                     larkbaike.NewEntityBuilder,
		"newSpanBuilder":                       larkbaike.NewSpanBuilder,
		"newGetEntityReqBuilder":               larkbaike.NewGetEntityReqBuilder,
		"UserIdTypeCreateEntityOpenId":         larkbaike.UserIdTypeCreateEntityOpenId,
		"UserIdTypeCreateEntityUserId":         larkbaike.UserIdTypeCreateEntityUserId,
		"UserIdTypeUpdateEntityUserId":         larkbaike.UserIdTypeUpdateEntityUserId,
		"newCorrectErrorBuilder":               larkbaike.NewCorrectErrorBuilder,
		"newDisplayStatusBuilder":              larkbaike.NewDisplayStatusBuilder,
		"newMatchEntityReqBodyBuilder":         larkbaike.NewMatchEntityReqBodyBuilder,
		"newUpdateEntityReqBuilder":            larkbaike.NewUpdateEntityReqBuilder,
		"UserIdTypeUpdateEntityOpenId":         larkbaike.UserIdTypeUpdateEntityOpenId,
		"newClassificationBuilder":             larkbaike.NewClassificationBuilder,
		"newCorrectInfoBuilder":                larkbaike.NewCorrectInfoBuilder,
		"newRefererBuilder":                    larkbaike.NewRefererBuilder,
		"newUpdateDraftReqBuilder":             larkbaike.NewUpdateDraftReqBuilder,
		"newWordInfoBuilder":                   larkbaike.NewWordInfoBuilder,
		"newAbbreviationBuilder":               larkbaike.NewAbbreviationBuilder,
		"newMatchEntityReqBuilder":             larkbaike.NewMatchEntityReqBuilder,
		"newOuterInfoBuilder":                  larkbaike.NewOuterInfoBuilder,
		"newPhraseBuilder":                     larkbaike.NewPhraseBuilder,
		"UserIdTypeGetEntityUserId":            larkbaike.UserIdTypeGetEntityUserId,
		"newExtractEntityReqBuilder":           larkbaike.NewExtractEntityReqBuilder,
		"newListEntityReqBuilder":              larkbaike.NewListEntityReqBuilder,
		"newSearchEntityPathReqBodyBuilder":    larkbaike.NewSearchEntityPathReqBodyBuilder,
		"UserIdTypeCreateEntityUnionId":        larkbaike.UserIdTypeCreateEntityUnionId,
		"newClassificationFilterBuilder":       larkbaike.NewClassificationFilterBuilder,
		"newMatchInfoBuilder":                  larkbaike.NewMatchInfoBuilder,
		"UserIdTypeGetEntityUnionId":           larkbaike.UserIdTypeGetEntityUnionId,
		"newListClassificationReqBuilder":      larkbaike.NewListClassificationReqBuilder,
		"newUploadFileReqBuilder":              larkbaike.NewUploadFileReqBuilder,
		"newBaikeImageBuilder":                 larkbaike.NewBaikeImageBuilder,
		"UserIdTypeUpdateDraftOpenId":          larkbaike.UserIdTypeUpdateDraftOpenId,
		"UserIdTypeSearchEntityOpenId":         larkbaike.UserIdTypeSearchEntityOpenId,
		"newDraftBuilder":                      larkbaike.NewDraftBuilder,
		"newExtractEntityPathReqBodyBuilder":   larkbaike.NewExtractEntityPathReqBodyBuilder,
		"newTermBuilder":                       larkbaike.NewTermBuilder,
		"UserIdTypeUpdateDraftUserId":          larkbaike.UserIdTypeUpdateDraftUserId,
		"UserIdTypeUpdateEntityUnionId":        larkbaike.UserIdTypeUpdateEntityUnionId,
		"newCreateDraftReqBuilder":             larkbaike.NewCreateDraftReqBuilder,
		"newMatchEntityPathReqBodyBuilder":     larkbaike.NewMatchEntityPathReqBodyBuilder,
		"UserIdTypeUpdateDraftUnionId":         larkbaike.UserIdTypeUpdateDraftUnionId,
		"newFileBuilder":                       larkbaike.NewFileBuilder,
		"newRelatedMetaBuilder":                larkbaike.NewRelatedMetaBuilder,
		"newHighlightEntityReqBuilder":         larkbaike.NewHighlightEntityReqBuilder,
		"newSearchEntityReqBodyBuilder":        larkbaike.NewSearchEntityReqBodyBuilder,
		"UserIdTypeListEntityOpenId":           larkbaike.UserIdTypeListEntityOpenId,
		"New":                                  larkbaike.New,
		"newEntityWordBuilder":                 larkbaike.NewEntityWordBuilder,
		"UserIdTypeListEntityUnionId":          larkbaike.UserIdTypeListEntityUnionId,
		"UserIdTypeListEntityUserId":           larkbaike.UserIdTypeListEntityUserId,
		"newDownloadFileReqBuilder":            larkbaike.NewDownloadFileReqBuilder,
		"newHighlightEntityReqBodyBuilder":     larkbaike.NewHighlightEntityReqBodyBuilder,
		"UserIdTypeSearchEntityUserId":         larkbaike.UserIdTypeSearchEntityUserId,
		"UserIdTypeUserId":                     larkbaike.UserIdTypeUserId,
		"newHighlightEntityPathReqBodyBuilder": larkbaike.NewHighlightEntityPathReqBodyBuilder,
		"UserIdTypeGetEntityOpenId":            larkbaike.UserIdTypeGetEntityOpenId,

		"emptyListEntityIterator":            engine.Empty[larkbaike.ListEntityIterator],
		"emptyRefListEntityIterator":         engine.EmptyRefer[larkbaike.ListEntityIterator],
		"refOfListEntityIterator":            engine.ReferOf[larkbaike.ListEntityIterator],
		"unRefListEntityIterator":            engine.UnRefer[larkbaike.ListEntityIterator],
		"emptyMatchEntityReq":                engine.Empty[larkbaike.MatchEntityReq],
		"emptyRefMatchEntityReq":             engine.EmptyRefer[larkbaike.MatchEntityReq],
		"refOfMatchEntityReq":                engine.ReferOf[larkbaike.MatchEntityReq],
		"unRefMatchEntityReq":                engine.UnRefer[larkbaike.MatchEntityReq],
		"emptyOuterInfo":                     engine.Empty[larkbaike.OuterInfo],
		"emptyRefOuterInfo":                  engine.EmptyRefer[larkbaike.OuterInfo],
		"refOfOuterInfo":                     engine.ReferOf[larkbaike.OuterInfo],
		"unRefOuterInfo":                     engine.UnRefer[larkbaike.OuterInfo],
		"emptyMatchEntityReqBody":            engine.Empty[larkbaike.MatchEntityReqBody],
		"emptyRefMatchEntityReqBody":         engine.EmptyRefer[larkbaike.MatchEntityReqBody],
		"refOfMatchEntityReqBody":            engine.ReferOf[larkbaike.MatchEntityReqBody],
		"unRefMatchEntityReqBody":            engine.UnRefer[larkbaike.MatchEntityReqBody],
		"emptyDraft":                         engine.Empty[larkbaike.Draft],
		"emptyRefDraft":                      engine.EmptyRefer[larkbaike.Draft],
		"refOfDraft":                         engine.ReferOf[larkbaike.Draft],
		"unRefDraft":                         engine.UnRefer[larkbaike.Draft],
		"emptyExtractEntityRespData":         engine.Empty[larkbaike.ExtractEntityRespData],
		"emptyRefExtractEntityRespData":      engine.EmptyRefer[larkbaike.ExtractEntityRespData],
		"refOfExtractEntityRespData":         engine.ReferOf[larkbaike.ExtractEntityRespData],
		"unRefExtractEntityRespData":         engine.UnRefer[larkbaike.ExtractEntityRespData],
		"emptyListEntityResp":                engine.Empty[larkbaike.ListEntityResp],
		"emptyRefListEntityResp":             engine.EmptyRefer[larkbaike.ListEntityResp],
		"refOfListEntityResp":                engine.ReferOf[larkbaike.ListEntityResp],
		"unRefListEntityResp":                engine.UnRefer[larkbaike.ListEntityResp],
		"emptyPhrase":                        engine.Empty[larkbaike.Phrase],
		"emptyRefPhrase":                     engine.EmptyRefer[larkbaike.Phrase],
		"refOfPhrase":                        engine.ReferOf[larkbaike.Phrase],
		"unRefPhrase":                        engine.UnRefer[larkbaike.Phrase],
		"emptyExtractEntityReq":              engine.Empty[larkbaike.ExtractEntityReq],
		"emptyRefExtractEntityReq":           engine.EmptyRefer[larkbaike.ExtractEntityReq],
		"refOfExtractEntityReq":              engine.ReferOf[larkbaike.ExtractEntityReq],
		"unRefExtractEntityReq":              engine.UnRefer[larkbaike.ExtractEntityReq],
		"emptyWordInfo":                      engine.Empty[larkbaike.WordInfo],
		"emptyRefWordInfo":                   engine.EmptyRefer[larkbaike.WordInfo],
		"refOfWordInfo":                      engine.ReferOf[larkbaike.WordInfo],
		"unRefWordInfo":                      engine.UnRefer[larkbaike.WordInfo],
		"emptyCorrectPair":                   engine.Empty[larkbaike.CorrectPair],
		"emptyRefCorrectPair":                engine.EmptyRefer[larkbaike.CorrectPair],
		"refOfCorrectPair":                   engine.ReferOf[larkbaike.CorrectPair],
		"unRefCorrectPair":                   engine.UnRefer[larkbaike.CorrectPair],
		"emptySpan":                          engine.Empty[larkbaike.Span],
		"emptyRefSpan":                       engine.EmptyRefer[larkbaike.Span],
		"refOfSpan":                          engine.ReferOf[larkbaike.Span],
		"unRefSpan":                          engine.UnRefer[larkbaike.Span],
		"emptyUpdateDraftReq":                engine.Empty[larkbaike.UpdateDraftReq],
		"emptyRefUpdateDraftReq":             engine.EmptyRefer[larkbaike.UpdateDraftReq],
		"refOfUpdateDraftReq":                engine.ReferOf[larkbaike.UpdateDraftReq],
		"unRefUpdateDraftReq":                engine.UnRefer[larkbaike.UpdateDraftReq],
		"emptyListClassificationReq":         engine.Empty[larkbaike.ListClassificationReq],
		"emptyRefListClassificationReq":      engine.EmptyRefer[larkbaike.ListClassificationReq],
		"refOfListClassificationReq":         engine.ReferOf[larkbaike.ListClassificationReq],
		"unRefListClassificationReq":         engine.UnRefer[larkbaike.ListClassificationReq],
		"emptyMatchEntityResp":               engine.Empty[larkbaike.MatchEntityResp],
		"emptyRefMatchEntityResp":            engine.EmptyRefer[larkbaike.MatchEntityResp],
		"refOfMatchEntityResp":               engine.ReferOf[larkbaike.MatchEntityResp],
		"unRefMatchEntityResp":               engine.UnRefer[larkbaike.MatchEntityResp],
		"emptySearchEntityReq":               engine.Empty[larkbaike.SearchEntityReq],
		"emptyRefSearchEntityReq":            engine.EmptyRefer[larkbaike.SearchEntityReq],
		"refOfSearchEntityReq":               engine.ReferOf[larkbaike.SearchEntityReq],
		"unRefSearchEntityReq":               engine.UnRefer[larkbaike.SearchEntityReq],
		"emptyMatchEntityRespData":           engine.Empty[larkbaike.MatchEntityRespData],
		"emptyRefMatchEntityRespData":        engine.EmptyRefer[larkbaike.MatchEntityRespData],
		"refOfMatchEntityRespData":           engine.ReferOf[larkbaike.MatchEntityRespData],
		"unRefMatchEntityRespData":           engine.UnRefer[larkbaike.MatchEntityRespData],
		"emptySearchEntityRespData":          engine.Empty[larkbaike.SearchEntityRespData],
		"emptyRefSearchEntityRespData":       engine.EmptyRefer[larkbaike.SearchEntityRespData],
		"refOfSearchEntityRespData":          engine.ReferOf[larkbaike.SearchEntityRespData],
		"unRefSearchEntityRespData":          engine.UnRefer[larkbaike.SearchEntityRespData],
		"emptyDepartmentId":                  engine.Empty[larkbaike.DepartmentId],
		"emptyRefDepartmentId":               engine.EmptyRefer[larkbaike.DepartmentId],
		"refOfDepartmentId":                  engine.ReferOf[larkbaike.DepartmentId],
		"unRefDepartmentId":                  engine.UnRefer[larkbaike.DepartmentId],
		"emptyUploadFileRespData":            engine.Empty[larkbaike.UploadFileRespData],
		"emptyRefUploadFileRespData":         engine.EmptyRefer[larkbaike.UploadFileRespData],
		"refOfUploadFileRespData":            engine.ReferOf[larkbaike.UploadFileRespData],
		"unRefUploadFileRespData":            engine.UnRefer[larkbaike.UploadFileRespData],
		"emptyHighlightEntityReqBody":        engine.Empty[larkbaike.HighlightEntityReqBody],
		"emptyRefHighlightEntityReqBody":     engine.EmptyRefer[larkbaike.HighlightEntityReqBody],
		"refOfHighlightEntityReqBody":        engine.ReferOf[larkbaike.HighlightEntityReqBody],
		"unRefHighlightEntityReqBody":        engine.UnRefer[larkbaike.HighlightEntityReqBody],
		"emptyListEntityReq":                 engine.Empty[larkbaike.ListEntityReq],
		"emptyRefListEntityReq":              engine.EmptyRefer[larkbaike.ListEntityReq],
		"refOfListEntityReq":                 engine.ReferOf[larkbaike.ListEntityReq],
		"unRefListEntityReq":                 engine.UnRefer[larkbaike.ListEntityReq],
		"emptyStatistics":                    engine.Empty[larkbaike.Statistics],
		"emptyRefStatistics":                 engine.EmptyRefer[larkbaike.Statistics],
		"refOfStatistics":                    engine.ReferOf[larkbaike.Statistics],
		"unRefStatistics":                    engine.UnRefer[larkbaike.Statistics],
		"emptyUploadFileReq":                 engine.Empty[larkbaike.UploadFileReq],
		"emptyRefUploadFileReq":              engine.EmptyRefer[larkbaike.UploadFileReq],
		"refOfUploadFileReq":                 engine.ReferOf[larkbaike.UploadFileReq],
		"unRefUploadFileReq":                 engine.UnRefer[larkbaike.UploadFileReq],
		"emptyCreateDraftReq":                engine.Empty[larkbaike.CreateDraftReq],
		"emptyRefCreateDraftReq":             engine.EmptyRefer[larkbaike.CreateDraftReq],
		"refOfCreateDraftReq":                engine.ReferOf[larkbaike.CreateDraftReq],
		"unRefCreateDraftReq":                engine.UnRefer[larkbaike.CreateDraftReq],
		"emptyHighlightEntityReq":            engine.Empty[larkbaike.HighlightEntityReq],
		"emptyRefHighlightEntityReq":         engine.EmptyRefer[larkbaike.HighlightEntityReq],
		"refOfHighlightEntityReq":            engine.ReferOf[larkbaike.HighlightEntityReq],
		"unRefHighlightEntityReq":            engine.UnRefer[larkbaike.HighlightEntityReq],
		"emptyHighlightEntityResp":           engine.Empty[larkbaike.HighlightEntityResp],
		"emptyRefHighlightEntityResp":        engine.EmptyRefer[larkbaike.HighlightEntityResp],
		"refOfHighlightEntityResp":           engine.ReferOf[larkbaike.HighlightEntityResp],
		"unRefHighlightEntityResp":           engine.UnRefer[larkbaike.HighlightEntityResp],
		"emptyHighlightEntityRespData":       engine.Empty[larkbaike.HighlightEntityRespData],
		"emptyRefHighlightEntityRespData":    engine.EmptyRefer[larkbaike.HighlightEntityRespData],
		"refOfHighlightEntityRespData":       engine.ReferOf[larkbaike.HighlightEntityRespData],
		"unRefHighlightEntityRespData":       engine.UnRefer[larkbaike.HighlightEntityRespData],
		"emptyGetEntityResp":                 engine.Empty[larkbaike.GetEntityResp],
		"emptyRefGetEntityResp":              engine.EmptyRefer[larkbaike.GetEntityResp],
		"refOfGetEntityResp":                 engine.ReferOf[larkbaike.GetEntityResp],
		"unRefGetEntityResp":                 engine.UnRefer[larkbaike.GetEntityResp],
		"emptyListClassificationRespData":    engine.Empty[larkbaike.ListClassificationRespData],
		"emptyRefListClassificationRespData": engine.EmptyRefer[larkbaike.ListClassificationRespData],
		"refOfListClassificationRespData":    engine.ReferOf[larkbaike.ListClassificationRespData],
		"unRefListClassificationRespData":    engine.UnRefer[larkbaike.ListClassificationRespData],
		"emptyDownloadFileReq":               engine.Empty[larkbaike.DownloadFileReq],
		"emptyRefDownloadFileReq":            engine.EmptyRefer[larkbaike.DownloadFileReq],
		"refOfDownloadFileReq":               engine.ReferOf[larkbaike.DownloadFileReq],
		"unRefDownloadFileReq":               engine.UnRefer[larkbaike.DownloadFileReq],
		"emptyListEntityRespData":            engine.Empty[larkbaike.ListEntityRespData],
		"emptyRefListEntityRespData":         engine.EmptyRefer[larkbaike.ListEntityRespData],
		"refOfListEntityRespData":            engine.ReferOf[larkbaike.ListEntityRespData],
		"unRefListEntityRespData":            engine.UnRefer[larkbaike.ListEntityRespData],
		"emptyUpdateDraftRespData":           engine.Empty[larkbaike.UpdateDraftRespData],
		"emptyRefUpdateDraftRespData":        engine.EmptyRefer[larkbaike.UpdateDraftRespData],
		"refOfUpdateDraftRespData":           engine.ReferOf[larkbaike.UpdateDraftRespData],
		"unRefUpdateDraftRespData":           engine.UnRefer[larkbaike.UpdateDraftRespData],
		"emptyUploadFileResp":                engine.Empty[larkbaike.UploadFileResp],
		"emptyRefUploadFileResp":             engine.EmptyRefer[larkbaike.UploadFileResp],
		"refOfUploadFileResp":                engine.ReferOf[larkbaike.UploadFileResp],
		"unRefUploadFileResp":                engine.UnRefer[larkbaike.UploadFileResp],
		"emptyGetEntityReq":                  engine.Empty[larkbaike.GetEntityReq],
		"emptyRefGetEntityReq":               engine.EmptyRefer[larkbaike.GetEntityReq],
		"refOfGetEntityReq":                  engine.ReferOf[larkbaike.GetEntityReq],
		"unRefGetEntityReq":                  engine.UnRefer[larkbaike.GetEntityReq],
		"emptyClassification":                engine.Empty[larkbaike.Classification],
		"emptyRefClassification":             engine.EmptyRefer[larkbaike.Classification],
		"refOfClassification":                engine.ReferOf[larkbaike.Classification],
		"unRefClassification":                engine.UnRefer[larkbaike.Classification],
		"emptyDisplayStatus":                 engine.Empty[larkbaike.DisplayStatus],
		"emptyRefDisplayStatus":              engine.EmptyRefer[larkbaike.DisplayStatus],
		"refOfDisplayStatus":                 engine.ReferOf[larkbaike.DisplayStatus],
		"unRefDisplayStatus":                 engine.UnRefer[larkbaike.DisplayStatus],
		"emptyListClassificationIterator":    engine.Empty[larkbaike.ListClassificationIterator],
		"emptyRefListClassificationIterator": engine.EmptyRefer[larkbaike.ListClassificationIterator],
		"refOfListClassificationIterator":    engine.ReferOf[larkbaike.ListClassificationIterator],
		"unRefListClassificationIterator":    engine.UnRefer[larkbaike.ListClassificationIterator],
		"emptyCreateDraftResp":               engine.Empty[larkbaike.CreateDraftResp],
		"emptyRefCreateDraftResp":            engine.EmptyRefer[larkbaike.CreateDraftResp],
		"refOfCreateDraftResp":               engine.ReferOf[larkbaike.CreateDraftResp],
		"unRefCreateDraftResp":               engine.UnRefer[larkbaike.CreateDraftResp],
		"emptyStatisticalReport":             engine.Empty[larkbaike.StatisticalReport],
		"emptyRefStatisticalReport":          engine.EmptyRefer[larkbaike.StatisticalReport],
		"refOfStatisticalReport":             engine.ReferOf[larkbaike.StatisticalReport],
		"unRefStatisticalReport":             engine.UnRefer[larkbaike.StatisticalReport],
		"emptyUpdateEntityReq":               engine.Empty[larkbaike.UpdateEntityReq],
		"emptyRefUpdateEntityReq":            engine.EmptyRefer[larkbaike.UpdateEntityReq],
		"refOfUpdateEntityReq":               engine.ReferOf[larkbaike.UpdateEntityReq],
		"unRefUpdateEntityReq":               engine.UnRefer[larkbaike.UpdateEntityReq],
		"emptyCreateEntityRespData":          engine.Empty[larkbaike.CreateEntityRespData],
		"emptyRefCreateEntityRespData":       engine.EmptyRefer[larkbaike.CreateEntityRespData],
		"refOfCreateEntityRespData":          engine.ReferOf[larkbaike.CreateEntityRespData],
		"unRefCreateEntityRespData":          engine.UnRefer[larkbaike.CreateEntityRespData],
		"emptyListClassificationResp":        engine.Empty[larkbaike.ListClassificationResp],
		"emptyRefListClassificationResp":     engine.EmptyRefer[larkbaike.ListClassificationResp],
		"refOfListClassificationResp":        engine.ReferOf[larkbaike.ListClassificationResp],
		"unRefListClassificationResp":        engine.UnRefer[larkbaike.ListClassificationResp],
		"emptyCreateEntityResp":              engine.Empty[larkbaike.CreateEntityResp],
		"emptyRefCreateEntityResp":           engine.EmptyRefer[larkbaike.CreateEntityResp],
		"refOfCreateEntityResp":              engine.ReferOf[larkbaike.CreateEntityResp],
		"unRefCreateEntityResp":              engine.UnRefer[larkbaike.CreateEntityResp],
		"emptyEntityWord":                    engine.Empty[larkbaike.EntityWord],
		"emptyRefEntityWord":                 engine.EmptyRefer[larkbaike.EntityWord],
		"refOfEntityWord":                    engine.ReferOf[larkbaike.EntityWord],
		"unRefEntityWord":                    engine.UnRefer[larkbaike.EntityWord],
		"emptyUpdateDraftResp":               engine.Empty[larkbaike.UpdateDraftResp],
		"emptyRefUpdateDraftResp":            engine.EmptyRefer[larkbaike.UpdateDraftResp],
		"refOfUpdateDraftResp":               engine.ReferOf[larkbaike.UpdateDraftResp],
		"unRefUpdateDraftResp":               engine.UnRefer[larkbaike.UpdateDraftResp],
		"emptyClassificationFilter":          engine.Empty[larkbaike.ClassificationFilter],
		"emptyRefClassificationFilter":       engine.EmptyRefer[larkbaike.ClassificationFilter],
		"refOfClassificationFilter":          engine.ReferOf[larkbaike.ClassificationFilter],
		"unRefClassificationFilter":          engine.UnRefer[larkbaike.ClassificationFilter],
		"emptyExtractEntityResp":             engine.Empty[larkbaike.ExtractEntityResp],
		"emptyRefExtractEntityResp":          engine.EmptyRefer[larkbaike.ExtractEntityResp],
		"refOfExtractEntityResp":             engine.ReferOf[larkbaike.ExtractEntityResp],
		"unRefExtractEntityResp":             engine.UnRefer[larkbaike.ExtractEntityResp],
		"emptyMatchInfo":                     engine.Empty[larkbaike.MatchInfo],
		"emptyRefMatchInfo":                  engine.EmptyRefer[larkbaike.MatchInfo],
		"refOfMatchInfo":                     engine.ReferOf[larkbaike.MatchInfo],
		"unRefMatchInfo":                     engine.UnRefer[larkbaike.MatchInfo],
		"emptyCorrectInfo":                   engine.Empty[larkbaike.CorrectInfo],
		"emptyRefCorrectInfo":                engine.EmptyRefer[larkbaike.CorrectInfo],
		"refOfCorrectInfo":                   engine.ReferOf[larkbaike.CorrectInfo],
		"unRefCorrectInfo":                   engine.UnRefer[larkbaike.CorrectInfo],
		"emptyRelatedMeta":                   engine.Empty[larkbaike.RelatedMeta],
		"emptyRefRelatedMeta":                engine.EmptyRefer[larkbaike.RelatedMeta],
		"refOfRelatedMeta":                   engine.ReferOf[larkbaike.RelatedMeta],
		"unRefRelatedMeta":                   engine.UnRefer[larkbaike.RelatedMeta],
		"emptyAbbreviation":                  engine.Empty[larkbaike.Abbreviation],
		"emptyRefAbbreviation":               engine.EmptyRefer[larkbaike.Abbreviation],
		"refOfAbbreviation":                  engine.ReferOf[larkbaike.Abbreviation],
		"unRefAbbreviation":                  engine.UnRefer[larkbaike.Abbreviation],
		"emptyBaikeImage":                    engine.Empty[larkbaike.BaikeImage],
		"emptyRefBaikeImage":                 engine.EmptyRefer[larkbaike.BaikeImage],
		"refOfBaikeImage":                    engine.ReferOf[larkbaike.BaikeImage],
		"unRefBaikeImage":                    engine.UnRefer[larkbaike.BaikeImage],
		"emptyExtractEntityReqBody":          engine.Empty[larkbaike.ExtractEntityReqBody],
		"emptyRefExtractEntityReqBody":       engine.EmptyRefer[larkbaike.ExtractEntityReqBody],
		"refOfExtractEntityReqBody":          engine.ReferOf[larkbaike.ExtractEntityReqBody],
		"unRefExtractEntityReqBody":          engine.UnRefer[larkbaike.ExtractEntityReqBody],
		"emptyGetEntityRespData":             engine.Empty[larkbaike.GetEntityRespData],
		"emptyRefGetEntityRespData":          engine.EmptyRefer[larkbaike.GetEntityRespData],
		"refOfGetEntityRespData":             engine.ReferOf[larkbaike.GetEntityRespData],
		"unRefGetEntityRespData":             engine.UnRefer[larkbaike.GetEntityRespData],
		"emptyV1":                            engine.Empty[larkbaike.V1],
		"emptyRefV1":                         engine.EmptyRefer[larkbaike.V1],
		"refOfV1":                            engine.ReferOf[larkbaike.V1],
		"unRefV1":                            engine.UnRefer[larkbaike.V1],
		"emptyCreateDraftRespData":           engine.Empty[larkbaike.CreateDraftRespData],
		"emptyRefCreateDraftRespData":        engine.EmptyRefer[larkbaike.CreateDraftRespData],
		"refOfCreateDraftRespData":           engine.ReferOf[larkbaike.CreateDraftRespData],
		"unRefCreateDraftRespData":           engine.UnRefer[larkbaike.CreateDraftRespData],
		"emptyEntity":                        engine.Empty[larkbaike.Entity],
		"emptyRefEntity":                     engine.EmptyRefer[larkbaike.Entity],
		"refOfEntity":                        engine.ReferOf[larkbaike.Entity],
		"unRefEntity":                        engine.UnRefer[larkbaike.Entity],
		"emptySearchEntityReqBody":           engine.Empty[larkbaike.SearchEntityReqBody],
		"emptyRefSearchEntityReqBody":        engine.EmptyRefer[larkbaike.SearchEntityReqBody],
		"refOfSearchEntityReqBody":           engine.ReferOf[larkbaike.SearchEntityReqBody],
		"unRefSearchEntityReqBody":           engine.UnRefer[larkbaike.SearchEntityReqBody],
		"emptySearchEntityResp":              engine.Empty[larkbaike.SearchEntityResp],
		"emptyRefSearchEntityResp":           engine.EmptyRefer[larkbaike.SearchEntityResp],
		"refOfSearchEntityResp":              engine.ReferOf[larkbaike.SearchEntityResp],
		"unRefSearchEntityResp":              engine.UnRefer[larkbaike.SearchEntityResp],
		"emptyCreateEntityReq":               engine.Empty[larkbaike.CreateEntityReq],
		"emptyRefCreateEntityReq":            engine.EmptyRefer[larkbaike.CreateEntityReq],
		"refOfCreateEntityReq":               engine.ReferOf[larkbaike.CreateEntityReq],
		"unRefCreateEntityReq":               engine.UnRefer[larkbaike.CreateEntityReq],
		"emptyDownloadFileResp":              engine.Empty[larkbaike.DownloadFileResp],
		"emptyRefDownloadFileResp":           engine.EmptyRefer[larkbaike.DownloadFileResp],
		"refOfDownloadFileResp":              engine.ReferOf[larkbaike.DownloadFileResp],
		"unRefDownloadFileResp":              engine.UnRefer[larkbaike.DownloadFileResp],
		"emptySearchEntityIterator":          engine.Empty[larkbaike.SearchEntityIterator],
		"emptyRefSearchEntityIterator":       engine.EmptyRefer[larkbaike.SearchEntityIterator],
		"refOfSearchEntityIterator":          engine.ReferOf[larkbaike.SearchEntityIterator],
		"unRefSearchEntityIterator":          engine.UnRefer[larkbaike.SearchEntityIterator],
		"emptyReferer":                       engine.Empty[larkbaike.Referer],
		"emptyRefReferer":                    engine.EmptyRefer[larkbaike.Referer],
		"refOfReferer":                       engine.ReferOf[larkbaike.Referer],
		"unRefReferer":                       engine.UnRefer[larkbaike.Referer],
		"emptyTerm":                          engine.Empty[larkbaike.Term],
		"emptyRefTerm":                       engine.EmptyRefer[larkbaike.Term],
		"refOfTerm":                          engine.ReferOf[larkbaike.Term],
		"unRefTerm":                          engine.UnRefer[larkbaike.Term],
		"emptyUpdateEntityResp":              engine.Empty[larkbaike.UpdateEntityResp],
		"emptyRefUpdateEntityResp":           engine.EmptyRefer[larkbaike.UpdateEntityResp],
		"refOfUpdateEntityResp":              engine.ReferOf[larkbaike.UpdateEntityResp],
		"unRefUpdateEntityResp":              engine.UnRefer[larkbaike.UpdateEntityResp],
		"emptyUpdateEntityRespData":          engine.Empty[larkbaike.UpdateEntityRespData],
		"emptyRefUpdateEntityRespData":       engine.EmptyRefer[larkbaike.UpdateEntityRespData],
		"refOfUpdateEntityRespData":          engine.ReferOf[larkbaike.UpdateEntityRespData],
		"unRefUpdateEntityRespData":          engine.UnRefer[larkbaike.UpdateEntityRespData],
		"emptyFile":                          engine.Empty[larkbaike.File],
		"emptyRefFile":                       engine.EmptyRefer[larkbaike.File],
		"refOfFile":                          engine.ReferOf[larkbaike.File],
		"unRefFile":                          engine.UnRefer[larkbaike.File]}
)

func init() {
	engine.RegisterModule(GithubComLarksuiteOapiSdkGo3ServiceBaike1Module{})
}

type GithubComLarksuiteOapiSdkGo3ServiceBaike1Module struct{}

func (S GithubComLarksuiteOapiSdkGo3ServiceBaike1Module) Identity() string {
	return "github.com/larksuite/oapi-sdk-go/v3/service/baike/v1"
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBaike1Module) TypeDefine() []byte {
	return GithubComLarksuiteOapiSdkGo3ServiceBaike1Define
}
func (S GithubComLarksuiteOapiSdkGo3ServiceBaike1Module) Exports() map[string]any {
	return GithubComLarksuiteOapiSdkGo3ServiceBaike1Declared
}
