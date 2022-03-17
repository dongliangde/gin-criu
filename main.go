// @title 热迁移
// @version 1.0
// @description swagger server api
package main

import (
	_ "go-criu/docs"
	"go-criu/routes"
	"log"
)

func main() {
	log.Printf("Server started api address 0.0.0.0:swagger/index.html")
	routes.InitRouter()
}
