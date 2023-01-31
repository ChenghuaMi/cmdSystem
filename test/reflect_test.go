/**
 * @author mch
 */

package test

import (
	"fmt"
	"reflect"
	"testing"
)

type stu struct {
	A int
}
func TestZeroA(t *testing.T) {
	s := stu{
		A: 10,
	}
	v := reflect.ValueOf(s)
	field := v.Field(0)
	fmt.Printf("%T\n",field)
	fmt.Printf("%T\n",field.Interface())
	fmt.Printf("%T\n",reflect.Zero(field.Type()))
	fmt.Println(reflect.Zero(field.Type()))
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>.")
	fmt.Println(field.Interface(),reflect.Zero(field.Type()))
	fmt.Println(reflect.DeepEqual(field.Interface(),reflect.Zero(field.Type())))
}
func TestZeroB(t *testing.T) {
	s := struct {
		A int
	}{0}
	field := reflect.ValueOf(s).Field(0)
	fmt.Println(field.Interface())

	// use Zero()
	fmt.Println(reflect.Zero(field.Type()))
	fmt.Println(reflect.TypeOf(field.Interface()))

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(reflect.TypeOf(reflect.Zero(field.Type())))
	fmt.Println(">>>>>>>>>")
	fmt.Println(reflect.DeepEqual(field.Interface(),int(reflect.Zero(field.Type()).Int())))

	fmt.Println(">>>>????????")
	fmt.Println(reflect.DeepEqual(s, struct {
		A int
	}{}))
	fmt.Println(reflect.DeepEqual(s, struct {
		A int
	}{0}))
}
