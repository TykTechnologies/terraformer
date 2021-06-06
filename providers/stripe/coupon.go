package stripe

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/client"
)

type CouponGenerator struct {
	StripeService
}

func (g *CouponGenerator) InitResources() error {
	c, gcErr := g.getClient()
	if gcErr != nil {
		return gcErr
	}

	if err := g.createCouponResources(c); err != nil {
		return err
	}

	return nil
}

// createCouponResources generates resources from existing Stripe coupons
func (g *CouponGenerator) createCouponResources(c *client.API) error {
	cIter := c.Coupons.List(&stripe.CouponListParams{})

	for cIter.Next() {
		coupon := cIter.Coupon()

		g.Resources = append(g.Resources, terraformutils.NewSimpleResource(
			coupon.ID,
			coupon.ID,
			"stripe_coupon",
			g.GetProviderName(),
			[]string{},
		))
	}

	if err := cIter.Err(); err != nil {
		return err
	}

	return nil
}
