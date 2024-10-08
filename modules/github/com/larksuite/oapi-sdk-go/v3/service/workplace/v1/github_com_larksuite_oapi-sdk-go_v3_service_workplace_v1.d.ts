// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/workplace/v1'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {Struct,bool,error,int,Ref} from 'go'
	export interface AccessData extends Struct<AccessData>{

			pv:Ref<int>
			uv:Ref<int>
	}
	export interface AccessDataBuilder extends Struct<AccessDataBuilder>{

			pv(pv:int):Ref<AccessDataBuilder>
			uv(uv:int):Ref<AccessDataBuilder>
			build():Ref<AccessData>
	}
	export interface BlockAccessData extends Struct<BlockAccessData>{

			date:Ref<string>
			blockId:Ref<string>
			accessData:Ref<AccessData>
	}
	export interface BlockAccessDataBuilder extends Struct<BlockAccessDataBuilder>{

			date(date:string):Ref<BlockAccessDataBuilder>
			blockId(blockId:string):Ref<BlockAccessDataBuilder>
			accessData(accessData:Ref<AccessData>):Ref<BlockAccessDataBuilder>
			build():Ref<BlockAccessData>
	}
	export interface CustomWorkplaceAccessData extends Struct<CustomWorkplaceAccessData>{

			customWorkplaceId:Ref<string>
			accessData:Ref<AccessData>
			date:Ref<string>
			customWorkplaceName:Ref<I18nName>[]
	}
	export interface CustomWorkplaceAccessDataBuilder extends Struct<CustomWorkplaceAccessDataBuilder>{

			customWorkplaceId(customWorkplaceId:string):Ref<CustomWorkplaceAccessDataBuilder>
			accessData(accessData:Ref<AccessData>):Ref<CustomWorkplaceAccessDataBuilder>
			date(date:string):Ref<CustomWorkplaceAccessDataBuilder>
			customWorkplaceName(customWorkplaceName:Ref<I18nName>[]):Ref<CustomWorkplaceAccessDataBuilder>
			build():Ref<CustomWorkplaceAccessData>
	}
	export interface DepartmentId extends Struct<DepartmentId>{

			departmentId:Ref<string>
			openDepartmentId:Ref<string>
	}
	export interface DepartmentIdBuilder extends Struct<DepartmentIdBuilder>{

			departmentId(departmentId:string):Ref<DepartmentIdBuilder>
			openDepartmentId(openDepartmentId:string):Ref<DepartmentIdBuilder>
			build():Ref<DepartmentId>
	}
	export interface I18nName extends Struct<I18nName>{

			language:Ref<string>
			name:Ref<string>
	}
	export interface I18nNameBuilder extends Struct<I18nNameBuilder>{

			language(language:string):Ref<I18nNameBuilder>
			name(name:string):Ref<I18nNameBuilder>
			build():Ref<I18nName>
	}
	export function New(config:Ref<larkcore.Config>):Ref<V1>

	export function newAccessDataBuilder():Ref<AccessDataBuilder>

	export function newBlockAccessDataBuilder():Ref<BlockAccessDataBuilder>

	export function newCustomWorkplaceAccessDataBuilder():Ref<CustomWorkplaceAccessDataBuilder>

	export function newDepartmentIdBuilder():Ref<DepartmentIdBuilder>

	export function newI18nNameBuilder():Ref<I18nNameBuilder>

	export function newRuleBuilder():Ref<RuleBuilder>

	export function newSearchCustomWorkplaceAccessDataReqBuilder():Ref<SearchCustomWorkplaceAccessDataReqBuilder>

	export function newSearchWorkplaceAccessDataReqBuilder():Ref<SearchWorkplaceAccessDataReqBuilder>

	export function newSearchWorkplaceBlockAccessDataReqBuilder():Ref<SearchWorkplaceBlockAccessDataReqBuilder>

	export function newTemplateWorkplaceAccessDataBuilder():Ref<TemplateWorkplaceAccessDataBuilder>

	export function newWorkplaceAccessDataBuilder():Ref<WorkplaceAccessDataBuilder>

	export function newWorkplaceTenantNotificationBuilder():Ref<WorkplaceTenantNotificationBuilder>

	export function newWorkplaceUserNotificationBuilder():Ref<WorkplaceUserNotificationBuilder>

	export interface Rule extends Struct<Rule>{

			isAllVisible:Ref<bool>
			visibleDepartmentIds:string[]
	}
	export interface RuleBuilder extends Struct<RuleBuilder>{

			isAllVisible(isAllVisible:bool):Ref<RuleBuilder>
			visibleDepartmentIds(visibleDepartmentIds:string[]):Ref<RuleBuilder>
			build():Ref<Rule>
	}
	export interface SearchCustomWorkplaceAccessDataReq extends Struct<SearchCustomWorkplaceAccessDataReq>{

	}
	export interface SearchCustomWorkplaceAccessDataReqBuilder extends Struct<SearchCustomWorkplaceAccessDataReqBuilder>{

			fromDate(fromDate:string):Ref<SearchCustomWorkplaceAccessDataReqBuilder>
			toDate(toDate:string):Ref<SearchCustomWorkplaceAccessDataReqBuilder>
			pageSize(pageSize:int):Ref<SearchCustomWorkplaceAccessDataReqBuilder>
			pageToken(pageToken:string):Ref<SearchCustomWorkplaceAccessDataReqBuilder>
			customWorkplaceId(customWorkplaceId:string):Ref<SearchCustomWorkplaceAccessDataReqBuilder>
			build():Ref<SearchCustomWorkplaceAccessDataReq>
	}
	export interface SearchCustomWorkplaceAccessDataResp extends Struct<SearchCustomWorkplaceAccessDataResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<SearchCustomWorkplaceAccessDataRespData>
			success():bool
	}
	export interface SearchCustomWorkplaceAccessDataRespData extends Struct<SearchCustomWorkplaceAccessDataRespData>{

			items:Ref<CustomWorkplaceAccessData>[]
			hasMore:Ref<bool>
			pageToken:Ref<string>
	}
	export interface SearchWorkplaceAccessDataReq extends Struct<SearchWorkplaceAccessDataReq>{

	}
	export interface SearchWorkplaceAccessDataReqBuilder extends Struct<SearchWorkplaceAccessDataReqBuilder>{

			fromDate(fromDate:string):Ref<SearchWorkplaceAccessDataReqBuilder>
			toDate(toDate:string):Ref<SearchWorkplaceAccessDataReqBuilder>
			pageSize(pageSize:int):Ref<SearchWorkplaceAccessDataReqBuilder>
			pageToken(pageToken:string):Ref<SearchWorkplaceAccessDataReqBuilder>
			build():Ref<SearchWorkplaceAccessDataReq>
	}
	export interface SearchWorkplaceAccessDataResp extends Struct<SearchWorkplaceAccessDataResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<SearchWorkplaceAccessDataRespData>
			success():bool
	}
	export interface SearchWorkplaceAccessDataRespData extends Struct<SearchWorkplaceAccessDataRespData>{

			items:Ref<WorkplaceAccessData>[]
			hasMore:Ref<bool>
			pageToken:Ref<string>
	}
	export interface SearchWorkplaceBlockAccessDataReq extends Struct<SearchWorkplaceBlockAccessDataReq>{

	}
	export interface SearchWorkplaceBlockAccessDataReqBuilder extends Struct<SearchWorkplaceBlockAccessDataReqBuilder>{

			fromDate(fromDate:string):Ref<SearchWorkplaceBlockAccessDataReqBuilder>
			toDate(toDate:string):Ref<SearchWorkplaceBlockAccessDataReqBuilder>
			pageSize(pageSize:int):Ref<SearchWorkplaceBlockAccessDataReqBuilder>
			pageToken(pageToken:string):Ref<SearchWorkplaceBlockAccessDataReqBuilder>
			blockId(blockId:string):Ref<SearchWorkplaceBlockAccessDataReqBuilder>
			build():Ref<SearchWorkplaceBlockAccessDataReq>
	}
	export interface SearchWorkplaceBlockAccessDataResp extends Struct<SearchWorkplaceBlockAccessDataResp>{

			apiResp:Ref<larkcore.ApiResp>
			codeError:larkcore.CodeError
			data:Ref<SearchWorkplaceBlockAccessDataRespData>
			success():bool
	}
	export interface SearchWorkplaceBlockAccessDataRespData extends Struct<SearchWorkplaceBlockAccessDataRespData>{

			items:Ref<BlockAccessData>[]
			hasMore:Ref<bool>
			pageToken:Ref<string>
	}
	export interface TemplateWorkplaceAccessData extends Struct<TemplateWorkplaceAccessData>{

			tplId:Ref<string>
			accessData:Ref<AccessData>
	}
	export interface TemplateWorkplaceAccessDataBuilder extends Struct<TemplateWorkplaceAccessDataBuilder>{

			tplId(tplId:string):Ref<TemplateWorkplaceAccessDataBuilder>
			accessData(accessData:Ref<AccessData>):Ref<TemplateWorkplaceAccessDataBuilder>
			build():Ref<TemplateWorkplaceAccessData>
	}
	export interface V1 extends Struct<V1>{

			customWorkplaceAccessData:Ref<{
			
				search(ctx:context.Context,req:Ref<SearchCustomWorkplaceAccessDataReq>,...options:larkcore.RequestOptionFunc[]):Ref<SearchCustomWorkplaceAccessDataResp>
			}>
			workplaceAccessData:Ref<{
			
				search(ctx:context.Context,req:Ref<SearchWorkplaceAccessDataReq>,...options:larkcore.RequestOptionFunc[]):Ref<SearchWorkplaceAccessDataResp>
			}>
			workplaceBlockAccessData:Ref<{
			
				search(ctx:context.Context,req:Ref<SearchWorkplaceBlockAccessDataReq>,...options:larkcore.RequestOptionFunc[]):Ref<SearchWorkplaceBlockAccessDataResp>
			}>
	}
	export interface WorkplaceAccessData extends Struct<WorkplaceAccessData>{

			date:Ref<string>
			allWorkplace:Ref<AccessData>
			defaultWorkplace:Ref<AccessData>
	}
	export interface WorkplaceAccessDataBuilder extends Struct<WorkplaceAccessDataBuilder>{

			date(date:string):Ref<WorkplaceAccessDataBuilder>
			allWorkplace(allWorkplace:Ref<AccessData>):Ref<WorkplaceAccessDataBuilder>
			defaultWorkplace(defaultWorkplace:Ref<AccessData>):Ref<WorkplaceAccessDataBuilder>
			build():Ref<WorkplaceAccessData>
	}
	export interface WorkplaceTenantNotification extends Struct<WorkplaceTenantNotification>{

			notificationId:Ref<string>
			content:Ref<string>
			expireTime:Ref<string>
			rule:Ref<Rule>
	}
	export interface WorkplaceTenantNotificationBuilder extends Struct<WorkplaceTenantNotificationBuilder>{

			notificationId(notificationId:string):Ref<WorkplaceTenantNotificationBuilder>
			content(content:string):Ref<WorkplaceTenantNotificationBuilder>
			expireTime(expireTime:string):Ref<WorkplaceTenantNotificationBuilder>
			rule(rule:Ref<Rule>):Ref<WorkplaceTenantNotificationBuilder>
			build():Ref<WorkplaceTenantNotification>
	}
	export interface WorkplaceUserNotification extends Struct<WorkplaceUserNotification>{

			notificationId:Ref<string>
			content:Ref<string>
			expireTime:Ref<string>
	}
	export interface WorkplaceUserNotificationBuilder extends Struct<WorkplaceUserNotificationBuilder>{

			notificationId(notificationId:string):Ref<WorkplaceUserNotificationBuilder>
			content(content:string):Ref<WorkplaceUserNotificationBuilder>
			expireTime(expireTime:string):Ref<WorkplaceUserNotificationBuilder>
			build():Ref<WorkplaceUserNotification>
	}
	export function emptySearchWorkplaceBlockAccessDataReq():SearchWorkplaceBlockAccessDataReq
	export function emptyRefSearchWorkplaceBlockAccessDataReq():Ref<SearchWorkplaceBlockAccessDataReq>
	export function refOfSearchWorkplaceBlockAccessDataReq(x:SearchWorkplaceBlockAccessDataReq,v:Ref<SearchWorkplaceBlockAccessDataReq>)
	export function unRefSearchWorkplaceBlockAccessDataReq(v:Ref<SearchWorkplaceBlockAccessDataReq>):SearchWorkplaceBlockAccessDataReq
	export function emptyWorkplaceAccessData():WorkplaceAccessData
	export function emptyRefWorkplaceAccessData():Ref<WorkplaceAccessData>
	export function refOfWorkplaceAccessData(x:WorkplaceAccessData,v:Ref<WorkplaceAccessData>)
	export function unRefWorkplaceAccessData(v:Ref<WorkplaceAccessData>):WorkplaceAccessData
	export function emptySearchWorkplaceAccessDataReq():SearchWorkplaceAccessDataReq
	export function emptyRefSearchWorkplaceAccessDataReq():Ref<SearchWorkplaceAccessDataReq>
	export function refOfSearchWorkplaceAccessDataReq(x:SearchWorkplaceAccessDataReq,v:Ref<SearchWorkplaceAccessDataReq>)
	export function unRefSearchWorkplaceAccessDataReq(v:Ref<SearchWorkplaceAccessDataReq>):SearchWorkplaceAccessDataReq
	export function emptySearchWorkplaceAccessDataRespData():SearchWorkplaceAccessDataRespData
	export function emptyRefSearchWorkplaceAccessDataRespData():Ref<SearchWorkplaceAccessDataRespData>
	export function refOfSearchWorkplaceAccessDataRespData(x:SearchWorkplaceAccessDataRespData,v:Ref<SearchWorkplaceAccessDataRespData>)
	export function unRefSearchWorkplaceAccessDataRespData(v:Ref<SearchWorkplaceAccessDataRespData>):SearchWorkplaceAccessDataRespData
	export function emptyCustomWorkplaceAccessData():CustomWorkplaceAccessData
	export function emptyRefCustomWorkplaceAccessData():Ref<CustomWorkplaceAccessData>
	export function refOfCustomWorkplaceAccessData(x:CustomWorkplaceAccessData,v:Ref<CustomWorkplaceAccessData>)
	export function unRefCustomWorkplaceAccessData(v:Ref<CustomWorkplaceAccessData>):CustomWorkplaceAccessData
	export function emptyWorkplaceTenantNotification():WorkplaceTenantNotification
	export function emptyRefWorkplaceTenantNotification():Ref<WorkplaceTenantNotification>
	export function refOfWorkplaceTenantNotification(x:WorkplaceTenantNotification,v:Ref<WorkplaceTenantNotification>)
	export function unRefWorkplaceTenantNotification(v:Ref<WorkplaceTenantNotification>):WorkplaceTenantNotification
	export function emptySearchWorkplaceBlockAccessDataResp():SearchWorkplaceBlockAccessDataResp
	export function emptyRefSearchWorkplaceBlockAccessDataResp():Ref<SearchWorkplaceBlockAccessDataResp>
	export function refOfSearchWorkplaceBlockAccessDataResp(x:SearchWorkplaceBlockAccessDataResp,v:Ref<SearchWorkplaceBlockAccessDataResp>)
	export function unRefSearchWorkplaceBlockAccessDataResp(v:Ref<SearchWorkplaceBlockAccessDataResp>):SearchWorkplaceBlockAccessDataResp
	export function emptyWorkplaceUserNotification():WorkplaceUserNotification
	export function emptyRefWorkplaceUserNotification():Ref<WorkplaceUserNotification>
	export function refOfWorkplaceUserNotification(x:WorkplaceUserNotification,v:Ref<WorkplaceUserNotification>)
	export function unRefWorkplaceUserNotification(v:Ref<WorkplaceUserNotification>):WorkplaceUserNotification
	export function emptyAccessData():AccessData
	export function emptyRefAccessData():Ref<AccessData>
	export function refOfAccessData(x:AccessData,v:Ref<AccessData>)
	export function unRefAccessData(v:Ref<AccessData>):AccessData
	export function emptyBlockAccessData():BlockAccessData
	export function emptyRefBlockAccessData():Ref<BlockAccessData>
	export function refOfBlockAccessData(x:BlockAccessData,v:Ref<BlockAccessData>)
	export function unRefBlockAccessData(v:Ref<BlockAccessData>):BlockAccessData
	export function emptySearchWorkplaceBlockAccessDataRespData():SearchWorkplaceBlockAccessDataRespData
	export function emptyRefSearchWorkplaceBlockAccessDataRespData():Ref<SearchWorkplaceBlockAccessDataRespData>
	export function refOfSearchWorkplaceBlockAccessDataRespData(x:SearchWorkplaceBlockAccessDataRespData,v:Ref<SearchWorkplaceBlockAccessDataRespData>)
	export function unRefSearchWorkplaceBlockAccessDataRespData(v:Ref<SearchWorkplaceBlockAccessDataRespData>):SearchWorkplaceBlockAccessDataRespData
	export function emptyV1():V1
	export function emptyRefV1():Ref<V1>
	export function refOfV1(x:V1,v:Ref<V1>)
	export function unRefV1(v:Ref<V1>):V1
	export function emptySearchCustomWorkplaceAccessDataReq():SearchCustomWorkplaceAccessDataReq
	export function emptyRefSearchCustomWorkplaceAccessDataReq():Ref<SearchCustomWorkplaceAccessDataReq>
	export function refOfSearchCustomWorkplaceAccessDataReq(x:SearchCustomWorkplaceAccessDataReq,v:Ref<SearchCustomWorkplaceAccessDataReq>)
	export function unRefSearchCustomWorkplaceAccessDataReq(v:Ref<SearchCustomWorkplaceAccessDataReq>):SearchCustomWorkplaceAccessDataReq
	export function emptySearchCustomWorkplaceAccessDataResp():SearchCustomWorkplaceAccessDataResp
	export function emptyRefSearchCustomWorkplaceAccessDataResp():Ref<SearchCustomWorkplaceAccessDataResp>
	export function refOfSearchCustomWorkplaceAccessDataResp(x:SearchCustomWorkplaceAccessDataResp,v:Ref<SearchCustomWorkplaceAccessDataResp>)
	export function unRefSearchCustomWorkplaceAccessDataResp(v:Ref<SearchCustomWorkplaceAccessDataResp>):SearchCustomWorkplaceAccessDataResp
	export function emptySearchWorkplaceAccessDataResp():SearchWorkplaceAccessDataResp
	export function emptyRefSearchWorkplaceAccessDataResp():Ref<SearchWorkplaceAccessDataResp>
	export function refOfSearchWorkplaceAccessDataResp(x:SearchWorkplaceAccessDataResp,v:Ref<SearchWorkplaceAccessDataResp>)
	export function unRefSearchWorkplaceAccessDataResp(v:Ref<SearchWorkplaceAccessDataResp>):SearchWorkplaceAccessDataResp
	export function emptyDepartmentId():DepartmentId
	export function emptyRefDepartmentId():Ref<DepartmentId>
	export function refOfDepartmentId(x:DepartmentId,v:Ref<DepartmentId>)
	export function unRefDepartmentId(v:Ref<DepartmentId>):DepartmentId
	export function emptyI18nName():I18nName
	export function emptyRefI18nName():Ref<I18nName>
	export function refOfI18nName(x:I18nName,v:Ref<I18nName>)
	export function unRefI18nName(v:Ref<I18nName>):I18nName
	export function emptyRule():Rule
	export function emptyRefRule():Ref<Rule>
	export function refOfRule(x:Rule,v:Ref<Rule>)
	export function unRefRule(v:Ref<Rule>):Rule
	export function emptySearchCustomWorkplaceAccessDataRespData():SearchCustomWorkplaceAccessDataRespData
	export function emptyRefSearchCustomWorkplaceAccessDataRespData():Ref<SearchCustomWorkplaceAccessDataRespData>
	export function refOfSearchCustomWorkplaceAccessDataRespData(x:SearchCustomWorkplaceAccessDataRespData,v:Ref<SearchCustomWorkplaceAccessDataRespData>)
	export function unRefSearchCustomWorkplaceAccessDataRespData(v:Ref<SearchCustomWorkplaceAccessDataRespData>):SearchCustomWorkplaceAccessDataRespData
	export function emptyTemplateWorkplaceAccessData():TemplateWorkplaceAccessData
	export function emptyRefTemplateWorkplaceAccessData():Ref<TemplateWorkplaceAccessData>
	export function refOfTemplateWorkplaceAccessData(x:TemplateWorkplaceAccessData,v:Ref<TemplateWorkplaceAccessData>)
	export function unRefTemplateWorkplaceAccessData(v:Ref<TemplateWorkplaceAccessData>):TemplateWorkplaceAccessData
}