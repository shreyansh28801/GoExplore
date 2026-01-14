PROTO_DIR=proto
PROTO_FILES=$(PROTO_DIR)/coffee_shop.proto

.PHONY: proto

proto:
	protoc \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		$(PROTO_FILES)
