package stripe

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type PriceGenerator struct {
	StripeService
}

func (g *PriceGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createPriceResources(c); err != nil {
		return err
	}

	return nil
}

// createPriceResources generates resources from existing Stripe prices
func (g *PriceGenerator) createPriceResources(c *client.API) error {
	pIter := c.Prices.List(&stripe.PriceListParams{})

	for pIter.Next() {
		p := pIter.Price()

		if !strings.HasPrefix(p.ID, "price_") {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			p.ID,
			p.ID,
			"stripe_price",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := pIter.Err(); err != nil {
		return err
	}

	return nil
}
