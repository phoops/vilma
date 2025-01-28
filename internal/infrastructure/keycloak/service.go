package keycloak

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2/clientcredentials"

	"bitbucket.org/phoops/vilma/internal/core/entities"
)

type KeycloakIdentity struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Service struct {
	logger     *zap.SugaredLogger
	httpClient *http.Client
	baseURL    string
	realm      string
}

func NewService(
	logger *zap.SugaredLogger,
	tokenURL string,
	clientID string,
	clientSecret string,
	realm string,
	baseURL string,
) *Service {
	l := logger.With("component", "keycloakService")

	credentialsConfig := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
	}

	return &Service{
		logger:     l,
		httpClient: credentialsConfig.Client(context.Background()),
		baseURL:    baseURL,
		realm:      realm,
	}
}

func (s *Service) GetIdentity(ctx context.Context, identityId string) (*entities.Identity, error) {
	l := s.logger.With("method", "GetIdentity")
	var identity KeycloakIdentity

	l.Debugf("getting identity from identity: %s", identityId)

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s",
			s.baseURL,
			fmt.Sprintf("admin/realms/%s/users/%s", s.realm, identityId),
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"invalid status code for get user details, received: %d, with body of :%s",
			res.StatusCode,
			string(rawBody),
		)
	}

	err = json.Unmarshal(rawBody, &identity)
	if err != nil {
		return nil, errors.Wrap(err, "could not unmarshal the user get details response")
	}

	l.Debugf("identity: %+v", identity)

	entity := &entities.Identity{
		Email:     identity.Email,
		FirstName: identity.FirstName,
		LastName:  identity.LastName,
	}

	err = entity.Valid()
	if err != nil {
		return nil, errors.Wrap(err, "invalid identity")
	}

	return entity, nil
}
