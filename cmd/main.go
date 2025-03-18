package main

import "log"

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal("Error to connect with database")
	}
}
