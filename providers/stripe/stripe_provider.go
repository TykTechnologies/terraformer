package stripe

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type StripeProvider struct { // nolint
	terraformutils.Provider

	// apiToken is a Stripe API apiToken
	apiToken string
}

func (p *StripeProvider) Init(args []string) error {
	p.apiToken = os.Getenv("STRIPE_APITOKEN")
	if p.apiToken == "" && len(args) > 0 {
		p.apiToken = args[0]
	}

	if p.apiToken == "" {
		return errors.New("stripe api token is required")
	}

	return nil
}

func (p *StripeProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}

	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"api_token": p.apiToken,
	})

	return nil
}

func (p *StripeProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

func (p *StripeProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *StripeProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_token": cty.StringVal(p.apiToken),
	})
}

func (p *StripeProvider) GetName() string {
	return "stripe"
}

func (p *StripeProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"product":          &ProductGenerator{},
		"price":            &PriceGenerator{},
		"plan":             &PlanGenerator{},
		"coupon":           &CouponGenerator{},
		"webhook_endpoint": &WebhookEndpointGenerator{},
		"tax_rate":         &TaxRateGenerator{},
	}
}
