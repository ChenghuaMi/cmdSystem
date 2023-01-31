/**
 * @author mch
 */

package service

import (
	"cmdSystem/util"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	next string
	view  string
)
func Run() {
	next = "index::Welcome"
	for  {
		flag := util.CReturn(util.Cfunc(dispatch))
		if flag {
			break
		}

	}
	fmt.Println("end.......")


}
func dispatch() (bool,error) {
	args := strings.Split(next,"::")
	controller,ok := controllers[args[0]]
	if ok != true {
		return false,errors.New("控制器不存在" + args[0])
	}
	ctr := reflect.ValueOf(controller)
	ctr.MethodByName(args[1]).Call([]reflect.Value{})

	opers,ok :=  views[view]
	if ok !=  true {
		return false,errors.New("获取不到 试图:"  + view)
	}

	methods,desc := toModelFormat(opers)
	util.Coper(desc)

	for{
		input := util.CRead()
		if input == "x" {
			return true,nil
		}
		flag,err := strconv.Atoi(input)
		if err == nil && flag < len(methods) {
			next = methods[flag]
			break
		}
		fmt.Println("输入有误，请重新输入")
	}
	return false,nil

}