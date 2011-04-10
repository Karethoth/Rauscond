include $(GOROOT)/src/Make.inc

TARG = rauscond

SRC = src

GOFILES = $(wildcard $(SRC)/*.go)

include $(GOROOT)/src/Make.cmd

