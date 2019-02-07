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

	// e.g. localhost:3000/employees/962134/vacation
	r.GET("/employees/:id/vacation", func(c *gin.Context) {
		id := c.Param("id")
		timesOff, ok := TimesOff[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Page Not Found")
			return
		}

		c.HTML(http.StatusOK, "vacation-overview.html",
			map[string]interface{}{
				"TimesOff": timesOff,
			})
	})
	// http://localhost:3000/admin
	admin := r.Group("/admin")
	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html", nil)
	})

	r.Static("/public", "../public")
	//   or
	// r.StaticFS("/public", http.Dir("./public"))

	return r
}
