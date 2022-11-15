// Code generated by goa v3.10.2, DO NOT EDIT.
//
// TestService endpoints
//
// Command:
// $ goa gen example.com/goa_issue/design

package testservice

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "TestService" service endpoints.
type Endpoints struct {
	Update goa.Endpoint
	Set    goa.Endpoint
}

// NewEndpoints wraps the methods of the "TestService" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Update: NewUpdateEndpoint(s),
		Set:    NewSetEndpoint(s),
	}
}

// Use applies the given middleware to all the "TestService" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Update = m(e.Update)
	e.Set = m(e.Set)
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "TestService".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*WorkspaceUpdatePayload)
		return nil, s.Update(ctx, p)
	}
}

// NewSetEndpoint returns an endpoint function that calls the method "set" of
// service "TestService".
func NewSetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Workspace)
		return nil, s.Set(ctx, p)
	}
}
