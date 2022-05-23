/*
 * @Author: alexander.huang
 * @Date:   2022-05-19 20:14:50
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-19 20:14:50
 */
package app

import (
	"github.com/Anzz-bot/DouYin_demo/app/common/request"
	"github.com/Anzz-bot/DouYin_demo/app/common/response"
	"github.com/Anzz-bot/DouYin_demo/app/services"
	"github.com/Anzz-bot/DouYin_demo/global"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var form request.Register
	username := c.Query("username")
	password := c.Query("password")
	form = request.Register{
		Name:     username,
		Password: password,
	}
	if err, user := services.UserService.Register(form); err != nil {
		response.RegisterFail(c, global.Errors.BusinessError.ErrorCode, err.Error())
	} else {
		response.RegisterSuccess(c, user)
	}
}

func Login(c *gin.Context) {
	var form request.Login
	username := c.Query("username")
	password := c.Query("password")
	form = request.Login{
		Name:     username,
		Password: password,
	}
	if err, user := services.UserService.Login(form); err != nil {
		response.LoginFail(c, global.Errors.BusinessError.ErrorCode, err.Error())
	} else {
		//tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		//if err != nil {
		//	response.LoginFail(c, global.Errors.TokenError.ErrorCode, err.Error())
		//	return
		//}
		response.LoginSuccess(c, *user)
	}
}
