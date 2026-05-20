package main

import (
	"context"
	"fmt"
	"os"

	cato_go_sdk "github.com/catonetworks/cato-go-sdk"
	cato_models "github.com/catonetworks/cato-go-sdk/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading .env")
		os.Exit(1)
	}

	url := "https://api.catonetworks.com/api/v1/graphql2"
	token := os.Getenv("CATO_TOKEN")
	accountID := os.Getenv("CATO_ACCOUNT_ID")

	catoClient, err := cato_go_sdk.New(url, token, accountID, nil, nil)
	if err != nil {
		fmt.Println("error creating Cato clientv2")
		os.Exit(1)
	}

	fmt.Println("Cato Client authenticated successfully")

	ctx := context.Background()

	result, err := catoClient.EntityLookup(
		ctx,
		accountID,
		cato_models.EntityType("site"),
		nil,        // parent
		nil,        // search
		nil,        // entityInput (filter by specific IDs in shape {id, name, type})
		nil,        // sort
		[]string{}, // entityIDs (empty = all)
		nil,        // from
		nil,        // limit
		nil,        // helperFields
	)
	if err != nil {
		fmt.Println("EntityLookup error:", err)
		os.Exit(1)
	}

	for _, item := range result.EntityLookup.GetItems() {
		// item.Entity has ID, Name, Type; Description and HelperFields are siblings
		id := item.Entity.ID
		name := ""
		if item.Entity.Name != nil {
			name = *item.Entity.Name
		}
		fmt.Printf("%s\t%s\n", id, name)
	}
}
