// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/base-sdk-go/v3'{

	// @ts-ignore
	import * as http from 'golang/net/http'
	// @ts-ignore
	import * as time from 'golang/time'
	// @ts-ignore
	import * as larkbase from 'github.com/larksuite/base-sdk-go/v3/service/base/v1'
	// @ts-ignore
	import * as larkdrive from 'github.com/larksuite/base-sdk-go/v3/service/drive/v1'
	// @ts-ignore
	import * as context from 'golang/context'
	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/base-sdk-go/v3/core'
	// @ts-ignore
	import type {Ref,Struct,error,Alias,bool} from 'go'
	export interface Client extends Struct<Client>{

			base:Ref<larkbase.BaseService>
			drive:Ref<larkdrive.DriveService>
			post(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			Do(ctx:context.Context,apiReq:Ref<larkcore.ApiReq>,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			get(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			delete(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			put(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
			patch(ctx:context.Context,httpPath:string,body:any,accessTokeType:larkcore.AccessTokenType,...options:larkcore.RequestOptionFunc[]):Ref<larkcore.ApiResp>
	}
	export interface ClientOptionFunc extends Alias<(config:Ref<larkcore.Config>)=>void>{

	}
	export const FeishuBaseUrl:string
	export const LarkBaseUrl:string
	export function newClient(personalBaseToken:string,appToken:string,...options:ClientOptionFunc[]):Ref<Client>

	export function withHeaders(header:http.Header):ClientOptionFunc

	export function withHttpClient(httpClient:larkcore.HttpClient):ClientOptionFunc

	export function withLogLevel(logLevel:larkcore.LogLevel):ClientOptionFunc

	export function withLogReqAtDebug(printReqRespLog:bool):ClientOptionFunc

	export function withLogger(logger:larkcore.Logger):ClientOptionFunc

	export function withOpenBaseUrl(baseUrl:string):ClientOptionFunc

	export function withReqTimeout(reqTimeout:time.Duration):ClientOptionFunc

	export function withSerialization(serializable:larkcore.Serializable):ClientOptionFunc

}