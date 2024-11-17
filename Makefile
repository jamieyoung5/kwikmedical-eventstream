PROTO_DIR := .
OUT_DIR := ./gen

# Go code generation
generate-go:
	mkdir -p $(OUT_DIR)
	protoc --go_out=$(OUT_DIR) --go-grpc_out=$(OUT_DIR) \
	       $(PROTO_DIR)/**/*.proto
