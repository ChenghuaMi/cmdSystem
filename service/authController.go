/**
 * @author mch
 */

package service

import (
	"cmdSystem/model"
	"cmdSystem/util"
	"fmt"
)

type  AuthController struct {

}
func (c *AuthController) Login() {
	view = "login_view"

	fmt.Println("请输入用户名:")
	username := util.CRead()

	fmt.Println("请输入密码：")
	password := util.CRead()

	user :=  model.GetUser(username)

	if user == nil {
		fmt.Println("用户不存在:",username)
		return
	}
	if user.GetPassword() == password {
		fmt.Println("登陆成功")
		view = "index_view"
		return
	}else {
		fmt.Println("密码错误")
		return
	}

}

func(c *AuthController) Register() {

}