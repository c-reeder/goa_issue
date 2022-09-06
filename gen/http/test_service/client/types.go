// Code generated by goa v3.8.5, DO NOT EDIT.
//
// TestService HTTP client types
//
// Command:
// $ goa gen example.com/goa_issue/design

package client

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

// NewUpdateRequestBody builds the HTTP request body from the payload of the
// "update" endpoint of the "TestService" service.
func NewUpdateRequestBody(p *testservice.WorkspaceUpdatePayload) *UpdateRequestBody {
	body := &UpdateRequestBody{
		Name:        p.Name,
		Description: p.Description,
		LogoURL:     p.LogoURL,
	}
	if p.DestinationIds != nil {
		body.DestinationIds = make([]string, len(p.DestinationIds))
		for i, val := range p.DestinationIds {
			body.DestinationIds[i] = val
		}
	}
	return body
}

// NewSetRequestBody builds the HTTP request body from the payload of the "set"
// endpoint of the "TestService" service.
func NewSetRequestBody(p *testservice.Workspace) *SetRequestBody {
	body := &SetRequestBody{
		Name:        p.Name,
		Description: p.Description,
		LogoURL:     p.LogoURL,
	}
	if p.DestinationIds != nil {
		body.DestinationIds = make([]string, len(p.DestinationIds))
		for i, val := range p.DestinationIds {
			body.DestinationIds[i] = val
		}
	}
	return body
}
