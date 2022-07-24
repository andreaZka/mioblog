/* eslint-disable */
import { Params } from "../mioblog/params";
import { Post } from "../mioblog/post";
import { NextPost } from "../mioblog/next_post";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "andreazka.mioblog.mioblog";

/** GenesisState defines the mioblog module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  postList: Post[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  nextPost: NextPost | undefined;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.postList) {
      Post.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.nextPost !== undefined) {
      NextPost.encode(message.nextPost, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.postList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.postList.push(Post.decode(reader, reader.uint32()));
          break;
        case 3:
          message.nextPost = NextPost.decode(reader, reader.uint32());
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
    message.postList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.postList !== undefined && object.postList !== null) {
      for (const e of object.postList) {
        message.postList.push(Post.fromJSON(e));
      }
    }
    if (object.nextPost !== undefined && object.nextPost !== null) {
      message.nextPost = NextPost.fromJSON(object.nextPost);
    } else {
      message.nextPost = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.postList) {
      obj.postList = message.postList.map((e) =>
        e ? Post.toJSON(e) : undefined
      );
    } else {
      obj.postList = [];
    }
    message.nextPost !== undefined &&
      (obj.nextPost = message.nextPost
        ? NextPost.toJSON(message.nextPost)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.postList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.postList !== undefined && object.postList !== null) {
      for (const e of object.postList) {
        message.postList.push(Post.fromPartial(e));
      }
    }
    if (object.nextPost !== undefined && object.nextPost !== null) {
      message.nextPost = NextPost.fromPartial(object.nextPost);
    } else {
      message.nextPost = undefined;
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
