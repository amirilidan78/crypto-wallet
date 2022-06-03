package httpServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	DefaultAddr  = "0.0.0.0:"
	ReadTimeout  = 10 * time.Second
	WriteTimeout = 10 * time.Second
)

type HttpServer interface {
	ListenAndServe() error
	GetEngine() *gin.Engine
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type httpServer struct {
	server *http.Server
	engine *gin.Engine
}

func (s *httpServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func (s *httpServer) GetEngine() *gin.Engine {
	return s.engine
}

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			response := Response{
				Code:    http.StatusNotFound,
				Message: "404 not found",
			}
			c.AbortWithStatusJSON(http.StatusNotFound, response)
		}
	}(c)
	c.Next()
}

func NewHttpServer(port string) HttpServer {

	r := gin.New()

	r.Use(globalRecover)

	gin.SetMode(gin.ReleaseMode)

	server := &http.Server{
		Addr:           DefaultAddr + port,
		Handler:        r,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s := httpServer{server, r}

	return &s
}
