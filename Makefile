export GO111MODULE=on
export GOPROXY=https://goproxy.s.o3.ru
export GOSUMDB=off
LOCAL_BIN:=$(CURDIR)/bin
MIMIR_CFG="mimir.yaml"
MIMIR_BIN:=$(LOCAL_BIN)/mimir-cli
BUF_BIN:=$(LOCAL_BIN)/buf

##################### GOLANG-CI RELATED CHECKS #####################
# Check global GOLANGCI-LINT
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
GOLANGCI_TAG:=1.58.0

# Check local bin version
# Check local bin version
ifneq ($(wildcard $(GOLANGCI_BIN)),)
GOLANGCI_BIN_VERSION:=$(shell $(GOLANGCI_BIN) --version)
ifneq ($(GOLANGCI_BIN_VERSION),)
GOLANGCI_BIN_VERSION_SHORT:=$(shell echo "$(GOLANGCI_BIN_VERSION)" | sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_BIN_VERSION_SHORT:=0
endif
ifneq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_BIN_VERSION_SHORT)))"
GOLANGCI_BIN:=
endif
endif

# Check global bin version
ifneq (, $(shell which golangci-lint))
GOLANGCI_VERSION:=$(shell golangci-lint --version 2> /dev/null )
ifneq ($(GOLANGCI_VERSION),)
GOLANGCI_VERSION_SHORT:=$(shell echo "$(GOLANGCI_VERSION)"|sed -E 's/.* version (.*) built from .* on .*/\1/g')
else
GOLANGCI_VERSION_SHORT:=0
endif
ifeq "$(GOLANGCI_TAG)" "$(word 1, $(sort $(GOLANGCI_TAG) $(GOLANGCI_VERSION_SHORT)))"
GOLANGCI_BIN:=$(shell which golangci-lint)
endif
endif
##################### GOLANG-CI RELATED CHECKS #####################

## run full lint like in pipeline
.PHONY: lint
lint: install-lint
	$(GOLANGCI_BIN) run --config=.golangci.pipeline.yaml ./...


.PHONY: install-lint
install-lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info v$(GOLANGCI_TAG))  # Downloading golangci-lint
	tmp=$$(mktemp -d) && cd $$tmp && pwd && go mod init temp && go get -d github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG) && \
		go build -ldflags "-X 'main.version=$(GOLANGCI_TAG)' -X 'main.commit=test' -X 'main.date=test'" -o $(LOCAL_BIN)/golangci-lint github.com/golangci/golangci-lint/cmd/golangci-lint && \
		rm -rf $$tmp
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: update
update:
	go get -u gitlab.ozon.ru/asinyaev/allure-testify gitlab.ozon.ru/marketplace/qa/tools/qa_pack

.PHONY: install-lint
install-lint: ## install golangci-lint binary
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	$(info Downloading golangci-lint v$(GOLANGCI_TAG))
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)
GOLANGCI_BIN:=$(LOCAL_BIN)/golangci-lint
endif

.PHONY: lint
lint: .lint ## run golangci-lint only for files that differ from master


.PHONY: .lint
.lint: install-lint
	$(info Running lint...)
	$(GOLANGCI_BIN) run --new-from-rev=origin/master --config=.golangci.pipeline.yaml ./...

bin-deps:
	mkdir -p $(LOCAL_BIN)

	ls $(LOCAL_BIN)/protoc-gen-go &> /dev/null || \
        GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1

	ls $(LOCAL_BIN)/protoc-gen-go-grpc &> /dev/null || \
        GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

	ls $(CURDIR)/bin/mimir-cli &> /dev/null || \
        GOBIN=$(LOCAL_BIN) go install gitlab.ozon.ru/platform/mimir-cli/cmd/mimir-cli@latest

	ls $(CURDIR)/bin/buf &> /dev/null || \
        GOBIN=$(LOCAL_BIN) go install github.com/bufbuild/buf/cmd/buf@latest

.PHONY: .deps-pb
.deps-pb:
	$(info Install proto dependencies...)
	rm -rf $(CURDIR)/vendor.protogen
	$(MIMIR_BIN) vendor --config $(MIMIR_CFG)

.PHONY: deps-pb
deps-pb: .deps-pb

.PHONY: .generate
generate: bin-deps deps-pb .generate

# generate code from proto
.PHONY: generate
.generate:
	@[ -f "buf.gen.yaml" ] || (echo "ERROR: buf.gen.yaml not found. Run 'scratch update --confirm' to fix." &&  exit 1)
	$(info Generating code...)
		@# С помощью Mimir запускаем генерацию кода на основе proto-файлов.
			$(info Generating code based on .proto files...)
	$(MIMIR_BIN) generate \
	$(MIMIR_GEN_ARGS) \
	--config $(MIMIR_CFG) \
	--buf-bin $(BUF_BIN)
	@# Удаляем *.pb.scratch.go во избежание поломки кода
	@# на этапе генерации, связанной с переездом с esc на embed.
	$(info Removing *.pb.scratch.go files...)
	find $(CURDIR) \
	-type f \
	-name "*.pb.scratch.go" \
	-exec dirname {} \; | \
	PATH="$(LOCAL_BIN):$(PATH)" \
	xargs -I {} bash -c 'cd {} && go generate -run "(scratch|esc)" ./...'
