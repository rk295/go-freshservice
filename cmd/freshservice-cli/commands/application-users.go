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
	applicationUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "List application users",
		Run:   applicationUserRun,
	}
)

func init() {
	applicationUsersCmd.Flags().Int64VarP(&appID, "application", "a", 0, "Get users of a specific application")
	applicationCmd.AddCommand(applicationUsersCmd)
}

func applicationUserRun(cmd *cobra.Command, args []string) {

	if appID == 0 {
		fmt.Println("Error: application id is required")
		os.Exit(1)
	}

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
		t, err := api.Applications().ListUsers(ctx, appID)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		output = t
	}

	txt, err := printJSON(output)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(txt)
}
