/*
Copyright 2019 louis.
@Time : 2019/10/29 22:09
@Author : louis
@File : http
@Software: GoLand

*/

package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)


func defaultPage(g *gin.Context) {
	firstName := g.DefaultQuery("firstName", "test")
	lastName := g.Query("lastName")
	g.String(http.StatusOK, "Hello %s %s, This is My deploy Server~", firstName, lastName)
}


func main() {
	// Disable Console Color, you don't need console color when writing the logs to file
	//gin.DisableConsoleColor()
	// Logging to a file.
	f, _ := os.Create("/logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/", defaultPage)
	_ = router.Run(":8000")
}
