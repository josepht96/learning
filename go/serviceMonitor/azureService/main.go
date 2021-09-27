package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/storage/mgmt/storage"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	fmt.Println("AZURE_TENANT_ID:", os.Getenv("AZURE_TENANT_ID"))
	fmt.Println("AZURE_CLIENT_ID:", os.Getenv("AZURE_CLIENT_ID"))
	subscriptionID := os.Args[1]
	ctx := context.Background()
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Auth: Successful")
	}
	//getResourceGroups(ctx, &authorizer, subscriptionID)
	//getVirtualNetworks(ctx, &authorizer, subscriptionID)
	getStorageAccounts(ctx, &authorizer, subscriptionID)
}

func getResourceGroups(ctx context.Context, authorizer *autorest.Authorizer, subscriptionID string) {
	groupsClient := resources.NewGroupsClient(subscriptionID)
	groupsClient.Authorizer = *authorizer
	for list, err := groupsClient.ListComplete(context.Background(), "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Println(err)
		}
		rgName := *list.Value().Tags["test"]
		fmt.Println(rgName)
	}
}
func getVirtualNetworks(ctx context.Context, authorizer *autorest.Authorizer, subscriptionID string) {
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)
	vnetClient.Authorizer = *authorizer
	for list, err := vnetClient.ListComplete(context.Background(), "default"); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Println(err)
		}
		vnets := *list.Value().Name
		fmt.Println(vnets)
	}
}

func getStorageAccounts(ctx context.Context, authorizer *autorest.Authorizer, subscriptionID string) {
	storageAccountsClient := storage.NewAccountsClient(subscriptionID)
	storageAccountsClient.Authorizer = *authorizer
	for list, err := storageAccountsClient.ListComplete(context.Background()); list.NotDone(); err = list.Next() {
		if err != nil {
			fmt.Println(err)
		}
		vnets := *list.Value().Name
		fmt.Println(vnets)
	}
	resp, err := storageAccountsClient.ListKeys(context.Background(), "default", "defaultstracc96", "kerb")
	if err != nil {
		fmt.Println(err)
	}
	for _, key := range *resp.Keys {
		fmt.Printf("name: %s, value %s\n", *(key).KeyName, *(key).Value)
	}
}

//AzureAuth func
// func AzureAuth(subscriptionID string) compute.VirtualMachinesClient {
// 	vmClient := compute.NewVirtualMachinesClient(subscriptionID)

// 	authorizer, err := auth.NewAuthorizerFromEnvironment()
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Auth: Successful")
// 		vmClient.Authorizer = authorizer
// 	}
// 	return vmClient
// }
