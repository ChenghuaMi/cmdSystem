/**
 * @author mch
 */

package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
	"unsafe"
	"context"
)


type RetData struct {
	Result struct{RETCODE int}
	Name string
	RESULTLIST map[string]struct{Price string}
}
func TestJsonUnmarshal(t *testing.T) {
	url := "http://wbt-api.zjfh.top/api/property/get_cq_price"
	resp,err := http.Get(url)
	if err != nil {
		log.Fatal("err:",err)
	}
	defer resp.Body.Close() // resp.Body => 实现了 io.ReadCloser
	var ret RetData
	body,_ := ioutil.ReadAll(resp.Body) // body => []byte
	err = json.Unmarshal(body,&ret) // 将 []byte 写入结构体
	fmt.Println("er:",err)
	fmt.Println(ret)
}
func TestJsonDecoder(t *testing.T) {
	url := "http://wbt-api.zjfh.top/api/property/get_cq_price"
	resp,err := http.Get(url)
	if err != nil {
		log.Fatal("err:",err)
	}
	defer resp.Body.Close()
	var ret RetData
	err = json.NewDecoder(resp.Body).Decode(&ret) // 直接从流里面 转化
	fmt.Println("errrrr:",err)
	fmt.Println(ret)
}
func TestFunc(t *testing.T) {
	var a int = 4
	var b *int = &a
	fmt.Printf("b ptr = %p\n",b)
	fmt.Printf("p p ptr = %p\n",&b)
	do(b)
	fmt.Println(a)
}
func do(p *int) {
	*p = 5
}
func TestSafe(t *testing.T) {
	var a int = 4
	var b int64 = int64(a)
	var c *int = (*int)(unsafe.Pointer(&b))
	fmt.Println("&c=",c)
	fmt.Println("c=",*c)
}
func TestPoint(t *testing.T) {
	var a int64 = 10
	fmt.Printf("%#v\n",a)
	fmt.Printf("%#v\n",&a)
	c := (*int)(unsafe.Pointer(&a))
	fmt.Printf("%#v\n",c)
	fmt.Printf("%#v\n",*c)
	fmt.Println(">>>>>>>>>>>>>>>>.")
	b := (*string)(unsafe.Pointer(&a))
	fmt.Printf("%#v\n",b)
	//fmt.Printf("%#v\n",*b)
	*c = 20
	fmt.Println(a)
}

func TestWithDeadline(t *testing.T) {

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second * 5))
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("err:", ctx.Err())
			return
		case <-normal(4):
			fmt.Println("xxx")
			return
		}

	}

}
func normal(my int) chan string {
	result := make(chan string)
	go func() {
		time.Sleep(time.Duration(my) * time.Second)
		result <- "abc"
		return
	}()
	return result

}

func TestSprint(t *testing.T) {
	s := fmt.Sprint("a",1,"c")   // 可以 拼接不同类型 的元素，底层 使用 []byte 字节切片
	fmt.Println(s)
	s1 := fmt.Sprintf("%s%d","a",1)
	fmt.Println(s1)
	s2 := fmt.Sprintln("a",1,"b")
	fmt.Println(s2)
}