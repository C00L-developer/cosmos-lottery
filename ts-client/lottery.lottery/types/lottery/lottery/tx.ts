/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "lottery.lottery";

export interface MsgAddBet {
  creator: string;
  amount: string;
  suffixHash: string;
}

export interface MsgAddBetResponse {
}

export interface MsgRevealBet {
  creator: string;
  suffix: string;
}

export interface MsgRevealBetResponse {
}

function createBaseMsgAddBet(): MsgAddBet {
  return { creator: "", amount: "", suffixHash: "" };
}

export const MsgAddBet = {
  encode(message: MsgAddBet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.suffixHash !== "") {
      writer.uint32(26).string(message.suffixHash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.suffixHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddBet {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      suffixHash: isSet(object.suffixHash) ? String(object.suffixHash) : "",
    };
  },

  toJSON(message: MsgAddBet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.suffixHash !== undefined && (obj.suffixHash = message.suffixHash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBet>, I>>(object: I): MsgAddBet {
    const message = createBaseMsgAddBet();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.suffixHash = object.suffixHash ?? "";
    return message;
  },
};

function createBaseMsgAddBetResponse(): MsgAddBetResponse {
  return {};
}

export const MsgAddBetResponse = {
  encode(_: MsgAddBetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddBetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddBetResponse();
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

  fromJSON(_: any): MsgAddBetResponse {
    return {};
  },

  toJSON(_: MsgAddBetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddBetResponse>, I>>(_: I): MsgAddBetResponse {
    const message = createBaseMsgAddBetResponse();
    return message;
  },
};

function createBaseMsgRevealBet(): MsgRevealBet {
  return { creator: "", suffix: "" };
}

export const MsgRevealBet = {
  encode(message: MsgRevealBet, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.suffix !== "") {
      writer.uint32(18).string(message.suffix);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevealBet {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevealBet();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.suffix = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRevealBet {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      suffix: isSet(object.suffix) ? String(object.suffix) : "",
    };
  },

  toJSON(message: MsgRevealBet): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.suffix !== undefined && (obj.suffix = message.suffix);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevealBet>, I>>(object: I): MsgRevealBet {
    const message = createBaseMsgRevealBet();
    message.creator = object.creator ?? "";
    message.suffix = object.suffix ?? "";
    return message;
  },
};

function createBaseMsgRevealBetResponse(): MsgRevealBetResponse {
  return {};
}

export const MsgRevealBetResponse = {
  encode(_: MsgRevealBetResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRevealBetResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRevealBetResponse();
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

  fromJSON(_: any): MsgRevealBetResponse {
    return {};
  },

  toJSON(_: MsgRevealBetResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRevealBetResponse>, I>>(_: I): MsgRevealBetResponse {
    const message = createBaseMsgRevealBetResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddBet(request: MsgAddBet): Promise<MsgAddBetResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RevealBet(request: MsgRevealBet): Promise<MsgRevealBetResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddBet = this.AddBet.bind(this);
    this.RevealBet = this.RevealBet.bind(this);
  }
  AddBet(request: MsgAddBet): Promise<MsgAddBetResponse> {
    const data = MsgAddBet.encode(request).finish();
    const promise = this.rpc.request("lottery.lottery.Msg", "AddBet", data);
    return promise.then((data) => MsgAddBetResponse.decode(new _m0.Reader(data)));
  }

  RevealBet(request: MsgRevealBet): Promise<MsgRevealBetResponse> {
    const data = MsgRevealBet.encode(request).finish();
    const promise = this.rpc.request("lottery.lottery.Msg", "RevealBet", data);
    return promise.then((data) => MsgRevealBetResponse.decode(new _m0.Reader(data)));
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
