// Code generated by codegen2; DO NOT EDIT.
package authorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ProviderOperationsMetadata() *schema.Table {
	return &schema.Table{
		Name:      "azure_authorization_provider_operations_metadata",
		Resolver:  fetchProviderOperationsMetadata,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_authorization),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "display_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DisplayName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "operations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Operations"),
			},
			{
				Name:     "resource_types",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResourceTypes"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchProviderOperationsMetadata(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc, err := armauthorization.NewProviderOperationsMetadataClient(cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
