package commands

import (
	"context"
	"fmt"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var (
	assetCmd = &cobra.Command{
		Use:   "assets",
		Short: "List assets",
		Run:   assetRun,
	}

	assetID    int
	trashed    bool
	typeFields bool
)

func init() {
	assetCmd.Flags().IntVarP(&assetID, "asset", "a", 0, "Get a specific asset (must be display_id NOT ID)")
	assetCmd.Flags().BoolVarP(&trashed, "trashed", "t", false, "Will return all the assets that are in trash")
	assetCmd.Flags().BoolVarP(&typeFields, "type-fields", "f", false, "Will return all fields that are specific to each asset type")
	rootCmd.AddCommand(assetCmd)
}

func assetRun(cmd *cobra.Command, args []string) {
	token, domain, err := getConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	var output interface{}

	page := "page=1"
	assets := &fs.Assets{}
	if assetID != 0 {
		t, err := api.Assets().Get(ctx, assetID)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		output = t
	} else {
		for {

			f := &fs.AssetListOptions{
				PageQuery: page,
				Embed: &fs.AssetEmbedOptions{
					TypeFields: typeFields,
					Trashed:    trashed,
				},
			}

			t, nextPage, err := api.Assets().List(ctx, f)
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}

			assets.List = append(assets.List, t...)

			if nextPage == "" {
				break
			}

			page = nextPage
		}

		output = assets
	}

	txt, err := printJSON(output)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println(txt)
}
