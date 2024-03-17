COVER_DIR=coverage
MKDIR := mkdir -p

test:
	go test ./...

pre-cover:
	@$(MKDIR) $(COVER_DIR)

test-cover: pre-cover
	go test -coverprofile=$(COVER_DIR)/cover.out ./...
	go tool cover -func=$(COVER_DIR)/cover.out