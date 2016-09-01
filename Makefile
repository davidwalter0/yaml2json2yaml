# statically link go 

DIR=$(GOPATH)
.PHONY: yaml2json json2yaml 
# targets:=$(patsubst %.go,bin/%,$(wildcard *.go))
targets:=$(patsubst %.go,bin/%,yaml2json.go json2yaml.go)

all: init build test
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


init: get save

%: bin/%

bin/%: %.go
	@echo "Building via % rule for $@ from $<"
	@mkdir -p bin
	./build

get:
	godep get $(libdep)
save:
	godep save

test: 
	echo No unit tests written, see transform package
#  $(GOPATH)/bin/godep go test -v

test-filter:  test-yaml2json test-json2yaml
	touch $(tmplist)
	bin/json2yaml --version
	bin/json2yaml < spec.image.json.yaml2json.unformatted > spec.image.json.json2yaml.reformatted

test-json2yaml: bin/json2yaml
	bin/json2yaml < spec.image.json.yaml2json > spec.image.json.json2yaml

test-yaml2json: bin/yaml2json
	bin/yaml2json --version
	bin/yaml2json < spec.image.json > spec.image.json.yaml2json
	bin/yaml2json --compress < spec.image.json > spec.image.json.yaml2json.unformatted

clean:
	@echo cleaning up temporary files
	@rm -f bin/json2yaml bin/yaml2json
	@rm -rf bin
	@rm -f $(tmplist)

