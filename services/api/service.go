package api

import (
	"crypto-wallet/config"
	"crypto-wallet/pkg/httpServer"
	"crypto-wallet/services/api/handlers"
	"crypto-wallet/services/api/middlewares"
	"crypto-wallet/services/database"
	"crypto-wallet/services/wallet"
	"fmt"
)

type Api interface {
	Serve()
	registerRoutes(server httpServer.HttpServer)
	getServer() httpServer.HttpServer
}

type api struct {
	c  config.Config
	db database.Database
	w  wallet.Wallet
}

func (a *api) Serve() {

	server := a.getServer()

	a.registerRoutes(server)

	server.ListenAndServe()

	select {}
}

func (a *api) getServer() httpServer.HttpServer {

	port := a.c.GetString("server.port")

	return httpServer.NewHttpServer(port)
}

func (a *api) registerRoutes(server httpServer.HttpServer) {

	r := server.GetEngine()

	r.Use(middlewares.IpMiddleware(a.c))
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/address")
		{
			auth.POST("/new", handlers.NewAddress(a.db, a.w))
		}
	}

	fmt.Println("----- Registered routes -----")
	fmt.Println(r.Routes())
	fmt.Println("----- Registered routes -----")
}

func NewApiService(c config.Config, db database.Database, w wallet.Wallet) Api {
	return &api{c, db, w}
}
