libdep="gopkg.in/yaml.v2" 
all:
	@echo make targets init to initialize godeps, get, save, test and build
build: 
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep go build -a -ldflags '-s'

init: get save

get:
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep get $(libdep)
save:
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep save
test: 
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep go test -v

