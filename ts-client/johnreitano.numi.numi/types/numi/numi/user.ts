/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "johnreitano.numi.numi";

export interface User {
  userId: string;
  firstName: string;
  lastName: string;
  countryCode: string;
  subnationalEntity: string;
  city: string;
  bio: string;
  creator: string;
  referrer: string;
  accountAddress: string;
}

function createBaseUser(): User {
  return {
    userId: "",
    firstName: "",
    lastName: "",
    countryCode: "",
    subnationalEntity: "",
    city: "",
    bio: "",
    creator: "",
    referrer: "",
    accountAddress: "",
  };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userId !== "") {
      writer.uint32(10).string(message.userId);
    }
    if (message.firstName !== "") {
      writer.uint32(18).string(message.firstName);
    }
    if (message.lastName !== "") {
      writer.uint32(26).string(message.lastName);
    }
    if (message.countryCode !== "") {
      writer.uint32(34).string(message.countryCode);
    }
    if (message.subnationalEntity !== "") {
      writer.uint32(42).string(message.subnationalEntity);
    }
    if (message.city !== "") {
      writer.uint32(50).string(message.city);
    }
    if (message.bio !== "") {
      writer.uint32(58).string(message.bio);
    }
    if (message.creator !== "") {
      writer.uint32(66).string(message.creator);
    }
    if (message.referrer !== "") {
      writer.uint32(74).string(message.referrer);
    }
    if (message.accountAddress !== "") {
      writer.uint32(82).string(message.accountAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userId = reader.string();
          break;
        case 2:
          message.firstName = reader.string();
          break;
        case 3:
          message.lastName = reader.string();
          break;
        case 4:
          message.countryCode = reader.string();
          break;
        case 5:
          message.subnationalEntity = reader.string();
          break;
        case 6:
          message.city = reader.string();
          break;
        case 7:
          message.bio = reader.string();
          break;
        case 8:
          message.creator = reader.string();
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

  fromJSON(object: any): User {
    return {
      userId: isSet(object.userId) ? String(object.userId) : "",
      firstName: isSet(object.firstName) ? String(object.firstName) : "",
      lastName: isSet(object.lastName) ? String(object.lastName) : "",
      countryCode: isSet(object.countryCode) ? String(object.countryCode) : "",
      subnationalEntity: isSet(object.subnationalEntity) ? String(object.subnationalEntity) : "",
      city: isSet(object.city) ? String(object.city) : "",
      bio: isSet(object.bio) ? String(object.bio) : "",
      creator: isSet(object.creator) ? String(object.creator) : "",
      referrer: isSet(object.referrer) ? String(object.referrer) : "",
      accountAddress: isSet(object.accountAddress) ? String(object.accountAddress) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.userId !== undefined && (obj.userId = message.userId);
    message.firstName !== undefined && (obj.firstName = message.firstName);
    message.lastName !== undefined && (obj.lastName = message.lastName);
    message.countryCode !== undefined && (obj.countryCode = message.countryCode);
    message.subnationalEntity !== undefined && (obj.subnationalEntity = message.subnationalEntity);
    message.city !== undefined && (obj.city = message.city);
    message.bio !== undefined && (obj.bio = message.bio);
    message.creator !== undefined && (obj.creator = message.creator);
    message.referrer !== undefined && (obj.referrer = message.referrer);
    message.accountAddress !== undefined && (obj.accountAddress = message.accountAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.userId = object.userId ?? "";
    message.firstName = object.firstName ?? "";
    message.lastName = object.lastName ?? "";
    message.countryCode = object.countryCode ?? "";
    message.subnationalEntity = object.subnationalEntity ?? "";
    message.city = object.city ?? "";
    message.bio = object.bio ?? "";
    message.creator = object.creator ?? "";
    message.referrer = object.referrer ?? "";
    message.accountAddress = object.accountAddress ?? "";
    return message;
  },
};

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
