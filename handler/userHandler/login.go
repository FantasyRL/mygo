package userHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/UserService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserLogin(ctx context.Context, c *app.RequestContext) (*UserService.UserService, error, model.Response) {
	var userLogin UserService.UserService
	if err := c.BindAndValidate(&userLogin); err == nil { //相当于gin shouldBind 将请求携带的参数和后端的结构体绑定起来
		res := userLogin.Login()
		return &userLogin, nil, res
	} else {
		return &userLogin, err, model.Response{
			Status: e.InvalidParams,
			Msg:    "参数出错",
		}
	}
}