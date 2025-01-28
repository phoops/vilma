// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var identity_pb = require('./identity_pb.js');

function serialize_com_phoops_vilma_GetIdentityByIdRequest(arg) {
  if (!(arg instanceof identity_pb.GetIdentityByIdRequest)) {
    throw new Error('Expected argument of type com.phoops.vilma.GetIdentityByIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_phoops_vilma_GetIdentityByIdRequest(buffer_arg) {
  return identity_pb.GetIdentityByIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_com_phoops_vilma_Identity(arg) {
  if (!(arg instanceof identity_pb.Identity)) {
    throw new Error('Expected argument of type com.phoops.vilma.Identity');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_com_phoops_vilma_Identity(buffer_arg) {
  return identity_pb.Identity.deserializeBinary(new Uint8Array(buffer_arg));
}


var VilmaIdentityPoolService = exports.VilmaIdentityPoolService = {
  getIdentityByIdentityId: {
    path: '/com.phoops.vilma.VilmaIdentityPool/GetIdentityByIdentityId',
    requestStream: false,
    responseStream: false,
    requestType: identity_pb.GetIdentityByIdRequest,
    responseType: identity_pb.Identity,
    requestSerialize: serialize_com_phoops_vilma_GetIdentityByIdRequest,
    requestDeserialize: deserialize_com_phoops_vilma_GetIdentityByIdRequest,
    responseSerialize: serialize_com_phoops_vilma_Identity,
    responseDeserialize: deserialize_com_phoops_vilma_Identity,
  },
};

exports.VilmaIdentityPoolClient = grpc.makeGenericClientConstructor(VilmaIdentityPoolService);
