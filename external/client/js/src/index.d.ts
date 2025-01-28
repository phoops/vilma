export interface VilmaResponse {
    id: string
    email: string
    firstName: string
    lastName: string
}

export interface VilmaClient {
    getIdentity(identityId: string): Promise<VilmaResponse>
}

export function newVilmaClient(vilmaBaseUrl: string): Promise<VilmaClient>

