package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// Data contains configuration for this stack
type Data struct {
	Accounts []Account
}

// Account contains properties of the AWS accounts to be created
type Account struct {
	Name string
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var data Data
		cfg := config.New(ctx, "")
		cfg.RequireObject("data", &data)

		_, err := organizations.NewOrganization(ctx, "personal-projects", &organizations.OrganizationArgs{})
		if err != nil {
			return err
		}

		for _, account := range data.Accounts {
			_, err := organizations.NewAccount(ctx, account.Name, &organizations.AccountArgs{
				Email: pulumi.Sprintf("matthewcanty+aws-account+%s@gmail.com", account.Name),
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}
