package controller

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pascallin/gin-template/conn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthController struct{}

type User struct {
	Username string `bson:"username" json:"username"`
	Nickname string `bson:"nickname" json:"nickname"`
	Password string `bson:"password" json:"password"`
}

type UsernameAndPasswordRequest struct {
	Username string `json:"username" form:"username" xml:"username" binding:"required"`
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}
type RegisterRequest struct {
	UsernameAndPasswordRequest
	Nickname string `json:"nickname" form:"nickname"`
}
type PatchPasswordRequest struct {
	UsernameAndPasswordRequest
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}

// @Summary user register
// @Description user register
// @Tags user
// @Accept  json
// @Param user body UsernameAndPasswordRequest true "register"
// @Produce  json
// @Success 200 {object} User
// @Router /user/register [post]
func (a AuthController) RegisterRoute(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, id := a.register(request.Username, request.Password, request.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary user login
// @Description user login
// @Tags user
// @Accept  json
// @Param user body UsernameAndPasswordRequest true "login"
// @Produce  json
// @Success 200 {object} User
// @Router /user/login [post]
func (a AuthController) LoginRoute(ctx *gin.Context) {
	var request UsernameAndPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err, token := a.login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary user patch password
// @Description user patch password
// @Tags user
// @Accept  json
// @Param user body PatchPasswordRequest true "login"
// @Produce  json
// @Success 200 {object} User
// @Router /user/password [patch]
func (a AuthController) PatchPasswordRoute(ctx *gin.Context) {
	var request PatchPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := a.updatePassword(request.Username, request.Password, request.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func findUserByUserName(username string) (error, *User) {
	user := &User{}
	err := conn.MongoDB.DB.Collection("users").
		FindOne(context.Background(), bson.M{"username": username}).Decode(user)
	if err != nil {
		return err, nil
	}
	return nil, user
}

func (a AuthController) login(username string, password string) (err error, token string) {
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

func (a AuthController) register(username, password, nickname string) (error, primitive.ObjectID) {
	_, existUser := findUserByUserName(username)
	fmt.Println(existUser)
	if existUser != nil {
		return errors.New("username existed"), primitive.NilObjectID
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	p := md5.Sum([]byte(password))
	insertResult, err := conn.MongoDB.DB.Collection("users").InsertOne(ctx, User{
		username,
		nickname,
		fmt.Sprintf("%x", p),
	})
	if err != nil {
		return err, primitive.NilObjectID
	}
	return nil, insertResult.InsertedID.(primitive.ObjectID)
}

func (a AuthController) updatePassword(username, password, newPassword string) error {
	var user User
	p := md5.Sum([]byte(password))
	matchUser := conn.MongoDB.DB.Collection("users").
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
	err := conn.MongoDB.DB.Collection("users").
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
