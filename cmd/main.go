package main

import (
	"clothing-store/config"
	"fmt"
)

func main() {

	config.ConnectDB()

	fmt.Println("Server 6688")
}
