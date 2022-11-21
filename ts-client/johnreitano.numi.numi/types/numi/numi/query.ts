/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { User } from "./user";
import { UserAccountAddress } from "./user_account_address";

export const protobufPackage = "johnreitano.numi.numi";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetUserRequest {
  userId: string;
}

export interface QueryGetUserResponse {
  user: User | undefined;
}

export interface QueryAllUserRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserResponse {
  user: User[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserAccountAddressRequest {
  accountAddress: string;
}

export interface QueryGetUserAccountAddressResponse {
  userAccountAddress: UserAccountAddress | undefined;
}

export interface QueryAllUserAccountAddressRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserAccountAddressResponse {
  userAccountAddress: UserAccountAddress[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserRequest(): QueryGetUserRequest {
  return { userId: "" };
}

export const QueryGetUserRequest = {
  encode(message: QueryGetUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserRequest {
    return { userId: isSet(object.userId) ? String(object.userId) : "" };
  },

  toJSON(message: QueryGetUserRequest): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserRequest>, I>>(object: I): QueryGetUserRequest {
    const message = createBaseQueryGetUserRequest();
    message.userId = object.userId ?? "";
    return message;
  },
};

function createBaseQueryGetUserResponse(): QueryGetUserResponse {
  return { user: undefined };
}

export const QueryGetUserResponse = {
  encode(message: QueryGetUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.user !== undefined) {
      User.encode(message.user, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user = User.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserResponse {
    return { user: isSet(object.user) ? User.fromJSON(object.user) : undefined };
  },

  toJSON(message: QueryGetUserResponse): unknown {
    const obj: any = {};
    message.user !== undefined && (obj.user = message.user ? User.toJSON(message.user) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserResponse>, I>>(object: I): QueryGetUserResponse {
    const message = createBaseQueryGetUserResponse();
    message.user = (object.user !== undefined && object.user !== null) ? User.fromPartial(object.user) : undefined;
    return message;
  },
};

function createBaseQueryAllUserRequest(): QueryAllUserRequest {
  return { pagination: undefined };
}

export const QueryAllUserRequest = {
  encode(message: QueryAllUserRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserRequest>, I>>(object: I): QueryAllUserRequest {
    const message = createBaseQueryAllUserRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserResponse(): QueryAllUserResponse {
  return { user: [], pagination: undefined };
}

export const QueryAllUserResponse = {
  encode(message: QueryAllUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.user) {
      User.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.user.push(User.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserResponse {
    return {
      user: Array.isArray(object?.user) ? object.user.map((e: any) => User.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserResponse): unknown {
    const obj: any = {};
    if (message.user) {
      obj.user = message.user.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.user = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserResponse>, I>>(object: I): QueryAllUserResponse {
    const message = createBaseQueryAllUserResponse();
    message.user = object.user?.map((e) => User.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserAccountAddressRequest(): QueryGetUserAccountAddressRequest {
  return { accountAddress: "" };
}

export const QueryGetUserAccountAddressRequest = {
  encode(message: QueryGetUserAccountAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountAddress !== "") {
      writer.uint32(10).string(message.accountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserAccountAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserAccountAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserAccountAddressRequest {
    return { accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "" };
  },

  toJSON(message: QueryGetUserAccountAddressRequest): unknown {
    const obj: any = {};
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserAccountAddressRequest>, I>>(
    object: I,
  ): QueryGetUserAccountAddressRequest {
    const message = createBaseQueryGetUserAccountAddressRequest();
    message.accountAddress = object.accountAddress ?? "";
    return message;
  },
};

function createBaseQueryGetUserAccountAddressResponse(): QueryGetUserAccountAddressResponse {
  return { userAccountAddress: undefined };
}

export const QueryGetUserAccountAddressResponse = {
  encode(message: QueryGetUserAccountAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userAccountAddress !== undefined) {
      UserAccountAddress.encode(message.userAccountAddress, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserAccountAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserAccountAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAccountAddress = UserAccountAddress.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserAccountAddressResponse {
    return {
      userAccountAddress: isSet(object.userAccountAddress)
        ? UserAccountAddress.fromJSON(object.userAccountAddress)
        : undefined,
    };
  },

  toJSON(message: QueryGetUserAccountAddressResponse): unknown {
    const obj: any = {};
    message.userAccountAddress !== undefined && (obj.userAccountAddress = message.userAccountAddress
      ? UserAccountAddress.toJSON(message.userAccountAddress)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserAccountAddressResponse>, I>>(
    object: I,
  ): QueryGetUserAccountAddressResponse {
    const message = createBaseQueryGetUserAccountAddressResponse();
    message.userAccountAddress = (object.userAccountAddress !== undefined && object.userAccountAddress !== null)
      ? UserAccountAddress.fromPartial(object.userAccountAddress)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserAccountAddressRequest(): QueryAllUserAccountAddressRequest {
  return { pagination: undefined };
}

export const QueryAllUserAccountAddressRequest = {
  encode(message: QueryAllUserAccountAddressRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserAccountAddressRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserAccountAddressRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserAccountAddressRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserAccountAddressRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserAccountAddressRequest>, I>>(
    object: I,
  ): QueryAllUserAccountAddressRequest {
    const message = createBaseQueryAllUserAccountAddressRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserAccountAddressResponse(): QueryAllUserAccountAddressResponse {
  return { userAccountAddress: [], pagination: undefined };
}

export const QueryAllUserAccountAddressResponse = {
  encode(message: QueryAllUserAccountAddressResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.userAccountAddress) {
      UserAccountAddress.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserAccountAddressResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserAccountAddressResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAccountAddress.push(UserAccountAddress.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllUserAccountAddressResponse {
    return {
      userAccountAddress: Array.isArray(object?.userAccountAddress)
        ? object.userAccountAddress.map((e: any) => UserAccountAddress.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserAccountAddressResponse): unknown {
    const obj: any = {};
    if (message.userAccountAddress) {
      obj.userAccountAddress = message.userAccountAddress.map((e) => e ? UserAccountAddress.toJSON(e) : undefined);
    } else {
      obj.userAccountAddress = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserAccountAddressResponse>, I>>(
    object: I,
  ): QueryAllUserAccountAddressResponse {
    const message = createBaseQueryAllUserAccountAddressResponse();
    message.userAccountAddress = object.userAccountAddress?.map((e) => UserAccountAddress.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a User by index. */
  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse>;
  /** Queries a list of User items. */
  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse>;
  /** Queries a UserAccountAddress by index. */
  UserAccountAddress(request: QueryGetUserAccountAddressRequest): Promise<QueryGetUserAccountAddressResponse>;
  /** Queries a list of UserAccountAddress items. */
  UserAccountAddressAll(request: QueryAllUserAccountAddressRequest): Promise<QueryAllUserAccountAddressResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.User = this.User.bind(this);
    this.UserAll = this.UserAll.bind(this);
    this.UserAccountAddress = this.UserAccountAddress.bind(this);
    this.UserAccountAddressAll = this.UserAccountAddressAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  User(request: QueryGetUserRequest): Promise<QueryGetUserResponse> {
    const data = QueryGetUserRequest.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Query", "User", data);
    return promise.then((data) => QueryGetUserResponse.decode(new _m0.Reader(data)));
  }

  UserAll(request: QueryAllUserRequest): Promise<QueryAllUserResponse> {
    const data = QueryAllUserRequest.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Query", "UserAll", data);
    return promise.then((data) => QueryAllUserResponse.decode(new _m0.Reader(data)));
  }

  UserAccountAddress(request: QueryGetUserAccountAddressRequest): Promise<QueryGetUserAccountAddressResponse> {
    const data = QueryGetUserAccountAddressRequest.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Query", "UserAccountAddress", data);
    return promise.then((data) => QueryGetUserAccountAddressResponse.decode(new _m0.Reader(data)));
  }

  UserAccountAddressAll(request: QueryAllUserAccountAddressRequest): Promise<QueryAllUserAccountAddressResponse> {
    const data = QueryAllUserAccountAddressRequest.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Query", "UserAccountAddressAll", data);
    return promise.then((data) => QueryAllUserAccountAddressResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
