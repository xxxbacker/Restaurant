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
			admin.GET("/getAccount/:id", server.getAccountForAdmin)
			admin.POST("/listAccount", server.listAccount)
			admin.POST("/createCourier", server.createCourier)
			admin.GET("/getCourier/:id", server.getCourier)
			admin.POST("/listCourier", server.listCourier)
			admin.GET("/getOrder/:id", server.getOrdForAdmin)
			admin.POST("/listOrder", server.listOrd)

		}
		user := api.Group("/user")
		{
			user.GET("/getAccount", server.getAccountForUser)
			user.POST("/createOrder", server.createOrd)
			user.GET("/getOrder", server.getOrdForUser)
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
