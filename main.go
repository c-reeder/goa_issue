package main

import (
	"context"
	"net/http"

	_ "goa.design/goa/v3/codegen/generator"

	goahttp "goa.design/goa/v3/http"

	"example.com/goa_issue/gen/http/test_service/server"
	testservice "example.com/goa_issue/gen/test_service"
	"example.com/goa_issue/gen/types"
)

var _ testservice.Service = (*Service)(nil)

type Service struct {
}

func (s Service) Create(ctx context.Context, firstType *types.FirstType) (err error) {
	//TODO implement me
	panic("implement me")
}

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
