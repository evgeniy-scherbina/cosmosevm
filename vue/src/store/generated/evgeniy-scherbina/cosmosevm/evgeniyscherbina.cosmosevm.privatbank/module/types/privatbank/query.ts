/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Balance } from "../privatbank/balance";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "evgeniyscherbina.cosmosevm.privatbank";

export interface QueryGetBalanceRequest {
  address: string;
}

export interface QueryGetBalanceResponse {
  balance: Balance | undefined;
}

export interface QueryAllBalanceRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBalanceResponse {
  balance: Balance[];
  pagination: PageResponse | undefined;
}

const baseQueryGetBalanceRequest: object = { address: "" };

export const QueryGetBalanceRequest = {
  encode(
    message: QueryGetBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBalanceRequest {
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetBalanceRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBalanceRequest>
  ): QueryGetBalanceRequest {
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetBalanceResponse: object = {};

export const QueryGetBalanceResponse = {
  encode(
    message: QueryGetBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.balance !== undefined) {
      Balance.encode(message.balance, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balance = Balance.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBalanceResponse {
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = Balance.fromJSON(object.balance);
    } else {
      message.balance = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBalanceResponse): unknown {
    const obj: any = {};
    message.balance !== undefined &&
      (obj.balance = message.balance
        ? Balance.toJSON(message.balance)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBalanceResponse>
  ): QueryGetBalanceResponse {
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = Balance.fromPartial(object.balance);
    } else {
      message.balance = undefined;
    }
    return message;
  },
};

const baseQueryAllBalanceRequest: object = {};

export const QueryAllBalanceRequest = {
  encode(
    message: QueryAllBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
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

  fromJSON(object: any): QueryAllBalanceRequest {
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBalanceRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBalanceRequest>
  ): QueryAllBalanceRequest {
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBalanceResponse: object = {};

export const QueryAllBalanceResponse = {
  encode(
    message: QueryAllBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.balance) {
      Balance.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balance.push(Balance.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllBalanceResponse {
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    if (object.balance !== undefined && object.balance !== null) {
      for (const e of object.balance) {
        message.balance.push(Balance.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBalanceResponse): unknown {
    const obj: any = {};
    if (message.balance) {
      obj.balance = message.balance.map((e) =>
        e ? Balance.toJSON(e) : undefined
      );
    } else {
      obj.balance = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBalanceResponse>
  ): QueryAllBalanceResponse {
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    if (object.balance !== undefined && object.balance !== null) {
      for (const e of object.balance) {
        message.balance.push(Balance.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

export interface Query {
  /** Queries a Balance by index. */
  Balance(request: QueryGetBalanceRequest): Promise<QueryGetBalanceResponse>;
  /** Queries a list of Balance items. */
  BalanceAll(request: QueryAllBalanceRequest): Promise<QueryAllBalanceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Balance(request: QueryGetBalanceRequest): Promise<QueryGetBalanceResponse> {
    const data = QueryGetBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "evgeniyscherbina.cosmosevm.privatbank.Query",
      "Balance",
      data
    );
    return promise.then((data) =>
      QueryGetBalanceResponse.decode(new Reader(data))
    );
  }

  BalanceAll(
    request: QueryAllBalanceRequest
  ): Promise<QueryAllBalanceResponse> {
    const data = QueryAllBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "evgeniyscherbina.cosmosevm.privatbank.Query",
      "BalanceAll",
      data
    );
    return promise.then((data) =>
      QueryAllBalanceResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
