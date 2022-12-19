// Code generated by codegen1; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devhub/armdevhub"

func init() {
	tables := []Table{
		{
			Service:        "armdevhub",
			Name:           "workflow",
			Struct:         &armdevhub.Workflow{},
			ResponseStruct: &armdevhub.WorkflowClientListResponse{},
			Client:         &armdevhub.WorkflowClient{},
			ListFunc:       (&armdevhub.WorkflowClient{}).NewListPager,
			NewFunc:        armdevhub.NewWorkflowClient,
			URL:            "/subscriptions/{subscriptionId}/providers/Microsoft.DevHub/workflows",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_devhub)`,
			ExtraColumns:   DefaultExtraColumns,
		},
	}
	Tables = append(Tables, tables...)
}
