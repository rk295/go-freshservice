//go:build generate
// +build generate

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	fEndpoint    = flag.String("endpoint", "", "endpoint to generate code for")
	fAPIEndpoint = flag.String("api-endpoint", "", "endpoint of the API endpoint, defaults endpoint")
	fJSONKey     = flag.String("json-key", "", "json key to use for the API endpoint")
	fGenList     = flag.Bool("generate-list", false, "should we generate a List() function")
)

func Plural(s string) string {
	if strings.HasSuffix(s, "ss") {
		return s + "es"
	} else if strings.HasSuffix(s, "y") {
		return s[:len(s)-1] + "ies"
	} else if strings.HasSuffix(s, "s") {
		return s
	} else {
		return s + "s"
	}
}

func Export(s string) string {
	return strings.Title(s)
}

func StructTag(s string) string {
	return fmt.Sprintf("`json:\"%s\"`", s)
}

func main() {
	log.SetFlags(0)
	flag.Usage = func() {
		log.Println("Usage:")
		log.Println("\tgenerate", "<endpoint>")
		log.Println("Flags:")
		flag.PrintDefaults()
		log.Println("Example:")
		log.Fatalln("\tgenerate", "locations")
	}
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	outputFilename := flag.Arg(0)

	if *fAPIEndpoint == "" {
		*fAPIEndpoint = *fEndpoint
	}

	if *fJSONKey == "" {
		*fJSONKey = *fEndpoint
	}

	fmt.Printf("Generating code for %s into file %s\n", *fEndpoint, outputFilename)

	t := template.New("").Funcs(template.FuncMap{
		"Plural":    Plural,
		"Export":    Export,
		"StructTag": StructTag,
	})

	if _, err := t.Parse(funcTemplate); err != nil {
		log.Fatalln("error parsing in template:", err)
	}

	args := map[string]interface{}{
		"Endpoint":     *fEndpoint,
		"APIEndpoint":  *fAPIEndpoint,
		"JSONKey":      *fJSONKey,
		"GenerateList": *fGenList,
	}

	o := bytes.NewBuffer(nil)

	if err := t.Execute(o, args); err != nil {
		log.Fatalln("error in executing template:", err)
	}

	formatted, err := format.Source(o.Bytes())
	if err != nil {
		log.Fatalln("invalid generated go reported by formatter:", err)
	}

	if err := ioutil.WriteFile(outputFilename, formatted, 0644); err != nil {
		log.Fatalln("failed to write to file:", err)
	}
}

var funcTemplate = `package freshservice

// Generated Code DO NOT EDIT
{{if .GenerateList}}
import (
	"context"
	"net/http"
	"net/url"
)
{{end}}

const {{ .Endpoint }}URL = "/api/v2/{{Plural .APIEndpoint}}"


// {{Export (Plural .Endpoint) }} holds a list of Freshservice {{Export .Endpoint }} details
type {{Export (Plural .Endpoint) }} struct {
	List []{{Export .Endpoint }}Details {{StructTag (Plural .JSONKey) }}

}

// {{Export .Endpoint }} holds the details of a specific Freshservice {{Export .Endpoint }}
type {{Export .Endpoint }} struct {
	Details {{Export .Endpoint}}Details {{StructTag .JSONKey }}
}

// {{Export (Plural .Endpoint) }} is the interface between the HTTP client and the Freshservice {{ .Endpoint }} related endpoints
func (fs *Client) {{Export (Plural .Endpoint) }}() {{Export (Plural .Endpoint) }}Service {
	return &{{Export (Plural .Endpoint) }}ServiceClient{client: fs}
}

// {{Export (Plural .Endpoint) }}ServiceClient facilitates requests with the {{Export (Plural .Endpoint) }}Service methods
type {{Export (Plural .Endpoint) }}ServiceClient struct {
	client *Client
}

{{if .GenerateList}}
// List all {{Plural .Endpoint }}
func (d *{{Export (Plural .Endpoint) }}ServiceClient) List(ctx context.Context, filter QueryFilter) ([]{{Export .Endpoint}}Details, string, error) {

	url := &url.URL{
		Scheme: "https",
		Host:   d.client.Domain,
		Path:   {{ .Endpoint}}URL,
	}

	if filter != nil {
		url.RawQuery = filter.QueryString()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, "", err
	}

	res := &{{Export (Plural .Endpoint) }}{}
	resp, err := d.client.makeRequest(req, res)
	if err != nil {
		return nil, "", err
	}

	return res.List, HasNextPage(resp), nil
}
{{end}}
`
