protoc --go_out=../server/grpc --go_opt=paths=source_relative \
	--go-grpc_out=../server/grpc --go-grpc_opt=paths=source_relative \
	--go-grpc_opt require_unimplemented_servers=false \
	./v1/healthcheck.proto