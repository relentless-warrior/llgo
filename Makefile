include $(GOROOT)/src/Make.inc

TARG=llgo
GOFILES=llgo.go expr.go println.go stmt.go decl.go literals.go

include $(GOROOT)/src/Make.cmd
