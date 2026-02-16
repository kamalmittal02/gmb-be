package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
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
	//config, err := config.InitConfig(cmdArgs.configPath)
	//if err != nil {
	//	fmt.Printf("Failed to load config: %s \n", err.Error())
	//	panic("Failed to load config")
	//}
	di := injector.InitInjector()
	return &BaseApp{
		//Config: config,
		di: di,
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

	// CORS middleware - must be before routes
	app.Use(func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	SetupRouter(app, ba.Config, ba.di)
	httpServer := NewHttpServer(":8080", app)
	if err := httpServer.Serve(); err != nil {
		fmt.Printf("HTTP server stopped with error: %s", err.Error())
	}
}
