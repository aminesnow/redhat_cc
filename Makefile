NAME=object-store-service
VERSION=$(shell cat VERSION)

.PHONY: help

#help help target
help:
	@fgrep -h "#help" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/#help//'

#####################
# Tools targets     #
#####################
.PHONY: tools.clean tools.get tools

TOOLS_DIR=$(CURDIR)/tools/bin

#help tools.clean: remove everything in the tools/bin directory
tools.clean:
	rm -fr $(TOOLS_DIR)/*

#help tools.get: retrieve all the tools specified in gex
tools.get:
	cd $(CURDIR)/tools && go generate tools.go

#help tools: clean then retrieve all the tools specified in gex
tools: tools.clean tools.get


############################
# Model generation targets #
############################
.PHONY: generate.models

#help db.generate.models: generate models for the database
generate.models:
	$(TOOLS_DIR)/gen --sqltype=postgres \
	--gorm --no-json --no-xml --overwrite --mapping postgresql/mapping.json --out postgresql/ \
	--connstr "postgresql://bucket_store_admin:azerty@localhost:5432/bucket_store?sslmode=disable" --model=models --database bucket_store

#help generate.mocks: Generate/update mocks for testing
generate.repo.mocks:
	$(TOOLS_DIR)/mockgen -source=$(CURDIR)/internal/repo/interfaces.go -destination=$(CURDIR)/internal/repo/mock/repo.go

#####################
# Build targets     #
#####################
.PHONY: build.clean build.vendor build.vendor.full build.swagger build.local build.mocks

TOOLS_DIR=$(CURDIR)/tools/bin

#help build.clean: prepare target/ folder
build.clean:
	@mkdir -p $(CURDIR)/target
	@rm -f $(CURDIR)/target/$(NAME)

#help build.vendor: retrieve all the dependencies used for the project
build.vendor:
	go mod vendor

#help build.vendor.full: retrieve all the dependencies after cleaning the go.sum
build.vendor.full:
	@rm -fr $(CURDIR)/vendor
	go mod tidy
	go mod vendor

#help build.swagger: generate REST api files from swagger.yaml
build.swagger: build.clean
	cp swagger.yaml $(CURDIR)/target/swagger.yaml
	sed "s/#VERSION#/$(VERSION)/g" -i $(CURDIR)/target/swagger.yaml
	$(TOOLS_DIR)/swagger generate server -f $(CURDIR)/target/swagger.yaml --name=$(NAME) -m restapi/models -s restapi/server --exclude-main --regenerate-configureapi

#help build.local: build locally a binary, in target/ folder
build.local: build.clean
	go build -mod=vendor $(BUILD_ARGS) -o $(CURDIR)/target/run $(CURDIR)/cli/main.go


#####################
# Run targets     #
#####################
.PHONY: run.local run.infra run.infra.detach run.infra.stop run.upload.object run.read.object run.delete.object

BODY=$(CURDIR)/resources/test/body.json

#help run.infra: start docker-compose 
run.infra:
	docker-compose -f $(CURDIR)/docker-compose.yaml up

#help run.infra.detach: start docker-compose on detached mode
run.infra.detach:
	docker-compose -f $(CURDIR)/docker-compose.yaml up -d

#help run.infra.stop: stop docker-compose
run.infra.stop:
	docker-compose -f $(CURDIR)/docker-compose.yaml down

#help run.local: run the application locally
run.local:
	@$(CURDIR)/target/run run -c resources/config/.$(NAME).yaml

#help run.upload.object: upload object data form resources/test/body.json
run.upload.object:
	curl -X PUT -i localhost:8080/objects/bucket00/42 -H 'Content-Type: application/json' -d "@$(BODY)"

run.read.object:
	curl -X GET -i localhost:8080/objects/bucket00/42 -H 'Content-Type: application/json'

run.delete.object:
	curl -X DELETE -i localhost:8080/objects/bucket00/42 -H 'Content-Type: application/json'

#####################
# Test targets     #
#####################
.PHONY: test

#help test: execute go tests
test:
	go test -mod=vendor ./...
