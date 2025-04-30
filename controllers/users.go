package controllers

import (
	"dev-assignment/middlewares"
	"dev-assignment/models"
	"encoding/json"
	"time"

	"github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Register() {
	var req struct {
		Username string
		Password string
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	err := models.RegisterUser(req.Username, req.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
	} else {
		c.Data["json"] = map[string]string{"message": "registration successful"}
	}
	c.ServeJSON()
}

func (c *UserController) Login() {
	var req struct {
		Username string
		Password string
	}
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)

	err := models.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		c.Ctx.Output.SetStatus(401)
		c.Data["json"] = map[string]string{"error": "invalid credentials"}
		c.ServeJSON()
		return
	}

	claims := &jwt.RegisteredClaims{
		Subject:   req.Username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, _ := token.SignedString(middlewares.JwtKey)

	c.Data["json"] = map[string]string{"token": tokenStr}
	c.ServeJSON()
}
