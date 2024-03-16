/* eslint-disable */
import { Balance } from "../privatbank/balance";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "evgeniyscherbina.cosmosevm.privatbank";

export interface GenesisState {
  balanceList: Balance[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    for (const v of message.balanceList) {
      Balance.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.balanceList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.balanceList.push(Balance.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.balanceList = [];
    if (object.balanceList !== undefined && object.balanceList !== null) {
      for (const e of object.balanceList) {
        message.balanceList.push(Balance.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    if (message.balanceList) {
      obj.balanceList = message.balanceList.map((e) =>
        e ? Balance.toJSON(e) : undefined
      );
    } else {
      obj.balanceList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.balanceList = [];
    if (object.balanceList !== undefined && object.balanceList !== null) {
      for (const e of object.balanceList) {
        message.balanceList.push(Balance.fromPartial(e));
      }
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
