# note: call scripts from /scripts
gen:
	./scripts/protoc.sh

gb:
	./scripts/build.sh

up:
	./bin/swagger

install:
	go install \
		./vendor/github.com/gogo/protobuf/protoc-gen-gogo \
		./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		./vendor/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
		./vendor/github.com/rakyll/statik