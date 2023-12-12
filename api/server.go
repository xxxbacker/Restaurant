package api

import (
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	server.router = router

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", server.signUp)
		auth.POST("/sign-in", server.signIn)
	}

	api := router.Group("/api", server.accountIdentity)
	{
		admin := api.Group("/admin")
		{
			// вот здесь я просто накидал для примера
			admin.POST("/ListMenu", server.listMenuItem)
		}
		user := api.Group("/user")
		{
			// вот здесь я просто накидал для примера
			user.POST("/ListMenu", server.listMenuItem)
		}
	}

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) *gin.H {
	return &gin.H{"err": err.Error()}
}
