/**
 * @author mch
 */

package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputReader *bufio.Reader

type  Cfunc func() (bool,error)

func init() {
	inputReader = bufio.NewReader(os.Stdin)
}

func CRead() string {
	input,_ := inputReader.ReadString('\n')
	return strings.TrimSpace(strings.TrimSuffix(input,"\n"))
}
func CReturn(cf Cfunc) bool {
	fmt.Println(">>>>>>>>>>>>>start >>>>>>>>>>>>>")
	flag,err := cf()
	if err != nil {
		fmt.Println("系统异常，",err)
	}
	fmt.Println("<<<<<<<<<<<<<<<end <<<<<<<<<<<<<")
	return flag
}
func Coper(operate []string) {
	for key,value := range operate {
		fmt.Printf("(%d) : %s \n",key,value)
	}
	fmt.Println("退出请输入 x")
}