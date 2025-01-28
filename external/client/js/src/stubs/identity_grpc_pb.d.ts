// package: com.phoops.vilma
// file: identity.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as identity_pb from "./identity_pb";

interface IVilmaIdentityPoolService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getIdentityByIdentityId: IVilmaIdentityPoolService_IGetIdentityByIdentityId;
}

interface IVilmaIdentityPoolService_IGetIdentityByIdentityId extends grpc.MethodDefinition<identity_pb.GetIdentityByIdRequest, identity_pb.Identity> {
    path: "/com.phoops.vilma.VilmaIdentityPool/GetIdentityByIdentityId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<identity_pb.GetIdentityByIdRequest>;
    requestDeserialize: grpc.deserialize<identity_pb.GetIdentityByIdRequest>;
    responseSerialize: grpc.serialize<identity_pb.Identity>;
    responseDeserialize: grpc.deserialize<identity_pb.Identity>;
}

export const VilmaIdentityPoolService: IVilmaIdentityPoolService;

export interface IVilmaIdentityPoolServer extends grpc.UntypedServiceImplementation {
    getIdentityByIdentityId: grpc.handleUnaryCall<identity_pb.GetIdentityByIdRequest, identity_pb.Identity>;
}

export interface IVilmaIdentityPoolClient {
    getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
    getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
    getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
}

export class VilmaIdentityPoolClient extends grpc.Client implements IVilmaIdentityPoolClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
    public getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
    public getIdentityByIdentityId(request: identity_pb.GetIdentityByIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: identity_pb.Identity) => void): grpc.ClientUnaryCall;
}
