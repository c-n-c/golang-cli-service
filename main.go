package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"reflect"
	"strings"
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}


func RunCmd(c *gin.Context) {
	allowed := [...]string {"ls", "dir", "cat"}
	ctext := c.Query("text")
	arg := strings.Split(ctext, " ")
	cmd := exec.Command(arg[0], arg[1:]...)
	if itemExists(allowed,arg[0]) {
		out, err := cmd.CombinedOutput()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": string(out)})
		} else {
			c.JSON(http.StatusOK, gin.H{"data": string(out)})
		}
	} else {
		if arg[0] != "" {
			c.JSON(http.StatusForbidden, gin.H{"data": "command not allowed:" + arg[0]})
		} else {
			c.JSON(http.StatusForbidden, gin.H{"data": "empty upload"})
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
	r := gin.Default()

	r.GET("/", HomePage)

	r.GET("/runcmd", RunCmd)

	r.Run()
}

