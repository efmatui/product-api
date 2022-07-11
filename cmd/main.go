package main

import (
	"github.com/efmatui/product-api/internal/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/product", app.GetAllProduct())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
