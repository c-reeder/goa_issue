// Code generated by goa v3.8.5, DO NOT EDIT.
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
func BuildUpdatePayload(testServiceUpdateBody string, testServiceUpdateID string) (*testservice.WorkspaceUpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(testServiceUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Quibusdam eius.\",\n      \"destination_ids\": [\n         \"Eaque consectetur excepturi eaque.\",\n         \"Veritatis id iure.\"\n      ],\n      \"logo_url\": \"Accusantium culpa odit eaque.\",\n      \"name\": \"Quasi est eum ad.\"\n   }'")
		}
	}
	var id string
	{
		id = testServiceUpdateID
	}
	v := &testservice.WorkspaceUpdatePayload{
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
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Mollitia et quis consectetur.\",\n      \"destination_ids\": [\n         \"Neque consequatur maiores aut.\",\n         \"Et consequatur quis dolorum corporis officiis rerum.\",\n         \"Quia dolores consequatur cum itaque.\",\n         \"Nam deserunt et tempora atque.\"\n      ],\n      \"logo_url\": \"Ea ut enim aut.\",\n      \"name\": \"Omnis velit laboriosam unde omnis incidunt.\"\n   }'")
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
