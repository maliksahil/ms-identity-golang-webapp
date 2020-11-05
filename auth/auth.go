package auth

import (
	"context"
	"log"
	"os"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://login.microsoftonline.com/"+os.Getenv("tenantid")+"/v2.0")
	if err != nil {
		log.Printf("failed to get provider: %v", err)
		return nil, err
	}

	conf := oauth2.Config{
		ClientID:     os.Getenv("clientid"),
		ClientSecret: os.Getenv("clientsecret"),
		RedirectURL:  os.Getenv("callbackurl"),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	log.Print(conf)

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}
