package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
)

func main() {
	token := os.Getenv("TOKEN")
	domain := os.Getenv("DOMAIN")
	ctx := context.Background()
	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	page := "page=1"
	assetTypes := []fs.AssetTypeDetails{}
	for {

		f := &fs.AssetTypeListOptions{
			PageQuery: page,
		}

		t, nextPage, err := api.AssetTypes().List(ctx, f)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		assetTypes = append(assetTypes, t...)

		if nextPage == "" {
			break
		}

		page = nextPage
	}

	for _, a := range assetTypes {
		fmt.Println(a.ID, a.Name, a.Description)
	}
	fmt.Println("found ", len(assetTypes), " assets")

}
