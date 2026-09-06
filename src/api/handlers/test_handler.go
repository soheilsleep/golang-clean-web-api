package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soheilsleep/golang-clean-web-api/api/helper"
)

type TestHandler struct {
}
type header struct {
	UserId  string
	Browser string
}
type personData struct {
	FirstName    string `json:"first_name" binding:"required,min=4,alpha,max=10"`
	LastName     string `json:"last_name" binding:"required,min=5,alpha,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (handler *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Test", true, 0))
}
func (handler *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Users!", true, 0))
}
func (handler *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "UserById",
		"id":     id,
	}, true, 0))
}
func (handler *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "UserByUsername",
		"username":     username,
	}, true, 0) )
}
func (handler *TestHandler) Accounts(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "Accounts",
		"id":     id,
	}, true, 0))
}
func (handler *TestHandler) AddUser(c *gin.Context) {

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "Accounts",
	}, true, 0))
}

func (handler *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("userId")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "HeaderBinder1",
		"userId": userId,

	}, true, 0))

}
func (handler *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "HeaderBinder2",
		"userId": header,

	}, true, 0))

}
func (handler *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "QueryBinder1",
		"id": id,
		"name": name,

	}, true, 0))
}
func (handler *TestHandler) QueryBinder2(c *gin.Context) {
	name := c.Query("name")
	ids := c.QueryArray("id")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "QueryBinder2",
		"ids": ids,
		"name": name,

	}, true, 0))
}
func (handler *TestHandler) UriBinder(c *gin.Context) {
	name := c.Param("name")
	id := c.Param("id")

	c.JSON(http.StatusOK,helper.GenerateBaseResponse( gin.H{
		"result": "QueryBinder2",
		"id": id,
		"name": name,

	}, true, 0))
}
func (handler *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusBadRequest,helper.GenerateBaseResponseWithValidationError("validation error",false,1,err))
		return
	}

	c.JSON(http.StatusOK,helper.GenerateBaseResponse(gin.H{
		"result": "BodyBinder",
		"Person": p,
	},true,0))
}
func (handler *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.ShouldBind(&p)
	c.JSON(http.StatusOK,helper.GenerateBaseResponse(gin.H{
		"result": "FormBinder",
		"Person": p,
	},true,0) )

}
