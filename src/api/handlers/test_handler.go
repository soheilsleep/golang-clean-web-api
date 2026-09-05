package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}
type header struct {
	UserId  string
	Browser string
}
type personData struct {
	FirstName string
	LastName  string
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (handler *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}
func (handler *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Users",
	})
}
func (handler *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UserById",
		"id":     id,
	})
}
func (handler *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"result":   "UserByUsername",
		"username": username,
	})
}
func (handler *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
		"id":     id,
	})
}
func (handler *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": "Accounts",
	})
}

func (handler *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder1",
		"userId": userId,
	})

}
func (handler *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"result": "HeaderBinder2",
		"Header": header,
	})

}
func (handler *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder1",
		"id":     id,
		"name":   name,
	})
}
func (handler *TestHandler) QueryBinder2(c *gin.Context) {
	name := c.Query("name")
	ids := c.QueryArray("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "QueryBinder2",
		"ids":    ids,
		"name":   name,
	})
}
func (handler *TestHandler) UriBinder(c *gin.Context) {
	name := c.Param("name")
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "UriBinder",
		"id":     id,
		"name":   name,
	})
}
func (handler *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBindJSON(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "BodyBinder",
		"Person": p,
	})
}
func (handler *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBind(&p)
	c.JSON(http.StatusOK, gin.H{
		"result": "FormBinder",
		"Person": p,
	})
}
