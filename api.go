package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func apiFileList(c *gin.Context) {
	files, err := ioutil.ReadDir(documentDir)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err,
		})
	} else {
		c.JSON(200, gin.H{
			"files": createFileList(files),
		})
	}
}

func apiReadFile(c *gin.Context) {
	fileName := c.Param("path")
	filePath := filepath.Join(documentDir, fileName)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "No file found named %s.", filePath)
	} else {
		c.String(http.StatusOK, string(file))
	}
}

func apiCloneFile(c *gin.Context) {
	type Clone struct {
		NewName string `form:"newName" json:"newName" xml:"newName" binding:"required"`
	}
	var json Clone
	if err := c.ShouldBindJSON(&json); err != nil {
		c.String(http.StatusBadRequest, "Error: Bad parse of data, fill the field 'newName'")
		return;
	}
	originalFileName := c.Param("path")
	newFileName := filepath.Join(documentDir, json.NewName)
	filePath := filepath.Join(documentDir, originalFileName)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.String(http.StatusNotFound, "No file found named %s.", filePath)
		return
	}
	fmt.Printf("Output to %s", newFileName)
	ioutil.WriteFile(newFileName, file, 777)
	c.String(http.StatusOK, string(file))
}
