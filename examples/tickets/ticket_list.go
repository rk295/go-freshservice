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

	f := &fs.AssetListOptions{
		PageQuery: "page=1",
		Embed: &fs.AssetEmbedOptions{
			TypeFields: true,
			Trashed:    false,
		},
	}

	t, _, err := api.Assets().List(ctx, f)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println(t[0])

	// // List all tickets
	// // https://example.com/api/v2/tickets
	// t, np, err := api.Tickets().List(ctx, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// tList := []string{}
	// for _, tick := range t {
	// 	tList = append(tList, fmt.Sprintf("\n%d - %d", tick.ID, tick.ResponderID))
	// }

	// // example querying another page using the returned query parameter
	// if np != "" {
	// 	t, _, err := api.Tickets().List(ctx, &fs.TicketListOptions{PageQuery: np})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	for _, tick := range t {
	// 		tList = append(tList, fmt.Sprintf("\n%d - %d", tick.ID, tick.ResponderID))
	// 	}
	// }

	// fmt.Printf("All Tickets:\nCount: %d\nResults: %v\n", len(tList), tList)

	// // List tickets using a built in filer query and sort by
	// f := &fs.TicketListOptions{
	// 	FilterBy: &fs.TicketFilter{
	// 		RequesterEmail: fs.String("test-account@example.com"),
	// 	},
	// 	SortBy: &fs.SortOptions{
	// 		Descending: true,
	// 	},
	// }

	// // https://example.com/api/v2/tickets?email=test-account@example.com&order_type=desc
	// ftList := []string{}
	// ft, _, err := api.Tickets().List(ctx, f)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, ft := range ft {
	// 	ftList = append(ftList, fmt.Sprintf("\n%d - %d - %v", ft.ID, ft.RequesterID, ft.CreatedAt))
	// }

	// fmt.Printf("Filtered Tickets:\nCount: %d\nResults: %v\n", len(ft), ftList)
}
