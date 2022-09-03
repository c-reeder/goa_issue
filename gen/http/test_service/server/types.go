// Code generated by goa v3.8.4, DO NOT EDIT.
//
// TestService HTTP server types
//
// Command:
// $ goa gen example.com/goa_issue/design

package server

import (
	testservice "example.com/goa_issue/gen/test_service"
)

// UpdateRequestBody is the type of the "TestService" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	Name           *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Description    *string  `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	DestinationIds []string `form:"destination_ids,omitempty" json:"destination_ids,omitempty" xml:"destination_ids,omitempty"`
	LogoURL        *string  `form:"logo_url,omitempty" json:"logo_url,omitempty" xml:"logo_url,omitempty"`
}

// SetRequestBody is the type of the "TestService" service "set" endpoint HTTP
// request body.
type SetRequestBody struct {
	Name           *string  `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Description    *string  `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	DestinationIds []string `form:"destination_ids,omitempty" json:"destination_ids,omitempty" xml:"destination_ids,omitempty"`
	LogoURL        *string  `form:"logo_url,omitempty" json:"logo_url,omitempty" xml:"logo_url,omitempty"`
}

// NewUpdateWorkspace builds a TestService service update endpoint payload.
func NewUpdateWorkspace(body *UpdateRequestBody, id string) *testservice.Workspace {
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

	return v
}

// NewSetWorkspace builds a TestService service set endpoint payload.
func NewSetWorkspace(body *SetRequestBody, id string) *testservice.Workspace {
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

	return v
}
