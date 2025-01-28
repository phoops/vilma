package grpc

import (
	"context"

	"bitbucket.org/phoops/vilma/internal/core/interactors"
	"bitbucket.org/phoops/vilma/internal/infrastructure/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VilmaHandler struct {
	logger          *zap.SugaredLogger
	getIdentityById *interactors.GetIdentityById
	proto.UnimplementedVilmaIdentityPoolServer
}

func NewVilmaHandler(
	logger *zap.SugaredLogger,
	getIdentityById *interactors.GetIdentityById,
) *VilmaHandler {
	l := logger.With("component", "vilmaHandler")

	return &VilmaHandler{
		logger:          l,
		getIdentityById: getIdentityById,
	}
}

var _ proto.VilmaIdentityPoolServer = (*VilmaHandler)(nil)

func (h *VilmaHandler) GetIdentityByIdentityId(ctx context.Context, req *proto.GetIdentityByIdRequest) (*proto.Identity, error) {
	if len(req.IdentityId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "identity is is empty")
	}

	id, err := h.getIdentityById.Execute(ctx, req.IdentityId)
	if err != nil {
		h.logger.Errorw("error during identity retrieving", "errror", err)

		return nil, status.Error(codes.Internal, "error during identity retrieving try later")
	}

	return &proto.Identity{
		Id:        req.IdentityId,
		FirstName: id.FirstName,
		LastName:  id.LastName,
		Email:     id.Email,
	}, nil
}
