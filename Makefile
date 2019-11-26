PROGRAM := growser

CUR_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CMD_DIR := $(CUR_DIR)/cmd
BIN_DIR := $(CUR_DIR)/bin

growser:
	@echo $(CUR_DIR)
	@go build -o $(BIN_DIR)/$(PROGRAM) $(CMD_DIR)/$(PROGRAM)/main.go

.PHONY: growser

clean:
	@rm -f $(BIN_DIR)/$(PROGRAM)

.PHONY: clean