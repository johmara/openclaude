VERSION ?= dev
MODULE  := github.com/johmara/openclaude
LDFLAGS := -s -w -X $(MODULE)/cmd.version=$(VERSION)
BINARY  := openclaude
DIST    := dist

PLATFORMS := linux/amd64 linux/arm64 darwin/amd64 darwin/arm64 windows/amd64 windows/arm64

.PHONY: build vet test dist clean

build:
	go build -ldflags "$(LDFLAGS)" -o $(BINARY) .

vet:
	go vet ./...

test:
	go test -v ./...

dist: clean
	@mkdir -p $(DIST)
	@for platform in $(PLATFORMS); do \
		os=$${platform%/*}; \
		arch=$${platform#*/}; \
		ext=""; \
		if [ "$$os" = "windows" ]; then ext=".exe"; fi; \
		output="$(DIST)/$(BINARY)-$$os-$$arch$$ext"; \
		echo "Building $$output..."; \
		GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build -ldflags "$(LDFLAGS)" -o $$output . ; \
	done

clean:
	rm -rf $(DIST)
