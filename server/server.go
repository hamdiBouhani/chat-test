package server

import (
	"chat-test/internal"
	"flag"
	"fmt"
	"path/filepath"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server server instance
type HttpService struct {
	Log *logrus.Logger

	router     *gin.Engine
	ns         *gin.RouterGroup
	ws         *gin.RouterGroup
	port       string
	corsConfig cors.Config

	hub *Hub
}

// NewServer creates and initializes a new server.
func NewServer() (*HttpService, error) {

	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	//create new websocket hub
	hub := newHub()

	return &HttpService{
		Log: logger,
		hub: hub,
	}, nil
}

// Run runs the server
func (svc *HttpService) Run() error {
	//run hub
	go svc.hub.run()

	var port = flag.String("port", "8080", "Http port to serve application on")
	flag.Parse()
	svc.port = fmt.Sprintf(":%s", *port)

	svc.corsConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	err := svc.registerRoutes()
	if err != nil {
		return err
	}

	return svc.router.Run(svc.port) //Blocks
}

func (svc *HttpService) registerRoutes() error {

	svc.router = gin.Default()
	svc.router.Use(cors.New(svc.corsConfig))

	//Define normal routes
	svc.ns = svc.router.Group("/")
	svc.ws = svc.router.Group("/ws")
	svc.ns.GET("/ping", svc.ping) //PING/PONG

	svc.wsRoutes()

	return nil
}

func (svc *HttpService) ping(c *gin.Context) {
	c.JSON(200, "pong")
}

func (svc *HttpService) wsRoutes() {
	svc.router.LoadHTMLFiles(filepath.Join(internal.GetProjectPath(), "server", "view", "chat_client.html"))
	svc.ws.GET("/test", func(c *gin.Context) { c.HTML(200, "chat_client.html", nil) })
	svc.ws.GET("/chat-server", svc.getChatServer)
}
