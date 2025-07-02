package api

import (
	db "bobbabank/internal/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Сервер служит всем HTTP запросам для данного банк-сервиса
type Server struct {
	store db.Store
	// Помогает отправлять HTTP req к нужным handler'ам
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to a router
	router.POST("/accounts", server.createAccount)
	router.POST("/accounts/:id", server.deleteAccount)
	router.POST("/accounts/update/:id", server.updateAccount)
	// так Gin понимает, что id это URL парметр
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResp(err error) gin.H {
	return gin.H{"error": err.Error()}
}
