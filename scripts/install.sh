go build -v -mod=vendor  -o $GOPATH/bin/protoc-gen-gogo cmd/github.com/gogo/protobuf/protoc-gen-gogo/main.go
go build -v -mod=vendor  -o $GOPATH/bin/protoc-gen-swagger cmd/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/main.go
go build -v -mod=vendor  -o $GOPATH/bin/protoc-gen-govalidators cmd/github.com/mwitkow/go-proto-validators/protoc-gen-govalidators/main.go
go build -v -mod=vendor  -o $GOPATH/bin/statik cmd/github.com/rakyll/statik/main.go