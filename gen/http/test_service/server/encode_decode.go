// Code generated by goa v3.10.2, DO NOT EDIT.
//
// TestService HTTP server encoders and decoders
//
// Command:
// $ goa gen example.com/goa_issue/design

package server

import (
	"context"
	"io"
	"net/http"

	types "example.com/goa_issue/gen/types"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the
// TestService create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the TestService
// create endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewCreateFirstType(&body)

		return payload, nil
	}
}

// unmarshalSecondTypeRequestBodyToTypesSecondType builds a value of type
// *types.SecondType from a value of type *SecondTypeRequestBody.
func unmarshalSecondTypeRequestBodyToTypesSecondType(v *SecondTypeRequestBody) *types.SecondType {
	if v == nil {
		return nil
	}
	res := &types.SecondType{
		Description: v.Description,
	}

	return res
}
