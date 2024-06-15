package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koizumi7010/go-todo-api/domain/models"
	"github.com/koizumi7010/go-todo-api/usecase"
)

type ToDo interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetAll(c *gin.Context)
}

type todoHandler struct {
	usecase usecase.ToDo
}

func NewTodo(u usecase.ToDo) ToDo {
	return &todoHandler{u}
}

type createRequestParam struct {
	Task string `json:"task" binding:"required,max=60"`
}

func (th *todoHandler) Create(c *gin.Context) {
	var req createRequestParam
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.usecase.Create(req.Task); err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusCreated, nil)
}

type DeleteRequestParam struct {
	ID int `uri:"id"`
}

func (th *todoHandler) Delete(c *gin.Context) {
	var req DeleteRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.usecase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type UpdateRequestPathParam struct {
	ID int `uri:"id"`
}

type UpdateRequestBodyParam struct {
	Task   string            `json:"task" binding:"required,max=60"`
	Status models.TaskStatus `json:"status" binding:"required,tasks_status"`
}

func (th *todoHandler) Update(c *gin.Context) {
	var reqPathParam UpdateRequestPathParam
	if err := c.ShouldBindUri(&reqPathParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reqBody UpdateRequestBodyParam
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := th.usecase.Update(reqPathParam.ID, reqBody.Task, reqBody.Status); err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

type GetRequestParam struct {
	ID int `uri:"id" binding:"required"`
}

func (th *todoHandler) Get(c *gin.Context) {
	var req GetRequestParam
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	td, err := th.usecase.Get(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, td)
}

func (th *todoHandler) GetAll(c *gin.Context) {
	tds, err := th.usecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tds)
}
