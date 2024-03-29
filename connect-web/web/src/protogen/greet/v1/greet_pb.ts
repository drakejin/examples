// @generated by protoc-gen-es v1.1.1 with parameter "target=ts"
// @generated from file greet/v1/greet.proto (package greet.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Any, Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message greet.v1.GreetRequest
 */
export class GreetRequest extends Message<GreetRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<GreetRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "greet.v1.GreetRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GreetRequest {
    return new GreetRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GreetRequest {
    return new GreetRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GreetRequest {
    return new GreetRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GreetRequest | PlainMessage<GreetRequest> | undefined, b: GreetRequest | PlainMessage<GreetRequest> | undefined): boolean {
    return proto3.util.equals(GreetRequest, a, b);
  }
}

/**
 * @generated from message greet.v1.GreetResponse
 */
export class GreetResponse extends Message<GreetResponse> {
  /**
   * @generated from field: string greeting = 1;
   */
  greeting = "";

  /**
   * @generated from oneof greet.v1.GreetResponse.HelloOneOf
   */
  HelloOneOf: {
    /**
     * @generated from field: greet.v1.Foo foo = 2;
     */
    value: Foo;
    case: "foo";
  } | {
    /**
     * @generated from field: greet.v1.Bar bar = 3;
     */
    value: Bar;
    case: "bar";
  } | { case: undefined; value?: undefined } = { case: undefined };

  /**
   * @generated from field: google.protobuf.Any helloAny = 4;
   */
  helloAny?: Any;

  constructor(data?: PartialMessage<GreetResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "greet.v1.GreetResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "greeting", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "foo", kind: "message", T: Foo, oneof: "HelloOneOf" },
    { no: 3, name: "bar", kind: "message", T: Bar, oneof: "HelloOneOf" },
    { no: 4, name: "helloAny", kind: "message", T: Any },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GreetResponse {
    return new GreetResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GreetResponse {
    return new GreetResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GreetResponse {
    return new GreetResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GreetResponse | PlainMessage<GreetResponse> | undefined, b: GreetResponse | PlainMessage<GreetResponse> | undefined): boolean {
    return proto3.util.equals(GreetResponse, a, b);
  }
}

/**
 * @generated from message greet.v1.Foo
 */
export class Foo extends Message<Foo> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  constructor(data?: PartialMessage<Foo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "greet.v1.Foo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Foo {
    return new Foo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Foo {
    return new Foo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Foo {
    return new Foo().fromJsonString(jsonString, options);
  }

  static equals(a: Foo | PlainMessage<Foo> | undefined, b: Foo | PlainMessage<Foo> | undefined): boolean {
    return proto3.util.equals(Foo, a, b);
  }
}

/**
 * @generated from message greet.v1.Bar
 */
export class Bar extends Message<Bar> {
  /**
   * @generated from field: string key = 1;
   */
  key = "";

  constructor(data?: PartialMessage<Bar>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "greet.v1.Bar";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Bar {
    return new Bar().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Bar {
    return new Bar().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Bar {
    return new Bar().fromJsonString(jsonString, options);
  }

  static equals(a: Bar | PlainMessage<Bar> | undefined, b: Bar | PlainMessage<Bar> | undefined): boolean {
    return proto3.util.equals(Bar, a, b);
  }
}

