/* eslint-disable */
import Long from "long";
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

export interface MsgCreateLottery {
  creator: string;
  index: number;
  winner: string;
}

export interface MsgCreateLotteryResponse {
}

export interface MsgUpdateLottery {
  creator: string;
  index: number;
  winner: string;
}

export interface MsgUpdateLotteryResponse {
}

export interface MsgDeleteLottery {
  creator: string;
  index: number;
}

export interface MsgDeleteLotteryResponse {
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

function createBaseMsgCreateLottery(): MsgCreateLottery {
  return { creator: "", index: 0, winner: "" };
}

export const MsgCreateLottery = {
  encode(message: MsgCreateLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    if (message.winner !== "") {
      writer.uint32(26).string(message.winner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? Number(object.index) : 0,
      winner: isSet(object.winner) ? String(object.winner) : "",
    };
  },

  toJSON(message: MsgCreateLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = Math.round(message.index));
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLottery>, I>>(object: I): MsgCreateLottery {
    const message = createBaseMsgCreateLottery();
    message.creator = object.creator ?? "";
    message.index = object.index ?? 0;
    message.winner = object.winner ?? "";
    return message;
  },
};

function createBaseMsgCreateLotteryResponse(): MsgCreateLotteryResponse {
  return {};
}

export const MsgCreateLotteryResponse = {
  encode(_: MsgCreateLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateLotteryResponse();
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

  fromJSON(_: any): MsgCreateLotteryResponse {
    return {};
  },

  toJSON(_: MsgCreateLotteryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateLotteryResponse>, I>>(_: I): MsgCreateLotteryResponse {
    const message = createBaseMsgCreateLotteryResponse();
    return message;
  },
};

function createBaseMsgUpdateLottery(): MsgUpdateLottery {
  return { creator: "", index: 0, winner: "" };
}

export const MsgUpdateLottery = {
  encode(message: MsgUpdateLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    if (message.winner !== "") {
      writer.uint32(26).string(message.winner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? Number(object.index) : 0,
      winner: isSet(object.winner) ? String(object.winner) : "",
    };
  },

  toJSON(message: MsgUpdateLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = Math.round(message.index));
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateLottery>, I>>(object: I): MsgUpdateLottery {
    const message = createBaseMsgUpdateLottery();
    message.creator = object.creator ?? "";
    message.index = object.index ?? 0;
    message.winner = object.winner ?? "";
    return message;
  },
};

function createBaseMsgUpdateLotteryResponse(): MsgUpdateLotteryResponse {
  return {};
}

export const MsgUpdateLotteryResponse = {
  encode(_: MsgUpdateLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateLotteryResponse();
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

  fromJSON(_: any): MsgUpdateLotteryResponse {
    return {};
  },

  toJSON(_: MsgUpdateLotteryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateLotteryResponse>, I>>(_: I): MsgUpdateLotteryResponse {
    const message = createBaseMsgUpdateLotteryResponse();
    return message;
  },
};

function createBaseMsgDeleteLottery(): MsgDeleteLottery {
  return { creator: "", index: 0 };
}

export const MsgDeleteLottery = {
  encode(message: MsgDeleteLottery, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.index !== 0) {
      writer.uint32(16).uint64(message.index);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteLottery {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteLottery();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.index = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteLottery {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      index: isSet(object.index) ? Number(object.index) : 0,
    };
  },

  toJSON(message: MsgDeleteLottery): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.index !== undefined && (obj.index = Math.round(message.index));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteLottery>, I>>(object: I): MsgDeleteLottery {
    const message = createBaseMsgDeleteLottery();
    message.creator = object.creator ?? "";
    message.index = object.index ?? 0;
    return message;
  },
};

function createBaseMsgDeleteLotteryResponse(): MsgDeleteLotteryResponse {
  return {};
}

export const MsgDeleteLotteryResponse = {
  encode(_: MsgDeleteLotteryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteLotteryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteLotteryResponse();
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

  fromJSON(_: any): MsgDeleteLotteryResponse {
    return {};
  },

  toJSON(_: MsgDeleteLotteryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteLotteryResponse>, I>>(_: I): MsgDeleteLotteryResponse {
    const message = createBaseMsgDeleteLotteryResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AddBet(request: MsgAddBet): Promise<MsgAddBetResponse>;
  RevealBet(request: MsgRevealBet): Promise<MsgRevealBetResponse>;
  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse>;
  UpdateLottery(request: MsgUpdateLottery): Promise<MsgUpdateLotteryResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteLottery(request: MsgDeleteLottery): Promise<MsgDeleteLotteryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AddBet = this.AddBet.bind(this);
    this.RevealBet = this.RevealBet.bind(this);
    this.CreateLottery = this.CreateLottery.bind(this);
    this.UpdateLottery = this.UpdateLottery.bind(this);
    this.DeleteLottery = this.DeleteLottery.bind(this);
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

  CreateLottery(request: MsgCreateLottery): Promise<MsgCreateLotteryResponse> {
    const data = MsgCreateLottery.encode(request).finish();
    const promise = this.rpc.request("lottery.lottery.Msg", "CreateLottery", data);
    return promise.then((data) => MsgCreateLotteryResponse.decode(new _m0.Reader(data)));
  }

  UpdateLottery(request: MsgUpdateLottery): Promise<MsgUpdateLotteryResponse> {
    const data = MsgUpdateLottery.encode(request).finish();
    const promise = this.rpc.request("lottery.lottery.Msg", "UpdateLottery", data);
    return promise.then((data) => MsgUpdateLotteryResponse.decode(new _m0.Reader(data)));
  }

  DeleteLottery(request: MsgDeleteLottery): Promise<MsgDeleteLotteryResponse> {
    const data = MsgDeleteLottery.encode(request).finish();
    const promise = this.rpc.request("lottery.lottery.Msg", "DeleteLottery", data);
    return promise.then((data) => MsgDeleteLotteryResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
