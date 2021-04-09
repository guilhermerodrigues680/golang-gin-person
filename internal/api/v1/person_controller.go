package apiv1

import (
	"app/internal/person"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type personService interface {
	Create(name string, age int) (*person.Person, error)
	ReadAll() ([]*person.Person, error)
	Read(id int) (*person.Person, error)
	MakeBirthday(id int) (*person.Person, error)
}

type personController struct {
	s personService
}

func NewPersonController(s personService) *personController {
	return &personController{s: s}
}

func (pc *personController) HandleReadAll(c *gin.Context) {

	ps, err := pc.s.ReadAll()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Deu ruim: %s", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"persons": ps,
	})
}

func (pc *personController) HandleCreate(c *gin.Context) {
	type personRequest struct {
		Name string `json:"name" binding:"required"`
		Age  int    `json:"age" binding:"required"`
	}

	var pReq personRequest

	if err := c.BindJSON(&pReq); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("Deu ruim: %s", err))
		return
	}

	p, err := pc.s.Create(pReq.Name, pReq.Age)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Deu ruim: %s", err))
		return
	}

	c.JSON(http.StatusOK, p)
}

func (pc *personController) HandleRead(c *gin.Context) {
	type uriParams struct {
		Id int `uri:"id" binding:"required"`
	}

	var uriP uriParams
	if err := c.ShouldBindUri(&uriP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Deu ruim: %s", err)})
		return
	}

	p, err := pc.s.Read(uriP.Id)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Deu ruim: %s", err))
		return
	}

	c.JSON(http.StatusOK, p)
}

func (pc *personController) HandleMakeBirthday(c *gin.Context) {
	type uriParams struct {
		Id int `uri:"id" binding:"required"`
	}

	var uriP uriParams
	if err := c.ShouldBindUri(&uriP); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintf("Deu ruim: %s", err)})
		return
	}

	p, err := pc.s.MakeBirthday(uriP.Id)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Deu ruim: %s", err))
		return
	}

	c.JSON(http.StatusOK, p)
}
