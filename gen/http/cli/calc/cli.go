// Code generated by goa v3.8.4, DO NOT EDIT.
//
// calc HTTP client CLI support package
//
// Command:
// $ goa gen example.com/goa_issue/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	testservicec "example.com/goa_issue/gen/http/test_service/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `test-service (update|set)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` test-service update --body '{
      "description": "Non et.",
      "destination_ids": [
         "Itaque sit corrupti.",
         "Ad iusto.",
         "Sed repudiandae molestiae modi quo.",
         "Et eum aut adipisci temporibus."
      ],
      "logo_url": "Et consequatur.",
      "name": "Alias incidunt."
   }' --id "Ex quae facilis necessitatibus."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		testServiceFlags = flag.NewFlagSet("test-service", flag.ContinueOnError)

		testServiceUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		testServiceUpdateBodyFlag = testServiceUpdateFlags.String("body", "REQUIRED", "")
		testServiceUpdateIDFlag   = testServiceUpdateFlags.String("id", "REQUIRED", "Workspace ID")

		testServiceSetFlags    = flag.NewFlagSet("set", flag.ExitOnError)
		testServiceSetBodyFlag = testServiceSetFlags.String("body", "REQUIRED", "")
		testServiceSetIDFlag   = testServiceSetFlags.String("id", "REQUIRED", "Workspace ID")
	)
	testServiceFlags.Usage = testServiceUsage
	testServiceUpdateFlags.Usage = testServiceUpdateUsage
	testServiceSetFlags.Usage = testServiceSetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "test-service":
			svcf = testServiceFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "test-service":
			switch epn {
			case "update":
				epf = testServiceUpdateFlags

			case "set":
				epf = testServiceSetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "test-service":
			c := testservicec.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "update":
				endpoint = c.Update()
				data, err = testservicec.BuildUpdatePayload(*testServiceUpdateBodyFlag, *testServiceUpdateIDFlag)
			case "set":
				endpoint = c.Set()
				data, err = testservicec.BuildSetPayload(*testServiceSetBodyFlag, *testServiceSetIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// test-serviceUsage displays the usage of the test-service command and its
// subcommands.
func testServiceUsage() {
	fmt.Fprintf(os.Stderr, `The calc service performs operations on numbers.
Usage:
    %[1]s [globalflags] test-service COMMAND [flags]

COMMAND:
    update: Update an existing workspace information
    set: Set the workspace to be used

Additional help:
    %[1]s test-service COMMAND --help
`, os.Args[0])
}
func testServiceUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] test-service update -body JSON -id STRING

Update an existing workspace information
    -body JSON: 
    -id STRING: Workspace ID

Example:
    %[1]s test-service update --body '{
      "description": "Non et.",
      "destination_ids": [
         "Itaque sit corrupti.",
         "Ad iusto.",
         "Sed repudiandae molestiae modi quo.",
         "Et eum aut adipisci temporibus."
      ],
      "logo_url": "Et consequatur.",
      "name": "Alias incidunt."
   }' --id "Ex quae facilis necessitatibus."
`, os.Args[0])
}

func testServiceSetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] test-service set -body JSON -id STRING

Set the workspace to be used
    -body JSON: 
    -id STRING: Workspace ID

Example:
    %[1]s test-service set --body '{
      "description": "Labore eaque consectetur.",
      "destination_ids": [
         "Omnis veritatis id iure repellat.",
         "Culpa odit.",
         "Odit non."
      ],
      "logo_url": "Et facere aperiam ut rerum ipsam.",
      "name": "Quibusdam eius."
   }' --id "Consequatur omnis velit."
`, os.Args[0])
}
