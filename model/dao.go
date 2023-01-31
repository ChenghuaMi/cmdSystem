/**
 * @author mch
 */

package model

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var (
	path string = "E:\\go\\src\\cmdSystem\\data\\"
	suffix string  = ".sql"
	models map[string]interface{}
)
type Model interface {
	ToString() string
	Save() bool
}

func init() {
	models = make(map[string]interface{})
	models["user"] = NewUser

	userDatas = make(map[string]Model,0)
	rfdata("user","username",userDatas)
}

func rfdata(name,primary string,datas map[string]Model) error {
	f,err := os.Open(path + name + suffix)
	if err != nil {
		fmt.Println("文件读取异常:",err)
		return errors.New("error:" + err.Error())
	}
	defer f.Close()
	buf := bufio.NewReader(f)

	field := make([]string,0)
	for  {
		row,err := buf.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("读取文件异常,",err)
		}
		data := strings.Split(string(row),",")
		if len(field) ==  0 {
			field = data
			for k,v := range data {
				field[k] = strings.TrimSpace(strings.TrimSuffix(v,"\n"))
			}
		}else {
			toModel(name,primary,datas,data,field)
		}

	}
	return nil
}

func toModel(name,primary string,datas map[string]Model,data,field []string) error {
	if models[name] == nil {
		return errors.New(" 模型不存在:" + name)
	}
	modelV :=  reflect.ValueOf(models[name]).Call([]reflect.Value{})[0]

	var primaryValue string
	for k,v := range data{
		if field[k] == primary {
			primaryValue = v
		}
		fset := modelV.MethodByName("Set" + strings.Title(field[k]))
		fset.Call([]reflect.Value{
			reflect.ValueOf(toTypeValue(modelV,field[k],v)),
		})
	}
	datas[primaryValue] = modelV.Interface().(Model)
	return nil
}
func toTypeValue(modelV reflect.Value,field,value string) interface{}  {
	mtype := modelV.Elem().FieldByName(field).Type().Name()
	switch mtype {
	case "int":
		b,_ := strconv.Atoi(value)
		return b
	}
	return string(value)
}
func fwrite(name string,models map[string]Model) bool {
	content := getModelsToString(models)
	f,err := os.OpenFile(path + name + suffix,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Println("file not found ")
		return false
	}
	defer f.Close()
	out := bufio.NewWriter(f)
	out.WriteString(content + "\n")
	out.Flush()
	return true
}
func getModelsToString(models map[string]Model) string {
	var fields string
	var content string
	for _,model := range models{
		if fields == "" {
			rmodel := reflect.ValueOf(model)
			modelZ := reflect.Zero(rmodel.Type().Elem())
			for i := 0;i < modelZ.NumField();i++ {
				fields = fields + modelZ.Type().Field(i).Name + ","
			}
			fields = strings.TrimSuffix(fields,",")
		}
		content = content + model.ToString() + "\n"
	}
	return fields + "\n" + strings.TrimSuffix(content,"\n")
}