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

	//server all the HTML template quickly as soon as the pages load
	router.LoadHTMLGlob("templates/*.html")

	// CREATE END POINTS
	//Seller Login
	router.POST("/seller/signin", handlers.SellerLoginHandler) //
	//Buyer Sig-up
	router.POST("/buyer/signup", handlers.BuyerSignUpHandler) //
	//Buyer Login-in
	router.POST("/buyer/login", handlers.LoginHandler) //
	//buyer page router
	router.GET("/buyer/addproducttocart/:id", handlers.AddToCartHandler) //
	//Admin Post Product
	router.POST("/seller/addproducts", handlers.AdminPostProductHandler) //
	//Admin Launch Product to Market Place
	router.POST("/seller/postproduct", handlers.AdminPostInMarketHandler) //

	// RETRIEVE END POINTS
	//Welcome page router
	router.GET("/", handlers.WelcomepageHandler) //
	//Market place router
	router.GET("/seller/marketplace", handlers.MarketPlaceHandler) //
	//Account Details
	router.GET("/buyer/paymentdetail", handlers.PaymentHandler) //
	//Router To Render Buyer SignUp Page
	router.GET("/buyer/signup", handlers.BuyerSignUpPageHandler) //
	//Seller Edit Product
	router.GET("/seller/editPost/:id", handlers.SellerEditProductHandler) //
	//Buyer Page
	router.GET("/buyer/cartpage", handlers.BuyerPageHandler)
	//Admin Get Product
	router.GET("/seller/addproductspage", handlers.AdminGetProductHandler)  //
	router.GET("/seller/launchproduct", handlers.AdminLaunchProductHandler) //
	//SIGN UP AND LOGIN
	router.GET("/seller/signup", handlers.SellerLoginPageHandler)
	//Seller page router
	router.GET("/sellerpage", handlers.SellerPageHandler)
	//To search for product
	router.GET("/buyer/searchproduct", handlers.SearchProduct)
	//Adminstrator's Dashboard
	router.GET("/seller/dashboard", handlers.AdminDashBoard)

	// UPDATE END POINTS
	//Seller Update Product
	router.POST("/seller/update-product/:id", handlers.SellerUpdateProductHandler)

	// DELETE END POINTS
	//Admin Delete Product
	router.GET("/seller/deleteproduct/:id", handlers.AdminDeleteProductHandler)
	//Buyer Page
	router.GET("/buyer/removeproduct/:id", handlers.RemoveProductFromCartHandler)
	// logOut LogoutUser
	router.GET("/logout", handlers.LogoutUserHandler)

	////Welcome page router
	//router.GET("/", handlers.WelcomepageHandler)
	//
	////Market place router
	//router.GET("/marketplace", handlers.MarketPlaceHandler)
	//
	////buyer page router
	//router.GET("/addtocart/:id", handlers.AddToCartHandler)
	//
	////Buyer Page
	//router.GET("/buyerremoveproduct/:id", handlers.RemoveProductFromCartHandler)
	//
	////Account Details
	//router.GET("/buyer/accountdetails", handlers.PaymentHandler)
	//
	////Buyer SignUp router
	//router.GET("/buyerssignup", handlers.BuyerSignUpPageHandler)
	//
	////Seller Edit Product
	//router.GET("/sellers/editPost/:id", handlers.SellerEditProductHandler)
	//
	////Seller Update Product
	//router.POST("/update-product/:id", handlers.SellerUpdateProductHandler)
	//
	////Admin Post Product
	//router.POST("/sellers/addproducts", handlers.AdminPostProductHandler)
	//
	////Buyer Page
	//router.GET("/buyerpage", handlers.BuyerPageHandler)
	//
	////Admin Get Product
	//router.GET("/sellers/addproductspage", handlers.AdminGetProductHandler)
	//
	////Admin Delete Product
	//router.GET("/sellers/deleteproduct/:id", handlers.AdminDeleteProductHandler)
	//
	////Admin Launch Product to Market Place
	//router.POST("/seller/postproduct", handlers.AdminPostInMarketHandler)
	//router.GET("/sellers/launchproduct", handlers.AdminLaunchProductHandler)
	//
	////SIGN UP AND LOGIN
	//router.GET("/sellersignup", handlers.SellerLoginPageHandler)
	//
	//// logOut LogoutUser
	//router.GET("/logout", handlers.LogoutUserHandler)
	//
	////Seller page router
	//router.GET("/sellerpage", handlers.SellerPageHandler)
	//
	////Seller Login
	//router.POST("/sellers/signin", handlers.SellerLoginHandler)
	//
	////Buyer Sig-up
	//router.POST("/buyers/signup", handlers.BuyerSignUpHandler)
	////Buyer Login-in
	//router.POST("/buyers/login", handlers.LoginHandler)
	//
	////To search for product
	//router.GET("/searchproduct", handlers.SearchProduct)
	//
	////Adminstrator's Dashboard
	//router.GET("/admindashboard", handlers.AdminDashBoard)

	//start and run the server on port 8084
	port = ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8094"
	}
	router.Run(port)
}
