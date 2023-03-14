package helper

import (
	"context"
	"fmt"
	"jwt-auth/database"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct {
	Email string
	Uid string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string , userType string, uid string) (signedToken string, signedRefreshToken string , err error){

	claims := &SignedDetails{
		Email : email, 
		Uid: uid , 
		User_type: userType,
		StandarClaims : jwt.StandardClaims{
			ExpiresAt: time.Now.Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: time.Now.Local().Add(time.Hour * time.Duration(168)).Unix(),
		},

		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
		refreshToken, err := jwt.NewWithClaims(jew.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

		if err != nil {
			log.Panic(err)
			return
		}

		return token, refreshToken, err


	}
}
