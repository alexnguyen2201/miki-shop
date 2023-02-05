package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/nguyenvanson2201/miki-shop/db/sqlc"
	"github.com/nguyenvanson2201/miki-shop/token"
	"github.com/nguyenvanson2201/miki-shop/util"
)

// Server servers HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// for admin user
	router.POST("/users", server.createUser)

	//for normal user
	router.POST("/users/login", server.loginUser)

	// ProductType route
	router.POST("/product_types", server.createProductType)
	router.GET("/product_types", server.getProductTypeList)
	router.GET("/product_types/:id", server.getProductType)

	// Product route
	router.POST("/products", server.createProduct)
	router.GET("/products/:id", server.getProduct)

	// Auth route
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/:id", server.getUser)
	server.router = router
}

// Start run the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
