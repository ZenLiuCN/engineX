declare module 'golang.design/x/clipboard'{
    export function init():void

    /**
     * 0 : format text
     * 1: format image/png
     */
    export type  Format= 0|1

    export function read(format:Format) :Uint8Array
    export function readText() :string
    export function readImage() :Uint8Array
    export function write(format:Format,data:Uint8Array) :Promise<void>
    export function writeText(data:string) :Promise<void>
    export function writeImage(data:Uint8Array) :Promise<void>
}