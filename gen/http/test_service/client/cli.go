// Code generated by goa v3.8.4, DO NOT EDIT.
//
// TestService HTTP client CLI support package
//
// Command:
// $ goa gen example.com/goa_issue/design

package client

import (
	"encoding/json"
	"fmt"

	testservice "example.com/goa_issue/gen/test_service"
)

// BuildUpdatePayload builds the payload for the TestService update endpoint
// from CLI flags.
func BuildUpdatePayload(testServiceUpdateBody string, testServiceUpdateID string) (*testservice.Workspace, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(testServiceUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Non et.\",\n      \"destination_ids\": [\n         \"Itaque sit corrupti.\",\n         \"Ad iusto.\",\n         \"Sed repudiandae molestiae modi quo.\",\n         \"Et eum aut adipisci temporibus.\"\n      ],\n      \"logo_url\": \"Et consequatur.\",\n      \"name\": \"Alias incidunt.\"\n   }'")
		}
	}
	var id string
	{
		id = testServiceUpdateID
	}
	v := &testservice.Workspace{
		Name:        body.Name,
		Description: body.Description,
		LogoURL:     body.LogoURL,
	}
	if body.DestinationIds != nil {
		v.DestinationIds = make([]string, len(body.DestinationIds))
		for i, val := range body.DestinationIds {
			v.DestinationIds[i] = val
		}
	}
	v.ID = id

	return v, nil
}

// BuildSetPayload builds the payload for the TestService set endpoint from CLI
// flags.
func BuildSetPayload(testServiceSetBody string, testServiceSetID string) (*testservice.Workspace, error) {
	var err error
	var body SetRequestBody
	{
		err = json.Unmarshal([]byte(testServiceSetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Labore eaque consectetur.\",\n      \"destination_ids\": [\n         \"Omnis veritatis id iure repellat.\",\n         \"Culpa odit.\",\n         \"Odit non.\"\n      ],\n      \"logo_url\": \"Et facere aperiam ut rerum ipsam.\",\n      \"name\": \"Quibusdam eius.\"\n   }'")
		}
	}
	var id string
	{
		id = testServiceSetID
	}
	v := &testservice.Workspace{
		Name:        body.Name,
		Description: body.Description,
		LogoURL:     body.LogoURL,
	}
	if body.DestinationIds != nil {
		v.DestinationIds = make([]string, len(body.DestinationIds))
		for i, val := range body.DestinationIds {
			v.DestinationIds[i] = val
		}
	}
	v.ID = &id

	return v, nil
}