package app

import (
	"log"
	"net/http"
	"strconv"
	"todoapp/model"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
func createHandler(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	msg := c.PostForm("msg")
	data := model.Create(msg)
	c.JSON(http.StatusCreated, data)
}
func getlistHandler(c *gin.Context) {
	data := model.Read()
	c.JSON(http.StatusOK, data)
}
func deleteHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	success := model.Delete(id)
	c.JSON(http.StatusOK, success)
}

func NewApp(port string) {
	model.NewModel()
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.Group("/api")
	{
		router.GET("/", indexHandler)
		router.POST("/todos", createHandler)
		router.GET("/todos", getlistHandler)
		router.DELETE("/todos/:id", deleteHandler)
	}

	router.Run(port)
}
