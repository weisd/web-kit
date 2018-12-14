# note: call scripts from /scripts
gen:
	./scripts/protoc.sh

gb:
	./scripts/build.sh

up:
	./bin/swagger

install:
	go install \
		./cmd/github.com/gogo/protobuf/protoc-gen-gogo \
		./cmd/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		./cmd/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		./cmd/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		./cmd/github.com/rakyll/statik