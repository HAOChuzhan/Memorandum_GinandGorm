package controller

import (
	"bubble/models"
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端页面填写相关数据 点击提交 会发请求到这里
	// 1. 从请求中把数据拿出来
	valid := validation.Validation{} //数据验证
	//c.BindJSON(&todo)
	corporate_name := c.PostForm("corporate_name")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	advice := c.PostForm("advice")
	valid.Required(corporate_name, "corporate_name").Message("CorporateName is necessary!")
	valid.Required(name, "name").Message("Name is necessary!")
	valid.Required(phone, "phone").Message("Phone number is necessary!")
	valid.Phone(phone, "phone").Message("电话或固话格式不正确")

	valid.Required(email, "email").Message("Email is necessary!")
	valid.Email(email, "email").Message("邮箱格式不正确")

	if valid.HasErrors() { // 有错误
		for _, err := range valid.Errors { // 循环打印错误
			fmt.Printf("%s：%s\n", err.Key, err.Message)
		}
		return
	}
	//fmt.Println(models.FilteredSQLInject(corporate_name), models.FilteredSQLInject(name), models.FilteredSQLInject(phone), models.FilteredSQLInject(email), models.FilteredSQLInject(advice))

	if models.FilteredSQLInject(corporate_name) || models.FilteredSQLInject(name) || models.FilteredSQLInject(phone) || models.FilteredSQLInject(email) || models.FilteredSQLInject(advice) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "禁止发送违规字符！",
		})
		return
	}

	todo := models.Todo{
		CorporateName: corporate_name,
		Name:          name,
		Phone:         phone,
		Email:         email,
		Advice:        advice,
	}

	if err := models.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, todo)
	//models.SendMsg()

}

func GetTodoList(c *gin.Context) {
	// 查询todo这个表里的所有数据
	todoList, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "获取成功！",
			"data": todoList,
		})
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
