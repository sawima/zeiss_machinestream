package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/fatih/color"
)

func httpServer() (httpServer *http.Server) {
	httpServer = &http.Server{
		Addr:           *port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return
}

func routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "hello machinestream") })
	r.GET("/machines", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success fetch all machine status", "data": machineRecords()})
	})

	return r
}

func startHTTP() {
	httpServer := httpServer()
	httpServer.Handler = routers()
	c := color.New(color.FgCyan)
	c.Printf("HTTP Server is running on http://localhost%s \n", *port)
	httpServer.ListenAndServe()
}
