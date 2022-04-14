package routes

import (
	"github.com/decadevs/next_store/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

//THE ROUTING IS USED TO HANDLE VARIOUS URL
func CallRoutes(port string) {
	//set route as default one made by Gin
	router := gin.Default()

	//serve the static files
	router.StaticFS("static", http.Dir("./templates/static"))

	//sever all the HTML template quickly as soon as the pages load
	router.LoadHTMLGlob("templates/*.html")

	//Welcome page router
	router.GET("/", handlers.Welcomepage)
	//Market place router
	router.GET("/marketplace", handlers.MarketPlace)
	//buyer page router
	router.GET("/buyerspage", handlers.BuyerPage)
	//Buyer SignUp router
	router.GET("/buyerssignup", handlers.BuyerSignUp)
	//seller signin router
	router.GET("/sellersignup", handlers.SellerSignUp)
	//Seller page router
	router.GET("/sellerpage", handlers.SellerPage)

	//Buyer Edit Product
	router.GET("/editProduct", handlers.EditProduct)
	//Buyer Update Product
	router.POST("/updateProduct", handlers.UpdateProduct)

	router.POST("/addProduct", handlers.AddProductToMarket)
	//Search Products
	//router.GET("/searchProduct", handlers.SearchProduct)

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
