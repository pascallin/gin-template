package auth

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pascallin/gin-template/pkg"
)

func findUserByUserName(username string) (error, *User) {
	user := &User{}
	err := pkg.MongoDB.DB.Collection("users").
		FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func login(username string, password string) (err error, token string) {
	err, user := findUserByUserName(username)
	if err != nil {
		return err, ""
	}
	p := md5.Sum([]byte(password))
	if user.Password != fmt.Sprintf("%x", p) {
		return errors.New("wrong password"), ""
	}
	gentoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		//"nbf": time.Date(2020, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, err := gentoken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return errors.New("generate token error: " + err.Error()), ""
	}
	return nil, tokenString
}

func register(username, password, nickname string) (error, primitive.ObjectID) {
	_, existUser := findUserByUserName(username)
	fmt.Println(existUser)
	if existUser != nil {
		return errors.New("username existed"), primitive.NilObjectID
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	p := md5.Sum([]byte(password))
	insertResult, err := pkg.MongoDB.DB.Collection("users").InsertOne(ctx, User{
		username,
		nickname,
		fmt.Sprintf("%x", p),
	})
	if err != nil {
		return err, primitive.NilObjectID
	}
	return nil, insertResult.InsertedID.(primitive.ObjectID)
}

func updatePassword(username, password, newPassword string) error {
	var user User
	p := md5.Sum([]byte(password))
	matchUser := pkg.MongoDB.DB.Collection("users").
		FindOne(context.Background(), bson.M{
			"username": username,
			"password": fmt.Sprintf("%x", p),
		})
	if matchUser == nil {
		return errors.New("username and old password not match")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	np := md5.Sum([]byte(newPassword))
	after := options.After
	err := pkg.MongoDB.DB.Collection("users").
		FindOneAndUpdate(ctx,
			bson.M{"username": username},
			bson.M{"$set": bson.M{"password": fmt.Sprintf("%x", np)}},
			&options.FindOneAndUpdateOptions{
				ReturnDocument: &after,
			},
		).Decode(&user)
	if err != nil {
		return errors.New("update password fail: " + err.Error())
	}
	return nil
}
