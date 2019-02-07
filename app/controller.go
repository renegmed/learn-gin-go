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
	admin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
	}))

	admin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin-overview.html",
			map[string]interface{}{
				"Employees": employees,
			})
	})

	admin.GET("/employees/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "add" {
			c.HTML(http.StatusOK, "admin-employee-add.html", nil)
			return
		}

		employee, ok := employees[id]

		if !ok {
			c.String(http.StatusNotFound, "404 - Not Found")
			return
		}

		c.HTML(http.StatusOK, "admin-employee-edit.html",
			map[string]interface{}{
				"Employee": employee,
			})

	})

	//r.Static("/public", "../public")
	//   or
	// r.StaticFS("/public", http.Dir("./public"))

	return r
}
