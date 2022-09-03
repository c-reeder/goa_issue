// Code generated by goa v3.8.4, DO NOT EDIT.
//
// TestService service
//
// Command:
// $ goa gen example.com/goa_issue/design

package testservice

import (
	"context"
)

// The calc service performs operations on numbers.
type Service interface {
	// Update an existing workspace information
	Update(context.Context, *Workspace) (err error)
	// Set the workspace to be used
	Set(context.Context, *Workspace) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "TestService"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"update", "set"}

// Workspace is the payload type of the TestService service update method.
type Workspace struct {
	// Workspace ID
	ID             string
	Name           *string
	Description    *string
	DestinationIds []string
	LogoURL        *string
}
