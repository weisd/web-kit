package main

import (

	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"os"

	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	// "github.com/gogo/grpc-example/server"

	// Static files gen from third_party by statik
	_ "github.com/weisd/web-kit/statik"

	"github.com/weisd/web-kit/configs/insecure"

	"github.com/weisd/web-kit/api/protobuf/user"
	userServer "github.com/weisd/web-kit/internal/app/user"
)

var (
	gRPCPort    = flag.Int("grpc-port", 10000, "The gRPC server port")
	gatewayPort = flag.Int("gateway-port", 11000, "The gRPC-Gateway server port")
	deriver = flag.String("deriver","memory", "The  Date Model driver dsn incloud: menory,sql, default:memory")
	dsn = flag.String("dsn","", "The  Date Model driver dsn, eg: mysql:<username>:<password>@tcp(<localhost>)/<dbname>?charset=utf8mb4&parseTime=True&loc=Local")
)

var log grpclog.LoggerV2

func init() {
	log = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)


}

// serveOpenAPI serves an OpenAPI UI on /openapi-ui/
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func serveOpenAPI(mux *http.ServeMux) error {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	// Expose files in static on <host>/openapi-ui
	fileServer := http.FileServer(statikFS)
	prefix := "/openapi-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	return nil
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *gRPCPort)

	log.Infoln("model start with", *deriver, *dsn)


	userRPCServer, has := userServer.UserRPCServerRegister[*deriver]; 
	if !has {
		log.Fatalf("user server not register: %s\n", *deriver)
	}

	ctx, cancel := context.WithCancel(context.Background())

	if err := userRPCServer.Init(ctx, *dsn); err != nil {
		log.Fatalf("server %s init failed: %v", err)
	}


	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
		grpc.UnaryInterceptor(grpc_validator.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpc_validator.StreamServerInterceptor()),
	)

	// user rpc
	user.RegisterRPCServiceServer(s, userRPCServer)
	user.RegisterAPIServiceServer(s, userServer.NewAPIServer())

	// Serve gRPC Server
	log.Info("Serving gRPC on https://", addr)
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// See https://github.com/grpc/grpc/blob/master/doc/naming.md
	// for gRPC naming standard information.
	dialAddr := fmt.Sprintf("passthrough://localhost/%s", addr)
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	mux := http.NewServeMux()

	jsonpb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonpb),
		// This is necessary to get error details properly
		// marshalled in unary requests.
		runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
	)

	// user rpc gw
	err = user.RegisterRPCServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// user api gw
	err = user.RegisterAPIServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	mux.Handle("/", gwmux)
	err = serveOpenAPI(mux)
	if err != nil {
		log.Fatalln("Failed to serve OpenAPI UI")
	}

	gatewayAddr := fmt.Sprintf("localhost:%d", *gatewayPort)
	log.Info("Serving gRPC-Gateway on https://", gatewayAddr)
	log.Info("Serving OpenAPI Documentation on https://", gatewayAddr, "/openapi-ui/")
	gwServer := http.Server{
		Addr: gatewayAddr,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{insecure.Cert},
		},
		Handler: mux,
	}
	log.Fatalln(gwServer.ListenAndServeTLS("", ""))

	cancel()
}
