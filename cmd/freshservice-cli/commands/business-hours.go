package commands

import (
	"context"
	"fmt"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
	"github.com/spf13/cobra"
)

var (
	businessHoursCmd = &cobra.Command{
		Use:   "business-hours",
		Short: "List business hours",
		Run:   businessHoursRun,
	}
	businessHoursID int
)

func init() {
	businessHoursCmd.Flags().IntVarP(&businessHoursID, "id", "b", -1, "Get a specific business hours definition")
	rootCmd.AddCommand(businessHoursCmd)
}

func businessHoursRun(cmd *cobra.Command, args []string) {

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

	if businessHoursID >= 0 {
		t, err := api.BusinessHours().Get(ctx, businessHoursID)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		output = t
	} else {
		output, err = api.BusinessHours().List(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	txt, err := printJSON(output)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Println(txt)

}
