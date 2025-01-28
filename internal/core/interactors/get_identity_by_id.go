package interactors

import (
	"context"

	"bitbucket.org/phoops/vilma/internal/core/entities"
	"go.uber.org/zap"
)

type IdentityFetcher interface {
	GetIdentity(ctx context.Context, identityId string) (*entities.Identity, error)
}

type GetIdentityById struct {
	logger  *zap.SugaredLogger
	fetcher IdentityFetcher
}

func NewGetIdentityById(
	logger *zap.SugaredLogger,
	fetcher IdentityFetcher,
) *GetIdentityById {
	l := logger.With("component", "getIdentityById")

	return &GetIdentityById{
		logger:  l,
		fetcher: fetcher,
	}
}

func (u *GetIdentityById) Execute(ctx context.Context, identityId string) (*entities.Identity, error) {
	u.logger.Infow("fetching identity", "identity_id", identityId)

	return u.fetcher.GetIdentity(ctx, identityId)
}
