package commands

import (
	"context"
	"fmt"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var (
	agentCmd = &cobra.Command{
		Use:   "agents",
		Short: "List agents",
		Run:   agentRun,
	}

	agentID int
)

func init() {
	agentCmd.Flags().IntVarP(&agentID, "agent", "a", 0, "Get a specific agent")
	rootCmd.AddCommand(agentCmd)
}

func agentRun(cmd *cobra.Command, args []string) {
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

	if agentID != 0 {
		t, err := api.Agents().Get(ctx, agentID)
		if err != nil {
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		output = t
	} else {
		page := "page=1"
		agents := &fs.Agents{}

		for {

			f := &fs.AgentListFilter{
				PageQuery: page,
			}

			t, nextPage, err := api.Agents().List(ctx, f)
			if err != nil {
				fmt.Println("error: ", err)
				os.Exit(1)
			}
			agents.List = append(agents.List, t...)

			if nextPage == "" {
				break
			}

			page = nextPage
			output = t

		}

		output = agents
	}

	txt, err := printJSON(output)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println(txt)
}
