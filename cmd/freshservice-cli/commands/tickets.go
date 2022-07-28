package commands

import (
	"context"
	"fmt"
	"os"

	fs "github.com/CoreyGriffin/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "List tickets",
	Run:   ticketsRun,
}

func init() {
	rootCmd.AddCommand(ticketsCmd)
}

func ticketsRun(cmd *cobra.Command, args []string) {

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

	page := "page=1"
	tickets := &fs.Tickets{}

	for {
		t, nextPage, err := api.Tickets().List(ctx, &fs.TicketListOptions{PageQuery: page})
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}

		tickets.List = append(tickets.List, t...)

		if nextPage == "" {
			break
		}

		page = nextPage

	}

	txt, err := printJSON(tickets)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println(txt)

}
