package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	v0 := router.Group("/")
	{
		v0.GET("/", getRNG)
		v0.GET("/ping", getPong)
		v0.GET("/bing", getBong)
		v0.GET("/healthz", getHealthz)
	}

	return router
}

func getRNG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"random": 4})
}

func getPong(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func getBong(c *gin.Context) {
	c.String(http.StatusOK, "bong")
}

func getHealthz(c *gin.Context) {
	c.Status(http.StatusOK)
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := setupRouter()

	srv := &http.Server{
		Addr:         ":3000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Graceful shutdown initiated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Force shutdown initiated: ", err)
	}

	log.Println("Server exiting")
}
