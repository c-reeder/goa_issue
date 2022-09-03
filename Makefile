# Makefile for example goa issue


goa-gen:
	goa gen example.com/goa_issue/design

build:
	go build example.com/goa_issue
