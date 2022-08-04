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
	assets := []fs.AssetDetails{}
	for {

		assetState := "In Use"
		f := &fs.AssetListOptions{
			PageQuery: page,
			Embed: &fs.AssetEmbedOptions{
				TypeFields: true,
				Trashed:    false,
			},
			FilterBy: &fs.AssetFilter{
				AssetState: &assetState,
			},
		}

		t, nextPage, err := api.Assets().List(ctx, f)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		assets = append(assets, t...)

		if nextPage == "" {
			break
		}

		page = nextPage
	}

	for _, a := range assets {
		fmt.Println(a.Name, a.AssetTypeID, a.DepartmentID)
	}
	fmt.Println("found ", len(assets), " assets")

}
