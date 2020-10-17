package controllers

import (
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const UserCollection = "user"

var (
	errNotExist        = errors.New("Users are not Exist")
	errInvalidID       = errors.New("Invalid ID")
	errInvalidBody     = errors.New("Invalid Request Body")
	errInsertionFailed = errors.New("Error in the user Insertion")
	errUpdationFailed  = errors.New("Error in the user Updation")
	errDeletionFailed  = errors.New("Error in the user Deletion")
)

func ReadAll(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Read All Successfully!",
		"code":    1,
	})
}

func Read(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Read Successfully!",
		"code":    1,
	})
}

func Create(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Create Successfully!",
		"code":    1,
	})
}

func Update(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Update Successfully!",
		"code":    1,
	})
}

func Delete(c *gin.Context){
	c.JSON(http.StatusCreated, gin.H{
		"message": "User Delete Successfully!",
		"code":    1,
	})
}