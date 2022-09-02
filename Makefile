# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

# abigen is need in your $PATH
compile:
		go generate contracts/did/generator.go
		go generate contracts/resolver/generator.go
