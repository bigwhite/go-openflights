package main

import (
	"os"

	"go.pedge.io/env"
	"go.pedge.io/openflights"
	"go.pedge.io/pkg/log"
	"go.pedge.io/proto/server"

	"google.golang.org/grpc"
)

type appEnv struct {
	Port      uint16 `env:"PORT,default=1747"`
	HTTPPort  uint16 `env:"HTTP_PORT,default=8080"`
	DebugPort uint16 `env:"DEBUG_PORT"`
	LogEnv    pkglog.Env
}

func main() {
	env.Main(do, &appEnv{})
}

func do(appEnvObj interface{}) error {
	appEnv := appEnvObj.(*appEnv)
	pkglog.SetupLogging(os.Args[0], appEnv.LogEnv)
	client, err := openflights.NewDefaultServerClient()
	if err != nil {
		return err
	}
	return protoserver.Serve(
		appEnv.Port,
		func(s *grpc.Server) {
			openflights.RegisterAPIServer(s, openflights.NewAPIServer(client))
		},
		protoserver.ServeOptions{
			HTTPPort:         appEnv.HTTPPort,
			DebugPort:        appEnv.DebugPort,
			HTTPRegisterFunc: openflights.RegisterAPIHandler,
			Version:          openflights.Version,
		},
	)
}
