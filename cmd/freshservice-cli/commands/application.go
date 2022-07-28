package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "List applications",
	Run:   applicationRun,
}

func init() {
	rootCmd.AddCommand(applicationCmd)
}

func applicationRun(cmd *cobra.Command, args []string) {
	token := os.Getenv("TOKEN")
	domain := os.Getenv("DOMAIN")

	ctx := context.Background()
	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	t, _, err := api.Applications().List(ctx)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Print("All Applications:\n")
	for _, app := range t {
		fmt.Printf("%d - %s\n", app.ID, app.Name)
	}
}
