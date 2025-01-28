const {
    VilmaIdentityPoolClient
} = require('./stubs/identity_grpc_pb')

const { credentials } = require("@grpc/grpc-js")
const { GetIdentityByIdRequest } = require('./stubs/identity_pb')

const newVilmaClient = (vilmaBaseUrl) => {
    const grpcClient = new VilmaIdentityPoolClient(
        vilmaBaseUrl,
        credentials.createInsecure(),
    )

    return {
        getIdentity(identityId){
            return new Promise((resolve, reject) => {
                const request = new GetIdentityByIdRequest()
                request.setIdentityId(identityId)

                grpcClient.getIdentityByIdentityId(request, (error, response) => {
                    if (error){
                        reject(error)
                        return
                    }

                    resolve({
                        id: identityId,
                        email: response.getEmail(),
                        firstName: response.getFirstName(),
                        lastName: response.getLastName(),
                    })
                })
            })
        }
    }
}

module.exports = {
    newVilmaClient: newVilmaClient,
}