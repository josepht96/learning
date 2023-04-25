package main

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
)

func main() {
	// Initialize the Keycloak client
	client := gocloak.NewClient("http://localhost:30000")

	// Set the Keycloak realm and authenticate the client
	realm := "master"
	token, err := client.LoginAdmin(
		context.Background(),
		"admin",
		"password123",
		realm,
	)
	if err != nil {
		fmt.Printf("Failed to authenticate the Keycloak client: %v\n", err)
		return
	}
	clientID := "test-from-go"
	clientSecret := "reeee"
	publicClient := false
	redirectURIs := []string{"*"}

	// Set the Keycloak client data
	newClient := gocloak.Client{
		ClientID:     &clientID,
		Name:         &clientSecret,
		Secret:       &clientSecret,
		PublicClient: &publicClient,
		RedirectURIs: &redirectURIs,
	}

	// Create the new client using the Keycloak API
	_, err = client.CreateClient(
		context.Background(),
		token.AccessToken,
		realm,
		newClient,
	)
	if err != nil {
		fmt.Printf("Failed to create the new Keycloak client: %v\n", err)
		return
	}

	fmt.Println("New Keycloak client created successfully!")
}
