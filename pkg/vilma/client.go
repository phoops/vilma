package vilma

import (
	"context"
	"fmt"

	"bitbucket.org/phoops/vilma/internal/infrastructure/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VilmaIdentity struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

type VilmaClientOption func(*Vilma) error

func SetLogger(l *zap.SugaredLogger) VilmaClientOption {
	return func(m *Vilma) error {
		m.logger = l
		return nil
	}
}

type Vilma struct {
	logger     *zap.SugaredLogger
	baseURL    string
	gprcClient proto.VilmaIdentityPoolClient
	baseClient *grpc.ClientConn
}

func NewVilmaClient(baseURL string, opts ...VilmaClientOption) (*Vilma, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	c, err := grpc.Dial(baseURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("could not connect to Vilma instance: %s", baseURL)
	}

	mc := proto.NewVilmaIdentityPoolClient(c)

	m := Vilma{
		logger:     logger.Sugar().With("component", "VilmaClient"),
		baseURL:    baseURL,
		gprcClient: mc,
		baseClient: c,
	}

	// apply options
	for _, o := range opts {
		err := o(&m)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}

func (c *Vilma) GetIdentityById(ctx context.Context, identityId string) (*VilmaIdentity, error) {
	c.logger.Debug("getting identiy", "identity_id", identityId)

	res, err := c.gprcClient.GetIdentityByIdentityId(ctx, &proto.GetIdentityByIdRequest{
		IdentityId: identityId,
	})

	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve identity from vilma server")
	}

	return &VilmaIdentity{
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		ID:        identityId,
	}, nil
}

func (c *Vilma) Close() error {
	return c.baseClient.Close()
}
