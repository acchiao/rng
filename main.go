package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	v0 := router.Group("/")
	{
		v0.GET("/", getRNG)
		v0.GET("/rng", getRNG)
		v0.GET("/ping", getPong)
	}

	return router
}

func getRNG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"random": 4})
}

func getPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func main() {
	router := setupRouter()

	s := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
