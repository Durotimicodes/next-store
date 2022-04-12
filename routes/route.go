package routes

import (
	"github.com/decadevs/next_store/handlers"
	"github.com/gin-gonic/gin"
	"os"
)

func CallRoutes(port string) {
	//create a new gin router
	router := gin.Default()

	//define a single homepage endpoint
	router.GET("/", handlers.Welcomepage)
	//
	////define crud endpoints for sellers & buyers
	//
	////create
	//router.POST("/buyers/signup", handlers.SignUpHandler)
	//
	//router.POST("/buyers/login", handlers.LoginHandler)
	//
	//router.POST("/sellers/addproducts", handlers.AdminAddProductHandler)
	//
	//// retrieve
	//router.GET("/buyers/viewproducts", handlers.ViewProductsHandler)
	//
	//// update
	//router.GET("/sellers/getproductsbyId", handlers.GetProductsByIdHandler)
	//
	//router.GET("/buyers/returnproducts", handlers.ReturnProductsHandler)
	//
	//router.POST("/sellers/udateproducts", handlers.UpdateProductsHandler)
	//
	//// delete
	//
	//router.GET("/sellers/deleteproduct", handlers.DeleteProducthandler)

	//run the router
	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8082"
	}
	router.Run(port)
}