// package: drakejin.examples.http
// file: proto/http.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Request extends jspb.Message { 

    hasReqA(): boolean;
    clearReqA(): void;
    getReqA(): RequestA | undefined;
    setReqA(value?: RequestA): void;


    hasReqB(): boolean;
    clearReqB(): void;
    getReqB(): RequestB | undefined;
    setReqB(value?: RequestB): void;


    hasReqC(): boolean;
    clearReqC(): void;
    getReqC(): RequestC | undefined;
    setReqC(value?: RequestC): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
    export type AsObject = {
        reqA?: RequestA.AsObject,
        reqB?: RequestB.AsObject,
        reqC?: RequestC.AsObject,
    }

    export enum Name {
    UNKNOWN = 0,
    REQUEST_A = 1,
    REQUEST_B = 2,
    REQUEST_C = 3,
    REQUEST_D = 4,
    }

}

export class Response extends jspb.Message { 

    hasResA(): boolean;
    clearResA(): void;
    getResA(): ResponseA | undefined;
    setResA(value?: ResponseA): void;


    hasResB(): boolean;
    clearResB(): void;
    getResB(): ResponseB | undefined;
    setResB(value?: ResponseB): void;


    hasResC(): boolean;
    clearResC(): void;
    getResC(): ResponseC | undefined;
    setResC(value?: ResponseC): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
    export type AsObject = {
        resA?: ResponseA.AsObject,
        resB?: ResponseB.AsObject,
        resC?: ResponseC.AsObject,
    }
}

export class RequestA extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestA.AsObject;
    static toObject(includeInstance: boolean, msg: RequestA): RequestA.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestA, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestA;
    static deserializeBinaryFromReader(message: RequestA, reader: jspb.BinaryReader): RequestA;
}

export namespace RequestA {
    export type AsObject = {
        query: string,
    }
}

export class ResponseA extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResponseA.AsObject;
    static toObject(includeInstance: boolean, msg: ResponseA): ResponseA.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResponseA, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResponseA;
    static deserializeBinaryFromReader(message: ResponseA, reader: jspb.BinaryReader): ResponseA;
}

export namespace ResponseA {
    export type AsObject = {
        query: string,
    }
}

export class RequestB extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestB.AsObject;
    static toObject(includeInstance: boolean, msg: RequestB): RequestB.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestB, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestB;
    static deserializeBinaryFromReader(message: RequestB, reader: jspb.BinaryReader): RequestB;
}

export namespace RequestB {
    export type AsObject = {
        query: string,
    }
}

export class ResponseB extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResponseB.AsObject;
    static toObject(includeInstance: boolean, msg: ResponseB): ResponseB.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResponseB, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResponseB;
    static deserializeBinaryFromReader(message: ResponseB, reader: jspb.BinaryReader): ResponseB;
}

export namespace ResponseB {
    export type AsObject = {
        query: string,
    }
}

export class RequestC extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RequestC.AsObject;
    static toObject(includeInstance: boolean, msg: RequestC): RequestC.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RequestC, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RequestC;
    static deserializeBinaryFromReader(message: RequestC, reader: jspb.BinaryReader): RequestC;
}

export namespace RequestC {
    export type AsObject = {
        query: string,
    }
}

export class ResponseC extends jspb.Message { 
    getQuery(): string;
    setQuery(value: string): void;


    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ResponseC.AsObject;
    static toObject(includeInstance: boolean, msg: ResponseC): ResponseC.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ResponseC, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ResponseC;
    static deserializeBinaryFromReader(message: ResponseC, reader: jspb.BinaryReader): ResponseC;
}

export namespace ResponseC {
    export type AsObject = {
        query: string,
    }
}
