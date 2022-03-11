package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("----Use Environment Variable in your next Golang Project----")

	// 3 different ways to access environment variables
	// 1. os package
	// 2. godotenv package
	// 3. viper package

	// 1.First using os package
	// using_os_package()

	// 2.Second using godotenv package
	// using_godotenv_package()

	// 3.Third using viper package
	using_viper_package()
}

func using_os_package() {
	os.Setenv("Name", "Prince")
	value := os.Getenv("Name")
	fmt.Println("Name is", value)
	// if we type in what inside these square bracket [APP_ENV=test go run main.go] in the terminal then os set `APP_ENV` as environment variable and we can access his value below like os.Getenv("APP_ENV")
	fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))
}

func using_godotenv_package() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	s3Bucket := os.Getenv("S3_BUCKET")
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Println("S3_BUCKET:", s3Bucket)
	fmt.Println("SECRET_KEY:", secretKey)
}

func using_viper_package() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	value, ok := viper.Get("S3_BUCKET").(string)
	if !ok {
		log.Fatal("Invalid type assertion")
	}
	fmt.Println(value)
	// err = viper.ReadConfig(os.Stdin)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(viper.Get("NAME"))

	// -----------config.yaml-----------
	// 1.way
	// viper.SetConfigName("config")
	// viper.AddConfigPath(".")

	// 2.way
	viper.SetConfigFile("config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	value, ok = viper.Get("I_AM_INEVITABLE").(string)
	if !ok {
		log.Fatal("Invalid type assertion")
	}
	fmt.Println(value)
}
