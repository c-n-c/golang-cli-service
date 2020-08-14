package main

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"reflect"
	"strings"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", gin.H{"title": "Infra UI!",})
}


func RunCmd(c *gin.Context) {
	prohibited := [...]string {"ls"}
	ctext := c.Query("text")
	arg := strings.Split(ctext, " ")
	cmd := exec.Command(arg[0], arg[1:]...)
	if itemExists(prohibited, arg[0]) {
		c.JSON(http.StatusForbidden, gin.H{"data": "command not allowed:" + arg[0]})
	} else {
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": string(out)})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": string(out)})
		}
	}
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func main() {
	gin.ForceConsoleColor()

	r := gin.Default()

	r.HTMLRender = ginview.Default()

	r.Static("/assets", "./assets")

	r.GET("/", HomePage)

	r.GET("/runcmd", RunCmd)

	r.Run()
}

