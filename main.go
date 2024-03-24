package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("RSS Feed")

	godotenv.Load()
	
	portString := os.Getenv("PORT")
	usernameString := os.Getenv("USERNAME")
	fmt.Println(portString, usernameString)
}