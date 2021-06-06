package stripe

import (
	"strings"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type PlanGenerator struct {
	StripeService
}

func (g *PlanGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createPlanResources(c); err != nil {
		return err
	}

	return nil
}

// createPlanResources generates resources from existing Stripe plans
func (g *PlanGenerator) createPlanResources(c *client.API) error {
	pIter := c.Plans.List(&stripe.PlanListParams{})

	for pIter.Next() {
		p := pIter.Plan()

		if strings.HasPrefix(p.ID, "price_") {
			continue
		}

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			p.ID,
			p.ID,
			"stripe_plan",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := pIter.Err(); err != nil {
		return err
	}

	return nil
}
