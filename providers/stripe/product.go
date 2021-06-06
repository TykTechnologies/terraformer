package stripe

import (
	"strings"

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

		// Fetch and store product's prices and (deprecated) plans
		if err := g.createPriceResources(c, p.ID); err != nil {
			return err
		}
	}

	if err := pIter.Err(); err != nil {
		return err
	}

	return nil
}

// createPriceResources generates resources from existing Stripe prices and deprecated plans
func (g *ProductGenerator) createPriceResources(c *client.API, productID string) error {
	pIter := c.Prices.List(&stripe.PriceListParams{Product: stripe.String(productID)})

	for pIter.Next() {
		p := pIter.Price()

		if strings.HasPrefix(p.ID, "price_") {
			g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
				p.ID,
				p.ID,
				"stripe_price",
				g.GetProviderName(),
				[]string{},
			))
		} else {
			// Stripe price API is compatible with the deprecated plans as well, so the provider can handle the rest
			g.Resources = append(g.Resources, terraformutils.NewResource(
				p.ID,
				p.ID,
				"stripe_plan",
				g.GetProviderName(),
				map[string]string{
					"product": productID,
				},
				[]string{},
				map[string]interface{}{},
			))
		}
	}

	if err := pIter.Err(); err != nil {
		return err
	}

	return nil
}
