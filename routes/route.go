package routes

import (
	"github.com/decadevs/next_store/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
)

//THE ROUTING IS USED TO HANDLE VARIOUS URL
func CallRoutes(port string, db *gorm.DB) {
	//set route as default one made by Gin
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//serve the static files
	router.StaticFS("static", http.Dir("./templates/static"))

	//sever all the HTML template quickly as soon as the pages load
	router.LoadHTMLGlob("templates/*.html")

	//Welcome page router
	router.GET("/", handlers.WelcomepageHandler)

	//Market place router
	router.GET("/marketplace", handlers.MarketPlaceHandler)

	//buyer page router
	router.GET("/addtocart/:id", handlers.AddToCartHandler)

	//Buyer Page
	router.GET("/buyerremoveproduct/:id", handlers.RemoveProductFromCartHandler)

	//Account Details
	router.GET("/buyer/accountdetails", handlers.PaymentHandler)

	//Buyer SignUp router
	router.GET("/buyerssignup", handlers.BuyerSignUpPageHandler)

	//Seller Edit Product
	router.GET("/sellers/editPost/:id", handlers.SellerEditProductHandler)

	//Seller Update Product
	router.POST("/update-product/:id", handlers.SellerUpdateProductHandler)

	//Admin Post Product
	router.POST("/sellers/addproducts", handlers.AdminPostProductHandler)

	//Buyer Page
	router.GET("/buyerpage", handlers.BuyerPageHandler)

	//Admin Get Product
	router.GET("/sellers/addproductspage", handlers.AdminGetProductHandler)

	//Admin Delete Product
	router.GET("/sellers/deleteproduct/:id", handlers.AdminDeleteProductHandler)

	//Admin Launch Product to Market Place
	router.POST("/seller/postproduct", handlers.AdminPostInMarketHandler)
	router.GET("/sellers/launchproduct", handlers.AdminLaunchProductHandler)

	//SIGN UP AND LOGIN
	router.GET("/sellersignup", handlers.SellerLoginPageHandler)

	//Seller page router
	router.GET("/sellerpage", handlers.SellerPageHandler)

	//Seller Login
	router.POST("/sellers/signin", handlers.SellerLoginHandler)

	//Buyer Sig-up
	router.POST("/buyers/signup", handlers.BuyerSignUpHandler)
	//Buyer Login-in
	router.POST("/buyers/login", handlers.LoginHandler)

	//To search for product
	router.GET("/searchproduct", handlers.SearchProduct)

	//Adminstrator's Dashboard
	router.GET("/admindashboard", handlers.AdminDashBoard)

	//start and run the server on port 8084
	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8084"
	}
	router.Run(port)
}
