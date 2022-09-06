package main

import (
	"context"
	"net/http"

	_ "goa.design/goa/v3/codegen/generator"

	"example.com/goa_issue/gen/http/test_service/server"
	testservice "example.com/goa_issue/gen/test_service"
	goahttp "goa.design/goa/v3/http"
)

var _ testservice.Service = (*Service)(nil)

type Service struct {
}

// Set implements testservice.Service
func (*Service) Set(context.Context, *testservice.Workspace) (err error) {
	panic("unimplemented")
}

// Update implements testservice.Service
func (*Service) Update(context.Context, *testservice.WorkspaceUpdatePayload) (err error) {
	panic("unimplemented")
}

//Update implements testservice.Service
//func (*Service) Update(ctx context.Context, payload *testservice.Workspace) (err error) {
//return fmt.Errorf("Incoming ID: %s", payload.ID)
//}

//Set implements testservice.Service
//func (*Service) Set(ctx context.Context, payload *testservice.Workspace) (err error) {
//return fmt.Errorf("Incoming ID: %s", payload.ID)
//}

func main() {
	s := &Service{}
	endpoints := testservice.NewEndpoints(s)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil)
	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
