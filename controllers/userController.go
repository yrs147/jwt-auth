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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()


func HashPassword()

func VerifyPassword()

func Signup() gin.HandlerFunc{
	return func(c *gin.Context){
		var ctx , cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadrequest, gin.H{"error": validateErr.Error()})
		}

		count , err := userColletion.CountDocuments(ctx ,bson.M{"email":user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error":"error occured while checking for the Email "})
		}

		if count >0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error":"this email already exists"} )
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()
		token , refreshToken, _ := helper.GenerateAllTokens(*user.Email, *user.Username, *user.User_type, *&user.User_id)
		user.Token = &token
		user.Refersh_token = &refershToken

		resultInsertionNumber, insertErr := userCollection.InsertOne(ctx , user)
		if insertErr != nil {
			msg = fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cance()
		c.JSON(http.StatusOK, resultInsertionNumber)



	}
}

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User 
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err !=nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOk, user)

	}
}