package database

import (
	"fmt"
	"github.com/decadevs/next_store/models"
	_ "github.com/decadevs/next_store/models"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

//DECLARE A VARIABLE THAT CONNECTS WITH DB
var DBClient *gorm.DB

//FUNCTION TO OPEN AND MIGRATE
func OpenAndMigrateDb() (*gorm.DB, error) {

	//storing values from .env file into the variables
	var username = os.Getenv("DB_USERNAME")
	var password = os.Getenv("DB_PASSWORD")
	var dbname = os.Getenv("DB_NAME")

	//open Database and connect using user details
	db, err := gorm.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/"+dbname+"?charset=utf8mb4&parseTime=True&loc=Local")
	//error handling
	if err != nil {
		log.Println("checking database error", err)
	}

	DBClient = db

	//calling the automigrate function
	AutoMigrate(db)

	return DBClient, nil

}

//FUNCTION FOR AUTOMIGRATION USING GORM
func AutoMigrate(db *gorm.DB) {

	err := db.AutoMigrate(&models.User{},
		&models.Buyer{},
		&models.Order{},
		&models.Product{},
		&models.Status{},
		&models.Seller{},
		&models.Cart{},
	)

	log.Println("checking database error", err)
}

//FUNCTION TO CHECK THE DB & FIND USERS BY THEIR EMAIL ADDRESS
func FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	var err = DBClient.Where("email = ?", email).First(user).Error
	log.Println("second line")
	if err != nil {
		return nil, err
	}
	return user, nil
}

//FUNCTION TO FIND USER ID BY EMAIL
func FindUserIdByEmail(email string) (uint, error) {
	user := &models.User{}
	var err = DBClient.Where("email = ?", email).First(user).Error
	log.Println("second line")
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

//FUNCTION TO FIND SELLER BY EMAIL
func FindSellerByEmail(email string) (*models.Seller, error) {
	seller := &models.Seller{}
	var err = DBClient.Where("email = ?", email).First(seller).Error
	log.Println("second line")
	if err != nil {
		return nil, err
	}
	return seller, nil
}

//FUNCTION TO CREATE NEW USER
func CreateNewUser(user *models.User) error {
	err := DBClient.Create(user).Error
	return err
}

//CREATE A IN-MEMORY SELLER DATABASE-RECORD
func SellerDB() *gorm.DB {
	db, err := OpenAndMigrateDb()
	//error handling
	if err != nil {
		log.Println("checking database error", err)
	}
	defer db.Close()

	//Seller details
	UserOne := models.User{
		ID:            1,
		Name:          os.Getenv("SELLER_NAME"),
		Email:         os.Getenv("SELLER_EMAIL"),
		Username:      os.Getenv("SELLER_USERNAME"),
		Password:      os.Getenv("SELLER_PASSWORD"),
		Address:       os.Getenv("SELLER_ADDRESS"),
		AccountName:   os.Getenv("SELLER_ACCOUNTNAME"),
		AccountNumber: os.Getenv("SELLER_ACCOUNT_NUMBER"),
		Phonenumber:   os.Getenv("SELLER_PHONENUMBER"),
	}

	//CREATE SELLER
	var Seller = models.Seller{
		UserOne,
		1,
	}

	result := db.Create(&Seller)
	err = result.Error
	if err != nil {
		log.Printf("Error creating seller record : %v", err)
	}
	rowsAffected := result.RowsAffected
	fmt.Println("Number of rows affected", rowsAffected)

	return db
}
