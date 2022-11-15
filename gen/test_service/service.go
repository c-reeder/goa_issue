// Code generated by goa v3.10.2, DO NOT EDIT.
//
// TestService service
//
// Command:
// $ goa gen example.com/goa_issue/design

package testservice

import (
	"context"

	types "example.com/goa_issue/gen/types"
)

// The calc service performs operations on numbers.
type Service interface {
	// Create a new thing
	Create(context.Context, *types.FirstType) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "TestService"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"create"}

type SecondType struct {
	Description *string
}
