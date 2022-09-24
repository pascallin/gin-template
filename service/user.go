package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/pascallin/gin-template/conn"
	"github.com/pascallin/gin-template/model"
	"github.com/pascallin/gin-template/types"
)

type UserService interface {
	Login(username string, password string) (token string, err error)
	CreteUser(username, password, nickname string) (primitive.ObjectID, error)
	UpdateUserPassword(username, password, newPassword string) error
	FindUserByUserName(username string) (*model.User, error)
}

var (
	ErrWrongPassword = types.NewAppError(500, "WRONG_PASSWORD")
	ErrUserNotFound  = types.NewAppError(500, "USER_NOT_FOUND")
)

func Login(username string, password string) (token string, err error) {
	user, err := FindUserByUserName(username)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", ErrUserNotFound
		}
		return "", err
	}
	p := md5.Sum([]byte(password))
	if user.Password != fmt.Sprintf("%x", p) {
		return "", ErrWrongPassword
	}

	claims := model.CustomerClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(3600 * time.Second).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "test",
			NotBefore: 0,
			Subject:   "",
		},
	}

	gentoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := gentoken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("generate token error: " + err.Error())
	}
	return tokenString, nil
}

func CreteUser(username, password, nickname string) (primitive.ObjectID, error) {
	existUser, err := FindUserByUserName(username)
	if err != nil && err != mongo.ErrNoDocuments {
		return primitive.NilObjectID, err
	}
	if existUser != nil {
		return primitive.NilObjectID, errors.New("username existed")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	p := md5.Sum([]byte(password))
	insertResult, err := conn.GetMongo(ctx).DB.Collection("users").InsertOne(ctx, model.User{
		Username: username,
		Nickname: nickname,
		Password: fmt.Sprintf("%x", p),
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	return insertResult.InsertedID.(primitive.ObjectID), nil
}

func UpdateUserPassword(username, password, newPassword string) error {
	var user model.User
	p := md5.Sum([]byte(password))
	ctx := context.TODO()
	matchUser := conn.GetMongo(ctx).DB.Collection("users").
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
	err := conn.GetMongo(ctx).DB.Collection("users").
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

func FindUserByUserName(username string) (*model.User, error) {
	user := &model.User{}
	ctx := context.TODO()
	err := conn.GetMongo(ctx).DB.Collection("users").
		FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
