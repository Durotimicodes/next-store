package main

import (
	"github.com/decadevs/next_store/database"
	"github.com/decadevs/next_store/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/driver/mysql"
	_ "net/http"
)

func main() {

	//call routes /sever
	routes.CallRoutes("port")

	database.DB()
	////Database connection
	//db, err := gorm.Open("mysql", "root:OluwaTimi30@tcp(127.0.0.1:3306)/e-commerce_db?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	//DEPLOYING STATIC FILES
	//fs := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fs)

}
