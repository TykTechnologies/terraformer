package stripe

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type TaxRateGenerator struct {
	StripeService
}

func (g *TaxRateGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createTaxRateResources(c); err != nil {
		return err
	}

	return nil
}

// createTaxRateResources generates resources from existing Stripe tax rates
func (g *TaxRateGenerator) createTaxRateResources(c *client.API) error {
	trIter := c.TaxRates.List(&stripe.TaxRateListParams{})

	for trIter.Next() {
		tr := trIter.TaxRate()

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			tr.ID,
			tr.ID,
			"stripe_tax_rate",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := trIter.Err(); err != nil {
		return err
	}

	return nil
}
