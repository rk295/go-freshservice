package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
)

func main() {
	token := os.Getenv(fs.FreshserviceAPITokenEnvName)
	domain := os.Getenv(fs.FreshserviceCompanyDomainEnvName)
	ctx := context.Background()
	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	// List 1 page of tickets
	t, _, err := api.Tickets().List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, tick := range t {
		fmt.Printf("%d - %s\n", tick.ID, tick.Subject)
	}

}
