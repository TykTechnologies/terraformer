package stripe

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type StripeService struct { // nolint
	terraformutils.Service
}

// getClient initialises and returns the Stripe API client
func (s *StripeService) getClient() (*client.API, error) {
	iKey, ok := s.GetArgs()["api_token"]
	if !ok {
		return nil, errors.New("missing api_token")
	}

	key, ok := iKey.(string)
	if !ok {
		return nil, errors.New("api_token is not a string")
	}

	stripe.SetAppInfo(&stripe.AppInfo{
		Name: "terraformer-provider-stripe",
	})

	sc := client.New(key, nil)

	return sc, nil
}
