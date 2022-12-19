// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"

func init() {
	tables := []Table{
		{
			Service:        "armdatamigration",
			Name:           "services",
			Struct:         &armdatamigration.Service{},
			ResponseStruct: &armdatamigration.ServicesClientListResponse{},
			Client:         &armdatamigration.ServicesClient{},
			ListFunc:       (&armdatamigration.ServicesClient{}).NewListPager,
			NewFunc:        armdatamigration.NewServicesClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DataMigration/services",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_datamigration)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
