package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

type AddParams struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type PrintJob struct {
	JobId int `json:"jobId" binding:"required,gte=10000"`
	Pages int `json:"pages" binding:"required,gte=1,lte=100"`
}

func add(c *gin.Context) {
	var ap AddParams
	if err := c.ShouldBindJSON(&ap); err != nil {
		c.JSON(400, gin.H{"error": "Calculator error"})
		return
	}

	c.JSON(200, gin.H{"answer": ap.X + ap.Y})
}

func v1EndpointHandler(c *gin.Context) {
	c.String(200, "v1: %s %s", c.Request.Method, c.Request.URL.Path)
}

func v2EndpointHandler(c *gin.Context) {
	c.String(200, "v2: %s %s", c.Request.Method, c.Request.URL.Path)
}

func main() {
	router := gin.Default()

	// router test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/os", func(c *gin.Context) {
		c.String(200, runtime.GOOS)
	})

	v1 := router.Group("/v1")

	v1.GET("/products", v1EndpointHandler)
	v1.GET("/products/:productId", v1EndpointHandler)
	v1.POST("/products", v1EndpointHandler)
	v1.PUT("/products/:productId", v1EndpointHandler)
	v1.DELETE("/products/:productId", v1EndpointHandler)

	// v2 := router.Group("/v2")

	// v2.GET("/products", v1EndpointHandler)
	// v2.GET("/products/:productId", v1EndpointHandler)
	// v2.POST("/products", v1EndpointHandler)
	// v2.PUT("/products/:productId", v1EndpointHandler)
	// v2.DELETE("/products/:productId", v1EndpointHandler)

	router.GET("/add/:x/:y", add)
	router.POST("/add", add)

	router.POST("/print", func(c *gin.Context) {
		var p PrintJob
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input!"})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("PrintJob #%v started!", p.JobId)})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
