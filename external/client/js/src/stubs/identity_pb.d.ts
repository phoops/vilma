// package: com.phoops.vilma
// file: identity.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Identity extends jspb.Message { 
    getId(): string;
    setId(value: string): Identity;
    getFirstName(): string;
    setFirstName(value: string): Identity;
    getLastName(): string;
    setLastName(value: string): Identity;
    getEmail(): string;
    setEmail(value: string): Identity;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Identity.AsObject;
    static toObject(includeInstance: boolean, msg: Identity): Identity.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Identity, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Identity;
    static deserializeBinaryFromReader(message: Identity, reader: jspb.BinaryReader): Identity;
}

export namespace Identity {
    export type AsObject = {
        id: string,
        firstName: string,
        lastName: string,
        email: string,
    }
}

export class GetIdentityByIdRequest extends jspb.Message { 
    getIdentityId(): string;
    setIdentityId(value: string): GetIdentityByIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetIdentityByIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetIdentityByIdRequest): GetIdentityByIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetIdentityByIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetIdentityByIdRequest;
    static deserializeBinaryFromReader(message: GetIdentityByIdRequest, reader: jspb.BinaryReader): GetIdentityByIdRequest;
}

export namespace GetIdentityByIdRequest {
    export type AsObject = {
        identityId: string,
    }
}
