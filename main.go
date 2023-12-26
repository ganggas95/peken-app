package main

import (
	"peken-be/app"

	"github.com/joho/godotenv"
)

func Init() {
	app.InitLog()
	godotenv.Load()
}

func main() {
	Init()
	router := InitializedServer()

	router.Run(":8080")

}
