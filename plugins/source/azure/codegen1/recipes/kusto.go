// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"

func Armkusto() []*Table {
	tables := []*Table{
		{
			NewFunc:        armkusto.NewClustersClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto",
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Kusto/clusters",
			Namespace:      "microsoft.kusto",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_kusto)`,
			Pager:          `NewListPager`,
			ResponseStruct: "ClustersClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armkusto())
}
