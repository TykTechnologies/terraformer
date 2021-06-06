package stripe

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type WebhookEndpointGenerator struct {
	StripeService
}

func (g *WebhookEndpointGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createWebhookEndpointResources(c); err != nil {
		return err
	}

	return nil
}

// createWebhookEndpointResources generates resources from existing Stripe webhook endpoints
func (g *WebhookEndpointGenerator) createWebhookEndpointResources(c *client.API) error {
	weIter := c.WebhookEndpoints.List(&stripe.WebhookEndpointListParams{})

	for weIter.Next() {
		we := weIter.WebhookEndpoint()

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			we.ID,
			we.ID,
			"stripe_webhook_endpoint",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := weIter.Err(); err != nil {
		return err
	}

	return nil
}
