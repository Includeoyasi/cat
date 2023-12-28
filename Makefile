PHONY: proto
proto:
	mkdir -p pkg/cat
	protoc --go_out=pkg/cat --go_opt=paths=import \
			--go-grpc_out=pkg/cat --go-grpc_opt=paths=import \
			api/cat.proto
	mv pkg/cat/github.com/Includeoyasi/cat/pkg/grpc/* pkg/cat
	rm -rf pkg/cat/github.com/