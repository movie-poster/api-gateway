package entry

import (
	"api-gateway/cmd/config"
	"api-gateway/internal/infra/grpc"
	"api-gateway/internal/utils"
	"flag"

	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
)

var router *echo.Echo

func init() {
	var configPath = ""
	configPath = *flag.String("config", "", "")

	if configPath == "" {
		configPath = "../data/config.yml"
	}

	setConfiguration(configPath)
}

func setConfiguration(configPath string) {
	config.Setup(configPath)
}

func Run() {
	utils.SetupLoggerZap()
	conf := config.GetConfig()
	// config.BugsnagClient()
	grpc.NewGrpcMovieConnection()
	grpc.NewGrpcUserConnection()

	router = echo.New()
	server := socketio.NewServer(nil)

	go HandlerProfiling(router, conf.Server.TestPort)
	RouterHandler(router, server)

	go server.Serve()
	defer server.Close()

	router.Start(":" + conf.Server.Port)
}
