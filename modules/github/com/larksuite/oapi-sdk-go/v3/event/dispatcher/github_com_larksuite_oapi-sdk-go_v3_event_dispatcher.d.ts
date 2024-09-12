// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/event/dispatcher'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as larkcorehr from 'github.com/larksuite/oapi-sdk-go/v3/service/corehr/v2'
	// @ts-ignore
	import * as larkmeeting_room from 'github.com/larksuite/oapi-sdk-go/v3/service/meeting_room/v1'
	// @ts-ignore
	import * as larkevent from 'github.com/larksuite/oapi-sdk-go/v3/event'
	// @ts-ignore
	import * as json from 'golang/encoding/json'
	// @ts-ignore
	import * as larkapproval from 'github.com/larksuite/oapi-sdk-go/v3/service/approval/v4'
	// @ts-ignore
	import * as larkcontact from 'github.com/larksuite/oapi-sdk-go/v3/service/contact/v3'
	// @ts-ignore
	import * as larkhire from 'github.com/larksuite/oapi-sdk-go/v3/service/hire/v1'
	// @ts-ignore
	import * as larkim from 'github.com/larksuite/oapi-sdk-go/v3/service/im/v1'
	// @ts-ignore
	import * as larkhelpdesk from 'github.com/larksuite/oapi-sdk-go/v3/service/helpdesk/v1'
	// @ts-ignore
	import * as larktask from 'github.com/larksuite/oapi-sdk-go/v3/service/task/v1'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import * as larkacs from 'github.com/larksuite/oapi-sdk-go/v3/service/acs/v1'
	// @ts-ignore
	import * as larkapplication from 'github.com/larksuite/oapi-sdk-go/v3/service/application/v6'
	// @ts-ignore
	import * as larkcalendar from 'github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4'
	// @ts-ignore
	import * as larkdrive from 'github.com/larksuite/oapi-sdk-go/v3/service/drive/v1'
	// @ts-ignore
	import * as larkcorehr from 'github.com/larksuite/oapi-sdk-go/v3/service/corehr/v1'
	// @ts-ignore
	import * as larkvc from 'github.com/larksuite/oapi-sdk-go/v3/service/vc/v1'
	// @ts-ignore
	import type {error,GoError,Ref,Struct} from 'go'
	export interface AppTicketEvent extends json.Token,Struct<AppTicketEvent>{

			eventBase:Ref<larkevent.EventBase>
	}
	export interface CustomAppTicketEventHandler extends json.Token,Struct<CustomAppTicketEventHandler>,larkevent.EventHandler{

			event():any
			handle(ctx:context.Context,event:any)/*error*/
	}
	export interface EventDispatcher extends larkevent.IReqHandler,Struct<EventDispatcher>,json.Token{

			config:Ref<larkcore.Config>
			onP2AccessRecordCreatedV1(v1:(ctx:context.Context,event:Ref<larkacs.P2AccessRecordCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2UserUpdatedV1(v1:(ctx:context.Context,event:Ref<larkacs.P2UserUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationCreatedV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationCreatedV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationAppVersionAuditV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationAppVersionAuditV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationAppVersionPublishApplyV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationAppVersionPublishApplyV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationAppVersionPublishRevokeV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationAppVersionPublishRevokeV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationFeedbackCreatedV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationFeedbackCreatedV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationFeedbackUpdatedV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationFeedbackUpdatedV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationVisibilityAddedV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2ApplicationVisibilityAddedV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2BotMenuV6(v1:(ctx:context.Context,event:Ref<larkapplication.P2BotMenuV6>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApprovalUpdatedV4(v1:(ctx:context.Context,event:Ref<larkapproval.P2ApprovalUpdatedV4>)=>void/*error*/):Ref<EventDispatcher>
			onP2CalendarChangedV4(v1:(ctx:context.Context,event:Ref<larkcalendar.P2CalendarChangedV4>)=>void/*error*/):Ref<EventDispatcher>
			onP2CalendarAclCreatedV4(v1:(ctx:context.Context,event:Ref<larkcalendar.P2CalendarAclCreatedV4>)=>void/*error*/):Ref<EventDispatcher>
			onP2CalendarAclDeletedV4(v1:(ctx:context.Context,event:Ref<larkcalendar.P2CalendarAclDeletedV4>)=>void/*error*/):Ref<EventDispatcher>
			onP2CalendarEventChangedV4(v1:(ctx:context.Context,event:Ref<larkcalendar.P2CalendarEventChangedV4>)=>void/*error*/):Ref<EventDispatcher>
			onP2CustomAttrEventUpdatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2CustomAttrEventUpdatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentCreatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2DepartmentCreatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentDeletedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2DepartmentDeletedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentUpdatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2DepartmentUpdatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeTypeEnumActivedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2EmployeeTypeEnumActivedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeTypeEnumCreatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2EmployeeTypeEnumCreatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeTypeEnumDeactivatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2EmployeeTypeEnumDeactivatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeTypeEnumDeletedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2EmployeeTypeEnumDeletedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeTypeEnumUpdatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2EmployeeTypeEnumUpdatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2ScopeUpdatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2ScopeUpdatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2UserCreatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2UserCreatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2UserDeletedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2UserDeletedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2UserUpdatedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P2UserUpdatedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP2ContractCreatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ContractCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ContractDeletedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ContractDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ContractUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ContractUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentCreatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2DepartmentCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentDeletedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2DepartmentDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2DepartmentUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2DepartmentUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmploymentConvertedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmploymentConvertedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmploymentCreatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmploymentCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmploymentDeletedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmploymentDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmploymentResignedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmploymentResignedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmploymentUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmploymentUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobCreatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobDeletedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobChangeUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobChangeUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobDataChangedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobDataChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobDataEmployedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobDataEmployedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2OffboardingUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2OffboardingUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2OrgRoleAuthorizationUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2OrgRoleAuthorizationUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2PersonCreatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2PersonCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2PersonDeletedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2PersonDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2PersonUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2PersonUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2PreHireUpdatedV1(v1:(ctx:context.Context,event:Ref<larkcorehr.P2PreHireUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EmployeeDomainEventV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2EmployeeDomainEventV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2JobChangeUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2JobChangeUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2OffboardingChecklistUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2OffboardingChecklistUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2OffboardingStatusUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2OffboardingStatusUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2OffboardingUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2OffboardingUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2ProbationUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ProbationUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2ProcessUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ProcessUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2ProcessApproverUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ProcessApproverUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2ProcessCcUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ProcessCcUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			onP2ProcessNodeUpdatedV2(v1:(ctx:context.Context,event:Ref<larkcorehr.P2ProcessNodeUpdatedV2>)=>void/*error*/):Ref<EventDispatcher>
			logger():larkcore.Logger
			initConfig(...options:larkevent.OptionFunc[]):void
			handle(ctx:context.Context,req:Ref<larkevent.EventReq>):Ref<larkevent.EventResp>
			parseReq(ctx:context.Context,req:Ref<larkevent.EventReq>):string
			decryptEvent(ctx:context.Context,cipherEventJsonStr:string):string
			verifySign(ctx:context.Context,req:Ref<larkevent.EventReq>)/*error*/
			authByChallenge(ctx:context.Context,reqType:larkevent.ReqType,challenge:string,token:string):Ref<larkevent.EventResp>
			Do(ctx:context.Context,payload:Uint8Array)/*error*/
			doHandle(ctx:context.Context,reqType:larkevent.ReqType,eventType:string,challenge:string,token:string,plainEventJsonStr:string,path:string,req:Ref<larkevent.EventReq>):Ref<larkevent.EventResp>
			onP2FileBitableFieldChangedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileBitableFieldChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileBitableRecordChangedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileBitableRecordChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileDeletedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileEditV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileEditV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FilePermissionMemberAddedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FilePermissionMemberAddedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FilePermissionMemberRemovedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FilePermissionMemberRemovedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileReadV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileReadV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileTitleUpdatedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileTitleUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2FileTrashedV1(v1:(ctx:context.Context,event:Ref<larkdrive.P2FileTrashedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1OrderPaidV6(v1:(ctx:context.Context,event:Ref<larkapplication.P1OrderPaidV6>)=>void/*error*/):Ref<EventDispatcher>
			onP1AppUninstalledV6(v1:(ctx:context.Context,event:Ref<larkapplication.P1AppUninstalledV6>)=>void/*error*/):Ref<EventDispatcher>
			onP1AppStatusChangedV6(v1:(ctx:context.Context,event:Ref<larkapplication.P1AppStatusChangedV6>)=>void/*error*/):Ref<EventDispatcher>
			onP1MessageReadV1(v1:(ctx:context.Context,event:Ref<larkim.P1MessageReadV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1MessageReceiveV1(v1:(ctx:context.Context,event:Ref<larkim.P1MessageReceiveV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1UserStatusChangedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P1UserStatusChangedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP1UserChangedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P1UserChangedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP1DepartmentChangedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P1DepartmentChangedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP1ContactScopeChangedV3(v1:(ctx:context.Context,event:Ref<larkcontact.P1ContactScopeChangedV3>)=>void/*error*/):Ref<EventDispatcher>
			onP1AddBotV1(v1:(ctx:context.Context,event:Ref<larkim.P1AddBotV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1RemoveAddBotV1(v1:(ctx:context.Context,event:Ref<larkim.P1RemoveBotV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1UserInOutChatV1(v1:(ctx:context.Context,event:Ref<larkim.P1UserInOutChatV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1ChatDisbandV1(v1:(ctx:context.Context,event:Ref<larkim.P1ChatDisbandV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1GroupSettingUpdatedV1(v1:(ctx:context.Context,event:Ref<larkim.P1GroupSettingUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1AppOpenV6(v1:(ctx:context.Context,event:Ref<larkapplication.P1AppOpenV6>)=>void/*error*/):Ref<EventDispatcher>
			onP1P2PChatCreatedV1(v1:(ctx:context.Context,event:Ref<larkim.P1P2PChatCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1ThirdPartyMeetingRoomChangedV1(v1:(ctx:context.Context,event:Ref<larkmeeting_room.P1ThirdPartyMeetingRoomChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP1LeaveApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1LeaveApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onP1WorkApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1WorkApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onP1ShiftApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1ShiftApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onP1RemedyApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1RemedyApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onP1TripApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1TripApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onP1OutApprovalV4(v1:(ctx:context.Context,event:Ref<larkapproval.P1OutApprovalV4>)=>void/*error*/):Ref<EventDispatcher>
			onCustomizedEvent(eventType:string,v1:(ctx:context.Context,event:Ref<larkevent.EventReq>)=>void/*error*/):Ref<EventDispatcher>
			onAppTicketEvent(v1:(ctx:context.Context,event:Ref<AppTicketEvent>)=>void/*error*/):Ref<EventDispatcher>
			onP2NotificationApproveV1(v1:(ctx:context.Context,event:Ref<larkhelpdesk.P2NotificationApproveV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TicketCreatedV1(v1:(ctx:context.Context,event:Ref<larkhelpdesk.P2TicketCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TicketUpdatedV1(v1:(ctx:context.Context,event:Ref<larkhelpdesk.P2TicketUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TicketMessageCreatedV1(v1:(ctx:context.Context,event:Ref<larkhelpdesk.P2TicketMessageCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationDeletedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2ApplicationDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ApplicationStageChangedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2ApplicationStageChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EcoAccountCreatedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EcoAccountCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EcoBackgroundCheckCanceledV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EcoBackgroundCheckCanceledV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EcoBackgroundCheckCreatedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EcoBackgroundCheckCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EcoExamCreatedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EcoExamCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EhrImportTaskImportedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EhrImportTaskImportedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2EhrImportTaskForInternshipOfferImportedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2EhrImportTaskForInternshipOfferImportedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2OfferStatusChangedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2OfferStatusChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ReferralAccountAssetsUpdateV1(v1:(ctx:context.Context,event:Ref<larkhire.P2ReferralAccountAssetsUpdateV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TalentDeletedV1(v1:(ctx:context.Context,event:Ref<larkhire.P2TalentDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatDisbandedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatDisbandedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatUpdatedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatAccessEventBotP2pChatEnteredV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatAccessEventBotP2pChatEnteredV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatMemberBotAddedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatMemberBotAddedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatMemberBotDeletedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatMemberBotDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatMemberUserAddedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatMemberUserAddedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatMemberUserDeletedV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatMemberUserDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ChatMemberUserWithdrawnV1(v1:(ctx:context.Context,event:Ref<larkim.P2ChatMemberUserWithdrawnV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MessageReadV1(v1:(ctx:context.Context,event:Ref<larkim.P2MessageReadV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MessageRecalledV1(v1:(ctx:context.Context,event:Ref<larkim.P2MessageRecalledV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MessageReceiveV1(v1:(ctx:context.Context,event:Ref<larkim.P2MessageReceiveV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MessageReactionCreatedV1(v1:(ctx:context.Context,event:Ref<larkim.P2MessageReactionCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MessageReactionDeletedV1(v1:(ctx:context.Context,event:Ref<larkim.P2MessageReactionDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRoomCreatedV1(v1:(ctx:context.Context,event:Ref<larkmeeting_room.P2MeetingRoomCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRoomDeletedV1(v1:(ctx:context.Context,event:Ref<larkmeeting_room.P2MeetingRoomDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRoomStatusChangedV1(v1:(ctx:context.Context,event:Ref<larkmeeting_room.P2MeetingRoomStatusChangedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRoomUpdatedV1(v1:(ctx:context.Context,event:Ref<larkmeeting_room.P2MeetingRoomUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TaskUpdateTenantV1(v1:(ctx:context.Context,event:Ref<larktask.P2TaskUpdateTenantV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TaskUpdatedV1(v1:(ctx:context.Context,event:Ref<larktask.P2TaskUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2TaskCommentUpdatedV1(v1:(ctx:context.Context,event:Ref<larktask.P2TaskCommentUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingAllMeetingEndedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingAllMeetingEndedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingAllMeetingStartedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingAllMeetingStartedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingJoinMeetingV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingJoinMeetingV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingLeaveMeetingV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingLeaveMeetingV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingEndedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingEndedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingStartedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingStartedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRecordingEndedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingRecordingEndedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRecordingReadyV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingRecordingReadyV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingRecordingStartedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingRecordingStartedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingShareEndedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingShareEndedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2MeetingShareStartedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2MeetingShareStartedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2ReserveConfigUpdatedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2ReserveConfigUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomCreatedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomDeletedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomUpdatedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomLevelCreatedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomLevelCreatedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomLevelDeletedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomLevelDeletedV1>)=>void/*error*/):Ref<EventDispatcher>
			onP2RoomLevelUpdatedV1(v1:(ctx:context.Context,event:Ref<larkvc.P2RoomLevelUpdatedV1>)=>void/*error*/):Ref<EventDispatcher>
	}
	export function newEventDispatcher(verificationToken:string,eventEncryptKey:string):Ref<EventDispatcher>

	export interface NotFoundEventHandlerErr extends Struct<NotFoundEventHandlerErr>,Error,GoError{

			error():string
	}
	export function emptyAppTicketEvent():AppTicketEvent
	export function emptyRefAppTicketEvent():Ref<AppTicketEvent>
	export function refOfAppTicketEvent(x:AppTicketEvent,v:Ref<AppTicketEvent>)
	export function unRefAppTicketEvent(v:Ref<AppTicketEvent>):AppTicketEvent
	export function emptyCustomAppTicketEventHandler():CustomAppTicketEventHandler
	export function emptyRefCustomAppTicketEventHandler():Ref<CustomAppTicketEventHandler>
	export function refOfCustomAppTicketEventHandler(x:CustomAppTicketEventHandler,v:Ref<CustomAppTicketEventHandler>)
	export function unRefCustomAppTicketEventHandler(v:Ref<CustomAppTicketEventHandler>):CustomAppTicketEventHandler
}