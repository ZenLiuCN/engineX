// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/base-sdk-go/v3/core'{

	// @ts-ignore
	import * as json from 'golang/encoding/json'
	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as fmt from 'golang/fmt'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as io from 'golang/io'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import type {map,Alias,Struct,error,bool,Nothing,int,Ref,GoError} from 'go'
	export interface AccessTokenType extends Alias<string>{

	}
	export const AccessTokenTypeNone:AccessTokenType
	export const AccessTokenTypePersonal:AccessTokenType
	export interface ApiReq extends json.Token,Struct<ApiReq>{

			httpMethod:string
			apiPath:string
			body:any
			queryParams:QueryParams
			pathParams:PathParams
			supportedAccessTokenTypes:AccessTokenType[]
	}
	export interface ApiResp extends Struct<ApiResp>,json.Token,fmt.Stringer{

			statusCode:int
			header:http.Header
			rawBody:Uint8Array
			write(writer:http.ResponseWriter):void
			jsonUnmarshalBody(val:any,config:Ref<Config>)/*error*/
			requestId():string
			string():string
	}
	export interface ClientTimeoutError extends Struct<ClientTimeoutError>,Error,GoError{

			error():string
	}
	export interface CodeError extends Struct<CodeError>,Error,GoError{

			code:int
			msg:string
			err:Ref<Struct<{
			
				details:Ref<CodeErrorDetail>[]
				permissionViolations:Ref<CodeErrorPermissionViolation>[]
				fieldViolations:Ref<CodeErrorFieldViolation>[]
			}>>
			error():string
			string():string
	}
	export interface CodeErrorDetail extends Struct<CodeErrorDetail>,Error{

			key:string
			value:string
	}
	export interface CodeErrorFieldViolation extends Error,Struct<CodeErrorFieldViolation>{

			field:string
			value:string
			description:string
	}
	export interface CodeErrorPermissionViolation extends Struct<CodeErrorPermissionViolation>,Error{

			type:string
			subject:string
			description:string
	}
	export interface Config extends Struct<Config>,json.Token{

			baseUrl:string
			personalBaseToken:string
			appToken:string
			reqTimeout:time.Duration
			logLevel:LogLevel
			httpClient:HttpClient
			logger:Logger
			logReqAtDebug:bool
			header:http.Header
			serializable:Serializable
	}
	export interface DefaultSerialization extends Alias<Nothing>{

			serialize(v:any):Uint8Array
			deserialize(data:Uint8Array,v:any)/*error*/
	}
	export interface DialFailedError extends Struct<DialFailedError>,Error,GoError{

			error():string
	}
	export function file2Bytes(fileName:string):Uint8Array

	export function fileNameByHeader(header:http.Header):string

	export interface Formdata extends Struct<Formdata>,json.Token{

			addField(field:string,val:any):Ref<Formdata>
			addFile(field:string,r:io.Reader):Ref<Formdata>
	}
	export interface HttpClient{

			Do(v1:Ref<http.Request>):Ref<http.Response>
	}
	//"X-Tt-Logid"
	export const HttpHeaderKeyLogId:string
	//"X-Request-Id"
	export const HttpHeaderKeyRequestId:string
	export interface IllegalParamError extends GoError,Struct<IllegalParamError>,Error{

			error():string
	}
	export interface LogLevel extends Alias<int>{

	}
	export const LogLevelDebug:LogLevel
	export const LogLevelError:LogLevel
	export const LogLevelInfo:LogLevel
	export const LogLevelWarn:LogLevel
	export interface Logger{

			debug(v2:context.Context,...v1:any[]):void
			error(v2:context.Context,...v1:any[]):void
			info(v2:context.Context,...v1:any[]):void
			warn(v2:context.Context,...v1:any[]):void
	}
	export function newEventLogger():Logger

	export function newFormdata():Ref<Formdata>

	export function newHttpClient(config:Ref<Config>):void

	export function newLogger(config:Ref<Config>):void

	export function newSerialization(config:Ref<Config>):void

	export interface PathParams extends map<Alias<string>,Alias<string>>{

			get(key:string):string
			set(key:string,value:string):void
	}
	export function prettify(i:any):string

	export interface QueryParams extends map<Alias<string>,Array<string>>{

			get(key:string):string
			set(key:string,value:string):void
			encode():string
			add(key:string,value:string):void
	}
	export interface ReqTranslator extends Alias<Nothing>{

	}
	export function request(ctx:context.Context,req:Ref<ApiReq>,config:Ref<Config>,...options:RequestOptionFunc[]):Ref<ApiResp>

	export interface RequestOption extends Struct<RequestOption>,json.Token{

			personalToken:string
			requestId:string
			fileUpload:bool
			fileDownload:bool
			header:http.Header
	}
	export interface RequestOptionFunc extends Alias<(option:Ref<RequestOption>)=>void>{

	}
	export interface Serializable{

			deserialize(data:Uint8Array,v:any)/*error*/
			serialize(v:any):Uint8Array
	}
	export interface ServerTimeoutError extends Error,GoError,Struct<ServerTimeoutError>{

			error():string
	}
	export function withFileDownload():RequestOptionFunc

	export function withFileUpload():RequestOptionFunc

	export function withHeaders(header:http.Header):RequestOptionFunc

	export function withRequestId(requestId:string):RequestOptionFunc

	export function emptyRequestOption():RequestOption
	export function emptyRefRequestOption():Ref<RequestOption>
	export function refOfRequestOption(x:RequestOption,v:Ref<RequestOption>)
	export function unRefRequestOption(v:Ref<RequestOption>):RequestOption
	export function emptyApiResp():ApiResp
	export function emptyRefApiResp():Ref<ApiResp>
	export function refOfApiResp(x:ApiResp,v:Ref<ApiResp>)
	export function unRefApiResp(v:Ref<ApiResp>):ApiResp
	export function emptyCodeErrorPermissionViolation():CodeErrorPermissionViolation
	export function emptyRefCodeErrorPermissionViolation():Ref<CodeErrorPermissionViolation>
	export function refOfCodeErrorPermissionViolation(x:CodeErrorPermissionViolation,v:Ref<CodeErrorPermissionViolation>)
	export function unRefCodeErrorPermissionViolation(v:Ref<CodeErrorPermissionViolation>):CodeErrorPermissionViolation
	export function emptyCodeErrorFieldViolation():CodeErrorFieldViolation
	export function emptyRefCodeErrorFieldViolation():Ref<CodeErrorFieldViolation>
	export function refOfCodeErrorFieldViolation(x:CodeErrorFieldViolation,v:Ref<CodeErrorFieldViolation>)
	export function unRefCodeErrorFieldViolation(v:Ref<CodeErrorFieldViolation>):CodeErrorFieldViolation
	export function emptyConfig():Config
	export function emptyRefConfig():Ref<Config>
	export function refOfConfig(x:Config,v:Ref<Config>)
	export function unRefConfig(v:Ref<Config>):Config
	export function emptyApiReq():ApiReq
	export function emptyRefApiReq():Ref<ApiReq>
	export function refOfApiReq(x:ApiReq,v:Ref<ApiReq>)
	export function unRefApiReq(v:Ref<ApiReq>):ApiReq
	export function emptyCodeErrorDetail():CodeErrorDetail
	export function emptyRefCodeErrorDetail():Ref<CodeErrorDetail>
	export function refOfCodeErrorDetail(x:CodeErrorDetail,v:Ref<CodeErrorDetail>)
	export function unRefCodeErrorDetail(v:Ref<CodeErrorDetail>):CodeErrorDetail
}