package main

import (
	"context"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"

	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.New()

	router.Use(
		gintrace.Middleware("rng"),
		cors.Default(),
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/healthz"}}),
		gin.Recovery(),
	)

	router.Use(gin.Recovery())

	v0 := router.Group("/")
	{
		v0.GET("/", getRNG)
		v0.GET("/ping", getPong)
		v0.GET("/bing", getBong)
		v0.GET("/healthz", getHealthz)
	}

	v2 := router.Group("/v2")
	{
		v2.GET("/", getRNGv2)
	}

	return router
}

func getRNG(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"random": 4})
}

func getRNGv2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"random": 5})
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
	tracer.Start(tracer.WithRuntimeMetrics())
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService("rng"),
		profiler.WithEnv("production"),
		profiler.WithVersion("0.1.0"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer profiler.Stop()

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
