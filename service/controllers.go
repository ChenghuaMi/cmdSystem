/**
 * @author mch
 */

package service

var (

	views map[string][][3]string
	controllers map[string]interface{}
)
func init() {
	views = make(map[string][][3]string,0)
	controllers = make(map[string]interface{},0)
	initViews()
	initController()
}
func initViews() {
	views["login_view"] = [][3]string{
		0: {0: "auth",1: "Login",2: "登陆"},
		1: {0: "auth",1: "Register",2: "注册"},

	}
	views["index_view"] = [][3]string{
		0: {0: "index",1: "Index",2: "首页"},
		1: {0: "user",1: "List",2: "展示信息"},
	}
}
func initController() {
	controllers["index"] = &IndexController{}
	controllers["user"] = &UserController{}
	controllers["auth"] = &AuthController{}
}

func toModelFormat(opers [][3]string) ([]string,[]string) {
	var method []string = make([]string,len(opers))
	var desc []string = make([]string,len(opers))
	for k,v := range opers {
		method[k] = v[0] + "::" + v[1]
		desc[k] = v[2]
	}
	return method,desc
}