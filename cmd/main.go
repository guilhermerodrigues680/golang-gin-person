package main

import (
	apiv1 "app/internal/api/v1"
	"app/internal/person"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
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

	r.Use(static.Serve("/", static.LocalFile("../web", true)))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
