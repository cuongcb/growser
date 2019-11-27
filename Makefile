PROGRAM := growser

CUR_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CMD_DIR := $(CUR_DIR)/cmd
BIN_DIR := $(CUR_DIR)/bin
PKG_DIR := $(CUR_DIR)/pkg

BUILD_DIR := $(CUR_DIR)/build
OUT_BUILD_DIR := $(BUILD_DIR)/gen
BIN_BUILD_DIR := $(BUILD_DIR)/bin
SRC_BUILD_DIR := $(BUILD_DIR)/src

DEV_DIR := $(CUR_DIR)/dev
SCRIPT_DIR := $(DEV_DIR)/script

PROTOC := $(BIN_BUILD_DIR)/proto/bin/protoc

SRC_PROTO_DIR := $(PKG_DIR)/proto

growser:
	@go build -o $(BIN_DIR)/$(PROGRAM) $(CMD_DIR)/$(PROGRAM)/main.go

.PHONY: growser

clean:
	@rm -f $(BIN_DIR)/$(PROGRAM)

.PHONY: clean

proto:
	cd $(SCRIPT_DIR) && python3 install_proto.py $(SRC_BUILD_DIR) $(BIN_BUILD_DIR)/proto
	@go get -u github.com/golang/protobuf/protoc-gen-go

proto-clean:
	cd $(SCRIPT_DIR) && python3 install_proto.py $(SRC_BUILD_DIR) $(BIN_BUILD_DIR)/proto clean

proto-gen:
	$(PROTOC) -I=$(SRC_PROTO_DIR) --go_out=$(SRC_PROTO_DIR) $(SRC_PROTO_DIR)/project.proto

.PHONY: proto proto-clean proto-gen