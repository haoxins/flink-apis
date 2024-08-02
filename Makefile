LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p $(LOCALBIN)

CONTROLLER_GEN ?= $(LOCALBIN)/controller-gen
CONTROLLER_TOOLS_VERSION ?= v0.15.0 # https://github.com/kubernetes-sigs/controller-tools/tags

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN)
$(CONTROLLER_GEN):
	GOBIN=$(LOCALBIN) go install sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_TOOLS_VERSION)

.PHONY: gen
gen: controller-gen
	$(CONTROLLER_GEN) object paths="./..."

all: gen
	go mod tidy
	go fmt ./...
	go vet ./...
	go test -v --cover ./...
