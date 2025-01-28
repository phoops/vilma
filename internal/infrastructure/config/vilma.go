package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment string

type VilmaConfig struct {
	Environment        Environment `required:"true" split_words:"true"`
	ListenAddr         string      `required:"true" split_words:"true" default:":9095"`
	KeycloakBaseUrl    string      `required:"true" split_words:"true"`
	KeycloakRealm      string      `required:"true" split_words:"true"`
	OAuth2TokenURL     string      `required:"true" split_words:"true"`
	OAuth2ClientID     string      `required:"true" split_words:"true"`
	OAuth2ClientSecret string      `required:"true" split_words:"true"`
}

func (s VilmaConfig) String() string {
	return fmt.Sprintf(`
Environment: %s 
KeycloakBaseUrl: %s,
KeycloakRealm: %s,
OAuth2TokenURL: %s,
OAuth2ClientID: %s,
OAuth2ClientSecret: %s,
`,
		s.Environment,
		s.KeycloakBaseUrl,
		s.KeycloakRealm,
		s.OAuth2TokenURL,
		s.OAuth2ClientID,
		s.OAuth2ClientSecret,
	)
}

func LoadVilmaConfig() (*VilmaConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("could not load configuration from .env file: %v", err)
	}
	var c VilmaConfig
	err = envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	log.Printf("Loaded configuration\n%+s", c)
	return &c, nil
}
