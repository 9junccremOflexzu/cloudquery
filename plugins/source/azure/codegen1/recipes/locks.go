// Code generated by codegen0; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"

func Armlocks() []*Table {
	tables := []*Table{
		{
			NewFunc:        armlocks.NewAuthorizationOperationsClient,
			PkgPath:        "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks",
			URL:            "/providers/Microsoft.Authorization/operations",
			Namespace:      "microsoft.authorization",
			Multiplex:      `client.SubscriptionMultiplexRegisteredNamespace(client.Namespacemicrosoft_authorization)`,
			Pager:          `NewListPager`,
			ResponseStruct: "AuthorizationOperationsClientListResponse",
		},
	}
	return tables
}

func init() {
	Tables = append(Tables, Armlocks())
}
