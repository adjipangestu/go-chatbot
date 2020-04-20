package main

import (
	"wa-chattbot/controllers"
)

var server = controllers.Server{}


func main() {
	server.Run(":8000")
}