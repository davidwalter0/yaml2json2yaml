# statically link go 

DIR=$(GOPATH)
.PHONY: yaml2json json2yaml 
targets:=$(patsubst %.go,bin/%,$(wildcard *.go))

all:
	@echo make targets init to initialize godeps, get, save, test and build

build: $(targets)

# enable vendoring for go
GO15VENDOREXPERIMENT=1
export GO15VENDOREXPERIMENT
libdep="github.com/davidwalter0/transform"
tmplist= spec.image.json.json2yaml			\
     spec.image.json.json2yaml.reformatted		\
     spec.image.json.yaml2json				\
     spec.image.json.yaml2json.unformatted	


%: bin/%

bin/%: %.go
	@echo "Building via % rule for $@ from $<"
	@mkdir -p bin
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep go build -a -ldflags '-s' -o $@ $<

init: get save

get:
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep get $(libdep)
save:
	GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep save
test: 
	echo No unit tests written, see transform package
# GO15VENDOREXPERIMENT=1 GOPATH=${GOPATH} $(GOPATH)/bin/godep go test -v

test-filter:  test-yaml2json test-json2yaml
	touch $(tmplist)
	bin/json2yaml < spec.image.json.yaml2json.unformatted > spec.image.json.json2yaml.reformatted

test-json2yaml: all
	bin/json2yaml < spec.image.json.yaml2json > spec.image.json.json2yaml

test-yaml2json: yaml2json
	bin/yaml2json < spec.image.json > spec.image.json.yaml2json
	bin/yaml2json --compress < spec.image.json > spec.image.json.yaml2json.unformatted

clean:
	@echo cleaning up temporary files
	@rm -f bin/json2yaml bin/yaml2json
	@rm -rf bin
	@rm -f $(tmplist)

