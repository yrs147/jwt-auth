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

func ValidateToken(signToken string) (claims  *SignedDetails , msg string){
	token , err :=jwt.ParseWithClaims(
		signedToken , 
		&SignedDetails{},
		func (token *jwt.Token)(interface{}, error){
			return []byte(SECRET_KEY), nil 
		},

	)
	if err != nil {
		msg= err.Error()
		return
	}

	claims , ok := token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("The token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix(){
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims , msg
}

func UpdateAllTokens(signedToken string, signedRefreshToken string , userId string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj , bson.E{"token": signedToken})
	updateObj = append(updateObj , bson.E{"refresh_token": signedRefreshToken})

	Updated_at, + := time.Parse(time.RFC3339 , time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{"updated_at":Updated_at})

	upsert := true
	filter := bson.M{"user_id":userId}
	opt := options.UpdateOptions{
		Upsert : &upsert,
	}

	_, err := userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{"$set",updateObj},
		},
		&opt,

	)

	defer cancel()

	if err != nil {
		log.Panic(err)
		return 
	}
	return
}