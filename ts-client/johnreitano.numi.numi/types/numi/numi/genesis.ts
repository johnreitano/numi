/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Params } from "./params";
import { User } from "./user";
import { UserAccountAddress } from "./user_account_address";

export const protobufPackage = "johnreitano.numi.numi";

/** GenesisState defines the numi module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  userList: User[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  userAccountAddressList: UserAccountAddress[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, userList: [], userAccountAddressList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.userList) {
      User.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.userAccountAddressList) {
      UserAccountAddress.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.userList.push(User.decode(reader, reader.uint32()));
          break;
        case 3:
          message.userAccountAddressList.push(UserAccountAddress.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      userList: Array.isArray(object?.userList) ? object.userList.map((e: any) => User.fromJSON(e)) : [],
      userAccountAddressList: Array.isArray(object?.userAccountAddressList)
        ? object.userAccountAddressList.map((e: any) => UserAccountAddress.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.userList) {
      obj.userList = message.userList.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.userList = [];
    }
    if (message.userAccountAddressList) {
      obj.userAccountAddressList = message.userAccountAddressList.map((e) =>
        e ? UserAccountAddress.toJSON(e) : undefined
      );
    } else {
      obj.userAccountAddressList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.userList = object.userList?.map((e) => User.fromPartial(e)) || [];
    message.userAccountAddressList = object.userAccountAddressList?.map((e) => UserAccountAddress.fromPartial(e)) || [];
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
