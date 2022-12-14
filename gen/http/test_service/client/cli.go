// Code generated by goa v3.10.2, DO NOT EDIT.
//
// TestService HTTP client CLI support package
//
// Command:
// $ goa gen example.com/goa_issue/design

package client

import (
	"encoding/json"
	"fmt"

	types "example.com/goa_issue/gen/types"
)

// BuildCreatePayload builds the payload for the TestService create endpoint
// from CLI flags.
func BuildCreatePayload(testServiceCreateBody string) (*types.FirstType, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(testServiceCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"description\": \"Similique rerum tempora quos totam quisquam voluptatibus.\",\n      \"id\": \"Tempora aut.\",\n      \"name\": \"Repudiandae perspiciatis est sint.\",\n      \"thing\": {\n         \"Description\": \"Cumque laudantium aliquid ea.\"\n      }\n   }'")
		}
	}
	v := &types.FirstType{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
	}
	if body.Thing != nil {
		v.Thing = marshalSecondTypeRequestBodyToTypesSecondType(body.Thing)
	}

	return v, nil
}
