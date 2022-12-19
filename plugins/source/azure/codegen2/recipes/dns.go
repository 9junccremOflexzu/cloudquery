// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"

func init() {
	tables := []Table{
		{
			Service:        "armdns",
			Name:           "zones",
			Struct:         &armdns.Zone{},
			ResponseStruct: &armdns.ZonesClientListResponse{},
			Client:         &armdns.ZonesClient{},
			ListFunc:       (&armdns.ZonesClient{}).NewListPager,
			NewFunc:        armdns.NewZonesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.Network/dnszones",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_network)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
