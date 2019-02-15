package util_jsoncrack

import (
	"fmt"
	"github.com/fwhezfwhez/jsoncrack"
	"testing"
)

func TestCRUD(t *testing.T){
	buf := []byte(`
        {
            "name": "ft",
            "class": {
                "name": "grade 7 14th",
                "master": "Li Hua"
            }
        }
    `)
	result,e :=Jc.Add(buf, "age", 9)
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(string(result))

	result2,e :=Jc.Add(result, "number", 47, "class")
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(string(result2))

	result3,e := Jc.Update(result2, "number", 50, "class")
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(string(result3))

	result4,e := Jc.Get(jsoncrack.BYTES, result3, "class", "number")
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(string(result4.([]byte)))

	result5, e:= Jc.Delete(jsoncrack.BYTES, false,result3, "class")
	if e!=nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(string(result5.([]byte)))
}
