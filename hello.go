package main
import (
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
  "github.com/prometheus/client_golang/prometheus/promhttp"
  "log"
)

func main() {
  router := gin.Default()

// Hello world
  router.GET("/hello", func(c *gin.Context) {
	  c.String(200, "Hello Remarkable!")
  })

// Basic API
  api := router.Group("/api")

  api.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

// Static page
  router.Use(static.Serve("/", static.LocalFile("./files", true)))

// Metrics
  router.GET("/metrics", gin.WrapH(promhttp.Handler()))

// Run
  err := router.Run(":8080")
  if err != nil {
    log.Fatal(err.Error())
  }
}
