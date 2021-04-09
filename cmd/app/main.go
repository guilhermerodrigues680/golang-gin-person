package main

import (
	apiv1 "app/internal/api/v1"
	"app/internal/person"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/pkger"
)

func main() {
	start := time.Now()
	r := gin.Default()

	pRepo := person.NewPersonRepository()
	pServ := person.NewPersonService(pRepo)
	pContr := apiv1.NewPersonController(pServ)
	othContr := apiv1.NewOtherController()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/person", pContr.HandleReadAll)
		v1.POST("/person", pContr.HandleCreate)
		v1.GET("/person/:id", pContr.HandleRead)
		v1.PUT("/person/:id/make-birthday", pContr.HandleMakeBirthday)
		v1.GET("/ping", othContr.HandlePing)
	}

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/web")
	})

	// r.Static("/web", "../web")
	// r.StaticFS("/web", gin.Dir("../", false))
	r.StaticFS("/web", pkger.Dir("/web"))

	// "github.com/gin-contrib/static"
	// r.Use(static.Serve("/", static.LocalFile("../web", true)))

	fmt.Printf("started in %v\n", time.Since(start))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
