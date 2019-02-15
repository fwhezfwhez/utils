package utils

import (
	"encoding/xml"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetryHandler(t *testing.T) {
	// a failure example
	f := func() (bool, error) {
		fmt.Println("connect to mysql")
		return false, errors.New("time out")
	}
	if e := RetryHandler(3, f); e != nil {
		fmt.Println(e.Error())
	}
}

func TestSmartPrint(t *testing.T) {
	type User struct {
		Name    string
		Age     int
		Sal     float64
		Friends []User
	}
	user := User{
		Name:    "ft",
		Age:     9,
		Sal:     1000,
		Friends: []User{User{"f1", 11, 0., nil}, {"f2", 12, 0., nil}},
	}
	SmartPrint(user)
}

func TestToLimitOffset(t *testing.T) {
	limit, offset := ToLimitOffset("20", "2", 34)
	t.Log(limit, offset)
}

func TestMD5(t *testing.T) {
	t.Log(MD5(`JJJ`))
}

type D struct{
}
func(d D) String() string{
	return "ok"
}
func TestToString(t *testing.T){
	fmt.Println(ToString(time.Now()))
	fmt.Println(ToString(1111.01))
	fmt.Println(ToString(123))
	d := D{}
	fmt.Println(ToString(d))
	a :="a"
	fmt.Println(ToString(&a))
}

func TestVXXML_ToParam(t *testing.T) {
	data := struct {
		XMLName     xml.Name `xml:"xml"`
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}{
		Name: "ft",
		Age:  9,
	}
	t.Log(ToParam(data, "xml"))
}

func TestAddJSONFormTag(t *testing.T) {
	fmt.Println(AddJSONFormTag(`
		type  U struct{
			Name string
			Age int
			Sal float64
		}
	`))
}
