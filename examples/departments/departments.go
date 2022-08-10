package main

import (
	"context"
	"fmt"
	"log"
	"os"

	fs "github.com/rk295/go-freshservice/freshservice"
)

func main() {
	token := os.Getenv(fs.FreshserviceAPITokenEnvName)
	domain := os.Getenv(fs.FreshserviceCompanyDomainEnvName)
	ctx := context.Background()
	api, err := fs.New(ctx, domain, token, nil)
	if err != nil {
		log.Fatal(err)
	}

	opts := &fs.DepartmentListOptions{
		PageQuery: "page=1",
	}

	t, _, err := api.Departments().List(ctx, opts)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d departments\n", len(t))
	fmt.Println("ID\t\tName")
	fmt.Println("----------\t----------")
	for _, d := range t {
		fmt.Printf("%d\t%s\n", d.ID, d.Name)
	}

}
