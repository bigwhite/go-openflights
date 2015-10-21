package main

import (
	"golang.org/x/net/context"

	"go.pedge.io/env"
	"go.pedge.io/openflights"
	"go.pedge.io/proto/server"
	"go.pedge.io/protolog"

	"github.com/gengo/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var (
	defaultEnv = map[string]string{
		"PORT":      "1747",
		"HTTP_PORT": "8080",
	}
)

type appEnv struct {
	Port      int  `env:"PORT"`
	HTTPPort  int  `env:"HTTP_PORT"`
	DebugPort int  `env:"DEBUG_PORT"`
	DebugLog  bool `env:"DEBUG_LOG"`
}

func main() {
	env.Main(do, &appEnv{}, defaultEnv)
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	if appEnv.DebugLog {
		protolog.SetLevel(protolog.Level_LEVEL_DEBUG)
	}
	client, err := openflights.NewDefaultServerClient()
	if err != nil {
		return err
	}
	return protoserver.Serve(
		uint16(appEnv.Port),
		func(s *grpc.Server) {
			openflights.RegisterAPIServer(s, openflights.NewAPIServer(client))
		},
		protoserver.ServeOptions{
			HTTPPort:  uint16(appEnv.HTTPPort),
			DebugPort: uint16(appEnv.DebugPort),
			HTTPRegisterFunc: func(ctx context.Context, mux *runtime.ServeMux, clientConn *grpc.ClientConn) error {
				return openflights.RegisterAPIHandler(ctx, mux, clientConn)
			},
		},
	)
}
