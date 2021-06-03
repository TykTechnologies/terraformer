package cmd

import (
	stripe_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/stripe"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdStripeImporter(options ImportOptions) *cobra.Command {
	apiToken := ""

	cmd := &cobra.Command{
		Use:   "stripe",
		Short: "Import current state to Terraform configuration from Stripe",
		Long:  "Import current state to Terraform configuration from Stripe",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newStripeProvider()
			return Import(provider, options, []string{apiToken})
		},
	}

	cmd.AddCommand(listCmd(newStripeProvider()))

	baseProviderFlags(cmd.PersistentFlags(), &options, "product", "product=id1:id2:id4")
	cmd.PersistentFlags().StringVarP(&apiToken, "token", "t", "",
		"YOUR_STRIPE_TOKEN or env param STRIPE_APITOKEN")

	return cmd
}

func newStripeProvider() terraformutils.ProviderGenerator {
	return &stripe_terraforming.StripeProvider{}
}
