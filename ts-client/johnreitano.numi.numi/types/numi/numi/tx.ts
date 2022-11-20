/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "johnreitano.numi.numi";

export interface MsgCreateAndVerifyUser {
  creator: string;
  userId: string;
  firstName: string;
  lastName: string;
  countryCode: string;
  subnationalEntity: string;
  city: string;
  bio: string;
  referrer: string;
  accountAddress: string;
}

export interface MsgCreateAndVerifyUserResponse {
}

function createBaseMsgCreateAndVerifyUser(): MsgCreateAndVerifyUser {
  return {
    creator: "",
    userId: "",
    firstName: "",
    lastName: "",
    countryCode: "",
    subnationalEntity: "",
    city: "",
    bio: "",
    referrer: "",
    accountAddress: "",
  };
}

export const MsgCreateAndVerifyUser = {
  encode(message: MsgCreateAndVerifyUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.userId !== "") {
      writer.uint32(18).string(message.userId);
    }
    if (message.firstName !== "") {
      writer.uint32(26).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(34).string(message.lastName);
    }
    if (message.countryCode !== "") {
      writer.uint32(42).string(message.countryCode);
    }
    if (message.subnationalEntity !== "") {
      writer.uint32(50).string(message.subnationalEntity);
    }
    if (message.city !== "") {
      writer.uint32(58).string(message.city);
    }
    if (message.bio !== "") {
      writer.uint32(66).string(message.bio);
    }
    if (message.referrer !== "") {
      writer.uint32(74).string(message.referrer);
    }
    if (message.accountAddress !== "") {
      writer.uint32(82).string(message.accountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAndVerifyUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAndVerifyUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.userId = reader.string();
          break;
        case 3:
          message.firstName = reader.string();
          break;
        case 4:
          message.lastName = reader.string();
          break;
        case 5:
          message.countryCode = reader.string();
          break;
        case 6:
          message.subnationalEntity = reader.string();
          break;
        case 7:
          message.city = reader.string();
          break;
        case 8:
          message.bio = reader.string();
          break;
        case 9:
          message.referrer = reader.string();
          break;
        case 10:
          message.accountAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAndVerifyUser {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      userId: isSet(object.userId) ? String(object.userId) : "",
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      countryCode: isSet(object.countryCode) ? String(object.countryCode) : "",
      subnationalEntity: isSet(object.subnationalEntity) ? String(object.subnationalEntity) : "",
      city: isSet(object.city) ? String(object.city) : "",
      bio: isSet(object.bio) ? String(object.bio) : "",
      referrer: isSet(object.referrer) ? String(object.referrer) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
    };
  },

  toJSON(message: MsgCreateAndVerifyUser): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.userId !== undefined && (obj.userId = message.userId);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.countryCode !== undefined && (obj.countryCode = message.countryCode);
    message.subnationalEntity !== undefined && (obj.subnationalEntity = message.subnationalEntity);
    message.city !== undefined && (obj.city = message.city);
    message.bio !== undefined && (obj.bio = message.bio);
    message.referrer !== undefined && (obj.referrer = message.referrer);
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAndVerifyUser>, I>>(object: I): MsgCreateAndVerifyUser {
    const message = createBaseMsgCreateAndVerifyUser();
    message.creator = object.creator ?? "";
    message.userId = object.userId ?? "";
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.countryCode = object.countryCode ?? "";
    message.subnationalEntity = object.subnationalEntity ?? "";
    message.city = object.city ?? "";
    message.bio = object.bio ?? "";
    message.referrer = object.referrer ?? "";
    message.accountAddress = object.accountAddress ?? "";
    return message;
  },
};

function createBaseMsgCreateAndVerifyUserResponse(): MsgCreateAndVerifyUserResponse {
  return {};
}

export const MsgCreateAndVerifyUserResponse = {
  encode(_: MsgCreateAndVerifyUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAndVerifyUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAndVerifyUserResponse();
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

  fromJSON(_: any): MsgCreateAndVerifyUserResponse {
    return {};
  },

  toJSON(_: MsgCreateAndVerifyUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAndVerifyUserResponse>, I>>(_: I): MsgCreateAndVerifyUserResponse {
    const message = createBaseMsgCreateAndVerifyUserResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAndVerifyUser(request: MsgCreateAndVerifyUser): Promise<MsgCreateAndVerifyUserResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateAndVerifyUser = this.CreateAndVerifyUser.bind(this);
  }
  CreateAndVerifyUser(request: MsgCreateAndVerifyUser): Promise<MsgCreateAndVerifyUserResponse> {
    const data = MsgCreateAndVerifyUser.encode(request).finish();
    const promise = this.rpc.request("johnreitano.numi.numi.Msg", "CreateAndVerifyUser", data);
    return promise.then((data) => MsgCreateAndVerifyUserResponse.decode(new _m0.Reader(data)));
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
