package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/simple_bank_course/db/sqlc"
)

// Сервер служит всем HTTP запросам для данного банк-сервиса
type Server struct {
	store *db.Store
	// Помогает отправлять HTTP req к нужным handler'ам
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to a router
	router.POST("/accounts", server.createAccount)
	router.POST("/accounts/:id", server.deleteAccount)
	router.POST("/accounts/update/:id", server.updateAccount)
	// так Gin понимает, что id это URL парметр
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResp(err error) gin.H {
	return gin.H{"error": err.Error()}
}
