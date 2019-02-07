package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// http://localhost:3000/login
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	// http://localhost:3000/admin
	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "./public")
	//   or
	// r.StaticFS("/public", http.Dir("./public"))

	return r
}
