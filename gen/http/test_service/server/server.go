// Code generated by goa v3.10.2, DO NOT EDIT.
//
// TestService HTTP server
//
// Command:
// $ goa gen example.com/goa_issue/design

package server

import (
	"context"
	"net/http"

	testservice "example.com/goa_issue/gen/test_service"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the TestService service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Update http.Handler
	Set    http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the TestService service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *testservice.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Update", "PATCH", "/{id}"},
			{"Set", "PUT", "/{id}"},
		},
		Update: NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Set:    NewSetHandler(e.Set, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "TestService" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Update = m(s.Update)
	s.Set = m(s.Set)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return testservice.MethodNames[:] }

// Mount configures the mux to serve the TestService endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountUpdateHandler(mux, h.Update)
	MountSetHandler(mux, h.Set)
}

// Mount configures the mux to serve the TestService endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountUpdateHandler configures the mux to serve the "TestService" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/{id}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "TestService" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "TestService")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountSetHandler configures the mux to serve the "TestService" service "set"
// endpoint.
func MountSetHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/{id}", f)
}

// NewSetHandler creates a HTTP handler which loads the HTTP request and calls
// the "TestService" service "set" endpoint.
func NewSetHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSetRequest(mux, decoder)
		encodeResponse = EncodeSetResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "set")
		ctx = context.WithValue(ctx, goa.ServiceKey, "TestService")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
