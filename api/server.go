package api

import (
	"github.com/gin-gonic/gin"
	dbbank "github.com/k2madureira/gobank/db/sqlc"
)

type Server struct {
	store  dbbank.Store
	router *gin.Engine
}

func NewServer(store dbbank.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
