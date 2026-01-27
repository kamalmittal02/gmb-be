package main

import "github.com/joho/godotenv"

func main() {
	godotenv.Load()
	app := NewBaseApp()
	app.RunHttpServer()
}
