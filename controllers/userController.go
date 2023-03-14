package controllers

import(
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "jwt-auth/helpers"
	"jwt-auth/models"
	"jwt-auth/helpers"
	"golang.org/crypto/bcrypt"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()


func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			
		}
	}
}