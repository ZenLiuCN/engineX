// noinspection JSUnusedGlobalSymbols,SpellCheckingInspection
// Code generated by define_gene; DO NOT EDIT.
declare module 'github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg'{

	// @ts-ignore
	import * as larkcore from 'github.com/larksuite/oapi-sdk-go/v3/core'
	// @ts-ignore
	import * as larkgray_test_open_sg from 'github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1'
	// @ts-ignore
	import type {Struct,Ref} from 'go'
	export function newService(config:Ref<larkcore.Config>):Ref<Service>

	export interface Service extends Struct<Service>{

			V1:Ref<larkgray_test_open_sg.V1>
	}
}