APP     := stackr
CMD     := ./cmd/stackr
VERSION := $(shell grep 'Version = ' internal/config/config.go | sed 's/.*"\(.*\)".*/\1/')
LDFLAGS := -ldflags="-s -w"

.PHONY: build install clean release

build:
	@echo "→ Compilando stackr v$(VERSION)..."
	go build $(LDFLAGS) -o $(APP) $(CMD)
	@echo "✔ ./$(APP)"

install: build
	sudo mv $(APP) /usr/local/bin/$(APP)
	@echo "✔ Instalado em /usr/local/bin/$(APP)"

clean:
	rm -f $(APP) $(APP)-*

release:
	@echo "→ Release v$(VERSION)"
	@GOOS=linux   GOARCH=amd64  go build $(LDFLAGS) -o $(APP)-linux-amd64   $(CMD) && echo "  ✔ linux/amd64"
	@GOOS=linux   GOARCH=arm64  go build $(LDFLAGS) -o $(APP)-linux-arm64   $(CMD) && echo "  ✔ linux/arm64"
	@GOOS=darwin  GOARCH=amd64  go build $(LDFLAGS) -o $(APP)-darwin-amd64  $(CMD) && echo "  ✔ darwin/amd64"
	@GOOS=darwin  GOARCH=arm64  go build $(LDFLAGS) -o $(APP)-darwin-arm64  $(CMD) && echo "  ✔ darwin/arm64"
	@GOOS=windows GOARCH=amd64  go build $(LDFLAGS) -o $(APP)-windows.exe   $(CMD) && echo "  ✔ windows/amd64"
