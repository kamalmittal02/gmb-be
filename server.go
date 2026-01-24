package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/config"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/dtos"
	"github.com/kamalmittal01/girraj-sweet-showcase-BE/injector"
	"net/http"
)

type HttpServer struct { //implements shutdown.Callback by implementing close()
	httpServer   *http.Server
	closeErrChan chan error
}

type BaseApp struct {
	Config *dtos.Config
	di     *injector.Injector
}

type CmdArgs struct {
	configPath string
}

func parseArgs() *CmdArgs {
	var configFlag = flag.String("conf", "config/config.yaml", "config/config.yaml")
	flag.Parse()
	return &CmdArgs{
		configPath: *configFlag,
	}
}

func NewBaseApp() *BaseApp {
	cmdArgs := parseArgs()
	fmt.Printf("Parsing config: %s \n", cmdArgs.configPath)
	config, err := config.InitConfig(cmdArgs.configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %s \n", err.Error())
		panic("Failed to load config")
	}
	di := injector.InitInjector(config)
	return &BaseApp{
		Config: config,
		di:     di,
	}
}

func NewHttpServer(address string, app *gin.Engine) *HttpServer {
	h := &HttpServer{
		httpServer: &http.Server{
			Addr:    address,
			Handler: app,
		},
		closeErrChan: make(chan error),
	}

	return h
}

func (h *HttpServer) Serve() error {
	if err := h.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}
	return <-h.closeErrChan
}

func (ba *BaseApp) RunHttpServer() {
	app := gin.New()
	app.Use(gin.Recovery())

	SetupRouter(app, ba.Config, ba.di)
	httpServer := NewHttpServer(":8080", app)
	if err := httpServer.Serve(); err != nil {
		fmt.Printf("HTTP server stopped with error: %s", err.Error())
	}
}
