package stripe

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type ProductGenerator struct {
	StripeService
}

func (g *ProductGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createProductResources(c); err != nil {
		return err
	}

	return nil
}

// createProductResources generates resources from existing Stripe products
func (g *ProductGenerator) createProductResources(c *client.API) error {
	pIter := c.Products.List(&stripe.ProductListParams{})

	for pIter.Next() {
		p := pIter.Product()

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			p.ID,
			p.ID,
			"stripe_product",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := pIter.Err(); err != nil {
		return err
	}

	return nil
}
