package commands

import (
	"context"
	"fmt"
	"log"
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
	token := os.Getenv("TOKEN")
	domain := os.Getenv("DOMAIN")

	ctx := context.Background()
	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	t, np, err := api.Tickets().List(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	tList := []string{}
	for _, tick := range t {
		tList = append(tList, fmt.Sprintf("\n%d - %d", tick.ID, tick.ResponderID))
	}

	// example querying another page using the returned query parameter
	if np != "" {
		t, _, err := api.Tickets().List(ctx, &fs.TicketListOptions{PageQuery: np})
		if err != nil {
			log.Fatal(err)
		}
		for _, tick := range t {
			tList = append(tList, fmt.Sprintf("\n%s - %d", tick.Subject, tick.ResponderID))
		}
	}

	fmt.Printf("All Tickets:\nCount: %d\nResults: %v\n", len(tList), tList)

}
