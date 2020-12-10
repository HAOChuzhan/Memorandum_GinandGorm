package controller

import (
	"bubble/models"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

func Teacher(c *gin.Context) {
	valid := validation.Validation{}

	goal1 := c.PostForm("goal1")
	completed1 := c.PostForm("completed1")
	nextMonth_goal1 := c.DefaultPostForm("nextMonthGoal1", "0")

	goal2 := c.PostForm("goal2")
	completed2 := c.PostForm("completed2")
	nextMonth_goal2 := c.DefaultPostForm("nextMonthGoal1", "0")

	goal3 := c.PostForm("goal3")
	completed3 := c.PostForm("completed3")
	nextMonth_goal3 := c.DefaultPostForm("nextMonthGoal1", "0")

	goal4 := c.PostForm("goal3")
	completed4 := c.PostForm("completed3")
	nextMonth_goal4 := c.DefaultPostForm("nextMonthGoal1", "0")

	valid.Required(goal1, "goal1").Message("goal是必需的")
	valid.Required(completed1, "completed1").Message("completed是必需的")

	data1, data2 := models.TeacherScore("1", goal1, completed1, nextMonth_goal1)
	data3, data4 := models.TeacherScore("2", goal2, completed2, nextMonth_goal2)
	data5, data6 := models.TeacherScore("3", goal3, completed3, nextMonth_goal3)
	data7, data8 := models.TeacherScore("4", goal4, completed4, nextMonth_goal4)

	data := make(map[string]interface{})
	data["Rate1"] = data1
	data["Score1"] = data2
	data["Rate2"] = data3
	data["Score2"] = data4
	data["Rate3"] = data5
	data["Score3"] = data6
	data["Rate4"] = data7
	data["Score4"] = data8
	data["SumScore"] = data2 + data4 + data6 + data8
	c.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": data,
	})
}
