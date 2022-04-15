package main

import (
	"fmt"
	"github.com/decadevs/next_store/database"
	_ "github.com/decadevs/next_store/models"
	"github.com/decadevs/next_store/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/driver/mysql"
	_ "net/http"
)

func main() {
	//database connection
	database.DB()
	fmt.Println("Did you get here?")
	//seller's in-memory data
	database.SellerDB()
	fmt.Println("yes i did!")

	//call routes /sever
	routes.CallRoutes("port")

}
