package main

import (
	"fmt"
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func startCosmos(writeOutput func(msg string)) {
	endpoint := ""

	writeOutput("Current Status:\tStarting...")

	// <create_client>
	credential, _ := azidentity.NewDefaultAzureCredential(nil)
	
	client, _ := azcosmos.NewClient(endpoint, credential, nil)
	// </create_client>

	// <get_database>
	database, _ := client.NewDatabase("cosmicworks")
	// </get_database>

	// <get_container>
	container, _ := database.NewContainer("products")
	// </get_container>

	// <create_item>
	item := Item {
		Id:			"70b63682-b93a-4c77-aad2-65501347265f",
		Category:	"gear-surf-surfboards",
		Name:		"Yamba Surfboard",
		Quantity:	12,
		Sale:		false,
	}

	partitionKey := azcosmos.NewPartitionKeyString("gear-surf-surfboards")

	context := context.TODO()

	bytes, _ := json.Marshal(item)

	response, _ := container.CreateItem(context, partitionKey, bytes, nil)
	// </create_item>

	if response.RawResponse.StatusCode == 200 || response.RawResponse.StatusCode == 201 {
		created_item := Item{}
		json.Unmarshal(response.Value, &created_item)

		writeOutput(fmt.Sprintf("Upserted item:\t%v", created_item))
	}
	writeOutput(fmt.Sprintf("Request charge:\t%f", response.RequestCharge))
}
