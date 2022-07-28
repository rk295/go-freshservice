package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var (
	applicationCmd = &cobra.Command{
		Use:   "application",
		Short: "List applications",
		Run:   applicationRun,
	}

	appID int64
)

func init() {
	applicationCmd.Flags().Int64VarP(&appID, "application", "a", 0, "Get a specific application")
	rootCmd.AddCommand(applicationCmd)
}

func applicationRun(cmd *cobra.Command, args []string) {
	token, domain, err := getConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()

	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	var output interface{}

	if appID != 0 {
		t, err := api.Applications().Get(ctx, appID)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		output = t
	} else {
		page := "page=1"
		applications := &fs.Applications{}

		for {
			t, nextPage, err := api.Applications().List(ctx, &fs.ApplicationListOptions{PageQuery: page})
			if err != nil {
				log.Println(t)
				log.Fatal("Error: ", err)
			}

			applications.List = append(applications.List, t...)

			if nextPage == "" {
				break
			}

			page = nextPage
		}

		output = applications
	}

	txt, err := printJSON(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(txt)
}
