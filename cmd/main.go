package main

import (
	"log"

	"github.com/abdrakhmanovzh/notemaker/pkg/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	var routerGroup = router.Group("/")
	routes.RegisterNoteRoutes(routerGroup)

	router.LoadHTMLGlob("./ui/html/*")
	router.Static("/ui/static", "./ui/static/")
	log.Fatal(router.Run())
}
